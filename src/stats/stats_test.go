// Copyright 2014 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// stat wraps gsl statistics routines
//
// NOTE: Unless explicityly noted test target results have been computed
//       via R
package stats

import (
  "github.com/haskelladdict/go-gsl/util"
  "math"
  "testing"
)

// test set 1
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

// test set 2
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

  // XXX: The target values for l1cor have been computed via GSL and
  // were not tested by a third party tool
  l1corr := norm_data.Lag1_autocorrelation(1)
  if !util.FloatEqual(l1corr, 0.9500395402358169) {
    t.Error("Test 2: Failed to compute the lag1 autocorrelation.", l1corr)
  }

  l1corr_m_sd := norm_data.Lag1_autocorrelation_m_sd(1, mean)
  if !util.FloatEqual(l1corr_m_sd, l1corr) {
    t.Error("Test 2: Failed to compute lag1 autocorrelation with mean and stddev.", l1corr_m_sd)
  }
}

// test set 3
func Test_stats_3(t *testing.T) {

  data1 := FloatSlice{16.0, 99.0, 26.0, 85.0, 76.0, 50.0, 46.0, 11.0, 79.0,
    97.0, 24.0, 20.0, 100.0, 68.0, 22.0, 15.0, 5.0, 89.0, 45.0, 2.0}

  data2 := FloatSlice{39.0, 10.0, 34.0, 29.0, 82.0, 54.0, 30.0, 65.0, 56.0,
    55.0, 20.0, 52.0, 96.0, 95.0, 23.0, 51.0, 27.0, 59.0, 31.0, 99.0}

  mean1 := data1.Mean(1)
  mean2 := data2.Mean(1)

  cov := data1.Covariance(1, data2, 1)
  if !util.FloatEqual(cov, 130.93421052631578) {
    t.Error("Test 3: Failed to compute covariance.")
  }

  cov_m := data1.Covariance_m(1, data2, 1, mean1, mean2)
  if !util.FloatEqual(cov_m, cov) {
    t.Error("Test 3: Failed to compute covariance with mean1 and mean2.")
  }

  corr := data1.Correlation(1, data2, 1)
  if !util.FloatEqual(corr, 0.14269187753186113) {
    t.Error("Test 3: Failed to compute Pearson correlation.")
  }

  spear := data1.Spearman(1, data2, 1)
  if !util.FloatEqual(spear, 0.091729323308270688) {
    t.Error("Test 3: Failed to compute Pearson correlation.")
  }
}

