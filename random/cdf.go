// Copyright 2014 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// cdf wraps gsl cumulative random distribution functions
//
// XXX: All of the slice functions could be made more efficient
// by calling the underlying GSL routines directly rather than
// calling the go wrappers.
package random

// #cgo pkg-config: gsl
// #include <gsl/gsl_cdf.h>
import "C"

// GaussianP returns the cumulative distribution function P(x) for
// the lower tail of a Gaussian.
func GaussianP(x float64, sigma float64) float64 {
	return float64(C.gsl_cdf_gaussian_P(C.double(x), C.double(sigma)))
}

// GaussianQ returns the cumulative distribution function Q(x) for
// the upper tail of a Gaussian.
func GaussianQ(x float64, sigma float64) float64 {
	return float64(C.gsl_cdf_gaussian_Q(C.double(x), C.double(sigma)))
}

// GaussianPinv returns the inverse cumulative distribution function
// Pinv(x) for the lower tail of a Gaussian.
func GaussianPinv(P float64, sigma float64) float64 {
	return float64(C.gsl_cdf_gaussian_Pinv(C.double(P), C.double(sigma)))
}

// GaussianQinv returns the inverse cumulative distribution function
// Qinv(x) for the ujjpper tail of a Gaussian.
func GaussianQinv(Q float64, sigma float64) float64 {
	return float64(C.gsl_cdf_gaussian_Qinv(C.double(Q), C.double(sigma)))
}

// UgaussianP returns the cumulative distribution function P(x) for
// the lower tail of a unit Gaussian.
func UgaussianP(x float64) float64 {
	return float64(C.gsl_cdf_ugaussian_P(C.double(x)))
}

// UgaussianQ returns the cumulative distribution function Q(x) for
// the upper tail of a unit Gaussian.
func UgaussianQ(x float64) float64 {
	return float64(C.gsl_cdf_ugaussian_Q(C.double(x)))
}

// UgaussianPinv returns the inverse cumulative distribution function
// Pinv(x) for the lower tail of a unit Gaussian.
func UgaussianPinv(P float64) float64 {
	return float64(C.gsl_cdf_ugaussian_Pinv(C.double(P)))
}

// UgaussianQinv returns the inverse cumulative distribution function
// Qinv(x) for the upper tail of a unit Gaussian.
func UgaussianQinv(Q float64) float64 {
	return float64(C.gsl_cdf_ugaussian_Qinv(C.double(Q)))
}

// ExponentialP returns the cumulative distribution function P(x) for
// the lower tail of an exponential distribution.
func ExponentialP(x, mu float64) float64 {
	return float64(C.gsl_cdf_exponential_P(C.double(x), C.double(mu)))
}

// ExponentialQ returns the cumulative distribution function Q(x) for
// the upper tail of an exponential distribution.
func ExponentialQ(x, mu float64) float64 {
	return float64(C.gsl_cdf_exponential_Q(C.double(x), C.double(mu)))
}

// ExponentialPinv returns the inverse cumulative distribution function Pinv(x) for
// the lower tail of an exponential distribution.
func ExponentialPinv(P, mu float64) float64 {
	return float64(C.gsl_cdf_exponential_Pinv(C.double(P), C.double(mu)))
}

// ExponentialQinv returns the cumulative distribution function Qinv(x) for
// the upper tail of an exponential distribution.
func ExponentialQinv(Q, mu float64) float64 {
	return float64(C.gsl_cdf_exponential_Qinv(C.double(Q), C.double(mu)))
}

// LaplaceP returns the cumulative distribution function P(x) for
// the lower tail of a laplace distribution.
func LaplaceP(x, a float64) float64 {
	return float64(C.gsl_cdf_laplace_P(C.double(x), C.double(a)))
}

// LaplaceQ returns the cumulative distribution function Q(x) for
// the upper tail of a laplace distribution.
func LaplaceQ(x, a float64) float64 {
	return float64(C.gsl_cdf_laplace_Q(C.double(x), C.double(a)))
}

// LaplacePinv returns the inverse cumulative distribution function Pinv(x) for
// the lower tail of a laplace distribution.
func LaplacePinv(P, a float64) float64 {
	return float64(C.gsl_cdf_laplace_Pinv(C.double(P), C.double(a)))
}

// LaplaceQinv returns the cumulative distribution function Qinv(x) for
// the upper tail of a laplace distribution.
func LaplaceQinv(Q, a float64) float64 {
	return float64(C.gsl_cdf_laplace_Qinv(C.double(Q), C.double(a)))
}

// ExppowP returns the cumulative distribution function P(x) for
// the lower tail of a exponential power distribution.
func ExppowP(x, a, b float64) float64 {
	return float64(C.gsl_cdf_exppow_P(C.double(x), C.double(a), C.double(b)))
}

// ExppowQ returns the cumulative distribution function Q(x) for
// the upper tail of an exponential power distribution.
func ExppowQ(x, a, b float64) float64 {
	return float64(C.gsl_cdf_exppow_Q(C.double(x), C.double(a), C.double(b)))
}

// CauchyP returns the cumulative distribution function P(x) for
// the lower tail of a cauchy distribution.
func CauchyP(x, a float64) float64 {
	return float64(C.gsl_cdf_cauchy_P(C.double(x), C.double(a)))
}

