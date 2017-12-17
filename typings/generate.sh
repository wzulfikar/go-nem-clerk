#!/bin/bash

# This script will generate GO struct files 
# based on json files inside json-sample directory.

dir=$(dirname ${BASH_SOURCE[0]})
jsonFiles="json-sample/*.json"

msg1="// GO struct based on JSON file"
msg2="// "
msg3="// "
msg4="// This file was generated using gojson via go-nem-client/typings/generate.sh"
msg5="// DO NOT EDIT UNLESS YOU KNOW WHAT YOU ARE DOING"

count=0
echo "Generating .go files from ${jsonFiles}"

for filename in $dir/$jsonFiles; do
	gofile=$(basename $filename .json).go
	echo $msg1 > $dir/$gofile
	echo $msg2 $filename >> $dir/$gofile
	echo $msg3 >> $dir/$gofile
	echo $msg4 >> $dir/$gofile
	echo $msg5 >> $dir/$gofile

	cat $filename | gojson -pkg nemtypings -name $(basename "$filename" .json) >> $dir/$gofile && echo "✔ ${gofile}"

	count=$((count+1))
done

printf "\n→ [DONE] $count GO struct file(s) have been successfully generated."