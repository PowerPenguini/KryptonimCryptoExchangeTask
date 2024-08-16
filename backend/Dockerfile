FROM golang:1.23 AS build-stage
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./kryptonim

FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM alpine:latest AS release-stage
ENV GIN_MODE=release
WORKDIR /app
COPY --from=build-stage ./kryptonim ./kryptonim
RUN addgroup -S nonroot && adduser -S nonroot -G nonroot
RUN chown nonroot:nonroot ./kryptonim
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["./kryptonim"]

