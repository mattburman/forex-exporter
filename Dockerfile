FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git ca-certificates

# Create appuser.
ENV USER=appuser
ENV UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/forex-exporter


FROM scratch
# Import user and group files from builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
# Copy CA certs
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Copy binary
COPY --from=builder /go/bin/forex-exporter /go/bin/forex-exporter

USER appuser:appuser
ENTRYPOINT ["/go/bin/forex-exporter"]
