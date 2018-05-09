package rpc

import (
	"context"
	"net"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/master/env"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"github.com/spacelavr/pandora/pkg/utils/network"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Candidate(ctx context.Context, in *pb.Block) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *server) Node(ctx context.Context, in *pb.PublicKey) (*pb.Empty, error) {
	var (
		e = env.GetEvents()
	)

	e.PMasterBlock(&types.MasterBlock{
		Block: &types.Block{
			Index: 999,
		},
	})

	return &pb.Empty{}, nil
}

func Listen() error {
	listen, err := net.Listen(network.TCP, network.PortWithSemicolon(config.Viper.Master.Endpoint))
	if err != nil {
		log.Error(err)
		return err
	}
	defer listen.Close()

	s := grpc.NewServer()
	defer s.GracefulStop()

	pb.RegisterMasterServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
