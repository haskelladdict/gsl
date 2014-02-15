#!/bin/sh

# this shell scripts runs go test on all go-gsl test cases
# (C) 2014 Markus Dittrich <haskelladdict@gmail.com>

packages=("../src/stats")

for package in ${packages[*]}; do
  go test ${package}
done
