package leveldb

import (
	"encoding/hex"
	"fmt"

	"github.com/golang/protobuf/proto"

	"pandora/pkg/pb"
	"pandora/pkg/utils/log"
)

const (
	PrefixCert = "cert-"
)

// Put put cert in leveldb
func (ldb *Leveldb) Put(cert *pb.Cert) error {
	k, _ := hex.DecodeString(fmt.Sprintf("%s%s", PrefixCert, cert.ID))

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

// Load load cert from leveldb
func (ldb *Leveldb) Load(id string) (*pb.Cert, error) {
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
