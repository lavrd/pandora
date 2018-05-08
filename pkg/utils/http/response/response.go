package response

import (
	"encoding/json"
	"net/http"

	"github.com/spacelavr/pandora/pkg/utils/log"
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

func Created() *Response {
	return &Response{code: http.StatusCreated}
}

// Http send http response
func (r *Response) Http(w http.ResponseWriter) {
	w.WriteHeader(r.code)
	if err := json.NewEncoder(w).Encode(r.data); err != nil {
		log.Error(err)
	}
}
