package grpc

import (
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/niubir/logger"
	"github.com/niubir/user-service/config"
	"github.com/niubir/user-service/grpc/server"
	"github.com/niubir/user-service/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func Init(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	if !config.Config.GRPC.Enable {
		return
	}

	address := fmt.Sprintf(":%d", config.Config.GRPC.Port)

	setGRPCLogger()
	s := grpc.NewServer()
	server.RegisterUserServiceServer(s, &UserService{})
	listen, err := net.Listen("tcp", address)
	if err != nil {
		logs.System.Errorln("start grpc failed:", err)
		return
	}

	logs.System.Infof("start grpc, address(%s)\n", address)
	s.Serve(listen)
}

func setGRPCLogger() {
	grpcLogger := io.Discard
	if config.Config.GRPC.Debug {
		if w, err := logs.NewLogger("user_service_grpc"); err != nil {
			logs.System.Errorln("create user service grpc logger failed:", err)
		} else {
			grpcLogger = &logger.GRPCLogger{Logger: w}
		}
	}
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(grpcLogger, grpcLogger, grpcLogger))
}
