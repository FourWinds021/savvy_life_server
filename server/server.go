package server

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	assetFS "github.com/elazarl/go-bindata-assetfs"
	"github.com/facebookgo/grace/gracehttp"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"path"
	. "savvy_life/common"
	"savvy_life/middlewares"
	. "savvy_life/proto/api/v1"
	"savvy_life/server/api/user"
	"savvy_life/ui/data/swagger"
	"strings"
	"time"
)

func Run() {
	flag.Parse()

	svr := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				middlewares.Auth(),
				middlewares.Trace(),
				middlewares.Recovery()),
		))
	RegisterUserServer(svr, &grpcUserServer{})
	listen, err := net.Listen("tcp", GrpcAddr)
	if err != nil {
		Logger.Error("GRPC failed to listen. error:", err)
	}
	Logger.Info("Serving gRPC on 0.0.0.0" + GrpcAddr)
	fmt.Println("Serving gRPC on 0.0.0.0: ", GrpcAddr)
	go func() {
		if err := svr.Serve(listen); err != nil {
			Logger.Error("GRPC failed to serve. error:", err)
		}
	}()

	connect, err := grpc.Dial(
		GrpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	gatewayMux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(middlewares.CustomMatcher))

	err = RegisterUserHandler(context.Background(), gatewayMux, connect)
	if err != nil {
		Logger.Error("RegisterUserHandler failed. error:", err)
	}
	if SwaggerEnable {
		swaggerMux := http.NewServeMux()
		swaggerMux.Handle("/", gatewayMux)
		swaggerMux.HandleFunc("/swagger/", swaggerFile)
		swaggerUI(swaggerMux)

		Logger.Info("Serving Swagger on http://0.0.0.0" + SwaggerAddr)
		go func() {
			err = http.ListenAndServe(SwaggerAddr, swaggerMux)
			if err != nil {
				Logger.Error("Swagger failed to serve. error:", err)
			}
		}()
	}
	fmt.Println("Serving gRPC-Gateway on http://0.0.0.0 ", GatewayAddr)
	Logger.Info("Serving gRPC-Gateway on http://0.0.0.0" + GatewayAddr)
	cert, err := tls.LoadX509KeyPair(
		"../config/server.pem",
		"../config/server.key")
	if err != nil {
		Logger.Error("tls.LoadX509KeyPair failed. error:", err)
	}
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}
	err = gracehttp.Serve(
		&http.Server{
			Addr: *flag.String(
				"httpserver",
				GatewayAddr,
				"address to server."),
			ReadTimeout: time.Duration(ReadTimeout) * time.Second,
			Handler:     gatewayMux,
			TLSConfig:   tlsConfig})
	if err != nil {
		Logger.Error("GateWay failed to serve. error:", err)
	}
}

func swaggerFile(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		Logger.Error("Not Found swagger.json. path:", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	name := path.Join("../openapi/api/v1", p)
	Logger.Info("Serving swagger-file:", name)
	http.ServeFile(w, r, name)
}

func swaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(&assetFS.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

/*
 * User Server
 */
type grpcUserServer struct {
	UnimplementedUserServer
}

func (*grpcUserServer) GetUser(_ context.Context, req *GetUserReq) (
	*GetUserResp, error) {

	return user.GetUser(req)
}
