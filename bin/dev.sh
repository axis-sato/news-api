#! /bin/bash

cd "$(dirname "$0")/.."

case $1 in
  "generate!") docker-compose exec generator go generate ./...;;
  "generate") docker-compose exec generator go generate "${2}";;
esac