#!/bin/bash
set -e

repos=istio-demo-helloworld
tag=v0.1

go build -o  helloworld main.go

docker build -t $repos:$tag .

