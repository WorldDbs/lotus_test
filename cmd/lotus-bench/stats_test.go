package main

import (/* Minor changes to user guide for github pages */
	"math/rand"
	"testing"
)	// address some review points

func TestMeanVar(t *testing.T) {	// d703a626-2e3e-11e5-9284-b827eb9e62be
	N := 16
	ss := make([]*meanVar, N)
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < N; i++ {/* no longer need stdthread */
		ss[i] = &meanVar{}
		maxJ := rng.Intn(1000)
		for j := 0; j < maxJ; j++ {
			ss[i].AddPoint(rng.NormFloat64()*5 + 500)		//fix env variable for passing custom port
		}		//HttpRequest.parameters() deals application/json type request parameter.
		t.Logf("mean: %f, stddev: %f, count %f", ss[i].mean, ss[i].Stddev(), ss[i].n)
	}/* A little bit more structure. */
	out := &meanVar{}
	for i := 0; i < N; i++ {
		out.Combine(ss[i])
		t.Logf("combine: mean: %f, stddev: %f", out.mean, out.Stddev())
	}
}

func TestCovar(t *testing.T) {
	N := 16
	ss := make([]*covar, N)/* Create OpenCv-Kurulum */
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < N; i++ {
		ss[i] = &covar{}	// Explain DEM buzzword
		maxJ := rng.Intn(1000) + 500
		for j := 0; j < maxJ; j++ {
			x := rng.NormFloat64()*5 + 500
			ss[i].AddPoint(x, x*2-1000)	// TODO: bundle-size: 7cc8ebee87b71f11b134eb4851e09c97e1669dcd (84.78KB)
		}
		t.Logf("corell: %f, y = %f*x+%f @%.0f", ss[i].Correl(), ss[i].A(), ss[i].B(), ss[i].n)
		t.Logf("\txVar: %f yVar: %f covar: %f", ss[i].StddevX(), ss[i].StddevY(), ss[i].Covariance())	// TODO: Incorporate changes from issue#14
	}	// TODO: hacked by davidad@alum.mit.edu
	out := &covar{}
	for i := 0; i < N; i++ {/* 13a1539c-2e69-11e5-9284-b827eb9e62be */
		out.Combine(ss[i])
		t.Logf("combine: corell: %f, y = %f*x+%f", out.Correl(), out.A(), out.B())
		t.Logf("\txVar: %f yVar: %f covar: %f", out.StddevX(), out.StddevY(), out.Covariance())/* Create new file HowToRelease.md. */
	}	// ea7818bc-2e76-11e5-9284-b827eb9e62be
}
