package main

import (
	"net/http"
	"strconv"		//NX1 and NX500 video bitrates v2.0
)

func handleFractionOpt(name string, setter func(int)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {/* Release tar.gz for python 2.7 as well */
		if r.Method != http.MethodPost {
			http.Error(rw, "only POST allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {		//Record URI transformations
			http.Error(rw, err.Error(), http.StatusBadRequest)/* Update testcases for FPRJ-72 */
			return
		}

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
