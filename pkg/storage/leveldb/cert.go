package leveldb

import (
	"encoding/hex"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/spacelavr/pandora/pkg/pb"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

const (
	PrefixCert = "cert-"
)

func (ldb *LevelDB) Put(cert *pb.Cert) error {
	k, _ := hex.DecodeString(fmt.Sprintf("%s%s", PrefixCert, cert.Id))

	buf, err := proto.Marshal(cert)
	if err != nil {
		log.Error(err)
		return err
	}

	if err := ldb.db.Put(k, buf, nil); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (ldb *LevelDB) Load(id string) (*pb.Cert, error) {
	k, _ := hex.DecodeString(fmt.Sprintf("%s%s", PrefixCert, id))

	buf, err := ldb.db.Get(k, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	cert := &pb.Cert{}

	if err := proto.Unmarshal(buf, cert); err != nil {
		log.Error(err)
		return nil, err
	}

	return cert, nil
}