// test set 4
func Test_stats_4(t *testing.T) {

  data := FloatSlice{16.0, 99.0, 26.0, 85.0, 76.0, 50.0, 46.0, 11.0, 79.0,
    97.0, 24.0, 20.0, 100.0, 68.0, 22.0, 15.0, 5.0, 89.0, 45.0, 2.0}

  weights := FloatSlice{0.1, 0.2, 0.3, 0.4, 0.1, 0.2, 0.2, 0.1, 0.8,
    0.8, 0.1, 0.2, 0.1, 0.1, 0.2, 0.3, 0.3, 0.3, 0.2, 0.1}

  wmean := data.Wmean(1, weights, 1)
  if !util.FloatEqual(wmean, 59.098039215686278) {
    t.Error("Test 4: Failed to compute weighted mean.")
  }

  // XXX: All values below in this test set were computed via GSL
  // and not confirmened to be correct by a third party application
  wvar := data.Wvariance(1, weights, 1)
  if !util.FloatEqual(wvar, 1248.9347280334725) {
    t.Error("Test 4: Failed to compute weighted variance.")
  }

  wvar_m := data.Wvariance_m(1, weights, 1, wmean)
  if !util.FloatEqual(wvar_m, wvar) {
    t.Error("Test 4: Failed to compute weighted variance with mean.")
  }

  wsd := data.Wsd(1, weights, 1)
  if !util.FloatEqual(wsd, 35.34027062762073) {
    t.Error("Test 4: Failed to compute weighted stddev.")
  }

  wsd_m := data.Wsd_m(1, weights, 1, wmean)
  if !util.FloatEqual(wsd_m, wsd) {
    t.Error("Test 4: Failed to compute weighted stdev with mean.")
  }

  wvariance_fixed_m := data.Wvariance_with_fixed_mean(1, weights, 1, wmean)
  if !util.FloatEqual(wvariance_fixed_m, 1147.6178392925797) {
    t.Error("Test 4: Failed to compute variance with fixed mean.")
  }

  wsd_fixed_m := data.Wsd_with_fixed_mean(1, weights, 1, wmean)
  if !util.FloatEqual(wsd_fixed_m, math.Sqrt(wvariance_fixed_m)) {
    t.Error("Test 4: Failed to compute stddev with fixed mean.")
  }

  wtss := data.Wtss(1, weights, 1)
  if !util.FloatEqual(wtss, 5852.850980392157) {
    t.Error("Test 4: Failed to compute weighted sum of squares.")
  }

  wtss_m := data.Wtss_m(1, weights, 1, wmean)
  if !util.FloatEqual(wtss_m, wtss) {
    t.Error("Test 4: Failed to compute weighted sum of squares with mean.")
  }

  wabsdev := data.Wabsdev(1, weights, 1)
  if !util.FloatEqual(wabsdev, 31.46097654748174) {
    t.Error("Test 4: Failed to compute weighted absolute deviation.")
  }

  wabsdev_m := data.Wabsdev_m(1, weights, 1, wmean)
  if !util.FloatEqual(wabsdev_m, wabsdev) {
    t.Error("Test 4: Failed to compute weighted absolute deviation with mean.")
  }

  wskew := data.Wskew(1, weights, 1)
  if !util.FloatEqual(wskew, -0.28679295109648656) {
    t.Error("Test 4: Failed to compute weighted skew.")
  }

  wskew_m_sd := data.Wskew_m_sd(1, weights, 1, wmean, wsd)
  if !util.FloatEqual(wskew_m_sd, wskew) {
    t.Error("Test 4: Failed to compute weighted skew with mean and stddev.")
  }

  wkurtosis := data.Wkurtosis(1, weights, 1)
  if !util.FloatEqual(wkurtosis, -1.7414231965732037) {
    t.Error("Test 4: Failed to compute weighted kurtosis.")
  }

  wkurtosis_m_sd := data.Wkurtosis_m_sd(1, weights, 1, wmean, wsd)
  if !util.FloatEqual(wkurtosis_m_sd, wkurtosis) {
    t.Error("Test 4: Failed to compute weighted kurtosis with mean and stddev.")
  }
}

// test set 5
func Test_stats_5(t *testing.T) {

  data := FloatSlice{16.0, 99.0, 26.0, 85.0, 76.0, 50.0, 46.0, 11.0, 79.0,
    97.0, 24.0, 20.0, 100.0, 68.0, 22.0, 15.0, 5.0, 89.0, 45.0, 2.0}

  max := data.Max(1)
  if !util.FloatEqual(max, 100.0) {
    t.Error("Test 5: Failed to compute data max.")
  }

  min := data.Min(1)
  if !util.FloatEqual(min, 2.0) {
    t.Error("Test 5: Failed to compute data min.")
  }

  min1, max1 := data.MinMax(1)
  if !util.FloatEqual(min, min1) || !util.FloatEqual(max, max1) {
    t.Error("Test 5: Failed to compute data minmax.")
  }

  maxInd := data.MaxIndex(1)
  if maxInd != 12 {
    t.Error("Test 5: Failed to compute data max index.")
  }

  minInd := data.MinIndex(1)
  if minInd != 19 {
    t.Error("Test 5: Failed to compute min index.")
  }

  minInd1, maxInd1 := data.MinMaxIndex(1)
  if (minInd != minInd1) || (maxInd != maxInd1) {
    t.Error("Test 5: Failed to compute data minmax index.")
  }
}

// test set 6
func Test_stats_6(t *testing.T) {

  data := FloatSlice{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}

  median := data.MedianFromSortedData(1)
  if !util.FloatEqual(median, 5.5) {
    t.Error("Test 6: Failed to compute median value.")
  }

  quantile1 := data.QuantileFromSortedData(1, 0.15)
  if !util.FloatEqual(quantile1, 2.35) {
    t.Error("Test 6: Failed to compute quantile1 value.")
  }

  quantile2 := data.QuantileFromSortedData(1, 0.92)
  if !util.FloatEqual(quantile2, 9.28) {
    t.Error("Test 6: Failed to compute quantile2 value.")
  }
}
