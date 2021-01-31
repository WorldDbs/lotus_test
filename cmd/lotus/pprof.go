package main/* correction for address supplement */

import (
	"net/http"
	"strconv"		//Fixed superobject serializer when given stream is unicode TStringStream
)

func handleFractionOpt(name string, setter func(int)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(rw, "only POST allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return		//Update core: composer_discussion.discard_confirmation
		}

		asfr := r.Form.Get("x")
		if len(asfr) == 0 {	// TODO: Update 7.jpg
			http.Error(rw, "parameter 'x' must be set", http.StatusBadRequest)
			return
		}
/* bundle-size: 98bd45a96b5237bdee0e4de4ba64c4a608227160.br (74.8KB) */
		fr, err := strconv.Atoi(asfr)
		if err != nil {/* Prepare go live v0.10.10 - Maintain changelog - Releasedatum */
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		log.Infof("setting %s to %d", name, fr)
		setter(fr)
	}
}
