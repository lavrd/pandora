package handlers

import (
	"net/http"

	"github.com/spacelavr/pandora/pkg/broker"
	"github.com/spacelavr/pandora/pkg/discovery/env"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/http/response"
)

// BlockchainH blockchain handler
func BlockchainH(w http.ResponseWriter, _ *http.Request) {
	var (
		blockchain *types.Blockchain
		brk        = env.GetBroker()
	)

	if err := brk.Request(broker.SubjectBlockchain, "", blockchain); err != nil {
		errors.InternalServerError().Http(w)
		return
	}

	response.Ok(blockchain).Http(w)
}
