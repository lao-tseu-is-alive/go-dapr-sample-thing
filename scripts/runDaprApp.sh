#!/bin/bash
echo "NUM ARGS : " $#
if [ $# -eq 1 ]; then
  APP_NAME=${1,,}
else
  APP_NAME=main
fi
# check if the file exist
if [ ! -f $APP_NAME ]; then
  echo "💥💥💥 the file ${APP_NAME} does not exist"
  exit 1
fi
# get the name of the go file to run as first argument

dapr run --app-id example-service \
         --app-protocol grpc \
         --app-port 50001 \
         go run "${APP_NAME}"