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

// #cgo pkg-config: gsl
// #include <gsl/gsl_statistics.h>
import "C"

// type definitions
type FloatSlice []float64

// Mean returns the arithmetic mean of data with stride stride.
func (d FloatSlice) Mean(stride int) float64 {
  mean := C.gsl_stats_mean((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)))
  return float64(mean)
}

// Variance returns the variance of data with stride stride.
func (d FloatSlice) Variance(stride int) float64 {
  variance := C.gsl_stats_variance((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)))
  return float64(variance)
}

// Variance_m returns the variance of data with stride stride and a user
// supplied value for the mean
func (d FloatSlice) Variance_m(stride int, mean float64) float64 {
  variance_m := C.gsl_stats_variance_m((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)), C.double(mean))
  return float64(variance_m)
}

// Sd returns the standard deviation of data with stride stride.
// NOTE: The value of this function equals the square root of the variance
func (d FloatSlice) Sd(stride int) float64 {
  sd := C.gsl_stats_sd((*C.double)(&d[0]), C.size_t(stride), C.size_t(len(d)))
  return float64(sd)
}

// Sd_m returns the standard deviation of data with stride stride and a user
// supplied value for the mean.
// NOTE: The value of this function equals the square root of the variance
func (d FloatSlice) Sd_m(stride int, mean float64) float64 {
  sd_m := C.gsl_stats_sd_m((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)), C.double(mean))
  return float64(sd_m)
}

// Tss returns the total sum of squares (TSS) of data about the mean with
// stride stride.
func (d FloatSlice) Tss(stride int) float64 {
  tss := C.gsl_stats_tss((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)))
  return float64(tss)
}

// Tss_m returns the total sum of squares (TSS) of data about the mean with
// stride stride and a user supplied value for the mean.
func (d FloatSlice) Tss_m(stride int, mean float64) float64 {
  tss_m := C.gsl_stats_tss_m((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)), C.double(mean))
  return float64(tss_m)
}

// Variance_with_fixed_mean computes an unbiased estimate of the variance
// of data when the population mean mean of the underlying distribution is
// known a priori.
func (d FloatSlice) Variance_with_fixed_mean(stride int, mean float64) float64 {
  var_fixed := C.gsl_stats_variance_with_fixed_mean((*C.double)(&d[0]),
    C.size_t(stride), C.size_t(len(d)), C.double(mean))
  return float64(var_fixed)
}

// Sd_with_fixed_mean computes an unbiased estimate of the standard deviation
// of data when the population mean mean of the underlying distribution is
// known a priori.
// NOTE: This is equal to the square root of the Variance_with_fixed_mean
func (d FloatSlice) Sd_with_fixed_mean(stride int, mean float64) float64 {
  var_fixed := C.gsl_stats_sd_with_fixed_mean((*C.double)(&d[0]),
    C.size_t(stride), C.size_t(len(d)), C.double(mean))
  return float64(var_fixed)
}

// Absdev computes the absolute deviation from the mean of data with
// stride stride
func (d FloatSlice) Absdev(stride int) float64 {
  absdev := C.gsl_stats_absdev((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)))
  return float64(absdev)
}

// Absdev_m computes the absolute deviation from the mean of data with
// stride stride and given mean
func (d FloatSlice) Absdev_m(stride int, mean float64) float64 {
  absdev_m := C.gsl_stats_absdev_m((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)), C.double(mean))
  return float64(absdev_m)
}

// Skew computes the skewness of data with stride stride.
func (d FloatSlice) Skew(stride int) float64 {
  skew := C.gsl_stats_skew((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)))
  return float64(skew)
}

// Skew_m_sd computes the skewness of the dataset data using stride stride
// and the given values of the mean mean and standard deviation sd.
func (d FloatSlice) Skew_m_sd(stride int, mean float64, sd float64) float64 {
  skew_m_sd := C.gsl_stats_skew_m_sd((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)), C.double(mean), C.double(sd))
  return float64(skew_m_sd)
}

// Kurtosis computes the kurtosis of data with stride stride.
func (d FloatSlice) Kurtosis(stride int) float64 {
  kurt := C.gsl_stats_kurtosis((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)))
  return float64(kurt)
}

// Kurtosis computes the kurtosis of data with stride stride and the given
// values of mean and sd.
func (d FloatSlice) Kurtosis_m_sd(stride int, mean float64, sd float64) float64 {
  kurt_m_sd := C.gsl_stats_kurtosis_m_sd((*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)), C.double(mean), C.double(sd))
  return float64(kurt_m_sd)
}

// Lag1_autocorrelation computes the lag-1 autocorrelation of the dataset
// data with stride stride.
func (d FloatSlice) Lag1_autocorrelation(stride int) float64 {
  corr := C.gsl_stats_lag1_autocorrelation((*C.double)(&d[0]),
    C.size_t(stride), C.size_t(len(d)))
  return float64(corr)
}

