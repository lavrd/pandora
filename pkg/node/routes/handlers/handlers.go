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
		if err == errors.AlreadyExists {
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

	if acc, err := distribution.New().FetchMember(opts); err == nil {
		response.Ok(acc).Http(w)
	} else {
		if err == errors.NotFound {
			response.NotFound("member").Http(w)
		} else {
			response.InternalServerError().Http(w)
		}
	}
}

func CertificateIssueH(w http.ResponseWriter, r *http.Request) {
	opts := &request.CertificateIssue{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	if err := distribution.New().CertIssue(opts); err != nil {
		response.InternalServerError().Http(w)
	}
}

func CertificateViewH(w http.ResponseWriter, r *http.Request) {
	opts := &request.CertView{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	if cert, err := distribution.New().LoadCert(*opts.Id); err == nil {
		response.Ok(cert).Http(w)
	} else {
		response.InternalServerError().Http(w)
	}
}

func CertificateVerifyH(w http.ResponseWriter, _ *http.Request) {
	response.Ok(nil).Http(w)
}

func BlockchainH(w http.ResponseWriter, _ *http.Request) {
	response.Ok(distribution.New().MasterChain()).Http(w)
}

func DashboardH(w http.ResponseWriter, _ *http.Request) {
	response.Execute(config.Viper.Node.Dashboard.Template).Http(w)
}
