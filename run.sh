#!/bin/sh

docker build -t kubectl . && docker run -p 9191:8080 kubectl