package testregistry

import (
	"context"
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/hashicorp/consul/api"
	stdlog "log"
	"testing"

	pingv1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/resources"
	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"
	"github.com/ikaiguang/go-srv-kit/example/internal/setup"
	apputil "github.com/ikaiguang/go-srv-kit/kratos/app"
	logutil "github.com/ikaiguang/go-srv-kit/log"
)

// go test -v -count=1 ./example/integration-test/registry -test.run=Test_RegistryDiscovery
func Test_RegistryDiscovery(t *testing.T) {
	consulCli, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	r := consul.New(consulCli)

	// 引擎模块
	engineHandler, err := setup.GetEngine()
	if err != nil {
		stdlog.Fatalf("%+v\n", err)
		return
	}
	appID := apputil.ID(engineHandler.AppConfig())
	endpoint := "discovery:///" + appID

	// new grpc client
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = conn.Close() }()

	gClient := pingservicev1.NewSrvPingClient(conn)

	// new http client
	hConn, err := http.NewClient(
		context.Background(),
		http.WithMiddleware(
			recovery.Recovery(),
		),
		http.WithEndpoint(endpoint),
		http.WithDiscovery(r),
		// 解析
		http.WithResponseDecoder(apputil.ResponseDecoder),
	)
	if err != nil {
		logutil.Fatal(err)
	}
	defer func() { _ = hConn.Close() }()
	hClient := pingservicev1.NewSrvPingHTTPClient(hConn)

	//for {
	//	time.Sleep(time.Second)
	//	callGRPC(gClient)
	//	callHTTP(hClient)
	//}
	callGRPC(gClient)
	callHTTP(hClient)
}

func callGRPC(client pingservicev1.SrvPingClient) {
	reply, err := client.Ping(context.Background(), &pingv1.PingReq{Message: "grpc"})
	if err != nil {
		log.Fatal(err)
	}
	logutil.Infof("[grpc] SayHello %+v\n", reply)
}

func callHTTP(client pingservicev1.SrvPingHTTPClient) {
	reply, err := client.Ping(context.Background(), &pingv1.PingReq{Message: "http"})
	if err != nil {
		log.Fatal(err)
	}
	logutil.Printf("[http] SayHello %s\n", reply.Message)
}
