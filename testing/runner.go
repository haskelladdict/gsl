// Copyright 2014 Markus Dittrich. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// stat wraps gsl statistics routines
package main

import (
  "fmt"
  "log"
  "time"

  "github.com/haskelladdict/gsl/random"
  //  "github.com/haskelladdict/gsl/stats"
)

// helper functions for timing
func trace(s string) (string, time.Time) {
  log.Println("START:", s)
  return s, time.Now()
}

func un(s string, startTime time.Time) {
  endTime := time.Now()
  log.Println("  END:", s, "ElapsedTime in seconds:", endTime.Sub(startTime))
}

func time_rng(rngType random.RngType) {
  defer un(trace("TIMING " + rngType.Name()))

  rng_state := random.Rng_alloc(rngType)
  rng_state.UniformSlice(100000000)
}

func main() {

  //  rng_type := random.Ranlxd2
  //  rng_state := random.Rng_alloc(rng_type)

  /*
    data := random.FlatSlice(rng_state, 1.0, 3.0, 10000)
    for _, v := range data {
      fmt.Println(v)
    }
  */

  x := 0.0
  for i := 0; i < 10000; i++ {
    fmt.Println(x, random.FlatPdf(x, 1.0, 3.0))
    x = x + 0.001
  }
  //  for _, x := range random.BivariateGaussianSlice(rng_state, 1, 1, -1.0, 10000) {
  //    fmt.Println(x[0],x[1])
  //  }

  /*
     gaus := stats.FloatSlice(random.GaussianRatioMethodSlice(rng_state, 1, 10000))
     fmt.Println(gaus.Mean(1))
     fmt.Println(gaus.Sd(1))
     for _, i := range gaus {
       fmt.Println(i)
     }
  */

  /*
     data := stats.FloatSlice{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
     fmt.Println("mean is ", data.Mean(1), data.Variance(1),
       data.Variance_with_fixed_mean(1, 3.0), data.Absdev(1))
  */

  /*
     random.EnvSetup()
     rng_type := random.Default()

     for k, v := range random.TypesSetup() {
       if k == "vax" {
         rng_type = v
       }
     }


     fmt.Println(" the type is ", rng_type, rng_type.Name() == "mt19937")
     rng_state := random.Rng_alloc(random.Ranlxs0)
     fmt.Println("seed ", random.DefaultSeed())
     //rng_state.Set(3241)
     //for i := 0; i < 100000; i++ {
     //  fmt.Println(rng_state.UniformInt(100))
     //}


     fmt.Println(random.GaussianSlice(rng_state, 4, 10))
  */
  /*
     foo := rng_state.UniformIntSlice(10, 100000)
     fmt.Println(foo)

     //fmt.Println(rng_state.Name())
     //fmt.Println(rng_state)
     //fmt.Println(rng_state.Max(), rng_state.Min())
     //fmt.Println(rng_state.State(), rng_state.Size())
     //random.TypesSetup()
  */
  /*
     if rng_state.Fwrite("something") != nil {
       panic("failed to write state")
     }

     rng_state_new := random.Alloc(random.Ranlxs0)
     rng_state_new.Fread("something")

     fmt.Println(rng_state.Uniform(), rng_state_new.Uniform())
  */
  //fmt.Println(rng_state.UniformIntSlice(10, 100))

  /*time_rng(random.Taus)
    time_rng(random.Ranlux)
    time_rng(random.Ranlxs1)
    time_rng(random.Ranlxd2)
    time_rng(random.Mt19937)*/

  /*
     rng_state_q := random.Qrng_alloc(random.Halton, 2)
     for i := 0; i < 1024; i++ {
         point := rng_state_q.Get()
         fmt.Println(point[0], point[1])
       }
       rng_state.Free()
  */
}
