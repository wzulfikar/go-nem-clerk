#!/bin/bash

# This script will generate .go files from
# json files inside ../sample-json directory.

dir=$(dirname ${BASH_SOURCE[0]})
msg="// This file is auto-generated using go-nem-clerk/typings/generate.sh\n"
jsonFiles="../sample-json/*.json"
echo "Generating .go files from ${jsonFiles}"

for filename in $dir/$jsonFiles; do
	gofile=$(basename $filename .json).go
	echo $msg > $dir/$gofile

	cat $filename | gojson -pkg nemtypings -name $(basename "$filename" .json) >> $dir/$gofile && echo "✔ ${gofile}"
done
