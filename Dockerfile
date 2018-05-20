FROM amd64/golang
WORKDIR ./simplegoserver
COPY . .
ENTRYPOINT ["./startServer.sh"]
