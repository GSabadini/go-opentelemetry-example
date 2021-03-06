package usecase

import (
	"context"
	"log"

	otelcodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"go-opentelemetry-example/domain"
)

type CreateAccountUC interface {
	Execute(context.Context, CreateAccountInput) (CreateAccountOutput, error)
}

type CreateAccountInput struct {
	ID string `json:"id"`
}

type CreateAccountOutput struct {
	ID string `json:"id"`
}

type createAccount struct {
	repo   domain.UserCreator
	tracer trace.Tracer
}

func NewCreateAccount(repo domain.UserCreator, tracer trace.Tracer) CreateAccountUC {
	return createAccount{
		repo:   repo,
		tracer: tracer,
	}
}

func (c createAccount) Execute(
	ctx context.Context,
	input CreateAccountInput,
) (CreateAccountOutput, error) {
	ctx, span := c.tracer.Start(ctx, "usecase::create_user")
	defer span.End()

	var user = domain.NewUser(input.ID)

	if err := c.repo.Create(ctx, user); err != nil {
		log.Print("Usecase execute error", err)

		span.SetStatus(otelcodes.Error, "Handler execute error")
		span.RecordError(err)

		return CreateAccountOutput{}, err
	}

	log.Print("Usecase execute success")

	span.SetStatus(otelcodes.Ok, "Usecase execute success")

	return CreateAccountOutput{
		ID: user.ID(),
	}, nil
}
