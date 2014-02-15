#!/bin/sh

# this shell scripts runs go format on all go-gsl packages
# (C) 2014 Markus Dittrich <haskelladdict@gmail.com>

packages=("../src/stats" "../util")

options="-w=true -tabs=false -tabwidth=2"


for package in ${packages[*]}; do
  gofmt ${options} ${package}
done
