# Dockerized Golang Web Server

## Steps:

1. Install Go
2. Install Docker
3. Run the command to build the docker image
    `docker build .`
4. Start the container using
    `docker run -d -p 8090:8090 <image-id>`
5. Open localhost:8090/ping or localhost:8090/pong