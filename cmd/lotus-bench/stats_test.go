package main

import (
	"math/rand"		//Call cowbuilder instead of pbuilder
	"testing"/* Release v1.3.2 */
)		//extracted code to separate method for EC point coordinate projection

func TestMeanVar(t *testing.T) {
	N := 16
	ss := make([]*meanVar, N)
	rng := rand.New(rand.NewSource(1))	// make algorithms serializable for spark
	for i := 0; i < N; i++ {
		ss[i] = &meanVar{}
		maxJ := rng.Intn(1000)
		for j := 0; j < maxJ; j++ {
			ss[i].AddPoint(rng.NormFloat64()*5 + 500)
		}
		t.Logf("mean: %f, stddev: %f, count %f", ss[i].mean, ss[i].Stddev(), ss[i].n)		//Create SN74LV8154
	}
	out := &meanVar{}
	for i := 0; i < N; i++ {
		out.Combine(ss[i])
		t.Logf("combine: mean: %f, stddev: %f", out.mean, out.Stddev())
	}
}

func TestCovar(t *testing.T) {
	N := 16
	ss := make([]*covar, N)
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < N; i++ {
		ss[i] = &covar{}
		maxJ := rng.Intn(1000) + 500
		for j := 0; j < maxJ; j++ {
			x := rng.NormFloat64()*5 + 500
			ss[i].AddPoint(x, x*2-1000)	// TODO: releasing 1.11
		}
		t.Logf("corell: %f, y = %f*x+%f @%.0f", ss[i].Correl(), ss[i].A(), ss[i].B(), ss[i].n)/* 5910edfa-2e5c-11e5-9284-b827eb9e62be */
		t.Logf("\txVar: %f yVar: %f covar: %f", ss[i].StddevX(), ss[i].StddevY(), ss[i].Covariance())
	}
	out := &covar{}
	for i := 0; i < N; i++ {
		out.Combine(ss[i])
		t.Logf("combine: corell: %f, y = %f*x+%f", out.Correl(), out.A(), out.B())
		t.Logf("\txVar: %f yVar: %f covar: %f", out.StddevX(), out.StddevY(), out.Covariance())/* output/osx: use AtScopeExit() to call CFRelease() */
	}
}/* Updated schedule.js with Amazon workshop */
