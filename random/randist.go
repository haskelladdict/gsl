// Copyright 2014 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// randist wraps gsl random number distributions
//
// XXX: All of the slice functions could be made more efficient
// by calling the underlying GSL routines directly rather than
// calling the go wrappers.
package random

// #cgo pkg-config: gsl
// #include <gsl/gsl_randist.h>
import "C"

// pair encapsulates an array of two doubles
type Pair [2]float64

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

// Custom functions for unit Gaussian distribution

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

// GaussianTail provides random variates from the upper tail of a Gaussian
// distribution with standard deviation sigma. The values returned are
// larger than the lower limit a, which must be positive. The method is
// based on Marsaglia's famous rectangle-wedge-tail algorithm (Ann. Math.
// Stat. 32, 894­899 (1961)), with this aspect explained in Knuth, v2,
// 3rd ed, p139,586 (exercise 11).
func GaussianTail(rng RngState, a, sigma float64) float64 {
	return float64(C.gsl_ran_gaussian_tail(rng.state, C.double(a), C.double(sigma)))
}

// GaussianTailSlice returns a slice of length n from a Gaussian tail
// distribution.
func GaussianTailSlice(rng RngState, a, sigma float64, n uint64) []float64 {
	data := make([]float64, n)
	for i := uint64(0); i < n; i++ {
		data[i] = GaussianTail(rng, a, sigma)
	}
	return data
}

// GaussianTailPdf computes the probability density p(x) at x for a
// Gaussian tail distribution with standard deviation sigma and lower
// limit a.
func GaussianTailPdf(x, a, sigma float64) float64 {
	return float64(C.gsl_ran_gaussian_tail_pdf(C.double(x), C.double(a),
		C.double(sigma)))
}

// UgaussianTail provides random variates from a unit Gaussian tail
// distribution
func UgaussianTail(rng RngState, a float64) float64 {
	return float64(C.gsl_ran_ugaussian_tail(rng.state, C.double(a)))
}

// UgaussianTailSlice returns a slice of length n from a unit Gaussian tail
// distribution.
func UgaussianTailSlice(rng RngState, a float64, n uint64) []float64 {
	data := make([]float64, n)
	for i := uint64(0); i < n; i++ {
		data[i] = UgaussianTail(rng, a)
	}
	return data
}

// UgaussianTailPdf computes the probability density p(x) at x for a
// unit Gaussian tail distribution with lower limit a.
func UgaussianTailPdf(x, a float64) float64 {
	return float64(C.gsl_ran_ugaussian_tail_pdf(C.double(x), C.double(a)))
}

// BivariateGaussian generates a pair of correlated Gaussian variates,
// with mean zero, correlation coefficient rho and standard deviations
// sigma x and sigma y in the x and y directions. The correlation
// coefficient rho should lie between 1 and -1.
func BivariateGaussian(rng RngState, sigma_x, sigma_y, rho float64) (float64,
	float64) {
	var x, y float64
	C.gsl_ran_bivariate_gaussian(rng.state, C.double(sigma_x),
		C.double(sigma_y), C.double(rho), (*C.double)(&x), (*C.double)(&y))
	return x, y
}

// BivariateGaussianSlice generates a slice of length n with pairs of
// correlated Gaussian variates,
func BivariateGaussianSlice(rng RngState, sigma_x, sigma_y, rho float64,
	n uint64) []Pair {
	data := make([]Pair, n)
	for i := uint64(0); i < n; i++ {
		x, y := BivariateGaussian(rng, sigma_x, sigma_y, rho)
		data[i] = Pair{x, y}
	}
	return data
}

// BivariateGaussianPdf computes the probability density p(x,y) at (x,y)
// for a bivariate Gaussian distribution with standard deviations sigma x,
// sigma y and correlation coefficient rho.
func BivariateGaussianPdf(x, y, sigma_x, sigma_y, rho float64) float64 {
	return float64(C.gsl_ran_bivariate_gaussian_pdf(C.double(x), C.double(y),
		C.double(sigma_x), C.double(sigma_y), C.double(rho)))
}

// Exponential returns a random variate from the exponential distribution with mean mu.
func Exponential(rng RngState, mu float64) float64 {
	return float64(C.gsl_ran_exponential(rng.state, C.double(mu)))
}

// ExponentialSlice generates a slice of length n of exponentially distributed values
func ExponentialSlice(rng RngState, mu float64, n uint64) []float64 {
	data := make([]float64, n)
	for i := uint64(0); i < n; i++ {
		data[i] = Exponential(rng, mu)
	}
	return data
}

