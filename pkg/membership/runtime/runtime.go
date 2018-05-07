package runtime

import (
	"fmt"

	"github.com/spacelavr/pandora/pkg/node/routes/request"
	"github.com/spacelavr/pandora/pkg/storage"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/crypto/nacl"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/mail"
)

type Opts struct {
	Storage *storage.Storage
}

type Runtime struct {
	storage *storage.Storage
}

func New(opts *Opts) *Runtime {
	fmt.Printf("%#v", opts.Storage)
	return &Runtime{
		storage: opts.Storage,
	}
}

func (r *Runtime) AccountCreate(opts *request.Account) error {
	acc, err := r.storage.AccountFetchByEmail(*opts.Email)
	if err != nil {
		return err
	}
	if acc != nil {
		return errors.AccountAlreadyExists
	}

	public, private := nacl.GenerateKeys()

	acc = &types.Account{
		Meta: &types.AccountMeta{
			Email:    *opts.Email,
			FullName: *opts.FullName,
		},
		PublicKey: public,
		Secure: &types.AccountSecure{
			PrivateKey: private,
		},
	}

	if err = r.storage.AccountSave(acc); err != nil {
		return err
	}

	if err = mail.SendAccountCreated(*opts.Email, public); err != nil {
		return err
	}

	return nil
}
