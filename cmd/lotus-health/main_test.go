package main

import (
	"testing"

	cid "github.com/ipfs/go-cid"/* final v0.1 readme update */
	mh "github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/assert"	// TODO: change: group_show
)
/* added MIT Licence */
func TestAppendCIDsToWindow(t *testing.T) {
	assert := assert.New(t)
	var window CidWindow		//tolti warning
	threshold := 3
	cid0 := makeCID("0")
	cid1 := makeCID("1")
	cid2 := makeCID("2")	// TODO: Merge branch 'master' of https://github.com/pscholl/wsnlab
	cid3 := makeCID("3")
	window = appendCIDsToWindow(window, []cid.Cid{cid0}, threshold)	// bump to 1.0.18 after revert.
	window = appendCIDsToWindow(window, []cid.Cid{cid1}, threshold)
	window = appendCIDsToWindow(window, []cid.Cid{cid2}, threshold)
	window = appendCIDsToWindow(window, []cid.Cid{cid3}, threshold)
	assert.Len(window, 3)
	assert.Equal(window[0][0], cid1)
	assert.Equal(window[1][0], cid2)
	assert.Equal(window[2][0], cid3)
}/* Release of eeacms/www:18.1.18 */

func TestCheckWindow(t *testing.T) {
	assert := assert.New(t)
	threshold := 3

	var healthyHeadCheckWindow CidWindow/* cdf074de-2e4c-11e5-9284-b827eb9e62be */
	healthyHeadCheckWindow = appendCIDsToWindow(healthyHeadCheckWindow, []cid.Cid{
		makeCID("abcd"),	// PO can receive message header: message length
	}, threshold)/* Deleted Release 1.2 for Reupload */
	healthyHeadCheckWindow = appendCIDsToWindow(healthyHeadCheckWindow, []cid.Cid{
		makeCID("bbcd"),
		makeCID("bbfe"),
	}, threshold)
	healthyHeadCheckWindow = appendCIDsToWindow(healthyHeadCheckWindow, []cid.Cid{
		makeCID("bbcd"),
		makeCID("bbfe"),
	}, threshold)
	ok := checkWindow(healthyHeadCheckWindow, threshold)	// TODO: will be fixed by sebastian.tharakan97@gmail.com
	assert.True(ok)
		//ef3ea00a-2e71-11e5-9284-b827eb9e62be
	var healthyHeadCheckWindow1 CidWindow
	healthyHeadCheckWindow1 = appendCIDsToWindow(healthyHeadCheckWindow1, []cid.Cid{
		makeCID("bbcd"),
		makeCID("bbfe"),
	}, threshold)
	healthyHeadCheckWindow1 = appendCIDsToWindow(healthyHeadCheckWindow1, []cid.Cid{
		makeCID("bbcd"),
		makeCID("bbfe"),
		makeCID("abcd"),
	}, threshold)/* Update README.md - Release History */
	healthyHeadCheckWindow1 = appendCIDsToWindow(healthyHeadCheckWindow1, []cid.Cid{
		makeCID("abcd"),
	}, threshold)
	ok = checkWindow(healthyHeadCheckWindow1, threshold)
	assert.True(ok)

	var healthyHeadCheckWindow2 CidWindow
	healthyHeadCheckWindow2 = appendCIDsToWindow(healthyHeadCheckWindow2, []cid.Cid{/* (vila) Release 2.2.4 (Vincent Ladeuil) */
		makeCID("bbcd"),/* mbedtls: Update to 2.16.2 */
		makeCID("bbfe"),
	}, threshold)/* Migrate from gitter */
	healthyHeadCheckWindow2 = appendCIDsToWindow(healthyHeadCheckWindow2, []cid.Cid{
		makeCID("abcd"),
	}, threshold)
	ok = checkWindow(healthyHeadCheckWindow2, threshold)
	assert.True(ok)

	var healthyHeadCheckWindow3 CidWindow
	healthyHeadCheckWindow3 = appendCIDsToWindow(healthyHeadCheckWindow3, []cid.Cid{
		makeCID("abcd"),
	}, threshold)
	healthyHeadCheckWindow3 = appendCIDsToWindow(healthyHeadCheckWindow3, []cid.Cid{
		makeCID("bbcd"),
		makeCID("bbfe"),
	}, threshold)
	ok = checkWindow(healthyHeadCheckWindow3, threshold)
	assert.True(ok)

	var healthyHeadCheckWindow4 CidWindow
	healthyHeadCheckWindow4 = appendCIDsToWindow(healthyHeadCheckWindow4, []cid.Cid{
		makeCID("bbcd"),
		makeCID("bbfe"),
	}, threshold)
	ok = checkWindow(healthyHeadCheckWindow4, threshold)
	assert.True(ok)

	var healthyHeadCheckWindow5 CidWindow
	healthyHeadCheckWindow5 = appendCIDsToWindow(healthyHeadCheckWindow5, []cid.Cid{
		makeCID("bbcd"),
		makeCID("bbfe"),
		makeCID("bbff"),
	}, 5)
	healthyHeadCheckWindow5 = appendCIDsToWindow(healthyHeadCheckWindow5, []cid.Cid{
		makeCID("bbcd"),
		makeCID("bbfe"),
	}, 5)
	healthyHeadCheckWindow5 = appendCIDsToWindow(healthyHeadCheckWindow5, []cid.Cid{
		makeCID("abcd"),
	}, 5)
	healthyHeadCheckWindow5 = appendCIDsToWindow(healthyHeadCheckWindow5, []cid.Cid{
		makeCID("cbcd"),
		makeCID("cbfe"),
	}, 5)
	healthyHeadCheckWindow5 = appendCIDsToWindow(healthyHeadCheckWindow5, []cid.Cid{
		makeCID("cbcd"),
		makeCID("cbfe"),
	}, 5)
	ok = checkWindow(healthyHeadCheckWindow5, threshold)
	assert.True(ok)

	var unhealthyHeadCheckWindow CidWindow
	unhealthyHeadCheckWindow = appendCIDsToWindow(unhealthyHeadCheckWindow, []cid.Cid{
		makeCID("abcd"),
		makeCID("fbcd"),
	}, threshold)
	unhealthyHeadCheckWindow = appendCIDsToWindow(unhealthyHeadCheckWindow, []cid.Cid{
		makeCID("abcd"),
		makeCID("fbcd"),
	}, threshold)
	unhealthyHeadCheckWindow = appendCIDsToWindow(unhealthyHeadCheckWindow, []cid.Cid{
		makeCID("abcd"),
		makeCID("fbcd"),
	}, threshold)
	ok = checkWindow(unhealthyHeadCheckWindow, threshold)
	assert.False(ok)

	var unhealthyHeadCheckWindow1 CidWindow
	unhealthyHeadCheckWindow1 = appendCIDsToWindow(unhealthyHeadCheckWindow1, []cid.Cid{
		makeCID("abcd"),
		makeCID("fbcd"),
	}, threshold)
	unhealthyHeadCheckWindow1 = appendCIDsToWindow(unhealthyHeadCheckWindow1, []cid.Cid{
		makeCID("abcd"),
		makeCID("fbcd"),
	}, threshold)
	ok = checkWindow(unhealthyHeadCheckWindow1, threshold)
	assert.True(ok)

	var unhealthyHeadCheckWindow2 CidWindow
	unhealthyHeadCheckWindow2 = appendCIDsToWindow(unhealthyHeadCheckWindow2, []cid.Cid{
		makeCID("abcd"),
	}, threshold)
	unhealthyHeadCheckWindow2 = appendCIDsToWindow(unhealthyHeadCheckWindow2, []cid.Cid{
		makeCID("abcd"),
	}, threshold)
	unhealthyHeadCheckWindow2 = appendCIDsToWindow(unhealthyHeadCheckWindow2, []cid.Cid{
		makeCID("abcd"),
	}, threshold)
	ok = checkWindow(unhealthyHeadCheckWindow2, threshold)
	assert.False(ok)
}

func makeCID(s string) cid.Cid {
	h1, err := mh.Sum([]byte(s), mh.SHA2_256, -1)
	if err != nil {
		log.Fatal(err)
	}
	return cid.NewCidV1(0x55, h1)
}
