#!/bin/bash

docker build -t go-web .
docker run --restart=always --network=host go-web