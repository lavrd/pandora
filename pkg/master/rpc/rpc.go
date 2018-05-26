package rpc

import (
	"context"
	"net"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/master/env"
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

func (_ *gRPC) ConfirmCert(ctx context.Context, in *pb.Cert) (*pb.Empty, error) {
	var (
		evt = env.GetEvents()
		bc  = env.GetBlockchain()
	)

	evt.PCertBlock(bc.PrepareCertBlock(in))
	evt.PubCert(in)

	return &pb.Empty{}, nil
}

func (_ *gRPC) ConfirmNode(ctx context.Context, in *pb.PublicKey) (*pb.MasterChain, error) {
	var (
		bc = env.GetBlockchain()
	)

	bc.AddMasterBlock(in)

	return bc.MC(), nil
}

func (_ *gRPC) Listen() error {
	creds, err := credentials.NewServerTLSFromFile(config.Viper.TLS.Cert, config.Viper.TLS.Key)
	if err != nil {
		log.Error(err)
		return err
	}

	s := grpc.NewServer(grpc.Creds(creds))
	defer s.GracefulStop()

	pb.RegisterMasterServer(s, &gRPC{})

	listen, err := net.Listen(network.TCP, network.PortWithSemicolon(config.Viper.Master.Endpoint))
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
