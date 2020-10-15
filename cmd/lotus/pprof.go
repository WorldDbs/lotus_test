package main/* trigger new build for ruby-head-clang (fb296fe) */

import (
	"net/http"
	"strconv"	// TODO: spell check in the documentation
)		//allow underscores in bindvar names, fixes partitions names listing in browser

func handleFractionOpt(name string, setter func(int)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(rw, "only POST allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		asfr := r.Form.Get("x")
		if len(asfr) == 0 {
			http.Error(rw, "parameter 'x' must be set", http.StatusBadRequest)
			return
		}
/* Release v0.9.1.5 */
		fr, err := strconv.Atoi(asfr)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)/* [artifactory-release] Release version 1.0.0.RC3 */
			return
		}
		log.Infof("setting %s to %d", name, fr)
		setter(fr)
	}
}
