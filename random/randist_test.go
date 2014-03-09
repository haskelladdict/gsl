// Copyright 2014 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// randist wraps gsl random number distributions
package random

import (
  "testing"

  "github.com/haskelladdict/gsl/stats"
  "github.com/haskelladdict/gsl/util"
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

// test set 2
// XXX: This a are silly tests that only check that the function work
// but not that they work correctly. Not sure how to check correctness.
func Test_randist_2(t *testing.T) {

  // test gaussian tail distributions
  rng_type := Ranlxd2
  rng_state := Rng_alloc(rng_type)

  // gaussian
  tail := GaussianTailSlice(rng_state, 0, 1, numSamples)
  if len(tail) != int(numSamples) {
    t.Error("randist: error computing gaussian tail distribution")
  }

  // unit gaussian
  utail := UgaussianTailSlice(rng_state, 0, numSamples)
  if len(utail) != int(numSamples) {
    t.Error("randist: error computing unit gaussian tail distribution")
  }

  // pdfs
  g_pdf := GaussianTailPdf(1, 0, 1)
  ug_pdf := UgaussianTailPdf(1, 0)
  if !util.FloatEqual(g_pdf, ug_pdf) {
    t.Error("randist: missmatch in computing gaussian tail pdfs",
      g_pdf, ug_pdf)
  }
}

// test set 3
func Test_randist_3(t *testing.T) {

  // test bivariate gaussian distributions
  rng_type := Ranlxd2
  rng_state := Rng_alloc(rng_type)

  // bivariate gaussian
  tail := BivariateGaussianSlice(rng_state, 1, 1, 1, numSamples)

  first := stats.FloatSlice(make([]float64, numSamples))
  second := stats.FloatSlice(make([]float64, numSamples))
  for i, v := range tail {
    first[i], second[i] = v[0], v[1]
  }

  if first.Mean(1) > margin {
    t.Error("randist: Mean of 1st component of bivariate gaussian is not 0.")
  }

  if second.Mean(1) > margin {
    t.Error("randist: Mean of 2nd component of bivariate gaussian is not 0.")
  }

  if (first.Sd(1) - 1) > margin {
    t.Error("randist: Std of 1st component of bivariate gaussian is not 0.")
  }

  if (second.Sd(1) - 1) > margin {
    t.Error("randist: Std of 2nd component of bivariate gaussian is not 0.")
  }
}