// ExponentialPdf computes the probability density p(x) at x for an exponential
// distribution with mean mu.
func ExponentialPdf(x, mu float64) float64 {
	return float64(C.gsl_ran_exponential_pdf(C.double(x), C.double(mu)))
}

// Laplace returns a random variate from the exponential distribution with width a.
func Laplace(rng RngState, a float64) float64 {
	return float64(C.gsl_ran_laplace(rng.state, C.double(a)))
}

// LaplaceSlice generates a slice of length n of laplace distributed values
func LaplaceSlice(rng RngState, a float64, n uint64) []float64 {
	data := make([]float64, n)
	for i := uint64(0); i < n; i++ {
		data[i] = Laplace(rng, a)
	}
	return data
}

// LaplacePdf computes the probability density p(x) at x for a laplace
// distribution with width a.
func LaplacePdf(x, a float64) float64 {
	return float64(C.gsl_ran_laplace_pdf(C.double(x), C.double(a)))
}

// Exppow returns a random variate from the exponential power distribution with scale
// parameter a and exponent b.
func Exppow(rng RngState, a, b float64) float64 {
	return float64(C.gsl_ran_exppow(rng.state, C.double(a), C.double(b)))
}

// ExppowSlice generates a slice of length n of exponential power distributes values
func ExppowSlice(rng RngState, a, b float64, n uint64) []float64 {
	data := make([]float64, n)
	for i := uint64(0); i < n; i++ {
		data[i] = Exppow(rng, a, b)
	}
	return data
}

// ExppowPdf computes the probability density p(x) at x for a exponential power
// distribution with scale parameter a and exponent b
func ExppowPdf(x, a, b float64) float64 {
	return float64(C.gsl_ran_exppow_pdf(C.double(x), C.double(a), C.double(b)))
}

// Cauchy returns a random variate from the Cauchy distribution with scale parameter a.
func Cauchy(rng RngState, a float64) float64 {
	return float64(C.gsl_ran_cauchy(rng.state, C.double(a)))
}

// CauchySlice generates a slice of length n of cauchy distributes values
func CauchySlice(rng RngState, a float64, n uint64) []float64 {
	data := make([]float64, n)
	for i := uint64(0); i < n; i++ {
		data[i] = Cauchy(rng, a)
	}
	return data
}

// CauchyPdf computes the probability density p(x) at x for a cauchy
// distribution with scale parameter a
func CauchyPdf(x, a float64) float64 {
	return float64(C.gsl_ran_cauchy_pdf(C.double(x), C.double(a)))
}

// Rayleigh returns a random variate from the Rayleigh distribution with scale
// parameter sigma.
func Rayleigh(rng RngState, sigma float64) float64 {
	return float64(C.gsl_ran_rayleigh(rng.state, C.double(sigma)))
}

// RayleighSlice generates a slice of length n of rayleigh distributed values
func RayleighSlice(rng RngState, sigma float64, n uint64) []float64 {
	data := make([]float64, n)
	for i := uint64(0); i < n; i++ {
		data[i] = Rayleigh(rng, sigma)
	}
	return data
}

// RayleighPdf computes the probability density p(x) at x for a rayleigh
// distribution with scale parameter sigma
func RayleighPdf(x, sigma float64) float64 {
	return float64(C.gsl_ran_rayleigh_pdf(C.double(x), C.double(sigma)))
}

// RayleighTail returns a random variate from the tail of the Rayleigh distribution
// with scale parameter sigma and a lower limit of a.
func RayleighTail(rng RngState, a, sigma float64) float64 {
	return float64(C.gsl_ran_rayleigh_tail(rng.state, C.double(a), C.double(sigma)))
}

// RayleighTailSlice generates a slice of length n of rayleigh tail distributed values
func RayleighTailSlice(rng RngState, a, sigma float64, n uint64) []float64 {
	data := make([]float64, n)
	for i := uint64(0); i < n; i++ {
		data[i] = RayleighTail(rng, a, sigma)
	}
	return data
}

// RayleighTailPdf computes the probability density p(x) at x for a rayleigh
// tail distribution with scale parameter sigma and a lower limit of a
func RayleighTailPdf(x, a, sigma float64) float64 {
	return float64(C.gsl_ran_rayleigh_tail_pdf(C.double(x), C.double(a),
		C.double(sigma)))
}

// Landay returns a random variate from the Landau distribution.
func Landau(rng RngState) float64 {
	return float64(C.gsl_ran_landau(rng.state))
}

