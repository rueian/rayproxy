package main

import (
	"flag"
	"log"
	"math"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/rueian/rayproxy"
	"github.com/rueian/rayproxy/gen/src/ray/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func main() {
	bind := flag.String("listen", "0.0.0.0:10001", "listen address, default: 0.0.0.0:10001")
	addr := flag.String("target", "", "the proxy target address, ex: my.ray.server:10001")
	flag.Parse()

	ln, err := net.Listen("tcp", *bind)
	if err != nil {
		log.Fatalf("fail to listen on %s: %v\n", *bind, err)
	}

	client, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("fail to proxy on %s: %v\n", *addr, err)
	}

	server := grpc.NewServer(
		grpc.MaxConcurrentStreams(math.MaxUint32),
		grpc.MaxRecvMsgSize(math.MaxUint32),
		grpc.MaxSendMsgSize(math.MaxUint32),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    1000 * 30,
			Timeout: 1000 * 600,
		}))

	protobuf.RegisterRayletDriverServer(server, rayproxy.NewRayletDriverProxy(protobuf.NewRayletDriverClient(client)))
	protobuf.RegisterRayletDataStreamerServer(server, rayproxy.NewRayletDataStreamerProxy(protobuf.NewRayletDataStreamerClient(client)))
	protobuf.RegisterRayletLogStreamerServer(server, rayproxy.NewRayletLogStreamerProxy(protobuf.NewRayletLogStreamerClient(client)))

	go func() {
		log.Printf("server listens on %s to proxy %s\n", *bind, *addr)
		if err := server.Serve(ln); err != nil {
			log.Fatalf("server exit with err: %v\n", err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	server.GracefulStop()
	log.Printf("server stopped\n")
}
