// Copyright 2015 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// random wraps gsl random number generation routines
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


// GaussianSlice returns a slice on length N with Gaussian random 
// variates, with mean zero and standard deviation sigma.
func GaussianSlice(rng RngState, sigma float64, n uint64) []float64 {
  data := make([]float64, n)
  for i := uint64(0); i < n; i++ {
    data[i] = Gaussian(rng, sigma)
  }
  return data
}


// GaussianPDF computes the probability density p(x) at x for a 
// Gaussian distribution with standard deviation sigma, using the formula 
// given above.
func GaussianPdf(x float64, sigma float64) float64 {
  return float64(C.gsl_ran_gaussian_pdf(C.double(x), C.double(sigma)))
}


// GaussianZiggurat returns a Gaussian random variate, with mean zero and 
// standard deviation sigma computed via the Marsaglia-Zang ziggurat m
// method.
func GaussianZiggurat(rng RngState, sigma float64) float64 {
  return float64(C.gsl_ran_gaussian_ziggurat(rng.state, C.double(sigma)))
}


// GaussianSlice returns a slice on length N with Gaussian random 
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
// methods.
func GaussianRatioMethod(rng RngState, sigma float64) float64 {
  return float64(C.gsl_ran_gaussian_ratio_method(rng.state, C.double(sigma)))
}


// GaussianSlice returns a slice on length N with Gaussian random 
// variates, with mean zero and standard deviation sigma computed via the 
// Kinderman-Monahan-Leva ratio methods.
func GaussianRatioMethodSlice(rng RngState, sigma float64, n uint64) []float64 {
  data := make([]float64, n)
  for i := uint64(0); i < n; i++ {
    data[i] = GaussianRatioMethod(rng, sigma)
  }
  return data
}


