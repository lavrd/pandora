package rpc

import (
	"context"
	"net"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/master/distribution"
	"github.com/spacelavr/pandora/pkg/master/env"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"github.com/spacelavr/pandora/pkg/utils/network"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct{}

func (s *server) ConfirmCert(ctx context.Context, in *pb.Cert) (*pb.Empty, error) {
	block := distribution.PrepareBlock(in)
	env.GetEvents().PCertBlock(block)
	env.GetEvents().PubCert(in)
	return &pb.Empty{}, nil
}

func (s *server) ConfirmNode(ctx context.Context, in *pb.PublicKey) (*pb.MasterChain, error) {
	distribution.AddMasterBlock(in)
	return distribution.GetMasterChain(), nil
}

func Listen() error {
	listen, err := net.Listen(network.TCP, network.PortWithSemicolon(config.Viper.Master.Endpoint))
	if err != nil {
		log.Error(err)
		return err
	}
	defer listen.Close()

	creds, err := credentials.NewServerTLSFromFile("./contrib/cert.pem", "./contrib/key.pem")
	if err != nil {
		log.Error(err)
		return err
	}

	s := grpc.NewServer(grpc.Creds(creds))
	defer s.GracefulStop()

	pb.RegisterMasterServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
