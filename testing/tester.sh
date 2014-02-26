#!/bin/sh

# this shell scripts runs go test on all go-gsl test cases
# (C) 2014 Markus Dittrich <haskelladdict@gmail.com>


if [[ $1 != "" ]]; then
  packages=$1
else
  packages=$(find ../ -name "*test.go")
fi

for package in ${packages[*]}; do
  go test `dirname $package`
done
