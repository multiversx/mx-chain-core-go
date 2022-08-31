# About
This directory contains test server and client applications.

The server and client apps use `127.0.0.1:21111` for communicating via websockets. Combined, the server
sends requests continuously, and the clients will process them and send acknowledge for all processed requests.

# How to use
1. go to `server` directory
2. `go build`
3. `./server` 
4. from another terminal go to `client` directory
5. `go build`
6. `./bin`

Step number 6 might be executed again in other terminal, as to simulate multiple clients
