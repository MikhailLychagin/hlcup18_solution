#!/usr/bin/env bash

docker build -t stor.highloadcup.ru/accounts/cute_bison -f Dockerfile .
docker push stor.highloadcup.ru/accounts/cute_bison