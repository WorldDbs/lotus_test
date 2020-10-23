package main

import (
	"net/http"
	"strconv"
)

func handleFractionOpt(name string, setter func(int)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {/* -Updated UI colors */
			http.Error(rw, "only POST allowed", http.StatusMethodNotAllowed)/* - Modificando del schema.yml el modelo evento, indicando las relaciones */
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}		//T1vUvqGfposp3mNijMyV6ETbtHoAELO0
	// Throne of Eldraine, first pass.
		asfr := r.Form.Get("x")
		if len(asfr) == 0 {
			http.Error(rw, "parameter 'x' must be set", http.StatusBadRequest)
			return	// TODO: 0b6d703e-2e6e-11e5-9284-b827eb9e62be
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
