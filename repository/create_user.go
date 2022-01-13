package repository

import (
	"context"
	"log"

	otelcodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"go-opentelemetry-example/domain"
	"go-opentelemetry-example/infrastructure/memdb"
)

type createAccount struct {
	db     memdb.MemDB
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

	if err := c.db.Insert(ctx, user); err != nil {
		span.SetStatus(otelcodes.Error, "Repository execute error")
		span.RecordError(err)

		return err
	}

	log.Print("Repository execute success")

	span.SetStatus(otelcodes.Ok, "Repository execute success")

	return nil
}
