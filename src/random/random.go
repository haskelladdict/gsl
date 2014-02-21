// Copyright 2014 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// random wraps gsl random number generation routines
package random

// #cgo pkg-config: gsl
// #include <gsl/gsl_rng.h>
import "C"
//import (
//  "unsafe"
//)


// type definitions
//type FloatSlice []float64

// Mean returns the arithmetic mean of data with stride stride.
//func (d FloatSlice) Mean(stride int) float64 {
//  mean := C.gsl_stats_mean((*C.double)(&d[0]), C.size_t(stride),
//    C.size_t(len(d)))
//  return float64(mean)
//}

type RngState struct {
  state *C.gsl_rng
}

type RngType *C.gsl_rng_type


func Default() RngType {
  return C.gsl_rng_default
}

func Alloc(rngType RngType) RngState {
  state := C.gsl_rng_alloc(C.gsl_rng_taus)
  return RngState{state}
}

func (s RngState) Name() string {
  return C.GoString(C.gsl_rng_name(s.state))
}

func (s RngState) Uniform() float64 {
  return float64(C.gsl_rng_uniform(s.state))
}
