/* 
 * Copyright 2014 Markus Dittrich. All rights reserved.                       
 * Use of this source code is governed by a BSD-style                         
 * license that can be found in the LICENSE file. 
 *
 * this function provides additional gsl wrappers for go-gsl
 */


#ifndef CWRAP_H
#define CWRAP_H

#ifdef __cplusplus
extern "C" {
#endif


size_t rng_types_length();

int rng_fwrite(const char *fileName, const gsl_rng *r);
int rng_fread(const char *fileName, gsl_rng *r);


#ifdef __cplusplus
}
#endif

#endif
