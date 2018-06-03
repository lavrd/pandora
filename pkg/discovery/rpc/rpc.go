package rpc

import (
	"context"
	"net"

	"github.com/spacelavr/pandora/pkg/conf"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"github.com/spacelavr/pandora/pkg/utils/network"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type rpc struct {
	master     string
	membership string
	broker     *pb.BrokerOpts
}

func New() *rpc {
	return &rpc{
		broker: &pb.BrokerOpts{
			Endpoint: conf.Viper.Discovery.Broker.Endpoint,
			User:     conf.Viper.Discovery.Broker.User,
			Password: conf.Viper.Discovery.Broker.Password,
		},
	}
}

func (rpc *rpc) InitMaster(ctx context.Context, in *pb.Endpoint) (*pb.BrokerOpts, error) {
	rpc.master = in.Endpoint
	return rpc.broker, nil
}

func (rpc *rpc) InitMembership(ctx context.Context, in *pb.Endpoint) (*pb.InitNetworkOpts, error) {
	rpc.membership = in.Endpoint

	if rpc.master == "" {
		return &pb.InitNetworkOpts{}, status.Error(codes.Unavailable, codes.Unavailable.String())
	}

	return &pb.InitNetworkOpts{
		Broker: rpc.broker,
		Master: rpc.master,
	}, nil
}

func (rpc *rpc) InitNode(ctx context.Context, in *pb.Empty) (*pb.InitNetworkOpts, error) {
	if rpc.master == "" || rpc.membership == "" {
		return &pb.InitNetworkOpts{}, status.Error(codes.Unavailable, codes.Unavailable.String())
	}

	return &pb.InitNetworkOpts{
		Broker:     rpc.broker,
		Master:     rpc.master,
		Membership: rpc.membership,
	}, nil
}

func (rpc *rpc) Listen() error {
	creds, err := credentials.NewServerTLSFromFile(conf.Viper.TLS.Cert, conf.Viper.TLS.Key)
	if err != nil {
		log.Error(err)
		return err
	}

	s := grpc.NewServer(grpc.Creds(creds))
	defer s.GracefulStop()

	pb.RegisterDiscoveryServer(s, rpc)

	listen, err := net.Listen(network.TCP, network.PortWithSemicolon(conf.Viper.Discovery.Endpoint))
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
