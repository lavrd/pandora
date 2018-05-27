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

type rpc struct{}

func New() (*rpc, *pb.BrokerOpts, error) {
	creds, err := credentials.NewClientTLSFromFile(config.Viper.TLS.Cert, "")
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	cc, err := grpc.Dial(config.Viper.Discovery.Endpoint, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}
	defer cc.Close()

	c := pb.NewDiscoveryClient(cc)

	opts, err := c.InitMaster(context.Background(), &pb.Endpoint{Endpoint: config.Viper.Master.Endpoint})
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	return &rpc{}, opts, nil
}

func (_ *rpc) CommitCert(ctx context.Context, in *pb.Cert) (*pb.Empty, error) {
	var (
		evt = env.GetEvents()
		bc  = env.GetBlockchain()
	)

	evt.PCertBlock(bc.PrepareCertBlock(in))
	evt.PubCert(in)

	return &pb.Empty{}, nil
}

func (_ *rpc) InitNode(ctx context.Context, in *pb.PublicKey) (*pb.MasterChain, error) {
	bc := env.GetBlockchain()
	bc.CommitMasterBlock(bc.PrepareMasterBlock(in))

	return bc.GetMasterChain(), nil
}

func (_ *rpc) Listen() error {
	creds, err := credentials.NewServerTLSFromFile(config.Viper.TLS.Cert, config.Viper.TLS.Key)
	if err != nil {
		log.Error(err)
		return err
	}

	s := grpc.NewServer(grpc.Creds(creds))
	defer s.GracefulStop()

	pb.RegisterMasterServer(s, &rpc{})

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
