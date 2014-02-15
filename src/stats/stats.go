// Copyright 2014 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// stat wraps gsl statistics routines
//
// XXX: None of the statistics function check for the lenght of the
//      data string and will trigger a runtime exception if called
//      with an empty slice.
package stats

// #cgo pkg-config: gsl
// #include <gsl/gsl_statistics.h>
import "C"

// type definitions
type FloatSlice []float64

// Mean returns the arithmetic mean of data with stride stride.
func (d FloatSlice) Mean(stride int) float64 {
  mean := C.gsl_stats_mean((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)))
  return float64(mean)
}

// Variance returns the variance of data with stride stride.
func (d FloatSlice) Variance(stride int) float64 {
  variance := C.gsl_stats_variance((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)))
  return float64(variance)
}

// Variance_m returns the variance of data with stride stride and a user
// supplied value for the mean
func (d FloatSlice) Variance_m(stride int, mean float64) float64 {
  variance_m := C.gsl_stats_variance_m((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)), C.double(mean))
  return float64(variance_m)
}

// Sd returns the standard deviation of data with stride stride.
// NOTE: The value of this function equals the square root of the variance
func (d FloatSlice) Sd(stride int) float64 {
  sd := C.gsl_stats_sd((*C.double)(&d[0]), C.size_t(stride), C.size_t(len(d)))
  return float64(sd)
}

// Sd_m returns the standard deviation of data with stride stride and a user
// supplied value for the mean.
// NOTE: The value of this function equals the square root of the variance
func (d FloatSlice) Sd_m(stride int, mean float64) float64 {
  sd_m := C.gsl_stats_sd_m((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)), C.double(mean))
  return float64(sd_m)
}

// Tss returns the total sum of squares (TSS) of data about the mean with
// stride stride.
func (d FloatSlice) Tss(stride int) float64 {
  tss := C.gsl_stats_tss((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)))
  return float64(tss)
}

// Tss_m returns the total sum of squares (TSS) of data about the mean with
// stride stride and a user supplied value for the mean.
func (d FloatSlice) Tss_m(stride int, mean float64) float64 {
  tss_m := C.gsl_stats_tss_m((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)), C.double(mean))
  return float64(tss_m)
}

// Variance_with_fixed_mean computes an unbiased estimate of the variance
// of data when the population mean mean of the underlying distribution is
// known a priori.
func (d FloatSlice) Variance_with_fixed_mean(stride int, mean float64) float64 {
  var_fixed := C.gsl_stats_variance_with_fixed_mean((*C.double)(&d[0]),
    C.size_t(stride), C.size_t(len(d)), C.double(mean))
  return float64(var_fixed)
}

// Sd_with_fixed_mean computes an unbiased estimate of the standard deviation
// of data when the population mean mean of the underlying distribution is
// known a priori.
// NOTE: This is equal to the square root of the Variance_with_fixed_mean
func (d FloatSlice) Sd_with_fixed_mean(stride int, mean float64) float64 {
  var_fixed := C.gsl_stats_sd_with_fixed_mean((*C.double)(&d[0]),
    C.size_t(stride), C.size_t(len(d)), C.double(mean))
  return float64(var_fixed)
}


// Absdev computes the absolute deviation from the mean of data with 
// stride stride
func (d FloatSlice) Absdev(stride int) float64 {
  absdev := C.gsl_stats_absdev((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)))
  return float64(absdev)
}


// Absdev_m computes the absolute deviation from the mean of data with 
// stride stride and given mean
func (d FloatSlice) Absdev_m(stride int, mean float64) float64 {
  absdev_m := C.gsl_stats_absdev_m((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)), C.double(mean))
  return float64(absdev_m)
}


// Skew computes the skewness of data with stride stride.
func (d FloatSlice) Skew(stride int) float64 {
  skew := C.gsl_stats_skew((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)))
  return float64(skew)
}


// Skew_m_sd computes the skewness of the dataset data using stride stride
// and the given values of the mean mean and standard deviation sd.
func (d FloatSlice) Skew_m_sd(stride int, mean float64, sd float64) float64 {
  skew_m_sd := C.gsl_stats_skew_m_sd((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)), C.double(mean), C.double(sd))
  return float64(skew_m_sd)
}


// Kurtosis computes the kurtosis of data with stride stride.
func (d FloatSlice) Kurtosis(stride int) float64 {
  kurt := C.gsl_stats_kurtosis((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)))
  return float64(kurt)
}


// Kurtosis computes the kurtosis of data with stride stride.
func (d FloatSlice) Kurtosis_m_sd(stride int, mean float64, sd float64) float64 {
  kurt_m_sd := C.gsl_stats_kurtosis_m_sd((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)), C.double(mean), C.double(sd))
  return float64(kurt_m_sd)
}


