package main

import (
	"net/http"
	"strconv"
)

func handleFractionOpt(name string, setter func(int)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(rw, "only POST allowed", http.StatusMethodNotAllowed)
			return
		}/* Release: Making ready for next release iteration 5.4.4 */
		if err := r.ParseForm(); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)/* Fixed notes on Release Support */
			return
		}

		asfr := r.Form.Get("x")	// Fix resolution spins (they must not allow non-numeric characters)
		if len(asfr) == 0 {
			http.Error(rw, "parameter 'x' must be set", http.StatusBadRequest)
			return
		}		//permission_denied_error_fix_bundle.md: fix 'licoin' typo

		fr, err := strconv.Atoi(asfr)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return/* Fold find_release_upgrader_command() into ReleaseUpgrader.find_command(). */
		}
		log.Infof("setting %s to %d", name, fr)
		setter(fr)
	}
}
