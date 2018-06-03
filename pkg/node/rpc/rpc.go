package rpc

import (
	"context"
	"time"

	"github.com/spacelavr/pandora/pkg/conf"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type RPC struct {
	master       pb.MasterClient
	membership   pb.MembershipClient
	discovery    pb.DiscoveryClient
	discoveryCC  *grpc.ClientConn
	membershipCC *grpc.ClientConn
	masterCC     *grpc.ClientConn
}

func New() (*RPC, error) {
	creds, err := credentials.NewClientTLSFromFile(conf.Viper.TLS.Cert, "")
	if err != nil {
		log.Error(err)
		return nil, err
	}

	discoveryCC, err := grpc.Dial(conf.Viper.Discovery.Endpoint, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	discoveryC := pb.NewDiscoveryClient(discoveryCC)

	var (
		ino   = &pb.InitNetworkOpts{}
		tick  = time.NewTicker(time.Millisecond * 500).C
		timer = time.NewTimer(time.Second * 3).C
	)
loop:
	for {
		select {
		case <-tick:
			if ino, err = discoveryC.InitNode(context.Background(), &pb.Empty{}); err != nil {
				continue
			}
			break loop
		case <-timer:
			log.Error(err)
			return nil, err
		}
	}

	membershipCC, err := grpc.Dial(ino.Membership, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	membershipC := pb.NewMembershipClient(membershipCC)

	masterCC, err := grpc.Dial(ino.Master, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	masterC := pb.NewMasterClient(masterCC)

	return &RPC{
		master:       masterC,
		membership:   membershipC,
		discovery:    discoveryC,
		masterCC:     masterCC,
		discoveryCC:  discoveryCC,
		membershipCC: membershipCC,
	}, nil
}

func (rpc *RPC) Close() {
	if err := rpc.membershipCC.Close(); err != nil {
		log.Error(err)
	}
	if err := rpc.discoveryCC.Close(); err != nil {
		log.Error(err)
	}
	if err := rpc.masterCC.Close(); err != nil {
		log.Error(err)
	}
}

func (rpc *RPC) ProposeMember(candidate *pb.MemberMeta) (*pb.PublicKey, error) {
	key, err := rpc.membership.ProposeMember(context.Background(), candidate);
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return key, nil
}

func (rpc *RPC) SignCert(cert *pb.Cert) error {
	if _, err := rpc.membership.SignCert(context.Background(), cert); err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				return errors.ErrNotFound
			}
		}

		log.Error(err)
		return err
	}
	return nil
}

func (rpc *RPC) FetchMember(key *pb.PublicKey) (*pb.Member, error) {
	r, err := rpc.membership.FetchMember(context.Background(), key)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				return nil, errors.ErrNotFound
			}
		}

		log.Error(err)
		return nil, err
	}

	return r, nil
}

func (rpc *RPC) InitNode(key *pb.PublicKey) (*pb.MasterChain, *pb.BrokerOpts, error) {
	ino, err := rpc.discovery.InitNode(context.Background(), &pb.Empty{})
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	mc, err := rpc.master.InitNode(context.Background(), key)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	return mc, ino.Broker, nil
}
