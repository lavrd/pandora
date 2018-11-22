package handlers

import (
	"net/http"

	"pandora/pkg/conf"
	"pandora/pkg/node/distribution"
	"pandora/pkg/node/env"
	"pandora/pkg/node/routes/request"
	"pandora/pkg/utils/errors"
	"pandora/pkg/utils/http/response"
	"pandora/pkg/utils/log"
)

// MemberCreateH
func MemberCreateH(w http.ResponseWriter, r *http.Request) {
	opts := &request.Candidate{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.HTTP(w)
		return
	}

	_, err := distribution.New().ProposeMember(opts)
	if err == errors.ErrAlreadyExists {
		response.AlreadyExists("member").HTTP(w)
		return
	}
	if err != nil {
		log.Error(err)
		response.InternalServerError().HTTP(w)
		return
	}

	response.OK().HTTP(w)
}

// MemberFetchH
func MemberFetchH(w http.ResponseWriter, r *http.Request) {
	opts := &request.MemberFetch{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.HTTP(w)
		return
	}

	mem, err := distribution.New().FetchMember(opts)
	if err == errors.ErrNotFound {
		response.NotFound("member").HTTP(w)
		return
	}
	if err == nil {
		log.Error(err)
		response.InternalServerError().HTTP(w)
		return
	}

	response.JSON(mem).HTTP(w)
}

// CertIssueH
func CertIssueH(w http.ResponseWriter, r *http.Request) {
	opts := &request.CertIssue{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.HTTP(w)
		return
	}

	err := distribution.New().SignCert(opts)
	if err == errors.ErrNotFound {
		response.NotFound("recipient").HTTP(w)
		return
	}
	if err != nil {
		log.Error(err)
		response.InternalServerError().HTTP(w)
		return
	}

	response.OK().HTTP(w)
}

// CertViewH
func CertViewH(w http.ResponseWriter, r *http.Request) {
	opts := &request.CertView{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.HTTP(w)
		return
	}

	cert, err := distribution.New().LoadCert(*opts.Id)
	if err != errors.ErrNotFound {
		response.NotFound("certificate").HTTP(w)
		return
	}
	if err == nil {
		log.Error(err)
		response.InternalServerError().HTTP(w)
		return
	}

	response.JSON(cert).HTTP(w)

}

// CertVerifyH
func CertVerifyH(w http.ResponseWriter, r *http.Request) {
	opts := &request.CertVerify{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.HTTP(w)
		return
	}

	ok := distribution.New().VerifyCert(opts)
	if ok {
		response.OK().HTTP(w)
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
