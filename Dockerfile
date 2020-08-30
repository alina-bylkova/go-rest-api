# Uncomment, ff we built runable app (go-rest-api.exe) on host machine (on this computer)
# FROM alpine:3.12
# WORKDIR /usr/local/bin
# COPY go-rest-api ./
# EXPOSE 8080
# CMD ["./go-rest-api"]

# If we want to build executable file (go-rest-api) insidde the container
FROM golang:1.15-alpine
WORKDIR /usr/local/src
COPY . .
RUN go build .
RUN mv go-rest-api ../bin
RUN rm -r /usr/local/src
WORKDIR /usr/local/bin
EXPOSE 8080
CMD ["./go-rest-api"]