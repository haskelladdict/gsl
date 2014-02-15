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

func Test_stats_1(t *testing.T) {

  data := FloatSlice{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}

  mean := data.Mean(1)
  if !util.FloatEqual(mean, 5.5) {
    t.Error("Test 1: Failed to compute mean")
  }

  sigma2 := data.Variance(1)
  if !util.FloatEqual(sigma2, 9.166666666666666) {
    t.Error("Test 1: Failed to compute variance.")
  }

  sigma2_m := data.Variance_m(1, mean)
  if !util.FloatEqual(sigma2_m, 9.166666666666666) {
    t.Error("Test 1: Failed to compute variance.")
  }

  sigma := data.Sd(1)
  if !util.FloatEqual(sigma, math.Sqrt(sigma2)) {
    t.Error("Test 1: Failed to compute standard deviation.")
  }

  sigma_m := data.Sd_m(1, mean)
  if !util.FloatEqual(sigma_m, math.Sqrt(sigma2)) {
    t.Error("Test 1: Failed to compute standard deviation with mean.")
  }

  tss_comp := 0.0
  for _, v := range data {
    tss_comp += (v - mean) * (v - mean)
  }

  tss := data.Tss(1)
  if !util.FloatEqual(tss, tss_comp) {
    t.Error("Test 1: Failed to compute total sum of squares.")
  }

  tss_m := data.Tss_m(1, mean)
  if !util.FloatEqual(tss, tss_m) {
    t.Error("Test 1: Failed to compute total sum of squares with mean.")
  }

  sigma2_fixed := data.Variance_with_fixed_mean(1, mean)
  if !util.FloatEqual(sigma2_fixed, 8.25) {
    t.Error("Test 1: Failed to compute variance with fixed mean.")
  }

  sigma_fixed := data.Sd_with_fixed_mean(1, mean)
  if !util.FloatEqual(sigma_fixed, math.Sqrt(sigma2_fixed)) {
    t.Error("Test 1: Failed to compute standard deviation with fixed mean.")
  }

  absdev := data.Absdev(1)
  if !util.FloatEqual(absdev, 2.5) {
    t.Error("Test 1: Failed to compute absolute deviation.")
  }

  absdev_m := data.Absdev_m(1, mean)
  if !util.FloatEqual(absdev_m, absdev) {
    t.Error("Test 1: Failed to compute absolute deviation with mean.")
  }
}



func Test_stats_2(t *testing.T) {

  norm_data := FloatSlice{1.33830225764885e-04, 3.52595682367445e-04,
    8.72682695045760e-04, 2.02904805729977e-03, 4.43184841193801e-03,
    9.09356250159105e-03, 1.75283004935685e-02, 3.17396518356674e-02,
    5.39909665131881e-02, 8.62773188265115e-02, 1.29517595665892e-01,
    1.82649085389022e-01, 2.41970724519143e-01, 3.01137432154804e-01,
    3.52065326764300e-01, 3.86668116802849e-01, 3.98942280401433e-01,
    3.86668116802849e-01, 3.52065326764300e-01, 3.01137432154804e-01,
    2.41970724519143e-01, 1.82649085389022e-01, 1.29517595665892e-01,
    8.62773188265115e-02, 5.39909665131881e-02, 3.17396518356674e-02,
    1.75283004935685e-02, 9.09356250159105e-03, 4.43184841193801e-03,
    2.02904805729977e-03, 8.72682695045760e-04, 3.52595682367445e-04,
    1.33830225764885e-04}

  mean := norm_data.Mean(1)
  if !util.FloatEqual(mean, 0.12120783192361629654) {
    t.Error("Test 2: Failed to compute mean.")
  }

  sd := norm_data.Sd(1)
  if !util.FloatEqual(sd, 0.1418146888138111239) {
    t.Error("Test 2: Failed to compute standard deviation.")
  }

  skew := norm_data.Skew(1)
  if !util.FloatEqual(skew, 0.79446149715833858096) {
    t.Error("Test 2: Failed to compute skew.")
  }

  skew_m_sd := norm_data.Skew_m_sd(1, mean, sd)
  if !util.FloatEqual(skew_m_sd, skew) {
    t.Error("Test 2: Failed to compute skew with mean and stddev.")
  }

  kurt := norm_data.Kurtosis(1)
  if !util.FloatEqual(kurt, -0.98591331325429809596) {
    t.Error("Test 2: Failed to compute kurtosis.")
  }

  kurt_m_sd := norm_data.Kurtosis_m_sd(1, mean, sd)
  if !util.FloatEqual(kurt_m_sd, kurt) {
    t.Error("Test 2: Failed to compute kurtosis with mean and stddev.")
  }
}
