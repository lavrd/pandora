package rpc

import (
	"context"
	"net"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"github.com/spacelavr/pandora/pkg/utils/network"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type gRPC struct{}

func New() *gRPC {
	return &gRPC{}
}

func (_ *gRPC) Network(ctx context.Context, in *pb.Empty) (*pb.NetworkOpts, error) {
	return &pb.NetworkOpts{
		Broker: &pb.BrokerOpts{
			Endpoint: config.Viper.Discovery.Broker.Endpoint,
			User:     config.Viper.Discovery.Broker.User,
			Password: config.Viper.Discovery.Broker.Password,
		},
		Membership: &pb.MembershipOpts{
			Endpoint: config.Viper.Membership.Endpoint,
		},
		Master: &pb.MasterOpts{
			Endpoint: config.Viper.Master.Endpoint,
		},
	}, nil
}

func (_ *gRPC) Listen() error {
	creds, err := credentials.NewServerTLSFromFile(config.Viper.TLS.Cert, config.Viper.TLS.Key)
	if err != nil {
		log.Error(err)
		return err
	}

	s := grpc.NewServer(grpc.Creds(creds))
	defer s.GracefulStop()

	pb.RegisterDiscoveryServer(s, &gRPC{})

	listen, err := net.Listen(network.TCP, network.PortWithSemicolon(config.Viper.Discovery.Endpoint))
	if err != nil {
		log.Error(err)
		return err
	}
	defer listen.Close()

	if err := s.Serve(listen); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
