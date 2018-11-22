package rpc

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	"pandora/pkg/conf"
	"pandora/pkg/pb"
	"pandora/pkg/utils/errors"
	"pandora/pkg/utils/log"
)

// RPC
type RPC struct {
	master       pb.MasterClient
	membership   pb.MembershipClient
	discovery    pb.DiscoveryClient
	discoveryCC  *grpc.ClientConn
	membershipCC *grpc.ClientConn
	masterCC     *grpc.ClientConn
}

// New returns new rpc
func New() (*RPC, error) {
	creds, err := credentials.NewClientTLSFromFile(conf.Conf.TLS.Cert, "")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	discoveryCC, err := grpc.Dial(conf.Conf.Discovery.Endpoint, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, errors.WithStack(err)
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
			return nil, errors.WithStack(err)
		}
	}

	membershipCC, err := grpc.Dial(ino.Membership, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	membershipC := pb.NewMembershipClient(membershipCC)

	masterCC, err := grpc.Dial(ino.Master, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, errors.WithStack(err)
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

// Close close rpc connection with other rpc
func (rpc *RPC) Close() {
	if err := rpc.membershipCC.Close(); err != nil {
		log.Error(errors.WithStack(err))
	}
	if err := rpc.discoveryCC.Close(); err != nil {
		log.Error(errors.WithStack(err))
	}
	if err := rpc.masterCC.Close(); err != nil {
		log.Error(errors.WithStack(err))
	}
}

// ProposeMember propose member over rpc
func (rpc *RPC) ProposeMember(candidate *pb.MemberMeta) (*pb.PublicKey, error) {
	key, err := rpc.membership.ProposeMember(context.Background(), candidate);
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return key, nil
}

// SignCert sign cert over rpc
func (rpc *RPC) SignCert(cert *pb.Cert) error {
	if _, err := rpc.membership.SignCert(context.Background(), cert); err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				return errors.ErrNotFound
			}
		}

		return errors.WithStack(err)
	}
	return nil
}

// FetchMember fetch member over rpc
func (rpc *RPC) FetchMember(key *pb.PublicKey) (*pb.Member, error) {
	r, err := rpc.membership.FetchMember(context.Background(), key)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				return nil, errors.ErrNotFound
			}
		}

		return nil, errors.WithStack(err)
	}

	return r, nil
}

// InitNode init node in discovery service
func (rpc *RPC) InitNode(key *pb.PublicKey) (*pb.MasterChain, *pb.BrokerOpts, error) {
	ino, err := rpc.discovery.InitNode(context.Background(), &pb.Empty{})
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	mc, err := rpc.master.InitNode(context.Background(), key)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return mc, ino.Broker, nil
}
