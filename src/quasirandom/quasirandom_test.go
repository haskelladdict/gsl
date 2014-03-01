// Copyright 2014 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// random wraps gsl random number generation routines
package quasirandom

import (
  "fmt"
  "testing"

  "github.com/haskelladdict/go-gsl/src/stats"
)

// test set 1
func Test_random_1(t *testing.T) {

  // test sobol
  test_helper(t, Sobol)

  // test Niederreiter_2
  test_helper(t, Niederreiter_2)

  // test halton
  test_helper(t, Halton)

  // test reverse Halton
  test_helper(t, ReverseHalton)
}

// brief helper function
func test_helper(t *testing.T, rng_type QrngType) {

  numRands := 1000
  rng_state := Alloc(rng_type, 2)
  points1 := rng_state.GetSlice(numRands)

  if len(points1[0]) != 2 {
    msg := fmt.Sprintf("%s: Incorrect length of generated QrngPoint. "+
      "Expected 2 got %d", rng_state.Name(), len(points1[0]))
    t.Error(msg)
  }

  // test 2
  points1_slice := make(stats.FloatSlice, numRands)
  points2_slice := make(stats.FloatSlice, numRands)
  for i := 0; i < 100; i++ {
    points1_slice[i] = float64(points1[i][0])
    points2_slice[i] = float64(points1[i][1])
  }
  min1, max1 := points1_slice.MinMax(1)
  min2, max2 := points2_slice.MinMax(1)

  if min1 < 0 || max1 >= 1 {
    t.Error(rng_state.Name() + "Generated quasirandom numbers are out of range.")
  }

  if min2 < 0 || max2 >= 1 {
    t.Error(rng_state.Name() + "Generated quasirandom numbers are out of range.")
  }
  rng_state.Free()
}
