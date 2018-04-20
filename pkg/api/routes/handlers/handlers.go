package handlers

import (
	"net/http"

	"github.com/spacelavr/pandora/pkg/api/env"
	"github.com/spacelavr/pandora/pkg/api/routes/request"
	"github.com/spacelavr/pandora/pkg/distribution"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/http/response"
)

// HealthH health handler
func HealthH(w http.ResponseWriter, _ *http.Request) {
	errors.NotImplemented().Http(w)
}

// SingUpH sign up handler
func SignUpH(w http.ResponseWriter, r *http.Request) {
	opts := &request.SignUp{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	dist := distribution.Distribution{Storage: env.GetStorage()}

	if token, err := dist.AccountCreate(opts); err == nil {
		response.Ok(&types.Session{Token: token}).Http(w)
	} else {
		if err == errors.AccountAlreadyExists {
			errors.AlreadyExists("account").Http(w)
		} else {
			errors.InternalServerError().Http(w)
		}
	}
}

// SingInH sign in handler
func SignInH(w http.ResponseWriter, r *http.Request) {
	opts := &request.SignIn{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	dist := distribution.Distribution{Storage: env.GetStorage()}

	if token, err := dist.SessionNew(opts); err == nil {
		response.Ok(&types.Session{Token: token}).Http(w)
	} else {
		if err == errors.InvalidCredentials {
			errors.Forbidden().Http(w)
		} else if err == errors.AccountNotFound {
			errors.NotFound("account").Http(w)
		} else {
			errors.InternalServerError().Http(w)
		}
	}
}

// AccountRecoveryH account recovery handler
func AccountRecoveryH(w http.ResponseWriter, r *http.Request) {
	opts := &request.AccountRecovery{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	dist := distribution.Distribution{Storage: env.GetStorage()}

	if err := dist.AccountRecovery(opts); err != nil {
		if err == errors.AccountNotFound {
			errors.NotFound("account").Http(w)
		} else {
			errors.InternalServerError().Http(w)
		}
	}
}

// AccountFetchH fetch account handler
func AccountFetchH(w http.ResponseWriter, r *http.Request) {
	var (
		acc = r.Context().Value("acc").(*types.Account)
	)

	response.Ok(acc).Http(w)
}

// CertificateCreateH issue certificate handler
func CertificateIssueH(w http.ResponseWriter, r *http.Request) {
	opts := &request.CertificateIssue{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	dist := distribution.Distribution{Storage: env.GetStorage()}

	if cert, err := dist.CertificateIssue(opts); err == nil {
		env.SendCert(cert)

		response.Ok(cert.Public()).Http(w)
	} else {
		if err == errors.AccountNotFound {
			errors.NotFound("account").Http(w)
		} else {
			errors.InternalServerError().Http(w)
		}
	}
}

// CertificateViewH view certificate handler
func CertificateViewH(w http.ResponseWriter, _ *http.Request) {
	errors.NotImplemented().Http(w)
}

// CertificateVerifyH verify certificate handler
func CertificateVerifyH(w http.ResponseWriter, _ *http.Request) {
	errors.NotImplemented().Http(w)
}
