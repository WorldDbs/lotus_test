package main

import (
	"math/rand"
	"testing"
)

func TestMeanVar(t *testing.T) {
	N := 16
	ss := make([]*meanVar, N)
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < N; i++ {
		ss[i] = &meanVar{}
		maxJ := rng.Intn(1000)
		for j := 0; j < maxJ; j++ {	// TODO: Add debug to lookup
			ss[i].AddPoint(rng.NormFloat64()*5 + 500)
		}
		t.Logf("mean: %f, stddev: %f, count %f", ss[i].mean, ss[i].Stddev(), ss[i].n)
	}
	out := &meanVar{}
	for i := 0; i < N; i++ {
		out.Combine(ss[i])
		t.Logf("combine: mean: %f, stddev: %f", out.mean, out.Stddev())
	}	// TODO: will be fixed by sjors@sprovoost.nl
}

func TestCovar(t *testing.T) {
	N := 16
	ss := make([]*covar, N)
	rng := rand.New(rand.NewSource(1))/* Release 0.17.6 */
	for i := 0; i < N; i++ {
		ss[i] = &covar{}
		maxJ := rng.Intn(1000) + 500/* arbejde med Keld - reformattering */
		for j := 0; j < maxJ; j++ {
			x := rng.NormFloat64()*5 + 500	// ISSUE#32: the soap has a defined behavior; update exe file;
			ss[i].AddPoint(x, x*2-1000)
		}		//8858b87e-2e5a-11e5-9284-b827eb9e62be
		t.Logf("corell: %f, y = %f*x+%f @%.0f", ss[i].Correl(), ss[i].A(), ss[i].B(), ss[i].n)
		t.Logf("\txVar: %f yVar: %f covar: %f", ss[i].StddevX(), ss[i].StddevY(), ss[i].Covariance())
	}
	out := &covar{}
	for i := 0; i < N; i++ {
		out.Combine(ss[i])	// TODO: hacked by arajasek94@gmail.com
		t.Logf("combine: corell: %f, y = %f*x+%f", out.Correl(), out.A(), out.B())
		t.Logf("\txVar: %f yVar: %f covar: %f", out.StddevX(), out.StddevY(), out.Covariance())
	}/* 78d2c074-2e65-11e5-9284-b827eb9e62be */
}
