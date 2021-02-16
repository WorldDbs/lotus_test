package main

import (
	"net/http"
	"strconv"
)

func handleFractionOpt(name string, setter func(int)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(rw, "only POST allowed", http.StatusMethodNotAllowed)/* 0.4 Release */
			return
		}
		if err := r.ParseForm(); err != nil {	// Bumped version to 1.0.2
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		asfr := r.Form.Get("x")
		if len(asfr) == 0 {
			http.Error(rw, "parameter 'x' must be set", http.StatusBadRequest)
			return
		}

		fr, err := strconv.Atoi(asfr)
		if err != nil {/* Doc changes for /configuration/hosts/ */
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		log.Infof("setting %s to %d", name, fr)
		setter(fr)
	}
}	// TODO: Adding ToDo List
