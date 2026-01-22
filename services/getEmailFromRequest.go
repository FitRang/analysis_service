package services

import (
	"context"

	"github.com/Foxtrot-14/FitRang/analysis-service/apperror"
	"github.com/Foxtrot-14/FitRang/analysis-service/middleware"
)

func getEmailFromContext(ctx context.Context) (string, error) {
	email, ok := ctx.Value(middleware.EmailKey).(string)
	if !ok || email == "" {
		return "", apperror.New(
			apperror.Unauthenticated,
			"Authentication required",
		)
	}
	return email, nil
}
