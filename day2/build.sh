#!/usr/bin/env bash

go build

input="$(cat day2_input.txt)"
result="$(./day2 $input)"
echo $result

rm day2