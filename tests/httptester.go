package tests

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/environment-toolkit/grid/handler"
	"github.com/spf13/viper"

	"github.com/go-apis/utils/xservice"
)

type HttpTester interface {
	Url() string
	Do(req *http.Request) (*http.Response, error)
	Close()
}

type httpTester struct {
	srv *httptest.Server
}

func (h *httpTester) Do(req *http.Request) (*http.Response, error) {
	return h.srv.Client().Do(req)
}
func (h *httpTester) Url() string {
	return h.srv.URL
}
func (h *httpTester) Close() {
	if h.srv != nil {
		h.srv.Close()
	}
}

func NewHttpTester() (HttpTester, error) {
	ctx := context.Background()

	v := viper.New()
	cfg, err := xservice.NewService(ctx, v)
	if err != nil {
		return nil, err
	}
	h, err := handler.NewHandler(ctx, cfg, nil)
	if err != nil {
		return nil, err
	}
	srv := httptest.NewServer(h)

	return &httpTester{
		srv: srv,
	}, nil
}
