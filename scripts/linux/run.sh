#! /bin/bash

echo "Declaring environment variables"
source ./env-local.sh

echo "Starting to run server..."
go run ../../main.go
