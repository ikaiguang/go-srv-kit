package setuppkg

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	pkgerrors "github.com/pkg/errors"
	stdlog "log"
	"time"

	idutil "github.com/ikaiguang/go-srv-kit/kit/id"
	apputil "github.com/ikaiguang/go-srv-kit/kratos/app"
	clientutil "github.com/ikaiguang/go-srv-kit/kratos/client"
	apiv1 "github.com/ikaiguang/go-srv-kit/third_party/go-snowflake-node-id/api/node-id/v1/resources"
	servicev1 "github.com/ikaiguang/go-srv-kit/third_party/go-snowflake-node-id/api/node-id/v1/services"
)

// loadingSnowflakeWorker 加载雪花算法
func (s *engines) loadingSnowflakeWorker() error {
	workerConfig := s.SnowflakeWorkerConfig()
	if workerConfig == nil {
		stdlog.Println("|*** 加载：雪花算法：未初始化")
		return pkgerrors.WithStack(ErrUninitialized)
	}
	stdlog.Printf("|*** 加载：雪花算法")

	// http 选项
	var httpOptions = []http.ClientOption{
		http.WithMiddleware(
			recovery.Recovery(),
		),
		http.WithResponseDecoder(apputil.ResponseDecoder),
		http.WithEndpoint(workerConfig.Endpoint),
	}
	if workerConfig.WithDiscovery {
		consulClient, err := s.GetConsulClient()
		if err != nil {
			return err
		}
		r := consul.New(consulClient)
		httpOptions = append(httpOptions, http.WithDiscovery(r))
	}

	// http 链接
	httpConn, err := clientutil.NewHTTPClient(context.Background(), httpOptions...)
	if err != nil {
		return pkgerrors.WithStack(err)
	}
	httpClient := servicev1.NewSrvWorkerHTTPClient(httpConn)

	// 雪花算法ID
	appConfig := s.AppConfig()
	workerReq := &apiv1.GetNodeIdReq{
		InstanceId:   apputil.ID(appConfig),
		InstanceName: appConfig.Name,
		Endpoints:    appConfig.Endpoints,
		Metadata:     appConfig.Metadata,
	}
	workerResp, err := httpClient.GetNodeId(context.Background(), workerReq)
	if err != nil {
		return pkgerrors.WithStack(err)
	}

	// 雪花算法
	stdlog.Printf("|*** 加载：雪花算法：nodeId = %d", workerResp.NodeId)
	snowflakeNode, err := snowflake.NewNode(workerResp.NodeId)
	if err != nil {
		return pkgerrors.WithStack(err)
	}
	idutil.SetNode(snowflakeNode)

	// 续期
	extendReq := &apiv1.ExtendNodeIdReq{
		Id:         workerResp.Id,
		InstanceId: workerReq.InstanceId,
		NodeId:     workerResp.NodeId,
	}
	s.snowflakeStopChannel = make(chan int)
	go func() {
		ticker := time.NewTicker(time.Second * 3)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				//debugutil.Println("雪花算法：续期：", extendReq.String())
				_, _ = httpClient.ExtendNodeId(context.Background(), extendReq)
			case <-s.snowflakeStopChannel:
				return
			}
		}
	}()

	return err
}
