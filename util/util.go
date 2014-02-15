// Copyright 2014 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Package util contains helper function mostly used for testing
//
package util

import (
  "math"
)

// float_equal compares two float numbers for equality
// NOTE: the floating point comparison is based on an epsilon
//       which was chosen empirically so it's not rigorous
func FloatEqual(a1, a2 float64) bool {
  epsilon := 1e-13
  if math.Abs(a2-a1) > epsilon*math.Abs(a1) {
    return false
  }
  return true
}
