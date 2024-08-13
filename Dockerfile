# Stage 1 -> Installer && go build
FROM golang:1.21.3-alpine as builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./app ./main.go

# Stage 2 -> Run
FROM alpine:latest

RUN apk update && rm -rf /var/cache/apk/*

RUN mkdir -p /app
WORKDIR /app

# Copia la carpeta de assets desde el builder
COPY --from=builder /app/assets /app/assets

COPY --from=builder /app/app .

EXPOSE 8223

ENTRYPOINT [ "./app" ]