// CauchyQ returns the cumulative distribution function Q(x) for
// the upper tail of a cauchy distribution.
func CauchyQ(x, a float64) float64 {
	return float64(C.gsl_cdf_cauchy_Q(C.double(x), C.double(a)))
}

// CauchyPinv returns the inverse cumulative distribution function Pinv(x) for
// the lower tail of a cauchy distribution.
func CauchyPinv(P, a float64) float64 {
	return float64(C.gsl_cdf_cauchy_Pinv(C.double(P), C.double(a)))
}

// CauchyQinv returns the cumulative distribution function Qinv(x) for
// the upper tail of a cauchy distribution.
func CauchyQinv(Q, a float64) float64 {
	return float64(C.gsl_cdf_cauchy_Qinv(C.double(Q), C.double(a)))
}

// RayleighP returns the cumulative distribution function P(x) for
// the lower tail of a rayleigh distribution.
func RayleighP(x, sigma float64) float64 {
	return float64(C.gsl_cdf_rayleigh_P(C.double(x), C.double(sigma)))
}

// RayleighQ returns the cumulative distribution function Q(x) for
// the upper tail of a rayleigh distribution.
func RayleighQ(x, sigma float64) float64 {
	return float64(C.gsl_cdf_rayleigh_Q(C.double(x), C.double(sigma)))
}

// RayleighPinv returns the inverse cumulative distribution function Pinv(x) for
// the lower tail of a rayleigh distribution.
func RayleighPinv(P, sigma float64) float64 {
	return float64(C.gsl_cdf_rayleigh_Pinv(C.double(P), C.double(sigma)))
}

// RayleighQinv returns the cumulative distribution function Qinv(x) for
// the upper tail of a rayleigh distribution.
func RayleighQinv(Q, sigma float64) float64 {
	return float64(C.gsl_cdf_rayleigh_Qinv(C.double(Q), C.double(sigma)))
}

// GammaP returns the cumulative distribution function P(x) for
// the lower tail of a gamma distribution.
func GammaP(x, a, b float64) float64 {
	return float64(C.gsl_cdf_gamma_P(C.double(x), C.double(a), C.double(b)))
}

// GammaQ returns the cumulative distribution function Q(x) for
// the upper tail of a gamma distribution.
func GammaQ(x, a, b float64) float64 {
	return float64(C.gsl_cdf_gamma_Q(C.double(x), C.double(a), C.double(b)))
}

// GammaPinv returns the inverse cumulative distribution function Pinv(x) for
// the lower tail of a gamma distribution.
func GammaPinv(P, a, b float64) float64 {
	return float64(C.gsl_cdf_gamma_Pinv(C.double(P), C.double(a), C.double(b)))
}

// GammaQinv returns the cumulative distribution function Qinv(x) for
// the upper tail of a gamma distribution.
func GammaQinv(Q, a, b float64) float64 {
	return float64(C.gsl_cdf_gamma_Qinv(C.double(Q), C.double(a), C.double(b)))
}

// FlatP returns the cumulative distribution function P(x) for
// the lower tail of a flat distribution.
func FlatP(x, a, b float64) float64 {
	return float64(C.gsl_cdf_flat_P(C.double(x), C.double(a), C.double(b)))
}

// FlatQ returns the cumulative distribution function Q(x) for
// the upper tail of a flat distribution.
func FlatQ(x, a, b float64) float64 {
	return float64(C.gsl_cdf_flat_Q(C.double(x), C.double(a), C.double(b)))
}

// FlatPinv returns the inverse cumulative distribution function Pinv(x) for
// the lower tail of a flat distribution.
func FlatPinv(P, a, b float64) float64 {
	return float64(C.gsl_cdf_flat_Pinv(C.double(P), C.double(a), C.double(b)))
}

// FlatQinv returns the cumulative distribution function Qinv(x) for
// the upper tail of a flat distribution.
func FlatQinv(Q, a, b float64) float64 {
	return float64(C.gsl_cdf_flat_Qinv(C.double(Q), C.double(a), C.double(b)))
}

// LognormalP returns the cumulative distribution function P(x) for
// the lower tail of a lognormal distribution.
func LognormalP(x, zeta, sigma float64) float64 {
	return float64(C.gsl_cdf_lognormal_P(C.double(x), C.double(zeta), C.double(sigma)))
}

// LognormalQ returns the cumulative distribution function Q(x) for
// the upper tail of a lognormal distribution.
func LognormalQ(x, zeta, sigma float64) float64 {
	return float64(C.gsl_cdf_lognormal_Q(C.double(x), C.double(zeta), C.double(sigma)))
}

// LognormalPinv returns the inverse cumulative distribution function Pinv(x) for
// the lower tail of a lognormal distribution.
func LognormalPinv(P, zeta, sigma float64) float64 {
	return float64(C.gsl_cdf_lognormal_Pinv(C.double(P), C.double(zeta), C.double(sigma)))
}

// LognormalQinv returns the cumulative distribution function Qinv(x) for
// the upper tail of a lognormal distribution.
func LognormalQinv(Q, zeta, sigma float64) float64 {
	return float64(C.gsl_cdf_lognormal_Qinv(C.double(Q), C.double(zeta), C.double(sigma)))
}
