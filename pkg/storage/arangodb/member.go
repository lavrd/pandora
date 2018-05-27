package arangodb

import (
	"fmt"

	"github.com/spacelavr/pandora/pkg/pb"
)

const (
	CollectionMember = "member"
)

func (s *ArangoDB) MemberFetchByEmail(email string) (*pb.Member, error) {
	var (
		acc   = &pb.Member{}
		query = fmt.Sprintf(
			"for m in %s filter m.meta.email == @email return m",
			CollectionMember,
		)
		vars = map[string]interface{}{
			"email": email,
		}
	)

	if _, err := s.Exec(query, vars, acc); err != nil {
		return nil, err
	}

	return acc, nil
}

func (s *ArangoDB) MemberFetchByPublic(key *pb.PublicKey) (*pb.Member, error) {
	var (
		mem   = &pb.Member{}
		query = fmt.Sprintf(
			"for m in %s filter m.public_key.public_key == @public return m",
			CollectionMember,
		)
		vars = map[string]interface{}{
			"public": key.PublicKey,
		}
	)

	if _, err := s.Exec(query, vars, mem); err != nil {
		return nil, err
	}

	return mem, nil
}

func (s *ArangoDB) MemberSave(mem *pb.Member) error {
	_, err := s.Write(CollectionMember, mem)
	if err != nil {
		return err
	}
	return nil
}
