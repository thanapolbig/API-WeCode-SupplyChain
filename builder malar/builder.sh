#!/bin/bash

env=$1

echo "$env Building..."
GOOS=linux
GOARCH=amd64
cd .. && git checkout $env && git pull origin $env && go build .
echo "Copy file to server..."
cd builder\ malar
scp -r -i "sumate.pem" ../api-wecode-supplychain ubuntu@ec2-54-84-248-102.compute-1.amazonaws.com:/home/ubuntu/Projects/Proj1/
echo "Finish"
