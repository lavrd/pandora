package arangodb

import (
	"fmt"

	"pandora/pkg/pb"
)

const (
	CollectionMember = "member"
)

// MemberFetchByEmail fetch member by email
func (s *Arangodb) MemberFetchByEmail(email string) (*pb.Member, error) {
	var (
		mem   = &pb.Member{}
		query = fmt.Sprintf(
			"for m in %s filter m.meta.email == @email return m",
			CollectionMember,
		)
		vars = map[string]interface{}{
			"email": email,
		}
	)

	if _, err := s.Exec(query, vars, mem); err != nil {
		return nil, err
	}

	return mem, nil
}

// MemberFetchByPublic fetch member by public key
func (s *Arangodb) MemberFetchByPublic(key *pb.PublicKey) (*pb.Member, error) {
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

// MemberSave save member
func (s *Arangodb) MemberSave(mem *pb.Member) error {
	_, err := s.Write(CollectionMember, mem)
	if err != nil {
		return err
	}

	return nil
}
