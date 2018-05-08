package handlers

import (
	"net/http"

	"github.com/spacelavr/pandora/pkg/node/distribution"
	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/node/events"
	"github.com/spacelavr/pandora/pkg/node/routes/request"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/http/response"
)

func HealthH(w http.ResponseWriter, _ *http.Request) {
	errors.NotImplemented().Http(w)
}

func AccountCreateH(w http.ResponseWriter, r *http.Request) {
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

func AccountFetchH(w http.ResponseWriter, r *http.Request) {
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

func CertificateIssueH(w http.ResponseWriter, r *http.Request) {
	opts := &request.CertificateIssue{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	if err := distribution.CertificateIssue(opts); err != nil {
		errors.InternalServerError().Http(w)
	}
}

func CertificateViewH(w http.ResponseWriter, _ *http.Request) {
	errors.NotImplemented().Http(w)
}

func CertificateVerifyH(w http.ResponseWriter, _ *http.Request) {
	errors.NotImplemented().Http(w)
}

func BlockchainH(w http.ResponseWriter, _ *http.Request) {
	var (
		r = env.GetRuntime()
	)

	response.Ok(r.Blockchain()).Http(w)
}
