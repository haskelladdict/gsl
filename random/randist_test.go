// Copyright 2014 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// random wraps gsl random number generation routines
package random

import (
  "testing"

  "github.com/haskelladdict/gsl/stats"
)

// XXX this margin is used for comparing statistical properties
// and ad-hoc.
const margin float64 = 1e-3


// test set 1
func Test_randist_1(t *testing.T) {

  // test 1
  rng_type := Ranlxd2
  rng_state := Rng_alloc(rng_type)

  gaus := stats.FloatSlice(GaussianSlice(rng_state, 1, 1000000))
  if gaus.Mean(1) > margin {
    t.Error("Randist: Mean of gaussian distribution is not 0.")
  }

  if (gaus.Sd(1) - 1.0) > margin {
    t.Error("Randist: Stdev of gaussian distribution is not 0.")
  }

  gaus_zig := stats.FloatSlice(GaussianZigguratSlice(rng_state, 1, 1000000))
  if gaus_zig.Mean(1) > margin {
    t.Error("Randist: Mean of gaussian_zig distribution is not 0.")
  }

  if (gaus_zig.Sd(1) - 1.0) > margin {
    t.Error("Randist: Stdev of gaussian_zig distribution is not 0.")
  }

  gaus_rat := stats.FloatSlice(GaussianRatioMethodSlice(rng_state, 1, 1000000))
  if gaus_rat.Mean(1) > margin {
    t.Error("Randist: Mean of gaussian_zig distribution is not 0.")
  }

  if (gaus_rat.Sd(1) - 1.0) > margin {
    t.Error("Randist: Stdev of gaussian_zig distribution is not 0.")
  }



}

