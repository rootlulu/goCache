#!/bin/bash

data_path="/root/projs/docker/data/mysql"

if [ ! -d "${data_path}" ]; then
    mkdir -p "${data_path}"
    echo "the dir ${data_path} has created"
else
    echo "the dir ${data_path} already existed"
fi

container_name="mysql"
container_id=$(docker ps -a -f "name=${container_name}" -q)

if [ -n "$container_id" ]; then
    echo 1111111
    echo $container_id
    docker rm $container_id
fi

docker run -d \
  --name mysql \
  -e MYSQL_ROOT_PASSWORD=luban666. \
  -e MYSQL_USER=lulu \
  -e MYSQL_PASSWORD=luban666. \
  -p 3306:3306 \
  -v /root/projs/docker/data/mysql:/var/lib/mysql \
  mysql:5.7
