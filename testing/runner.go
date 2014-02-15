// Copyright 2014 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// stat wraps gsl statistics routines
package main

import (
  "fmt"
  "github.com/haskelladdict/go-gsl/src/stats"
)

func main() {

  data := stats.FloatSlice{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
  fmt.Println("mean is ", data.Mean(1), data.Variance(1),
    data.Variance_with_fixed_mean(1, 3.0), data.Absdev(1))
}
