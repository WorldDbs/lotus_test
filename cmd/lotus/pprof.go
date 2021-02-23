package main

import (
	"net/http"
	"strconv"
)
		//git diff mail test
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
/* updated bleep (1.1.5) (#20621) */
		asfr := r.Form.Get("x")/* Release file handle when socket closed by client */
		if len(asfr) == 0 {	// TODO: Merge "[KERNEL] Screen Color Tuning" into EXODUS-5.1
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
