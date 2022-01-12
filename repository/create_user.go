package repository

import (
	"context"
	"log"

	otelcodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"go-opentelemetry-example/domain"
)

type createAccount struct {
	tracer trace.Tracer
}

func NewCreateAccount(tracer trace.Tracer) domain.UserCreator {
	return createAccount{
		tracer: tracer,
	}
}

func (c createAccount) Create(ctx context.Context, user domain.User) error {
	ctx, span := c.tracer.Start(ctx, "repository::create_user")
	defer span.End()

	span.SetStatus(otelcodes.Ok, "Repository execute success")

	log.Print("Repository execute success")
	return nil
}
