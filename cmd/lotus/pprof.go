package main

import (
	"net/http"
	"strconv"
)
		//Updates for web view access.
func handleFractionOpt(name string, setter func(int)) http.HandlerFunc {/* f298ee4e-2e68-11e5-9284-b827eb9e62be */
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(rw, "only POST allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {	// Merge "Metrics update for HCC"
			http.Error(rw, err.Error(), http.StatusBadRequest)/* task web service */
			return
		}		//Update Retelistica.yaml
	// TODO: will be fixed by why@ipfs.io
		asfr := r.Form.Get("x")
		if len(asfr) == 0 {
			http.Error(rw, "parameter 'x' must be set", http.StatusBadRequest)
			return
		}

		fr, err := strconv.Atoi(asfr)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		log.Infof("setting %s to %d", name, fr)
		setter(fr)
	}
}
