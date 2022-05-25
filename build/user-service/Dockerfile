# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:stretch as builder

LABEL maintainer="Dionsio Paulo <dionsiopaulonamuetho@gmail.com>"

WORKDIR /app
COPY . ./

RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 9004

CMD ["./main"] 