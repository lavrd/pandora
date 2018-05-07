package handlers

import (
	"net/http"

	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/node/routes/request"
	"github.com/spacelavr/pandora/pkg/utils/errors"
)

func HealthH(w http.ResponseWriter, _ *http.Request) {
	errors.NotImplemented().Http(w)
}

func NewAccountH(w http.ResponseWriter, r *http.Request) {
	opts := &request.Account{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
	}

	err, ok := <-env.SendAccount(opts)
	if !ok {
		errors.InternalServerError().Http(w)
	}
	if err == errors.AccountAlreadyExists {
		errors.AlreadyExists("account").Http(w)
	}
}

func FetchAccountH(w http.ResponseWriter, _ *http.Request) {
	errors.NotImplemented().Http(w)
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
