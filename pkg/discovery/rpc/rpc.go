package rpc

import (
	"context"
	"net"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"github.com/spacelavr/pandora/pkg/utils/network"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Network(ctx context.Context, in *pb.Empty) (*pb.NetOpts, error) {
	return &pb.NetOpts{
		Broker: &pb.BrokerOpts{
			Endpoint: config.Viper.Broker.Endpoint,
			User:     config.Viper.Broker.User,
			Password: config.Viper.Broker.Password,
		},
		Membership: &pb.MembershipOpts{
			Endpoint: config.Viper.Membership.Endpoint,
		},
		Master: &pb.MasterOpts{
			Endpoint: config.Viper.Master.Endpoint,
		},
	}, nil
	return &pb.NetOpts{}, nil
}

func Listen() error {
	listen, err := net.Listen(network.TCP, network.PortWithSemicolon(config.Viper.Discovery.Endpoint))
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
