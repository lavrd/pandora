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
	return &pb.Net{
		Broker: &pb.Broker{
			Endpoint: config.Viper.Broker.Endpoint,
			User:     config.Viper.Broker.User,
			Password: config.Viper.Broker.Password,
		},
		Membership: &pb.Membership{
			Endpoint: config.Viper.Membership.Endpoint,
		},
		Master: &pb.Master{
			Endpoint: config.Viper.Master.Endpoint,
		},
	}, nil
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
