package rpc

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"pandora/pkg/conf"
	"pandora/pkg/master/env"
	"pandora/pkg/pb"
	"pandora/pkg/utils/errors"
	"pandora/pkg/utils/network"
)

type rpc struct{}

// New return new master rpc
func New() *rpc {
	return &rpc{}
}

// ProposeCert propose cert
func (*rpc) ProposeCert(ctx context.Context, in *pb.Cert) (*pb.Empty, error) {
	evt := env.GetEvents()
	bc := env.GetBlockchain()

	evt.PubCertBlock(bc.PrepareCertBlock(in))
	evt.PubCert(in)

	return &pb.Empty{}, nil
}

// InitNode init service node with blockchain
func (*rpc) InitNode(ctx context.Context, in *pb.PublicKey) (*pb.MasterChain, error) {
	evt := env.GetEvents()
	bc := env.GetBlockchain()

	b := bc.PrepareMasterBlock(in)
	bc.CommitMasterBlock(b)

	evt.PubMasterBlock(b)

	return bc.GetMasterChain(), nil
}

// Listen listen for rpc requests
func (rpc *rpc) Listen() error {
	creds, err := credentials.NewServerTLSFromFile(conf.Conf.TLS.Cert, conf.Conf.TLS.Key)
	if err != nil {
		return errors.WithStack(err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	defer s.GracefulStop()

	pb.RegisterMasterServer(s, rpc)

	listen, err := net.Listen(network.TCP, network.PortWithSemicolon(conf.Conf.Master.Endpoint))
	if err != nil {
		return errors.WithStack(err)
	}
	defer func() {
		_ = listen.Close()
	}()

	if err := s.Serve(listen); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// InitMaster init master rpc by discovery service
func (*rpc) InitMaster() (*pb.BrokerOpts, error) {
	creds, err := credentials.NewClientTLSFromFile(conf.Conf.TLS.Cert, "")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	cc, err := grpc.Dial(conf.Conf.Discovery.Endpoint, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer func() {
		_ = cc.Close()
	}()

	c := pb.NewDiscoveryClient(cc)

	opts := &pb.BrokerOpts{}
	tick := time.NewTicker(time.Millisecond * 500).C
	timer := time.NewTimer(time.Second * 3).C

loop:
	for {
		select {
		case <-tick:
			opts, err = c.InitMaster(context.Background(), &pb.Endpoint{Endpoint: conf.Conf.Master.Endpoint})
			if err != nil {
				continue
			}
			break loop
		case <-timer:
			return nil, errors.WithStack(err)
		}
	}

	return opts, nil
}
