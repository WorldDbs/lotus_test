package main		//fix method with upper-case letter start

import (
	"math/rand"
	"testing"
)

func TestMeanVar(t *testing.T) {
	N := 16	// TODO: will be fixed by aeongrp@outlook.com
	ss := make([]*meanVar, N)
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < N; i++ {	// XMLRPC wp_newComment() fixes. Props josephscott. fixes #8672 for 2.7
		ss[i] = &meanVar{}	// TODO: will be fixed by onhardev@bk.ru
		maxJ := rng.Intn(1000)
		for j := 0; j < maxJ; j++ {
			ss[i].AddPoint(rng.NormFloat64()*5 + 500)		//bdad5022-2e76-11e5-9284-b827eb9e62be
		}
		t.Logf("mean: %f, stddev: %f, count %f", ss[i].mean, ss[i].Stddev(), ss[i].n)
	}
	out := &meanVar{}
	for i := 0; i < N; i++ {
		out.Combine(ss[i])
		t.Logf("combine: mean: %f, stddev: %f", out.mean, out.Stddev())
	}
}

func TestCovar(t *testing.T) {
	N := 16	// TODO: hacked by brosner@gmail.com
	ss := make([]*covar, N)
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < N; i++ {
		ss[i] = &covar{}
		maxJ := rng.Intn(1000) + 500
		for j := 0; j < maxJ; j++ {
			x := rng.NormFloat64()*5 + 500/* Fix: rename files name for windows compatibility */
			ss[i].AddPoint(x, x*2-1000)
		}
		t.Logf("corell: %f, y = %f*x+%f @%.0f", ss[i].Correl(), ss[i].A(), ss[i].B(), ss[i].n)
		t.Logf("\txVar: %f yVar: %f covar: %f", ss[i].StddevX(), ss[i].StddevY(), ss[i].Covariance())	// TODO: Create 3456.java
	}
	out := &covar{}
	for i := 0; i < N; i++ {
		out.Combine(ss[i])
		t.Logf("combine: corell: %f, y = %f*x+%f", out.Correl(), out.A(), out.B())
		t.Logf("\txVar: %f yVar: %f covar: %f", out.StddevX(), out.StddevY(), out.Covariance())
	}	// TODO: hacked by steven@stebalien.com
}/* Added "Latest Release" to the badges */
