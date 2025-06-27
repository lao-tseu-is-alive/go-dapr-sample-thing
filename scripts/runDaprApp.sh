#!/bin/bash
echo "NUM ARGS : " $#
if [ $# -eq 1 ]; then
  APP_NAME=${1,,}
else
  APP_NAME=main
fi
# get the name of the go file to run as first argument

dapr run --app-id example-service \
         --app-protocol grpc \
         --app-port 50001 \
         go run "${APP_NAME}.go"