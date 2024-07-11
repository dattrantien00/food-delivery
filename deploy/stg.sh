#!/usr/bin/env bash

APP_NAME=food-delivery

docker load -i ${APP_NAME}.tar
docker rm -f ${APP_NAME}
# docker rmi $(docker images -qa -f 'dangling=true')

docker run -d --name ${APP_NAME} \
  --network demo \
  -e VIRTUAL_HOST="" \
  -e LETSENCRYPT_HOST="" \
  -e LETSENCRYPT_EMAIL="" \
  -e DBConnStr="" \
  -p 8080:8080 \
  ${APP_NAME}