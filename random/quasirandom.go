// Copyright 2014 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// random wraps gsl quasirandom number generation routines
package random

// #cgo pkg-config: gsl
// #include <gsl/gsl_qrng.h>
import "C"

// QrngState stores the quasi random number generator state
type QrngState struct {
  state *C.gsl_qrng
  dim   uint
}

// QrngType stores the type of qrng method used
type QrngType struct {
  qrng *C.gsl_qrng_type
}

// QrngPoint holds a single point of dimension d produced
// by the Qrng
type QrngPoint []float64

// list of available quasi random number generators. See gsl documentation
// for more detailed info on each of these.
var (
  Niederreiter_2 = QrngType{C.gsl_qrng_niederreiter_2}
  Sobol          = QrngType{C.gsl_qrng_sobol}
  Halton         = QrngType{C.gsl_qrng_halton}
  ReverseHalton  = QrngType{C.gsl_qrng_reversehalton}
)

// RNG initialization

// Qrng_Alloc creates a new quasirandom number generator of the
// requested type and dimension and returns it as a QrngState
// object.
// XXX: using SetFinalizer to release the generator doesn't work
// properly since the go runtime destroys the object prematurely.
func Qrng_alloc(qrngType QrngType, dim uint) QrngState {
  state := QrngState{C.gsl_qrng_alloc(qrngType.qrng, C.uint(dim)), dim}

  // make sure we get rid of any memory associated with the
  // rng within gsl
  //runtime.SetFinalizer(&state,
  //  func(qrng *QrngState) {C.gsl_qrng_free(qrng.state)})
  return state
}

// Free releases all the memory associated with the generator
// within the C part of gsl
func (s *QrngState) Free() {
  C.gsl_qrng_free(s.state)
  s.state = nil // to make sure we don't use after freeing
}

// Init reinitializes the generator q to its starting point. Note
// that quasirandom sequences do not use a seed and always produce the
// same set of values.
func (s *QrngState) Init() {
  C.gsl_qrng_init(s.state)
}

// RNG sampling functions

// Get returns the next point from the sequence generator s. The
// length of QrngPoint p matches the dimension of the underlying
// QrngState. Each element of QrngPoint p will lie in the range
// 0 < p_i < 1 for each p_i .
// XXX: The gsl manual does not say what the return value of
// gsl_qrng_get signifies so we ignore it for now.
func (s *QrngState) Get() QrngPoint {
  point := make(QrngPoint, s.dim)
  C.gsl_qrng_get(s.state, (*C.double)(&point[0]))
  return point
}

// GetSlice is a convenience function and returns a slice of length
// n of QrngPoints
func (s *QrngState) GetSlice(n uint64) []QrngPoint {
  slice := make([]QrngPoint, n)
  for i := uint64(0); i < n; i++ {
    slice[i] = s.Get()
  }
  return slice
}

// RNG auxiliary functions

// Name returns the name of the quasirandom number generator
func (s *QrngState) Name() string {
  return C.GoString(C.gsl_qrng_name(s.state))
}

// String provides a printable string representation for
// an QrngState
func (s *QrngState) String() string {
  return s.Name()
}

// Size returns the size of the generator
func (s *QrngState) Size() uint64 {
  return uint64(C.gsl_qrng_size(s.state))
}

// State returns a pointer to the underlying qrng state from gsl
func (s *QrngState) State() StatePointer {
  return StatePointer(C.gsl_qrng_state(s.state))
}

// Copying, cloning, writing and reading rng state

// Memcpy copies the quasi random number generator src into the
// pre-existing generator dest, making dest into an exact copy of src.
// The two generators must be of the same type.
// XXX: currently this ignores the return type of gsl_rng_memcpy since
// I don't know what it does (the manual is quiet on that)
func (s *QrngState) Memcpy(dest *QrngState) {
  C.gsl_qrng_memcpy(dest.state, s.state)
}

// Clone returns a newly created generator which is an exact copy
// of the generator r.
func (s *QrngState) Clone() QrngState {
  return QrngState{C.gsl_qrng_clone(s.state), s.dim}
}
