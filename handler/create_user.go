package handler

import (
	"encoding/json"
	"go.opentelemetry.io/otel/trace"
	"log"
	"net/http"

	otelcodes "go.opentelemetry.io/otel/codes"

	"go-opentelemetry-example/usecase"
)

type CreateUser struct {
	uc     usecase.CreateAccountUC
	tracer trace.Tracer
}

func NewCreateUser(uc usecase.CreateAccountUC, tracer trace.Tracer) CreateUser {
	return CreateUser{
		uc:     uc,
		tracer: tracer,
	}
}

func (c CreateUser) Handle(w http.ResponseWriter, r *http.Request) {
	ctx, span := c.tracer.Start(r.Context(), "handler::create_user")
	defer span.End()

	var input usecase.CreateAccountInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Print("Handler execute error", err)
		_ = response(w, err, http.StatusBadRequest)
	}
	defer r.Body.Close()

	output, err := c.uc.Execute(ctx, input)
	if err != nil {
		log.Print("Handler execute error", err)
		_ = response(w, err, http.StatusInternalServerError)
	}

	span.SetStatus(otelcodes.Ok, "Handler execute success")

	log.Print("Handler execute success")
	_ = response(w, output, http.StatusCreated)
}

func response(w http.ResponseWriter, output interface{}, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(output)
}
