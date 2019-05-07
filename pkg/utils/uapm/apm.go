package uapm

import (
	"context"

	"go.elastic.co/apm"
)

func StartSpan(ctx context.Context, name, spanType string) *apm.Span {
	span, _ := apm.StartSpan(ctx, name, spanType)
	return span
}
