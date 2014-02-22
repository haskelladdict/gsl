// Copyright 2014 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// stat wraps gsl statistics routines
package main

import (
  "fmt"
  "github.com/haskelladdict/go-gsl/src/random"
//  "github.com/haskelladdict/go-gsl/src/stats"
)

func main() {
/*
  data := stats.FloatSlice{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
  fmt.Println("mean is ", data.Mean(1), data.Variance(1),
    data.Variance_with_fixed_mean(1, 3.0), data.Absdev(1))
*/
  //rng_type := random.Default()
  rng_state := random.Alloc(random.Mt19937)
  rng_state.Set(3241)
  for i:=0; i < 10; i++ {
    fmt.Println(rng_state.UniformInt(100))
  }

  fmt.Println(rng_state.Name())
  fmt.Println(rng_state)
  fmt.Println(rng_state.Max(), rng_state.Min())
  fmt.Println(rng_state.State(), rng_state.Size())
}
