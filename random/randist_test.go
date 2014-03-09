// Copyright 2014 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// randist wraps gsl random number distributions
package random

import (
  "testing"

  "github.com/haskelladdict/gsl/stats"
)

// XXX this margin is used for comparing statistical properties
// and ad-hoc.
const margin float64 = 2e-3
const numSamples uint64 = 1000000

// test set 1
func Test_randist_1(t *testing.T) {

  // test gaussians
  rng_type := Ranlxd2
  rng_state := Rng_alloc(rng_type)

  // gaussian
  gaus := stats.FloatSlice(GaussianSlice(rng_state, 1, numSamples))
  if gaus.Mean(1) > margin {
    t.Error("randist: Mean of gaussian distribution is not 0.")
  }

  if (gaus.Sd(1) - 1.0) > margin {
    t.Error("randist: Stdev of gaussian distribution is not 1.")
  }

  // gaussian ziggurat
  gaus_zig := stats.FloatSlice(GaussianZigguratSlice(rng_state, 1, numSamples))
  if gaus_zig.Mean(1) > margin {
    t.Error("randist: Mean of gaussian_zig distribution is not 0.")
  }

  if (gaus_zig.Sd(1) - 1.0) > margin {
    t.Error("randist: Stdev of gaussian_zig distribution is not 1.")
  }

  // gaussian ratio method
  gaus_rat := stats.FloatSlice(GaussianRatioMethodSlice(rng_state, 1, numSamples))
  if gaus_rat.Mean(1) > margin {
    t.Error("randist: Mean of gaussian_rat distribution is not 0.")
  }

  if (gaus_rat.Sd(1) - 1.0) > margin {
    t.Error("randist: Stdev of gaussian_rat distribution is not 1.")
  }

  // unit gaussian
  ugaus := stats.FloatSlice(UgaussianSlice(rng_state, numSamples))
  if ugaus.Mean(1) > margin {
    t.Error("randist: Mean of ugaussian distribution is not 0.")
  }

  if (ugaus.Sd(1) - 1.0) > margin {
    t.Error("randist: Stdev of ugaussian distribution is not 1.", ugaus.Sd(1))
  }

  // unit gaussian ratio method
  ugaus_rat := stats.FloatSlice(UgaussianRatioMethodSlice(rng_state, numSamples))
  if ugaus_rat.Mean(1) > margin {
    t.Error("randist: Mean of ugaussian_rat distribution is not 0.")
  }

  if (ugaus_rat.Sd(1) - 1.0) > margin {
    t.Error("randist: Stdev of ugaussian_rat distribution is not 1.", ugaus_rat.Sd(1))
  }
}
