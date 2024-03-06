#!/bin/bash
rm /www/* -r
cp /root/goweb/* /www/ -r
cd /www
docker build --tag goweb .
docker rm goweb_1
docker run --name goweb_1 -p 80:8080 goweb:latest