package handlers

import (
	"net/http"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/node/distribution"
	"github.com/spacelavr/pandora/pkg/node/routes/request"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/http/response"
)

func HealthH(w http.ResponseWriter, _ *http.Request) {
	response.NotImplemented().Http(w)
}

func MemberCreateH(w http.ResponseWriter, r *http.Request) {
	opts := &request.Candidate{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	if _, err := distribution.New().ProposeMember(opts); err != nil {
		if err == errors.ErrAlreadyExists {
			response.AlreadyExists("member").Http(w)
		} else {
			response.InternalServerError().Http(w)
		}
	}
}

func MemberFetchH(w http.ResponseWriter, r *http.Request) {
	opts := &request.MemberFetch{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	if mem, err := distribution.New().FetchMember(opts); err == nil {
		response.Ok(mem).Http(w)
	} else {
		if err == errors.ErrNotFound {
			response.NotFound("member").Http(w)
		} else {
			response.InternalServerError().Http(w)
		}
	}
}

func CertIssueH(w http.ResponseWriter, r *http.Request) {
	opts := &request.CertIssue{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	if err := distribution.New().SignCert(opts); err != nil {
		if err == errors.ErrNotFound {
			response.NotFound("recipient").Http(w)
		} else {
			response.InternalServerError().Http(w)
		}
	}
}

func CertViewH(w http.ResponseWriter, r *http.Request) {
	opts := &request.CertView{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	if cert, err := distribution.New().LoadCert(*opts.Id); err == nil {
		response.Ok(cert).Http(w)
	} else {
		if err != errors.ErrNotFound {
			response.NotFound("certificate").Http(w)
		} else {
			response.InternalServerError().Http(w)
		}
	}
}

func CertVerifyH(w http.ResponseWriter, r *http.Request) {
	opts := &request.CertVerify{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	if ok := distribution.New().VerifyCert(opts); ok {
		response.Ok(nil).Http(w)
	} else {
		response.InternalServerError().Http(w)
	}
}

func BlockchainH(w http.ResponseWriter, _ *http.Request) {
	response.Ok(distribution.New().MasterChain()).Http(w)
}

func DashboardH(w http.ResponseWriter, _ *http.Request) {
	response.Execute(config.Viper.Node.Dashboard.Template).Http(w)
}
