/* 
 * Copyright 2014 Markus Dittrich. All rights reserved.                       
 * Use of this source code is governed by a BSD-style                         
 * license that can be found in the LICENSE file. 
 *
 * this function provides additional gsl wrappers for go-gsl
 *
 */

#include <gsl/gsl_rng.h>
#include <stdio.h>

size_t rng_types_length() {

  size_t length = 0;
  const gsl_rng_type **t0 = gsl_rng_types_setup();
  for (const gsl_rng_type **t = t0; *t != 0; t++) {
    length++;
  }

  return length;
}


