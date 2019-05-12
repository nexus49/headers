# Overview

Basic docker container that listens to get requests and prints out the received headers of the http request.

Useful for setting up hello world applications and troubleshooting.

# How to start

`docker run -p 8001:8001 nexus49/headers` 

and request `http://localhost:8001/`
