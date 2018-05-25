package rpc

import (
	"context"
	"net"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"github.com/spacelavr/pandora/pkg/utils/network"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type gRPC struct {
}

func New() *gRPC {
	return &gRPC{}
}

func (_ *gRPC) ConfirmMember(ctx context.Context, in *pb.Candidate) (*pb.PublicKey, error) {
	dist := &distribution.Distribution{
		Storage: env.GetStorage(),
		Runtime: env.GetRuntime(),
	}

	key, err := dist.AcceptCandidate(in)
	if err != nil {
		if err == errors.AlreadyExists {
			return key, status.Error(codes.AlreadyExists, codes.AlreadyExists.String())
		}
		return &pb.PublicKey{}, err
	}

	return key, nil
}

func (_ *gRPC) Node(ctx context.Context, in *pb.Candidate) (*pb.PublicKey, error) {
	dist := &distribution.Distribution{
		Storage: env.GetStorage(),
		Runtime: env.GetRuntime(),
	}

	key, err := dist.AcceptCandidate(in)
	if err != nil && err != errors.AlreadyExists {
		return &pb.PublicKey{}, err
	}

	return key, nil
}

func (_ *gRPC) SignCert(ctx context.Context, in *pb.Cert) (*pb.Empty, error) {
	cert, err := distribution.New().Issue(in)
	if err != nil {
		return &pb.Empty{}, err
	}

	if err := Issue(cert); err != nil {
		return &pb.Empty{}, err
	}

	return &pb.Empty{}, nil
}

func (_ *gRPC) Issue(cert *pb.Cert) error {
	creds, err := credentials.NewClientTLSFromFile("./contrib/cert.pem", "")
	if err != nil {
		log.Error(err)
		return err
	}

	cc, err := grpc.Dial(config.Viper.Master.Endpoint, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Error(err)
	}
	defer cc.Close()

	c := pb.NewMasterClient(cc)

	_, err = c.ConfirmCert(context.Background(), cert)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (_ *gRPC) FetchMember(ctx context.Context, in *pb.PublicKey) (*pb.Member, error) {
	dist := &distribution.Distribution{
		Storage: env.GetStorage(),
		Runtime: env.GetRuntime(),
	}

	acc, err := dist.AccountFetch(in)
	if err != nil {
		if err == errors.NotFound {
			// todo may be nil? and in all place
			return &pb.Member{}, status.Error(codes.NotFound, codes.NotFound.String())
		}
		return &pb.Member{}, err
	}
	return acc, nil
}

func (_ *gRPC) Listen() error {
	listen, err := net.Listen(network.TCP, network.PortWithSemicolon(config.Viper.Membership.Endpoint))
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

	pb.RegisterMembershipServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Error(err)
		return err
	}

	return nil
}
