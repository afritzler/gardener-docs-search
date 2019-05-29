FROM golang:1.12.5 as builder
WORKDIR /go/src/github.com/afritzler/gardener-docs-search
COPY . .
RUN make build-linux

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=builder /go/src/github.com/afritzler/gardener-docs-search/gardener-docs-search_linux_amd64 /gardener-docs-search
CMD ["./gardener-docs-search"]