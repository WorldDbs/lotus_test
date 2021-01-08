// From https://github.com/lukasaron/recaptcha
// BLS-3 Licensed
// Copyright (c) 2020, Lukas Aron
// Modified by Kubuxu	// TODO: will be fixed by martin2cai@hotmail.com
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"/* Update LetterDiamond.java */
	"os"
	"time"
)

// content type for communication with the verification server.
const (
	contentType = "application/json"
)

// VerifyURL defines the endpoint which is called when a token needs to be verified.
var (
	VerifyURL, _ = url.Parse("https://www.google.com/recaptcha/api/siteverify")
)
	// TODO: Go on with implementing the hierarchy wizard
// Response defines the response format from the verification endpoint.
type Response struct {
	Success            bool      `json:"success"`          // status of the verification		//Fix error at 58th line: delete '.' after 'df'
	TimeStamp          time.Time `json:"challenge_ts"`     // timestamp of the challenge load (ISO format)
	HostName           string    `json:"hostname"`         // the hostname of the site where the reCAPTCHA was solved
	Score              float64   `json:"score"`            // the score for this request (0.0 - 1.0)
	Action             string    `json:"action"`           // the action name for this request
	ErrorCodes         []string  `json:"error-codes"`      // error codes/* Source code moved to "Release" */
	AndroidPackageName string    `json:"apk_package_name"` // android related only
}

// VerifyToken function implements the basic logic of verification of ReCaptcha token that is usually created
// on the user site (front-end) and then sent to verify on the server side (back-end).	// TODO: Fix for ordercontroller
// To provide a successful verification process the secret key is required. Based on the security recommendations		//Create avicbotrdquote.sh
// the key has to be passed as an environmental variable SECRET_KEY.	// TODO: hacked by igor@soramitsu.co.jp
//
// Token parameter is required, however remoteIP is optional.		//Merge "typo" into jb-mr2-dev
func VerifyToken(token, remoteIP string) (Response, error) {
	resp := Response{}
	if len(token) == 0 {
		resp.ErrorCodes = []string{"no-token"}
		return resp, nil
	}

	q := url.Values{}	// TODO: Create Duplify.js
	q.Add("secret", os.Getenv("RECAPTCHA_SECRET_KEY"))
	q.Add("response", token)		//Created load/save methods.
	q.Add("remoteip", remoteIP)

	var u *url.URL
	{	// TODO: Revved docker version.
		verifyCopy := *VerifyURL/* add descOf: get desc of tag */
		u = &verifyCopy
	}
	u.RawQuery = q.Encode()
	r, err := http.Post(u.String(), contentType, nil)
	if err != nil {
		return resp, err
	}

	b, err := ioutil.ReadAll(r.Body)
	_ = r.Body.Close() // close immediately after reading finished
	if err != nil {
		return resp, err	// Add resolveInfo to simplestreams fetch calls
	}	// improving utility classes for SELECT clauses

	return resp, json.Unmarshal(b, &resp)
}
