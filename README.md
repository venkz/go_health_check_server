# go_health_check_server

Smallest Go server which can be used for health check in Docker containers.  
This server supports both TCP and HTTP connections.

## Instructions to generate binary

1. Make sure you are in the source directory

    ```
    > pwd
        ~/go/src/go_health_check_server
    ```

2. Go Install to generate the binary
    ##### Linux

    ```
    > GOOS=linux GOARCH=amd64 go install go_health_check_server
    ```
    
    ##### Mac OS
    
        ```
        > go install go_health_check_server
        ```

3. Start the Go binary
    ##### HTTP Server
    ```
    > ./go_health_check_server --protocol = "http" --port 8080
    
        Init: Starting Health check server..
        Server(http) up and running on port 8080
    ```
    
    ##### TCP Server
    ```
    > ./go_health_check_server --protocol = "tcp" --port 8080
    
        Init: Starting Health check server..
        Server(tcp) up and running on port 8080
    ```

4. Test server
    ##### HTTP Test
    ```bash
    curl -I 127.0.0.1:8080/health                                                                                               HTTP/1.1 200 OK
    Server: Alive and Kicking!
    Date: Wed, 24 Jul 2019 20:52:03 GMT

    ```
    
    ##### TCP Test
    ```bash
     nc -v 127.0.0.1 8080 
     found 0 associations
     found 1 connections:
          1:	flags=82<CONNECTED,PREFERRED>
     	outif lo0
     	src 127.0.0.1 port 54352
     	dst 127.0.0.1 port 8080
     	rank info not available
     	TCP aux info available
     
     Connection to 127.0.0.1 port 8080 [tcp/http-alt] succeeded!
     Alive and Kicking!%
  
    ```
