FROM golang:1.22.3 as builder
ARG CGO_ENABLED=0
WORKDIR /app

COPY api/go.mod api/go.sum ./
RUN go mod download
COPY ./api/ .

RUN go build

FROM scratch
COPY --from=builder /app/goride /server
ENTRYPOINT ["/goride"]
