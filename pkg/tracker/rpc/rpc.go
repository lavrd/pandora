package rpc

import (
	"context"
	"fmt"
	"net"

	"github.com/spacelavr/pandora/pkg/config"
	pb "github.com/spacelavr/pandora/pkg/proto"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"google.golang.org/grpc"
)

type server struct {
	IPs map[string]*pb.Empty
}

func (s *server) AddIP(ip string) {
	s.IPs[ip] = &pb.Empty{}
}

func (s *server) GetValidators(ctx context.Context, in *pb.GVRequest) (*pb.GVResponse, error) {
	return &pb.GVResponse{Ips: s.IPs}, nil
}

func (s *server) NewValidator(ctx context.Context, in *pb.NVRequest) (*pb.NVResponse, error) {
	s.AddIP(in.Ip)
	return &pb.NVResponse{}, nil
}

func Listen() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Viper.Tracker.Port))
	if err != nil {
		log.Error(err)
		return err
	}
	defer listen.Close()

	s := grpc.NewServer()
	defer s.GracefulStop()

	pb.RegisterTrackerServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
