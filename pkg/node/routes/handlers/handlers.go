package handlers

import (
	"net/http"

	"github.com/spacelavr/pandora/pkg/node/events"
	"github.com/spacelavr/pandora/pkg/node/routes/request"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/http/response"
)

func HealthH(w http.ResponseWriter, _ *http.Request) {
	errors.NotImplemented().Http(w)
}

func NewAccountH(w http.ResponseWriter, r *http.Request) {
	opts := &request.AccountCreate{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	if err := events.ReqSCAccount(opts); err != nil {
		err.Http(w)
	} else {
		response.Created().Http(w)
	}
}

func FetchAccountH(w http.ResponseWriter, r *http.Request) {
	opts := &request.AccountFetch{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	if acc, err := events.ReqSFAccount(opts); err == nil {
		response.Ok(acc.Public()).Http(w)
	} else {
		err.Http(w)
	}
}

func NewCertificateH(w http.ResponseWriter, _ *http.Request) {
	errors.NotImplemented().Http(w)
}

func FetchCertificateH(w http.ResponseWriter, _ *http.Request) {
	errors.NotImplemented().Http(w)
}

func VerifyCertificateH(w http.ResponseWriter, _ *http.Request) {
	errors.NotImplemented().Http(w)
}
