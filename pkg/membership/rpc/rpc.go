package rpc

import (
	"context"
	"net"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/membership/pb"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"github.com/spacelavr/pandora/pkg/utils/network"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Register(ctx context.Context, in *pb.Candidate) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *server) Fetch(ctx context.Context, in *pb.Public) (*pb.Account, error) {
	return &pb.Account{}, nil
}

func Listen() error {
	listen, err := net.Listen(network.TCP, config.Viper.Membership.Endpoint)
	if err != nil {
		log.Error(err)
		return err
	}
	defer listen.Close()

	s := grpc.NewServer()
	defer s.GracefulStop()

	pb.RegisterMembershipServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
