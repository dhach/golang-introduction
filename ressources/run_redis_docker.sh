#! /bin/bash

docker run -it --name redis-server --rm -p 6379:6379 -e "ALLOW_EMPTY_PASSWORD=yes" bitnami/redis

