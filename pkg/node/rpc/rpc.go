package rpc

import (
	"context"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type gRPC struct {
	master     pb.MasterClient
	membership pb.MembershipClient
}

func New() (*gRPC, *pb.BrokerOpts, error) {
	creds, err := credentials.NewClientTLSFromFile(config.Viper.TLS.Cert, "")
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	discoveryCC, err := grpc.Dial(config.Viper.Discovery.Endpoint, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}
	defer discoveryCC.Close()

	discoveryC := pb.NewDiscoveryClient(discoveryCC)

	ino, err := discoveryC.InitNode(context.Background(), &pb.Empty{})
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	membershipCC, err := grpc.Dial(ino.Membership, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}
	defer membershipCC.Close()

	membershipC := pb.NewMembershipClient(membershipCC)

	masterCC, err := grpc.Dial(ino.Master, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}
	defer masterCC.Close()

	masterC := pb.NewMasterClient(masterCC)

	return &gRPC{
		master:     masterC,
		membership: membershipC,
	}, ino.Broker, nil
}

func (g *gRPC) Register(candidate *pb.Candidate) error {
	if _, err := g.membership.ConfirmMember(context.Background(), candidate); err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return err
		}
		if st.Code() == codes.AlreadyExists {
			return errors.AlreadyExists
		}

		log.Error(err)
		return err
	}
	return nil
}

func (g *gRPC) Issue(cert *pb.Cert) error {
	if _, err := g.membership.SignCert(context.Background(), cert); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (g *gRPC) NodeReg(candidate *pb.Candidate) (*pb.PublicKey, error) {
	r, err := g.membership.ConfirmMember(context.Background(), candidate)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, err
		}
		if st.Code() == codes.AlreadyExists {
			log.Debug(r)
			return r, errors.AlreadyExists
		}

		log.Error(err)
		return nil, err
	}
	return r, nil
}

func (g *gRPC) FetchAccount(key *pb.PublicKey) (*pb.Member, error) {
	r, err := g.membership.FetchMember(context.Background(), key)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, err
		}
		if st.Code() == codes.NotFound {
			return nil, errors.NotFound
		}

		log.Error(err)
		return nil, err
	}
	return r, nil
}

func (g *gRPC) Node(key *pb.PublicKey) (*pb.MasterChain, error) {
	mc, err := g.master.ConfirmNode(context.Background(), key)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return mc, nil
}
