package rpc

import (
	"context"
	"net"

	"github.com/spacelavr/pandora/pkg/conf"
	"github.com/spacelavr/pandora/pkg/master/env"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"github.com/spacelavr/pandora/pkg/utils/network"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type rpc struct{}

func New() *rpc {
	return &rpc{}
}

func (_ *rpc) ProposeCert(ctx context.Context, in *pb.Cert) (*pb.Empty, error) {
	var (
		evt = env.GetEvents()
		bc  = env.GetBlockchain()
	)

	evt.PubCertBlock(bc.PrepareCBlock(in))
	evt.PubCert(in)

	return &pb.Empty{}, nil
}

func (_ *rpc) InitNode(ctx context.Context, in *pb.PublicKey) (*pb.MasterChain, error) {
	var (
		evt = env.GetEvents()
		bc  = env.GetBlockchain()
	)

	b := bc.PrepareMBlock(in)
	bc.CommitMBlock(b)

	evt.PubMasterBlock(b)

	return bc.MasterChain(), nil
}

func (rpc *rpc) Listen() error {
	creds, err := credentials.NewServerTLSFromFile(conf.Viper.TLS.Cert, conf.Viper.TLS.Key)
	if err != nil {
		log.Error(err)
		return err
	}

	s := grpc.NewServer(grpc.Creds(creds))
	defer s.GracefulStop()

	pb.RegisterMasterServer(s, rpc)

	listen, err := net.Listen(network.TCP, network.PortWithSemicolon(conf.Viper.Master.Endpoint))
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

func (_ *rpc) InitMaster() (*pb.BrokerOpts, error) {
	creds, err := credentials.NewClientTLSFromFile(conf.Viper.TLS.Cert, "")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	cc, err := grpc.Dial(conf.Viper.Discovery.Endpoint, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer cc.Close()

	c := pb.NewDiscoveryClient(cc)

	opts, err := c.InitMaster(context.Background(), &pb.Endpoint{Endpoint: conf.Viper.Master.Endpoint})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return opts, nil
}
