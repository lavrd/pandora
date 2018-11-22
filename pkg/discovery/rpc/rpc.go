package rpc

import (
	"context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"pandora/pkg/conf"
	"pandora/pkg/pb"
	"pandora/pkg/utils/errors"
	"pandora/pkg/utils/network"
)

type rpc struct {
	master     string
	membership string
	broker     *pb.BrokerOpts
}

// New returns new discovery rpc
func New() *rpc {
	return &rpc{
		broker: &pb.BrokerOpts{
			Endpoint: conf.Conf.NATS.Endpoint,
			User:     conf.Conf.NATS.User,
			Password: conf.Conf.NATS.Password,
		},
	}
}

// InitMaster init master service
func (rpc *rpc) InitMaster(ctx context.Context, in *pb.Endpoint) (*pb.BrokerOpts, error) {
	rpc.master = in.Endpoint
	return rpc.broker, nil
}

// InitMembership init membership service
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

// InitNode init node services
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

// Listen listen for rpc requests
func (rpc *rpc) Listen() error {
	creds, err := credentials.NewServerTLSFromFile(conf.Conf.TLS.Cert, conf.Conf.TLS.Key)
	if err != nil {
		return errors.WithStack(err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	defer s.GracefulStop()

	pb.RegisterDiscoveryServer(s, rpc)

	listen, err := net.Listen(network.TCP, network.PortWithSemicolon(conf.Conf.Discovery.Endpoint))
	if err != nil {
		return errors.WithStack(err)
	}
	defer listen.Close()

	if err := s.Serve(listen); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
