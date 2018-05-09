package handlers

import (
	"html/template"
	"net/http"

	"github.com/spacelavr/pandora/pkg/config"
	"github.com/spacelavr/pandora/pkg/node/distribution"
	"github.com/spacelavr/pandora/pkg/node/env"
	"github.com/spacelavr/pandora/pkg/node/routes/request"
	"github.com/spacelavr/pandora/pkg/utils/errors"
	"github.com/spacelavr/pandora/pkg/utils/http/response"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

func HealthH(w http.ResponseWriter, _ *http.Request) {
	response.NotImplemented().Http(w)
}

func CandidateH(w http.ResponseWriter, r *http.Request) {
	opts := &request.Candidate{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	dist := &distribution.Distribution{}

	if err := dist.Candidate(opts); err != nil {
		log.Debug(err)
		if err == errors.AlreadyExists {
			response.AlreadyExists("account").Http(w)
		} else {
			response.InternalServerError().Http(w)
		}
	}
}

func AccountFetchH(w http.ResponseWriter, r *http.Request) {
	opts := &request.AccountFetch{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}
}

func CertificateIssueH(w http.ResponseWriter, r *http.Request) {
	opts := &request.CertificateIssue{}
	if err := opts.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	if err := distribution.CertificateIssue(opts); err != nil {
		response.InternalServerError().Http(w)
	}
}

func CertificateViewH(w http.ResponseWriter, _ *http.Request) {
	response.NotImplemented().Http(w)
}

func CertificateVerifyH(w http.ResponseWriter, _ *http.Request) {
	response.NotImplemented().Http(w)
}

func BlockchainH(w http.ResponseWriter, _ *http.Request) {
	var (
		_ = env.GetRuntime()
	)

	// response.Ok(r.Blockchain()).Http(w)
}

func DashboardH(w http.ResponseWriter, _ *http.Request) {
	// todo move to http utils
	w.Header().Set("Content-Type", "text/html")
	tpl, err := template.ParseFiles(config.Viper.Dashboard.Template)
	if err != nil {
		response.InternalServerError().Http(w)
		return
	}

	tpl = template.Must(tpl, err)

	if err = tpl.Execute(w, nil); err != nil {
		response.InternalServerError().Http(w)
	}
}
