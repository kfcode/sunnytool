#!/bin/bash

name="protoc-gen-protorpcxtools"

go build -o $name

mv ./$name $GOPATH/bin/