// Lag1_autocorrelation_m computes the lag-1 autocorrelation of the dataset
// data with stride stride and the given values of mean and sd.
func (d FloatSlice) Lag1_autocorrelation_m_sd(stride int, mean float64) float64 {
  corr_m_sd := C.gsl_stats_lag1_autocorrelation_m((*C.double)(&d[0]),
    C.size_t(stride), C.size_t(len(d)), C.double(mean))
  return float64(corr_m_sd)
}

// Covariance computes the covariance of the dataset d with
// dataset data1 which must both be of the same length n using strides
// stride1 and stride2.
// NOTE: This function currently does not enforce that the datasets
//       have the same length.
func (d FloatSlice) Covariance(stride1 int, data1 FloatSlice, stride2 int) float64 {
  cov := C.gsl_stats_covariance((*C.double)(&d[0]), C.size_t(stride1),
    (*C.double)(&data1[0]), C.size_t(stride2), C.size_t(len(d)))
  return float64(cov)
}

// Covariance_m computes the covariance of the datasets d with dataset
// data1 using the given values of the means, mean1 and mean2.
func (d FloatSlice) Covariance_m(stride1 int, data1 FloatSlice, stride2 int,
  mean1 float64, mean2 float64) float64 {
  cov_m := C.gsl_stats_covariance_m((*C.double)(&d[0]), C.size_t(stride1),
    (*C.double)(&data1[0]), C.size_t(stride2), C.size_t(len(d)),
    C.double(mean1), C.double(mean1))
  return float64(cov_m)
}

// Correlation computes the Pearson correlation coefficient between the
// dataset d and dataset data1 which must both be of the same length n
// using strides stride1 and stride2.
func (d FloatSlice) Correlation(stride1 int, data1 FloatSlice, stride2 int) float64 {
  corr := C.gsl_stats_correlation((*C.double)(&d[0]), C.size_t(stride1),
    (*C.double)(&data1[0]), C.size_t(stride2), C.size_t(len(d)))
  return float64(corr)
}

// Spearman computes the Spearman rank correlation coefficient between the
// dataset d and dataset data2 which must both be of the same length n using
// strides stride1 and stride2. The Spearman rank correlation between
// vectors x and y is equivalent to the Pearson correlation between the
// ranked vectors x_R and y_R, where ranks are defined to be the average
// of the positions of an element in the ascending order of the values.
// NOTE: Additional workspace of size 2*n is required in work.
func (d FloatSlice) Spearman(stride1 int, data1 FloatSlice, stride2 int) float64 {
  work := make([]float64, len(d)*2)
  spear := C.gsl_stats_spearman((*C.double)(&d[0]), C.size_t(stride1),
    (*C.double)(&data1[0]), C.size_t(stride2), C.size_t(len(d)),
    (*C.double)(&work[0]))
  return float64(spear)
}

// Wmean computes the weighted mean of the dataset data with stride
// stride and length n, using the set of weights w with stride wstride and
// length n.
func (d FloatSlice) Wmean(stride int, weights FloatSlice, wstride int) float64 {
  wmean := C.gsl_stats_wmean((*C.double)(&weights[0]), C.size_t(wstride),
    (*C.double)(&d[0]), C.size_t(stride), C.size_t(len(d)))
  return float64(wmean)
}

// Wvariance computes the estimated variance of the dataset data with stride
// stride and length n, using the set of weights w with stride wstride and
// length n.
func (d FloatSlice) Wvariance(stride int, weights FloatSlice, wstride int) float64 {
  wvar := C.gsl_stats_wvariance((*C.double)(&weights[0]), C.size_t(wstride),
    (*C.double)(&d[0]), C.size_t(stride), C.size_t(len(d)))
  return float64(wvar)
}

// Wvariance_m computes the estimated variance of the dataset data with
// stride stride and length n, using the set of weights w with stride
// wstride and length n and weighted mean mean.
func (d FloatSlice) Wvariance_m(stride int, weights FloatSlice, wstride int,
  mean float64) float64 {
  wvar_m := C.gsl_stats_wvariance_m((*C.double)(&weights[0]),
    C.size_t(wstride), (*C.double)(&d[0]), C.size_t(stride), C.size_t(len(d)),
    C.double(mean))
  return float64(wvar_m)
}

// Wsd computes the standard deviation as the square root of the variance.
func (d FloatSlice) Wsd(stride int, weights FloatSlice, wstride int) float64 {
  wsd := C.gsl_stats_wsd((*C.double)(&weights[0]), C.size_t(wstride),
    (*C.double)(&d[0]), C.size_t(stride), C.size_t(len(d)))
  return float64(wsd)
}

// Wsd computes the standard deviation as the square root of the variance
// with given weighted mean.
func (d FloatSlice) Wsd_m(stride int, weights FloatSlice, wstride int,
  mean float64) float64 {
  wsd_m := C.gsl_stats_wsd_m((*C.double)(&weights[0]), C.size_t(wstride),
    (*C.double)(&d[0]), C.size_t(stride), C.size_t(len(d)), C.double(mean))
  return float64(wsd_m)
}

