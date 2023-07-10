FROM golang:1.20 AS builder
WORKDIR /go/src/BACKEND/
COPY go.mod go.sum ./
RUN go mod download -x
COPY cmd/server ./cmd/server
COPY config ./config
COPY internals ./internals
COPY docs ./docs
ARG _ENV
ENV ENV=${_ENV}
RUN CGO_ENABLED=0 GOOS=linux go build -a -v BACKEND/cmd/server

FROM gcr.io/distroless/static:latest
WORKDIR /
COPY --from=builder /go/src/BACKEND/cmd/server /server
COPY --from=builder /go/src/BACKEND/config /configs
EXPOSE 8080
ENTRYPOINT ["/server"]
