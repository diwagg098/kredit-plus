package utilities

import (
	"context"
	"diwa/kredit-plus/pkg/models"
	"fmt"
)

func Logf(ctx context.Context, message string, value ...interface{}) context.Context {
	v, ok := ctx.Value(models.LogKey).(*models.Data)
	if ok {
		msg := fmt.Sprintf(message, value...)
		v.Messages = append(v.Messages, msg)

		ctx = context.WithValue(ctx, models.LogKey, v)

		return ctx
	}
	return ctx
}

func GetDataCTX(ctx context.Context) *models.Data {
	v, ok := ctx.Value(models.LogKey).(*models.Data)
	if ok {
		return v
	}
	return nil
}
