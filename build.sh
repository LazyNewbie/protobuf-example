#!/bin/bash

rm -rf generated/go/*
rm -rf generated/php/*

find ./proto_files/ -iname "*.proto" | while read file
do
  ./protoc/bin/protoc --php_out=./generated/php/ --go_out=./generated/go/ "$file";
done