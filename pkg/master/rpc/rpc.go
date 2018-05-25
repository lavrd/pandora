package rpc

import (
	"context"
	"net"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/master/runtime"
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
		r = runtime.Get()
	)

	block := r.Chain.PrepareCertBlock(in)

	r.Events.PCertBlock(block)
	r.Events.PubCert(in)

	return &pb.Empty{}, nil
}

func (_ *gRPC) ConfirmNode(ctx context.Context, in *pb.PublicKey) (*pb.MasterChain, error) {
	var (
		r = runtime.Get()
	)

	r.Chain.AddMasterBlock(in)

	return r.Chain.GetMC(), nil
}

func (_ *gRPC) Listen() error {
	listen, err := net.Listen(network.TCP, network.PortWithSemicolon(config.Viper.Master.Endpoint))
	if err != nil {
		log.Error(err)
		return err
	}
	defer listen.Close()

	creds, err := credentials.NewServerTLSFromFile(config.Viper.TLS.Cert, config.Viper.TLS.Key)
	if err != nil {
		log.Error(err)
		return err
	}

	s := grpc.NewServer(grpc.Creds(creds))
	defer s.GracefulStop()

	pb.RegisterMasterServer(s, &gRPC{})

	if err := s.Serve(listen); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
