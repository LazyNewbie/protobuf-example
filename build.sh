#!/bin/bash

rm -r generated/go/*
rm -r generated/php/*
rm -r generated/avro/*

find ./proto_files/ -iname "*.proto" | while read file
do
#  ./protoc/bin/protoc --php_out=./generated/php/ --go_out=./generated/go/ "$file";
  ./protoc/bin/protoc --php_out=./generated/php/ --go_out=./generated/go/ --avro_out=./generated/avro/ --avro_opt=namespace_map=main:exads.avro "$file";
done