// Wvariance_with fixed mean computes an unbiased estimate of the variance
// of the weighted dataset d when the population mean mean of the underlying
// distribution is known a priori. In this case the estimator for the
// variance replaces the sample mean μ_hat by the known population mean μ.
func (d FloatSlice) Wvariance_with_fixed_mean(stride int, weights FloatSlice,
  wstride int, mean float64) float64 {
  wvar_fixed_m := C.gsl_stats_wvariance_with_fixed_mean(
    (*C.double)(&weights[0]), C.size_t(wstride), (*C.double)(&d[0]),
    C.size_t(stride), C.size_t(len(d)), C.double(mean))
  return float64(wvar_fixed_m)
}

// Wsd_with_fixed_mean computes the standard deviation of data d with fixed
// mean and is defined as the square root of Wvariance_with_fixed_mean.
func (d FloatSlice) Wsd_with_fixed_mean(stride int, weights FloatSlice,
  wstride int, mean float64) float64 {
  wsd_fixed_m := C.gsl_stats_wsd_with_fixed_mean((*C.double)(&weights[0]),
    C.size_t(wstride), (*C.double)(&d[0]), C.size_t(stride), C.size_t(len(d)),
    C.double(mean))
  return float64(wsd_fixed_m)
}

// Wtss computes the weighted total sum of squares (TSS) of data about the
// weighted mean.
func (d FloatSlice) Wtss(stride int, weights FloatSlice, wstride int) float64 {
  wtss := C.gsl_stats_wtss((*C.double)(&weights[0]), C.size_t(wstride),
    (*C.double)(&d[0]), C.size_t(stride), C.size_t(len(d)))
  return float64(wtss)
}

// Wtss_m computes the weighted total sum of squares (TSS) of data about the
// weighted mean supplied by the caller.
func (d FloatSlice) Wtss_m(stride int, weights FloatSlice, wstride int,
  mean float64) float64 {
  wtss_m := C.gsl_stats_wtss_m((*C.double)(&weights[0]), C.size_t(wstride),
    (*C.double)(&d[0]), C.size_t(stride), C.size_t(len(d)), C.double(mean))
  return float64(wtss_m)
}

// Wabsdev computes the weighted absolute deviation from the weighted mean
// of data with stride stride.
func (d FloatSlice) Wabsdev(stride int, weights FloatSlice, wstride int) float64 {
  wabsdev := C.gsl_stats_wabsdev((*C.double)(&weights[0]), C.size_t(wstride),
    (*C.double)(&d[0]), C.size_t(stride), C.size_t(len(d)))
  return float64(wabsdev)
}

// Wabsdev_m computes the weighted absolute deviation from the weighted mean
// of data with stride stride and user supplied mean.
func (d FloatSlice) Wabsdev_m(stride int, weights FloatSlice, wstride int,
  mean float64) float64 {
  wabsdev_m := C.gsl_stats_wabsdev_m((*C.double)(&weights[0]),
    C.size_t(wstride), (*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)), C.double(mean))
  return float64(wabsdev_m)
}

// Wskew computes the weighted skewness of the dataset d with stride stride.
func (d FloatSlice) Wskew(stride int, weights FloatSlice, wstride int) float64 {
  wskew := C.gsl_stats_wskew((*C.double)(&weights[0]), C.size_t(wstride),
    (*C.double)(&d[0]), C.size_t(stride), C.size_t(len(d)))
  return float64(wskew)
}

// Wskew_m_sd computes the weighted skewness of the dataset d with stride
// stride with user supplied mean and standard deviation sd.
func (d FloatSlice) Wskew_m_sd(stride int, weights FloatSlice, wstride int,
  mean float64, sd float64) float64 {
  wskew_m_sd := C.gsl_stats_wskew_m_sd((*C.double)(&weights[0]),
    C.size_t(wstride), (*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)), C.double(mean), C.double(sd))
  return float64(wskew_m_sd)
}

// Wkurtosis computes the weighted kurtosis of the dataset d with stride
// stride.
func (d FloatSlice) Wkurtosis(stride int, weights FloatSlice, wstride int) float64 {
  wkurtosis := C.gsl_stats_wkurtosis((*C.double)(&weights[0]),
    C.size_t(wstride), (*C.double)(&d[0]), C.size_t(stride), C.size_t(len(d)))
  return float64(wkurtosis)
}

// Wkurtosis_m_sd computes the weighted kurtosis of the dataset d with stride
// stride with user supplied mean and standard deviation sd.
func (d FloatSlice) Wkurtosis_m_sd(stride int, weights FloatSlice,
  wstride int, mean float64, sd float64) float64 {
  wkurtosis_m_sd := C.gsl_stats_wkurtosis_m_sd((*C.double)(&weights[0]),
    C.size_t(wstride), (*C.double)(&d[0]), C.size_t(stride),
    C.size_t(len(d)), C.double(mean), C.double(sd))
  return float64(wkurtosis_m_sd)
}
