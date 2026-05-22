FROM public.ecr.aws/docker/library/golang:1.22.1-alpine3.19 AS build

WORKDIR /app

COPY . .

RUN go mod init app

RUN go build -o main .

FROM public.ecr.aws/docker/library/alpine:3.19.1

WORKDIR /app

COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]

# docker buildx build -t myapp:v1 . --no-cache

# docker.io/library/myapp:v1

# docker run -it --name www -p 8011:8080 -d docker.io/library/myapp:v1

# http://localhost:8011