// LandauSlice generates a slice of length n of Landau distributed values
func LandauSlice(rng RngState, n uint64) []float64 {
	data := make([]float64, n)
	for i := uint64(0); i < n; i++ {
		data[i] = Landau(rng)
	}
	return data
}

// LandayPdf computes the probability density p(x) at x for a Landau
// distribution.
func LandauPdf(x float64) float64 {
	return float64(C.gsl_ran_landau_pdf(C.double(x)))
}

// Levy returns a random variate from the Levy symmetric stable distribution
// with scale c and exponent alpha.
// XXX: The algorithm only works for 0 < alpha <= 2
func Levy(rng RngState, c, alpha float64) float64 {
	return float64(C.gsl_ran_levy(rng.state, C.double(c), C.double(alpha)))
}

// LevySlice generates a slice of length n of Levy distributed values
func LevySlice(rng RngState, c, alpha float64, n uint64) []float64 {
	data := make([]float64, n)
	for i := uint64(0); i < n; i++ {
		data[i] = Levy(rng, c, alpha)
	}
	return data
}

// LevySkew returns a random variate from the Levy skew stable distribution with
// scale c, exponent alpha and skewness parameter beta. The skewness parameter must
// lie in the range [−1, 1].
// XXX: The algorithm only works for 0 < alpha <= 2
func LevySkew(rng RngState, c, alpha, beta float64) float64 {
	return float64(C.gsl_ran_levy_skew(rng.state, C.double(c), C.double(alpha),
		C.double(beta)))
}

// LevySkewSlice generates a slice of length n of Levy skew distributed values
func LevySkewSlice(rng RngState, c, alpha, beta float64, n uint64) []float64 {
	data := make([]float64, n)
	for i := uint64(0); i < n; i++ {
		data[i] = LevySkew(rng, c, alpha, beta)
	}
	return data
}

// Gamma returns a random variate from the gamma distribution.
// The gamma distribution with an integer parameter a is known as the Erlang
// distribution. The variates are computed using the Marsaglia-Tsang fast gamma method.
func Gamma(rng RngState, a, b float64) float64 {
	return float64(C.gsl_ran_gamma(rng.state, C.double(a), C.double(b)))
}

// GammaKnuth returns a random variate from the gamma distribution using the
// algorithms from Knuth.
func GammaKnuth(rng RngState, a, b float64) float64 {
	return float64(C.gsl_ran_gamma_knuth(rng.state, C.double(a), C.double(b)))
}

// GammaSlice generates a slice of length n of gamma distributed values
func GammaSlice(rng RngState, a, b float64, n uint64) []float64 {
	data := make([]float64, n)
	for i := uint64(0); i < n; i++ {
		data[i] = Gamma(rng, a, b)
	}
	return data
}

// GammaPdf computes the probability density p(x) at x for a gamma distribution.
func GammaPdf(x, a, b float64) float64 {
	return float64(C.gsl_ran_gamma_pdf(C.double(x), C.double(a), C.double(b)))
}

// Flat returns a random variate from the flat (uniform) distribution from a to b.
func Flat(rng RngState, a, b float64) float64 {
	return float64(C.gsl_ran_flat(rng.state, C.double(a), C.double(b)))
}

// FlatSlice generates a slice of length n of flat distributed values
func FlatSlice(rng RngState, a, b float64, n uint64) []float64 {
	data := make([]float64, n)
	for i := uint64(0); i < n; i++ {
		data[i] = Flat(rng, a, b)
	}
	return data
}

// FlatPdf computes the probability density p(x) at x for a flat distribution.
func FlatPdf(x, a, b float64) float64 {
	return float64(C.gsl_ran_flat_pdf(C.double(x), C.double(a), C.double(b)))
}

// Lognormal returns a random variate from the lognormal distribution
func Lognormal(rng RngState, zeta, sigma float64) float64 {
	return float64(C.gsl_ran_lognormal(rng.state, C.double(zeta), C.double(sigma)))
}

// LognormalSlice generates a slice of length n of lognormal distributed values
func LognormalSlice(rng RngState, zeta, sigma float64, n uint64) []float64 {
	data := make([]float64, n)
	for i := uint64(0); i < n; i++ {
		data[i] = Lognormal(rng, zeta, sigma)
	}
	return data
}

// LognormalPdf computes the probability density p(x) at x for a lognormal distribution.
func LognormalPdf(x, zeta, sigma float64) float64 {
	return float64(C.gsl_ran_lognormal_pdf(C.double(x), C.double(zeta), C.double(sigma)))
}
