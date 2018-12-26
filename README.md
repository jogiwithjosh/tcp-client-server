# tcp-client-server
Simple TCP Server runs on two ports

This repo has 4 Folders.
1. tcp-server
2. tcp-client-with-reply
3. tcp-client-only-send
4. system-call


# tcp-server runs on two ports:
6777 - Accepts JSON ex.{"id":"1", "name":"Jogi", "age":"26"}
7777 - Acceps any request and respond back with response

to run server: navigate to /tcp-server > go run main.go

# tcp-client-with-reply connects to tcp-server and request on 7777 port and prints response from server
to run: navigate to /tcp-client-with-reply > go run main.go

# tcp-client-only-send connects to tcp-server on 6777 and sends JSON data like {"id":"1", "name":"Jogi", "age":"26"}
to run: navigate to /tcp-client-only-send > go run main.go

# system-call

to run: navigate to /system-call > go run main.go
prints UID of the user that executed program and PID of the current process.

to test system-call
go test .



