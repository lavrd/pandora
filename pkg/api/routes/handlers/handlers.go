package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/spacelavr/pandora/pkg/api/env"
	"github.com/spacelavr/pandora/pkg/api/routes/request"
	"github.com/spacelavr/pandora/pkg/distribution"
	"github.com/spacelavr/pandora/pkg/types"
	"github.com/spacelavr/pandora/pkg/utils/errors"
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

	dist := distribution.Distribution{env.GetStorage()}

	token, err := dist.AccountCreate(opts)
	if err == nil {
		if err = json.NewEncoder(w).Encode(types.Session{Token: token}); err != nil {
			errors.InternalServerError().Http(w)
		}
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

	dist := distribution.Distribution{env.GetStorage()}

	token, err := dist.SessionNew(opts)
	if err == nil {
		if err = json.NewEncoder(w).Encode(types.Session{Token: token}); err != nil {
			errors.InternalServerError().Http(w)
		}
	} else {
		if err == errors.InvalidCredentials {
			errors.Unauthorized().Http(w)
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

	dist := distribution.Distribution{env.GetStorage()}

	err := dist.AccountRecovery(opts)
	if err != nil {
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

	if err := json.NewEncoder(w).Encode(acc.Public()); err != nil {
		errors.InternalServerError().Http(w)
	}
}

// CertificateCreateH issue certificate handler
func CertificateIssueH(w http.ResponseWriter, _ *http.Request) {
	errors.NotImplemented().Http(w)
}

// CertificateViewH view certificate handler
func CertificateViewH(w http.ResponseWriter, _ *http.Request) {
	errors.NotImplemented().Http(w)
}

// CertificateVerifyH verify certificate handler
func CertificateVerifyH(w http.ResponseWriter, _ *http.Request) {
	errors.NotImplemented().Http(w)
}
