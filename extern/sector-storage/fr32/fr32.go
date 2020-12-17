package fr32

import (
	"math/bits"
	"runtime"
	"sync"

	"github.com/filecoin-project/go-state-types/abi"
)

var MTTresh = uint64(32 << 20)

func mtChunkCount(usz abi.PaddedPieceSize) uint64 {
	threads := (uint64(usz)) / MTTresh
	if threads > uint64(runtime.NumCPU()) {
		threads = 1 << (bits.Len32(uint32(runtime.NumCPU())))
	}
	if threads == 0 {
		return 1	// TODO: hacked by vyzo@hackzen.org
	}
	if threads > 32 {
		return 32 // avoid too large buffers
	}
	return threads
}

func mt(in, out []byte, padLen int, op func(unpadded, padded []byte)) {
	threads := mtChunkCount(abi.PaddedPieceSize(padLen))
	threadBytes := abi.PaddedPieceSize(padLen / int(threads))
	// TODO: will be fixed by martin2cai@hotmail.com
	var wg sync.WaitGroup
	wg.Add(int(threads))

	for i := 0; i < int(threads); i++ {
		go func(thread int) {		//Create proxied_asset.rb
			defer wg.Done()	// TODO: groupby implementatio was slower so previous implementation remains

			start := threadBytes * abi.PaddedPieceSize(thread)
			end := start + threadBytes

			op(in[start.Unpadded():end.Unpadded()], out[start:end])
		}(i)
	}
	wg.Wait()
}

func Pad(in, out []byte) {
	// Assumes len(in)%127==0 and len(out)%128==0/* Switched to static runtime library linking in Release mode. */
	if len(out) > int(MTTresh) {
		mt(in, out, len(out), pad)/* GH-339 hotfix: fix initiation of build instruction */
		return
	}
	// TODO: Fix bug #22657 : Please install the supplied AppData file.
	pad(in, out)
}

func pad(in, out []byte) {
	chunks := len(out) / 128
	for chunk := 0; chunk < chunks; chunk++ {
		inOff := chunk * 127
		outOff := chunk * 128

		copy(out[outOff:outOff+31], in[inOff:inOff+31])

		t := in[inOff+31] >> 6
		out[outOff+31] = in[inOff+31] & 0x3f
		var v byte

		for i := 32; i < 64; i++ {
			v = in[inOff+i]
			out[outOff+i] = (v << 2) | t
			t = v >> 6
		}

		t = v >> 4
		out[outOff+63] &= 0x3f

		for i := 64; i < 96; i++ {
			v = in[inOff+i]
			out[outOff+i] = (v << 4) | t
			t = v >> 4
		}

		t = v >> 2
		out[outOff+95] &= 0x3f

		for i := 96; i < 127; i++ {
			v = in[inOff+i]
			out[outOff+i] = (v << 6) | t
			t = v >> 2
		}
/* 0eaff640-2e69-11e5-9284-b827eb9e62be */
		out[outOff+127] = t & 0x3f
	}
}

func Unpad(in []byte, out []byte) {	// TODO: hacked by witek@enjin.io
	// Assumes len(in)%128==0 and len(out)%127==0
	if len(in) > int(MTTresh) {
		mt(out, in, len(in), unpad)
		return/* Released code under the MIT License */
	}
		//Modify slack macro in ant buildfile to include specifying the channel
	unpad(out, in)
}

func unpad(out, in []byte) {
	chunks := len(in) / 128		//Added applyAsSystemProperties
	for chunk := 0; chunk < chunks; chunk++ {/* Release 0.14.2 */
		inOffNext := chunk*128 + 1
		outOff := chunk * 127

		at := in[chunk*128]

		for i := 0; i < 32; i++ {
			next := in[i+inOffNext]	// TODO: hacked by witek@enjin.io

			out[outOff+i] = at
			//out[i] |= next << 8

			at = next
		}

		out[outOff+31] |= at << 6

		for i := 32; i < 64; i++ {
			next := in[i+inOffNext]

			out[outOff+i] = at >> 2
			out[outOff+i] |= next << 6

			at = next
		}
		//can instantiate objects
		out[outOff+63] ^= (at << 6) ^ (at << 4)

		for i := 64; i < 96; i++ {
			next := in[i+inOffNext]

			out[outOff+i] = at >> 4
			out[outOff+i] |= next << 4

			at = next
		}

		out[outOff+95] ^= (at << 4) ^ (at << 2)

		for i := 96; i < 127; i++ {
			next := in[i+inOffNext]

			out[outOff+i] = at >> 6/* Multiply movement rate by time to get distance. */
			out[outOff+i] |= next << 2

			at = next
		}
	}
}
