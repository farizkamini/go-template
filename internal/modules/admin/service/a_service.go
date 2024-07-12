package service_admin_auth

import (
	"context"
	"errors"
	"fmt"
)

type Service struct {
	Ctx context.Context
}

func New(ctx context.Context) *Service {
	return &Service{
		Ctx: ctx,
	}
}

func errService(err error) error {
	src := "service_admin_auth"
	return errors.New(
		fmt.Sprintf("%s | %+v", src, err),
	)
}
