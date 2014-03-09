// Copyright 2014 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// cdf wraps gsl cumulative random distribution functions
package random

import (
  "testing"

  "github.com/haskelladdict/gsl/util"
)

// test set 1
func Test_cdf_1(t *testing.T) {

  // gaussian
  if !util.FloatEqual(GaussianP(0, 1), 0.5) {
    t.Error("cdf: error computing P(0).")
  }

  if !util.FloatEqual(GaussianP(0, 10), 0.5) {
    t.Error("cdf: error computing P(0).")
  }

  if !util.FloatEqual(GaussianQ(0, 1), 0.5) {
    t.Error("cdf: error computing P(0).")
  }

  if !util.FloatEqual(GaussianQ(0, 10), 0.5) {
    t.Error("cdf: error computing P(0).")
  }

  if !util.FloatEqual(GaussianPinv(0.5, 1), 0) {
    t.Error("cdf: error computing Pinv(0.5).")
  }

  if !util.FloatEqual(GaussianPinv(0.5, 10), 0) {
    t.Error("cdf: error computing Pinv(0.5).")
  }

  if !util.FloatEqual(GaussianQinv(0.5, 1), 0) {
    t.Error("cdf: error computing Qinv(0.5).")
  }

  if !util.FloatEqual(GaussianQinv(0.5, 10), 0) {
    t.Error("cdf: error computing Qinv(0.5).")
  }

  // unit gaussian
  if !util.FloatEqual(UgaussianP(0), 0.5) {
    t.Error("cdf: error computing P(0).")
  }

  if !util.FloatEqual(UgaussianQ(0), 0.5) {
    t.Error("cdf: error computing P(0).")
  }

  if !util.FloatEqual(UgaussianPinv(0.5), 0) {
    t.Error("cdf: error computing Pinv(0.5).")
  }

  if !util.FloatEqual(UgaussianQinv(0.5), 0) {
    t.Error("cdf: error computing Qinv(0.5).")
  }
}
