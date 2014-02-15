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

import (
  "github.com/haskelladdict/go-gsl/util"
  "math"
  "testing"
)

func Test_stats(t *testing.T) {

  data := FloatSlice{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}

  mean := data.Mean(1)
  if !util.FloatEqual(mean, 5.5) {
    t.Error("Failed to compute mean")
  }

  sigma2 := data.Variance(1)
  if !util.FloatEqual(sigma2, 9.166666666666666) {
    t.Error("Failed to compute variance.")
  }

  sigma2_m := data.Variance_m(1, mean)
  if !util.FloatEqual(sigma2_m, 9.166666666666666) {
    t.Error("Failed to compute variance.")
  }

  sigma := data.Sd(1)
  if !util.FloatEqual(sigma, math.Sqrt(sigma2)) {
    t.Error("Failed to compute standard deviation.")
  }

  sigma_m := data.Sd_m(1, mean)
  if !util.FloatEqual(sigma_m, math.Sqrt(sigma2)) {
    t.Error("Failed to compute standard deviation with mean.")
  }

  tss_comp := 0.0
  for _, v := range data {
    tss_comp += (v - mean) * (v - mean)
  }

  tss := data.Tss(1)
  if !util.FloatEqual(tss, tss_comp) {
    t.Error("Failed to compute total sum of squares.")
  }

  tss_m := data.Tss_m(1, mean)
  if !util.FloatEqual(tss, tss_m) {
    t.Error("Failed to compute total sum of squares with mean.")
  }

  sigma2_fixed := data.Variance_with_fixed_mean(1, mean)
  if !util.FloatEqual(sigma2_fixed, 8.25) {
    t.Error("Failed to compute variance with fixed mean.")
  }

  sigma_fixed := data.Sd_with_fixed_mean(1, mean)
  if !util.FloatEqual(sigma_fixed, math.Sqrt(sigma2_fixed)) {
    t.Error("Failed to compute standard deviation with fixed mean.")
  }

  absdev := data.Absdev(1)
  if !util.FloatEqual(absdev, 2.5) {
    t.Error("Failed to compute absolute deviation.")
  }

  absdev_m := data.Absdev_m(1, mean)
  if !util.FloatEqual(absdev_m, absdev) {
    t.Error("Failed to compute absolute deviation with mean.")
  }
}
