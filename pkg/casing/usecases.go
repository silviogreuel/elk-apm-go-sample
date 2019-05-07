package casing

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/silviogreuel/elk-apm-go-sample/pkg/utils/uhttp"
	"github.com/silviogreuel/elk-apm-go-sample/pkg/utils/urand"
	"github.com/silviogreuel/elk-apm-go-sample/pkg/utils/ures"
	"go.elastic.co/apm"
)

func CaseValue(ctx context.Context, casing, value string) (string, error) {
	span, _ := apm.StartSpan(ctx, "usecases.CaseValue", "usecases")
	defer span.End()

	var host string

	if casing == "random" {
		p := urand.RandPercent()
		if p < 10 {
			casing = "panic"
		} else if p < 20 {
			casing = "empty"
		} else if p < 60 {
			casing = "upper"
		} else if p < 100 {
			casing = "lower"
		}
	}

	switch casing {
	case "upper":
		host = "http://ucase:8080"
		break
	case "lower":
		host = "http://lcase:8080"
		break
	case "empty":
		host = ""
		break
	case "panic":
		panic(fmt.Sprintf("forced panic by casing %s and value %s", casing, value))
	}

	res, err := uhttp.Get(ctx, fmt.Sprintf("%s/casing/%s/%s", host, casing, value))
	if err != nil {
		return "", errors.WithStack(err)
	}

	return ures.BodyAsString(res)
}

func ToUpper(ctx context.Context, value string) string {
	span, _ := apm.StartSpan(ctx, "usecases.ToUpper", "usecases")
	defer span.End()

	return strings.ToUpper(value)
}

func ToLower(ctx context.Context, value string) string {
	span, _ := apm.StartSpan(ctx, "usecases.ToLower", "usecases")
	defer span.End()

	return strings.ToLower(value)
}
