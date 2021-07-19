#!/bin/sh
set -e

# prod or stg or local
command="$1"

# get environment variables
if [ "$command" = "prod" ];then
    /amadeus/main
elif [ "$command" = "stg" ];then
    /amadeus/main
else
    export PORT=80
    air
fi
