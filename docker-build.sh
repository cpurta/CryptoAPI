#!/bin/bash

if [ $# -ne 1 ]; then
    echo "tag required to build docker image"
    exit 1
fi

tag=$1

export GOOS=linux
export GOARCH=amd64

cd src

wgo build -o crypto-api
mv crypto-api ..

cd ..

docker build -f ./Dockerfile -t crypto-api:$tag .

rm ./crypto-api

exit 0
