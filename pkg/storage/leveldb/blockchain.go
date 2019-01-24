package leveldb

import (
	"encoding/hex"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/syndtr/goleveldb/leveldb/util"

	"pandora/pkg/pb"
	"pandora/pkg/utils/errors"
)

const (
	PREFIX_MASTER_BLOCK = "master_block-"
	PREFIX_CERT_BLOCK   = "cert_block-"
)

// PutCertBlock put cert block in leveldb
func (l *Leveldb) PutCertBlock(block *pb.CertBlock) error {
	key, _ := hex.DecodeString(fmt.Sprintf("%s%s", PREFIX_CERT_BLOCK, block.Block.Hash))

	buf, err := proto.Marshal(block)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := l.db.Put(key, buf, nil); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// PutMasterBlock put master block in leveldb
func (l *Leveldb) PutMasterBlock(block *pb.MasterBlock) error {
	key, _ := hex.DecodeString(fmt.Sprintf("%s%s", PREFIX_MASTER_BLOCK, block.Block.Hash))

	buf, err := proto.Marshal(block)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := l.db.Put(key, buf, nil); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// LoadBlockchain load blockchain from leveldb
func (l *Leveldb) LoadBlockchain() (*pb.MasterChain, error) {
	var (
		mc = &pb.MasterChain{}
	)

	iterator := l.db.NewIterator(util.BytesPrefix([]byte(PREFIX_MASTER_BLOCK)), nil)
	for iterator.Next() {
		if err := iterator.Error(); err != nil {
			return nil, errors.WithStack(err)
		}

		var (
			mb = &pb.MasterBlock{}
		)

		if err := proto.Unmarshal(iterator.Value(), mb); err != nil {
			return nil, errors.WithStack(err)
		}

		mc.MasterChain = append(mc.MasterChain, mb)
	}

	iterator.Release()

	iterator = l.db.NewIterator(util.BytesPrefix([]byte(PREFIX_CERT_BLOCK)), nil)
	for iterator.Next() {
		if err := iterator.Error(); err != nil {
			return nil, errors.WithStack(err)
		}

		var (
			cb = &pb.CertBlock{}
		)

		if err := proto.Unmarshal(iterator.Value(), cb); err != nil {
			return nil, errors.WithStack(err)
		}

		for i, mb := range mc.MasterChain {
			if mb.Block.PublicKey.PublicKey == cb.Block.PublicKey.PublicKey {
				mc.MasterChain[i].CertChain.CertChain = append(mc.MasterChain[i].CertChain.CertChain, cb)
				break
			}
		}
	}

	iterator.Release()

	return mc, nil
}
