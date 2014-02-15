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
  "testing"
  "github.com/haskelladdict/go-gsl/util"
)


func Test_stats(t *testing.T) {
  data := FloatSlice{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
  mean := data.Mean(1)
  if !util.FloatEqual(mean, 5.5) {
    t.Error("Failed to compute mean")
  }
}
