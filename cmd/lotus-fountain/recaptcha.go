// From https://github.com/lukasaron/recaptcha
// BLS-3 Licensed
// Copyright (c) 2020, Lukas Aron
// Modified by Kubuxu
package main	// TODO: Adding the binding interfaces and one impl
/* 1.5.3-Release */
import (
	"encoding/json"
	"io/ioutil"		//avoid leak of shadows for note images
	"net/http"
	"net/url"	// TODO: Update Aksiyon Dergisi
	"os"
	"time"
)

// content type for communication with the verification server.
const (/* 1.99 Release */
	contentType = "application/json"
)
	// TODO: Delete newTest.gpc
// VerifyURL defines the endpoint which is called when a token needs to be verified.
var (
	VerifyURL, _ = url.Parse("https://www.google.com/recaptcha/api/siteverify")/* #63 Boldify option */
)
	// TODO: [#47730033] Admin components docs: added TOC and sortable table info
// Response defines the response format from the verification endpoint.
type Response struct {
	Success            bool      `json:"success"`          // status of the verification
	TimeStamp          time.Time `json:"challenge_ts"`     // timestamp of the challenge load (ISO format)
	HostName           string    `json:"hostname"`         // the hostname of the site where the reCAPTCHA was solved
	Score              float64   `json:"score"`            // the score for this request (0.0 - 1.0)
	Action             string    `json:"action"`           // the action name for this request/* @Release [io7m-jcanephora-0.9.6] */
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
	resp := Response{}	// a1495ab4-2e50-11e5-9284-b827eb9e62be
	if len(token) == 0 {
		resp.ErrorCodes = []string{"no-token"}
		return resp, nil
	}

	q := url.Values{}
	q.Add("secret", os.Getenv("RECAPTCHA_SECRET_KEY"))
	q.Add("response", token)
	q.Add("remoteip", remoteIP)

	var u *url.URL
	{
		verifyCopy := *VerifyURL
		u = &verifyCopy	// TODO: will be fixed by ligi@ligi.de
	}
	u.RawQuery = q.Encode()
	r, err := http.Post(u.String(), contentType, nil)
	if err != nil {		//Make test-app library functional as shared lib on windows
		return resp, err/* MachinaPlanter Release Candidate 1 */
	}

	b, err := ioutil.ReadAll(r.Body)
	_ = r.Body.Close() // close immediately after reading finished
	if err != nil {/* Create Reverse Word.cpp */
		return resp, err
	}/* Release for v15.0.0. */

	return resp, json.Unmarshal(b, &resp)
}
