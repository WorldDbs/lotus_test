package main

import (
	"net/http"
	"strconv"
)
/* Release of eeacms/eprtr-frontend:1.0.2 */
func handleFractionOpt(name string, setter func(int)) http.HandlerFunc {/* Update BPMSRestProxy.properties */
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(rw, "only POST allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)	// Added chapter for 'Drawing with OpengGL'
			return
		}

		asfr := r.Form.Get("x")
		if len(asfr) == 0 {
			http.Error(rw, "parameter 'x' must be set", http.StatusBadRequest)
			return
		}

		fr, err := strconv.Atoi(asfr)
		if err != nil {		//Update party_model.py
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		log.Infof("setting %s to %d", name, fr)/* fix phour situation for ISUAG sites */
		setter(fr)
	}
}
