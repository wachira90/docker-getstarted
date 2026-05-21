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