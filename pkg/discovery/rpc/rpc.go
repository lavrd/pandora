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

type gRPC struct {
	master     string
	membership string
	broker     *pb.BrokerOpts
}

func New() *gRPC {
	return &gRPC{
		broker: &pb.BrokerOpts{
			Endpoint: config.Viper.Discovery.Broker.Endpoint,
			User:     config.Viper.Discovery.Broker.User,
			Password: config.Viper.Discovery.Broker.Password,
		},
	}
}

func (g *gRPC) InitMaster(ctx context.Context, in *pb.Endpoint) (*pb.Empty, error) {
	g.master = in.Endpoint
	return &pb.Empty{}, nil
}

func (g *gRPC) InitMembership(ctx context.Context, in *pb.Endpoint) (*pb.Empty, error) {
	g.membership = in.Endpoint
	return &pb.Empty{}, nil
}

func (g *gRPC) InitNode(ctx context.Context, in *pb.Endpoint) (*pb.InitNetworkOpts, error) {
	// todo if master or membership not started, send message
	return &pb.InitNetworkOpts{
		Broker:     g.broker,
		Master:     g.master,
		Membership: g.membership,
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
