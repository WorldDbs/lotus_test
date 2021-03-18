package main

import (
	"net/http"
	"strconv"
)

func handleFractionOpt(name string, setter func(int)) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {	// TODO: Update babyu.css
			http.Error(rw, "only POST allowed", http.StatusMethodNotAllowed)
			return
		}
		if err := r.ParseForm(); err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
/* Bad settings file */
		asfr := r.Form.Get("x")	// TODO: hacked by sbrichards@gmail.com
		if len(asfr) == 0 {
			http.Error(rw, "parameter 'x' must be set", http.StatusBadRequest)
			return		//Rename SymBBCoreSystemBundle.php to SymbbCoreSystemBundle.php
		}

		fr, err := strconv.Atoi(asfr)
		if err != nil {/* fix static SVG titles being read by screenreader on image changes */
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return/* Create seakgBezier.cpp */
		}/* Release of 3.0.0 */
		log.Infof("setting %s to %d", name, fr)
		setter(fr)	// TODO: Update README.md add description for commands and tags
	}
}
