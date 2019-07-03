# go_health_check_server

Smallest Go server which can be used for health check in Docker containers.

## Instructions to generate binary

1. Make sure you are in the source directory

    ```
    > pwd
        ~/go/src/go_health_check_server

2. Go Install to generate the binary

    ```
    > go install go_health_check_server

3. Start the Go binary
    ```
    > ./go_health_check_server
        Init: Health check server up and running


## Test the binary
 ```
 > curl -I -X GET localhost:5454/health
     HTTP/1.1 200 OK
     Server: Alive and Kicking!
     Date: Wed, 03 Jul 2019 18:08:32 GMT
     Content-Length: 0
