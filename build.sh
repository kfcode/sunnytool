#!/bin/bash

name="protoc-gen-sunnytool"

go build -o $name

mv ./$name $GOPATH/bin/
