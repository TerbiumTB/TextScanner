FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS builder
LABEL authors="terbium"

WORKDIR /app

COPY go.mod go.sum  ./

RUN go mod download

COPY . .


ARG TARGETOS
ARG TARGETARCH
#GOOS=$TARGETOS GOARCH=$TARGETARCH
RUN CGO_ENABLED=0 go build -o /app/gateway .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/gateway /app/gateway

EXPOSE 8000
ENTRYPOINT ["/app/gateway"]