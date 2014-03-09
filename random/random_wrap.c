/* 
 * Copyright 2014 Markus Dittrich. All rights reserved.                       
 * Use of this source code is governed by a BSD-style                         
 * license that can be found in the LICENSE file. 
 *
 * this function provides additional gsl wrappers for go-gsl
 */

#include <gsl/gsl_rng.h>
#include <stdio.h>


/* rng_types_length returns the number of rng types available */
size_t rng_types_length() {

  size_t length = 0;
  const gsl_rng_type **t0 = gsl_rng_types_setup();
  for (const gsl_rng_type **t = t0; *t != 0; t++) {
    length++;
  }

  return length;
}


/* rng_fwrite writes the state of the rng to a file with filename
 * NOTE: It would be much better to be able to use the gsl API
 * function gsl_rng_fwrite to write to any Go writer but I am not
 * sure how to do that or if it is even possible. */
int rng_fwrite(const char *fileName, const gsl_rng *r) {

  FILE *file = fopen(fileName, "w");
  if (file == NULL) {
    return 1;
  }

  int status = gsl_rng_fwrite(file, r);
 
  // we need to flush the stream since otherwise the stream
  // remains empty. I don't understand why.
  if (fflush(file) != 0) {
    return 1;
  }
  
  return status;
}


/* rng_fread read the state of the rng from a file with filename
 * NOTE: It would be much better to be able to use the gsl API
 * function gsl_rng_fread to read into any Go reader but I am not
 * sure how to do that or if it is even possible. */
int rng_fread(const char *fileName, gsl_rng *r) {

  FILE *file = fopen(fileName, "r");
  if (file == NULL) {
    return 1;
  }

  return gsl_rng_fread(file, r);
}



