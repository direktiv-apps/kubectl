#!/bin/sh

docker build -t kubetcl . && docker run -p 9191:8080 kubetcl