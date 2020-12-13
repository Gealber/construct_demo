#! /bin/bash

go mod vendor
docker build --tag Gealber/construct_demo -f dockerfile.local ../.
rm -r ../vendor
