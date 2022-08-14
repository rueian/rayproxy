package rayproxy

import (
	"context"
	"io"

	"github.com/rueian/rayproxy/gen/src/ray/protobuf"
	"google.golang.org/grpc/metadata"
)

var _ protobuf.RayletDriverServer = (*RayletDriverProxy)(nil)

func NewRayletDriverProxy(client protobuf.RayletDriverClient) *RayletDriverProxy {
	return &RayletDriverProxy{client: client}
}

type RayletDriverProxy struct {
	client protobuf.RayletDriverClient
}

func (r *RayletDriverProxy) Init(ctx context.Context, request *protobuf.InitRequest) (*protobuf.InitResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return r.client.Init(ctx, request)
}

func (r *RayletDriverProxy) PrepRuntimeEnv(ctx context.Context, request *protobuf.PrepRuntimeEnvRequest) (*protobuf.PrepRuntimeEnvResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return r.client.PrepRuntimeEnv(ctx, request)
}

func (r *RayletDriverProxy) GetObject(request *protobuf.GetRequest, server protobuf.RayletDriver_GetObjectServer) error {
	ctx := server.Context()
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	client, err := r.client.GetObject(ctx, request)
	if err != nil {
		return err
	}
	for {
		resp, err := client.Recv()
		if err == nil {
			err = server.Send(resp)
		}
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

func (r *RayletDriverProxy) PutObject(ctx context.Context, request *protobuf.PutRequest) (*protobuf.PutResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return r.client.PutObject(ctx, request)
}

func (r *RayletDriverProxy) WaitObject(ctx context.Context, request *protobuf.WaitRequest) (*protobuf.WaitResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return r.client.WaitObject(ctx, request)
}

func (r *RayletDriverProxy) Schedule(ctx context.Context, task *protobuf.ClientTask) (*protobuf.ClientTaskTicket, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return r.client.Schedule(ctx, task)
}

func (r *RayletDriverProxy) Terminate(ctx context.Context, request *protobuf.TerminateRequest) (*protobuf.TerminateResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return r.client.Terminate(ctx, request)
}

func (r *RayletDriverProxy) ClusterInfo(ctx context.Context, request *protobuf.ClusterInfoRequest) (*protobuf.ClusterInfoResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return r.client.ClusterInfo(ctx, request)
}

func (r *RayletDriverProxy) KVGet(ctx context.Context, request *protobuf.KVGetRequest) (*protobuf.KVGetResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return r.client.KVGet(ctx, request)
}

func (r *RayletDriverProxy) KVPut(ctx context.Context, request *protobuf.KVPutRequest) (*protobuf.KVPutResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return r.client.KVPut(ctx, request)
}

func (r *RayletDriverProxy) KVDel(ctx context.Context, request *protobuf.KVDelRequest) (*protobuf.KVDelResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return r.client.KVDel(ctx, request)
}

func (r *RayletDriverProxy) KVList(ctx context.Context, request *protobuf.KVListRequest) (*protobuf.KVListResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return r.client.KVList(ctx, request)
}

func (r *RayletDriverProxy) KVExists(ctx context.Context, request *protobuf.KVExistsRequest) (*protobuf.KVExistsResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return r.client.KVExists(ctx, request)
}

func (r *RayletDriverProxy) ListNamedActors(ctx context.Context, request *protobuf.ClientListNamedActorsRequest) (*protobuf.ClientListNamedActorsResponse, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return r.client.ListNamedActors(ctx, request)
}

var _ protobuf.RayletDataStreamerServer = (*RayletDataStreamerProxy)(nil)

func NewRayletDataStreamerProxy(client protobuf.RayletDataStreamerClient) *RayletDataStreamerProxy {
	return &RayletDataStreamerProxy{client: client}
}

type RayletDataStreamerProxy struct {
	client protobuf.RayletDataStreamerClient
}

func (r *RayletDataStreamerProxy) Datapath(server protobuf.RayletDataStreamer_DatapathServer) error {
	ctx := server.Context()
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	client, err := r.client.Datapath(ctx)
	if err != nil {
		return err
	}
	defer client.CloseSend()

	go func() {
		for {
			resp, err := client.Recv()
			if err == nil {
				err = server.Send(resp)
			}
			if err != nil {
				cancel()
				return
			}
		}
	}()

	for {
		req, err := server.Recv()
		if err == nil {
			err = client.Send(req)
		}
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

var _ protobuf.RayletLogStreamerServer = (*RayletLogStreamerProxy)(nil)

func NewRayletLogStreamerProxy(client protobuf.RayletLogStreamerClient) *RayletLogStreamerProxy {
	return &RayletLogStreamerProxy{client: client}
}

type RayletLogStreamerProxy struct {
	client protobuf.RayletLogStreamerClient
}

func (r *RayletLogStreamerProxy) Logstream(server protobuf.RayletLogStreamer_LogstreamServer) error {
	ctx := server.Context()
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	client, err := r.client.Logstream(ctx)
	if err != nil {
		return err
	}
	defer client.CloseSend()

	go func() {
		for {
			resp, err := client.Recv()
			if err == nil {
				err = server.Send(resp)
			}
			if err != nil {
				cancel()
				return
			}
		}
	}()

	for {
		req, err := server.Recv()
		if err == nil {
			err = client.Send(req)
		}
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}
