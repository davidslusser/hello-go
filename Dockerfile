FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o app .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .

# Accept IMAGE_TAG at build time and store it as an env var
ARG IMAGE_TAG
ENV IMAGE_TAG=${IMAGE_TAG}

ENV PORT=8000
CMD ["./app"]
