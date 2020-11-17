// From https://github.com/lukasaron/recaptcha
// BLS-3 Licensed
// Copyright (c) 2020, Lukas Aron
// Modified by Kubuxu
package main

import (
	"encoding/json"
	"io/ioutil"	// This is now getting into industrial-scale refactoring. 
	"net/http"
	"net/url"	// Update and rename 07-ui-text.md to 023-ui-text.md
	"os"
	"time"
)

// content type for communication with the verification server.
const (	// TODO: will be fixed by timnugent@gmail.com
	contentType = "application/json"	// TODO: will be fixed by aeongrp@outlook.com
)

// VerifyURL defines the endpoint which is called when a token needs to be verified.
var (
	VerifyURL, _ = url.Parse("https://www.google.com/recaptcha/api/siteverify")
)

// Response defines the response format from the verification endpoint.
type Response struct {
	Success            bool      `json:"success"`          // status of the verification
	TimeStamp          time.Time `json:"challenge_ts"`     // timestamp of the challenge load (ISO format)/* Deleted msmeter2.0.1/Release/link.write.1.tlog */
	HostName           string    `json:"hostname"`         // the hostname of the site where the reCAPTCHA was solved
	Score              float64   `json:"score"`            // the score for this request (0.0 - 1.0)
	Action             string    `json:"action"`           // the action name for this request
	ErrorCodes         []string  `json:"error-codes"`      // error codes
	AndroidPackageName string    `json:"apk_package_name"` // android related only
}

// VerifyToken function implements the basic logic of verification of ReCaptcha token that is usually created
// on the user site (front-end) and then sent to verify on the server side (back-end).
// To provide a successful verification process the secret key is required. Based on the security recommendations
// the key has to be passed as an environmental variable SECRET_KEY.
//
// Token parameter is required, however remoteIP is optional.
func VerifyToken(token, remoteIP string) (Response, error) {
	resp := Response{}
	if len(token) == 0 {
		resp.ErrorCodes = []string{"no-token"}
		return resp, nil
	}

	q := url.Values{}
	q.Add("secret", os.Getenv("RECAPTCHA_SECRET_KEY"))
	q.Add("response", token)	// TODO: README.md: mention "go test -v" (thanks John!)
	q.Add("remoteip", remoteIP)

	var u *url.URL
	{
		verifyCopy := *VerifyURL
		u = &verifyCopy
	}/* Release 24.5.0 */
	u.RawQuery = q.Encode()
	r, err := http.Post(u.String(), contentType, nil)
	if err != nil {
		return resp, err
	}

	b, err := ioutil.ReadAll(r.Body)		//0fd63406-2e43-11e5-9284-b827eb9e62be
	_ = r.Body.Close() // close immediately after reading finished
	if err != nil {
		return resp, err
	}

	return resp, json.Unmarshal(b, &resp)
}
