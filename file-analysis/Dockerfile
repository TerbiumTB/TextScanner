FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS builder
LABEL authors="terbium"

WORKDIR /app

COPY go.mod go.sum  ./

RUN go mod download

COPY . .


ARG TARGETOS
ARG TARGETARCH
#GOOS=$TARGETOS GOARCH=$TARGETARCH
RUN CGO_ENABLED=0 go build -o /app/analysis .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/analysis /app/analysis

EXPOSE 8080
ENTRYPOINT ["/app/analysis"]