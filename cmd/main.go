package main

import (
	"fmt"

	"github.com/hatlonely/go-kit/flag"
	"github.com/hatlonely/go-kit/logger"
	"github.com/hatlonely/go-kit/refx"
	"github.com/hatlonely/go-kit/rpcx"

	"github.com/hatlonely/rpc-ops/internal/service"
)

var Version string

type Options struct {
	flag.Options

	Service     service.Options
	GrpcGateway rpcx.GrpcGatewayOptions

	Logger struct {
		Info logger.Options
		Grpc logger.Options
	}
}

func main() {
	var options Options
	refx.Must(flag.Struct(&options, refx.WithCamelName()))
	refx.Must(flag.Parse(flag.WithJsonVal()))
	if options.Help {
		fmt.Println(flag.Usage())
		return
	}
	if options.Version {
		fmt.Println(Version)
		return
	}
}
