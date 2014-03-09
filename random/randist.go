// Copyright 2015 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// randist wraps gsl random number distributions
//
// XXX: All of the slice functions could be made more efficient
// by calling the underlying GSL routines directly rathen than
// calling the go wrappers.
package random

// #cgo pkg-config: gsl
// #include <gsl/gsl_randist.h>
import "C"

// Gaussian returns a Gaussian random variate, with mean zero and
// standard deviation sigma.
func Gaussian(rng RngState, sigma float64) float64 {
  return float64(C.gsl_ran_gaussian(rng.state, C.double(sigma)))
}

// GaussianSlice returns a slice of length N with Gaussian random
// variates, with mean zero and standard deviation sigma.
func GaussianSlice(rng RngState, sigma float64, n uint64) []float64 {
  data := make([]float64, n)
  for i := uint64(0); i < n; i++ {
    data[i] = Gaussian(rng, sigma)
  }
  return data
}

// GaussianPpf computes the probability density p(x) at x for a
// Gaussian distribution with standard deviation sigma.
func GaussianPdf(x float64, sigma float64) float64 {
  return float64(C.gsl_ran_gaussian_pdf(C.double(x), C.double(sigma)))
}

// GaussianZiggurat returns a Gaussian random variate, with mean zero and
// standard deviation sigma computed via the Marsaglia-Zang ziggurat m
// method.
func GaussianZiggurat(rng RngState, sigma float64) float64 {
  return float64(C.gsl_ran_gaussian_ziggurat(rng.state, C.double(sigma)))
}

// GaussianSlice returns a slice of length N with Gaussian random
// variates, with mean zero and standard deviation sigma computed via
// the Marsaglia-Zang ziggurat method.
func GaussianZigguratSlice(rng RngState, sigma float64, n uint64) []float64 {
  data := make([]float64, n)
  for i := uint64(0); i < n; i++ {
    data[i] = GaussianZiggurat(rng, sigma)
  }
  return data
}

// GaussianRatioMethod returns a Gaussian random variate, with mean zero and
// standard deviation sigma computed via the Kinderman-Monahan-Leva ratio
// method.
func GaussianRatioMethod(rng RngState, sigma float64) float64 {
  return float64(C.gsl_ran_gaussian_ratio_method(rng.state, C.double(sigma)))
}

// GaussianSlice returns a slice of length N with Gaussian random
// variates, with mean zero and standard deviation sigma computed via the
// Kinderman-Monahan-Leva ratio method.
func GaussianRatioMethodSlice(rng RngState, sigma float64, n uint64) []float64 {
  data := make([]float64, n)
  for i := uint64(0); i < n; i++ {
    data[i] = GaussianRatioMethod(rng, sigma)
  }
  return data
}

// Custom functions for unit Gaussion distribution

// UGaussian returns a unit Gaussian random variate with mean zero.
func Ugaussian(rng RngState) float64 {
  return float64(C.gsl_ran_ugaussian(rng.state))
}

// UgaussianSlice returns a slice of length N with unit Gaussian random
// variates with mean zero.
func UgaussianSlice(rng RngState, n uint64) []float64 {
  data := make([]float64, n)
  for i := uint64(0); i < n; i++ {
    data[i] = Ugaussian(rng)
  }
  return data
}

// UGaussian returns a unit Gaussian random variate with mean zero
// computed with the Kinderman-Mohanan-Leva ratio method.
func UgaussianRatioMethod(rng RngState) float64 {
  return float64(C.gsl_ran_ugaussian_ratio_method(rng.state))
}

// UgaussianSlice returns a slice of length N with unit Gaussian random
// variates with mean zero computed with the Kinderman-Mohanan-Leva ratio
// method.
func UgaussianRatioMethodSlice(rng RngState, n uint64) []float64 {
  data := make([]float64, n)
  for i := uint64(0); i < n; i++ {
    data[i] = UgaussianRatioMethod(rng)
  }
  return data
}

// UgaussianPpf computes the probability density p(x) at x for a
// unit Gaussian distribution.
func UgaussianPdf(x float64) float64 {
  return float64(C.gsl_ran_ugaussian_pdf(C.double(x)))
}
