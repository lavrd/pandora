package rpc

import (
	"context"
	"net"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/discovery/pb"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"github.com/spacelavr/pandora/pkg/utils/network"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Network(ctx context.Context, in *pb.Empty) (*pb.Net, error) {
	return &pb.Net{}, nil
}

func Listen() error {
	listen, err := net.Listen(network.TCP, config.Viper.Discovery.Endpoint)
	if err != nil {
		log.Error(err)
		return err
	}
	defer listen.Close()

	s := grpc.NewServer()
	defer s.GracefulStop()

	pb.RegisterDiscoveryServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
