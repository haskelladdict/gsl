// Copyright 2014 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// stat wraps gsl statistics routines
package stats


// #cgo pkg-config: gsl
// #include <gsl/gsl_statistics.h>
import "C"


// type definitions
type FloatSlice []float64


// Mean returns the arithmetic mean of data with stride
func (f FloatSlice) Mean(stride int) float64 {
  mean := C.gsl_stats_mean((*C.double)(&f[0]), C.size_t(stride),
    C.size_t(len(f)))
  return float64(mean)
}


// Variance returns the variance of data with stride
func (f FloatSlice) Variance(stride int) float64 {
  variance := C.gsl_stats_variance((*C.double)(&f[0]), C.size_t(stride),
    C.size_t(len(f)))
  return float64(variance)
}


// Variance_m returns the variance of data with stride and a user
// supplied value for the mean
func (f FloatSlice) Variance_m(stride int, mean float64) float64 {
  variance_m := C.gsl_stats_variance_m((*C.double)(&f[0]), C.size_t(stride),
    C.size_t(len(f)), C.double(mean))
  return float64(variance_m)
}


