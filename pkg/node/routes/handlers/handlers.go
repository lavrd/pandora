package handlers

import (
	"net/http"

	"pandora/pkg/conf"
	"pandora/pkg/node/distribution"
	"pandora/pkg/node/env"
	"pandora/pkg/node/routes/request"
	"pandora/pkg/utils/errors"
	"pandora/pkg/utils/http/response"
)

// MemberCreateH
func MemberCreateH(w http.ResponseWriter, r *http.Request) {
	opts := &request.Candidate{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.HTTP(w)
		return
	}

	if _, err := distribution.New().ProposeMember(opts); err != nil {
		if err == errors.ErrAlreadyExists {
			response.AlreadyExists("member").HTTP(w)
		} else {
			response.InternalServerError().HTTP(w)
		}
	}
}

// MemberFetchH
func MemberFetchH(w http.ResponseWriter, r *http.Request) {
	opts := &request.MemberFetch{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.HTTP(w)
		return
	}

	if mem, err := distribution.New().FetchMember(opts); err == nil {
		response.JSON(mem).HTTP(w)
	} else {
		if err == errors.ErrNotFound {
			response.NotFound("member").HTTP(w)
		} else {
			response.InternalServerError().HTTP(w)
		}
	}
}

// CertIssueH
func CertIssueH(w http.ResponseWriter, r *http.Request) {
	opts := &request.CertIssue{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.HTTP(w)
		return
	}

	if err := distribution.New().SignCert(opts); err != nil {
		if err == errors.ErrNotFound {
			response.NotFound("recipient").HTTP(w)
		} else {
			response.InternalServerError().HTTP(w)
		}
	}
}

// CertViewH
func CertViewH(w http.ResponseWriter, r *http.Request) {
	opts := &request.CertView{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.HTTP(w)
		return
	}

	if cert, err := distribution.New().LoadCert(*opts.Id); err == nil {
		response.JSON(cert).HTTP(w)
	} else {
		if err != errors.ErrNotFound {
			response.NotFound("certificate").HTTP(w)
		} else {
			response.InternalServerError().HTTP(w)
		}
	}
}

// CertVerifyH
func CertVerifyH(w http.ResponseWriter, r *http.Request) {
	opts := &request.CertVerify{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.HTTP(w)
		return
	}

	if ok := distribution.New().VerifyCert(opts); ok {
		response.JSON(nil).HTTP(w)
	} else {
		response.InternalServerError().HTTP(w)
	}
}

// BlockchainH
func BlockchainH(w http.ResponseWriter, _ *http.Request) {
	response.JSON(env.GetBlockchain().GetMasterChain()).HTTP(w)
}

// DashboardH
func DashboardH(w http.ResponseWriter, _ *http.Request) {
	response.Execute(conf.Conf.Node.Dashboard).HTTP(w)
}
