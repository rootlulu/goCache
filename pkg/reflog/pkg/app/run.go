package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"reflog.com/reflog/pkg/pb"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	go initReflogServer()
}

type server struct {
	pb.UnimplementedReflogServiceServer
}

func initReflogServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 29991))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterReflogServiceServer(s, &server{})
	if logrus.StandardLogger().Level == logrus.DebugLevel {
		reflection.Register(s)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Reflog(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	if in.GetReq() == "" {
		return nil, fmt.Errorf("empty request")
	}
	reqs := strings.Split(in.GetReq(), " ")
	logrus.Infof("reqs: %v", reqs)
	// funcName reqs[0]
	// args reqs[1:]
	var err error
	var resp string

	if len(reqs) == 1 {
		resp, err = funcMap[reqs[0]]()
	} else {
		resp, err = funcMap[reqs[0]](reqs[1:]...)
	}
	if err != nil {
		return &pb.Response{
			Resp: err.Error(),
			Code: 1,
		}, nil
	}
	return &pb.Response{
		Resp: resp,
		Code: 0,
	}, nil
}
