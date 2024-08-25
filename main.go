package main

import (
	"context"

	"github.com/environment-toolkit/grid/handler"
	"github.com/spf13/viper"

	"github.com/go-apis/utils/xgraceful"
	"github.com/go-apis/utils/xservice"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	v := viper.New()
	svc, err := xservice.NewService(ctx, v)
	if err != nil {
		panic(err)
	}

	handler, err := handler.NewHandler(ctx, svc, nil)
	if err != nil {
		panic(err)
	}

	xgraceful.Serve(ctx, svc.ServiceConfig, handler)
	cancel()

	<-ctx.Done()
}
