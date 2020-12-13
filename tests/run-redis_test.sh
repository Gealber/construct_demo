#! /bin/bash

echo "Environments variables"
source ../env-local.sh
go test ../repository/redis/ 
