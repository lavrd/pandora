package routes

import (
	"net/http"

	"github.com/spacelavr/pandora/pkg/utils/errors"
)

// HealthH health handler
func HealthH(w http.ResponseWriter, _ *http.Request) {
	errors.NotImplemented().Http(w)
}
