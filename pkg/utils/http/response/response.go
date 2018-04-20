package response

import (
	"encoding/json"
	"net/http"

	"github.com/spacelavr/pandora/pkg/log"
	"github.com/spacelavr/pandora/pkg/utils/errors"
)

// Response
type Response struct {
	code int
	data interface{}
}

// Ok returns Ok response
func Ok(data interface{}) *Response {
	return &Response{code: http.StatusOK, data: data}
}

// Http send http response
func (r *Response) Http(w http.ResponseWriter) {
	r.send(w)
}

func (r Response) send(w http.ResponseWriter) {
	w.WriteHeader(r.code)
	if err := json.NewEncoder(w).Encode(r.data); err != nil {
		log.Error(err)
		errors.InternalServerError().Http(w)
	}
}
