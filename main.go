package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel"

	"go-opentelemetry-example/handler"
	"go-opentelemetry-example/infrastructure/opentelemetry"
	"go-opentelemetry-example/repository"
	"go-opentelemetry-example/usecase"
)

const (
	PORT           = 8080
	DefaultAppName = ""
)

func main() {
	shutdownTracer := opentelemetry.NewTracer()
	defer shutdownTracer()

	//shutdownMetric := opentelemetry.InitMetric()
	//defer shutdownMetric()

	//meter := global.Meter("demo-server-meter")
	//serverAttribute := attribute.String("server-attribute", "foo")
	//commonLabels := []attribute.KeyValue{serverAttribute}
	//requestCount := metric.Must(meter).NewInt64Counter(
	//	"demo_server/request_counts",
	//	metric.WithDescription("The number of requests received"),
	//)

	r := mux.NewRouter()
	r.Use(otelmux.Middleware(DefaultAppName))

	tracer := otel.Tracer(DefaultAppName)

	createUserHandler := handler.NewCreateUser(
		usecase.NewCreateAccount(
			repository.NewCreateAccount(tracer),
			tracer,
		),
		tracer,
	)

	r.HandleFunc("/users", createUserHandler.Handle).Methods(http.MethodPost)

	log.Print("Start server in port:", PORT)

	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), r)
	if err != nil {
		log.Fatalln("Error start server", err)
		return
	}
}
