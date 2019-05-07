package uhttp

import (
	"context"
	"net/http"

	"go.elastic.co/apm/module/apmhttp"
	"golang.org/x/net/context/ctxhttp"
)

var client = apmhttp.WrapClient(http.DefaultClient)

func Get(ctx context.Context, url string) (*http.Response, error) {
	return ctxhttp.Get(ctx, client, url)
}
