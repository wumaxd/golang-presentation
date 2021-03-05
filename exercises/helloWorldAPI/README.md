# Hello World API

## Run locally

To run the program for testing use `go run ./cmd/api/main.go`.

To create a native executable run `go build ./...` and then execute the program `./api`.

## Build container image

| Option | Build command |
| ------ | ------------- |
| Alpine Linux | `docker build -t api:alpine .` |
| Universial Base Image 8 | `docker build -t api:ubi8 -f Dockerfile.ubi .` |

## Run container

To run the container use

`docker run --rm -p 8080:8080 api:alpine` or `docker run --rm -p 8080:8080 api:ubi8`.

To test the API you can use

`curl localhost:8080/api/v1/hello`

## Comparison of the image size

```bash
docker images
REPOSITORY        TAG             IMAGE ID       CREATED          SIZE
api               ubi8            0c67dd9cc33e   23 seconds ago   108MB
api               alpine          4e83eccdbbb3   10 minutes ago   10.5MB
```
