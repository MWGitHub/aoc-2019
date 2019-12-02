#!/usr/bin/env bash

go build

inputs=$(cat day1_input.txt | tr '\n' ' ')
result="$(./day1 $inputs)"
echo "$result"

rm day1