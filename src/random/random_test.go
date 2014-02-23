// Copyright 2014 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// random wraps gsl random number generation routines
package random

import (
  "github.com/haskelladdict/go-gsl/src/stats"
  "testing"
)

// test set 1
func Test_random_1(t *testing.T) {

  // test 1
  var rng_type RngType
  for k, v := range TypesSetup() {
    if k == "ranlxd2" {
      rng_type = v
    }
  }

  if rng_type.Name() != "ranlxd2" {
    t.Error("Test 1: Failed to select ranlxd2 rng type.")
  }
  rng_state := Alloc(rng_type)

  // test 2
  nums1 := rng_state.GetSlice(100)
  nums1_slice := make(stats.FloatSlice, 100)
  for i := 0; i < 100; i++ {
    nums1_slice[i] = float64(nums1[i])
  }
  min, max := nums1_slice.MinMax(1)

  if min < float64(rng_state.Min()) || max > float64(rng_state.Max()) {
    t.Error("Test 2: Generated random numbers are out of range.")
  }

  // test 3
  nums2 := stats.FloatSlice(rng_state.UniformSlice(100))
  min, max = nums2.MinMax(1)

  if min < 0 || max >= 1 {
    t.Error("Test 3: Generated uniform random numbers are out of range.")
  }

  // test 4
  nums3 := rng_state.UniformIntSlice(900, 100)
  nums3_slice := make(stats.FloatSlice, 100)
  for i := 0; i < 100; i++ {
    nums3_slice[i] = float64(nums3[i])
  }
  min, max = nums3_slice.MinMax(1)

  if min < 0 || max >= 900 {
    t.Error("Test 4: Generated uniform int random numbers are out of range.")
  }

  // test 5
  rng_state_1 := rng_state.Clone()
  if rng_state_1.Get() != rng_state.Get() {
    t.Error("Test 5: Failed to clone rng state.")
  }

  // test 6
  rng_state_2 := Alloc(rng_type)
  if rng_state_2.Get() == rng_state.Get() {
    t.Error("Test 6: Hmmm, this is very unlikely to happen.")
  }

  rng_state.Memcpy(rng_state_2)
  if rng_state_2.Get() != rng_state.Get() {
    t.Error("Test 6: Failed to memcpy rng state.")
  }
}
