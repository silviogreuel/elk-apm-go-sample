FROM golang:1.12-alpine AS base 
RUN apk add bash ca-certificates git gcc g++ libc-dev make

FROM base AS dependencies 
WORKDIR /workspace
COPY go.mod .
COPY go.sum .
COPY Makefile .
RUN make deps

FROM dependencies AS builder
ENV GOOS=linux
ENV GOARCH=amd64
COPY . .
RUN make build

FROM alpine AS runner
COPY --from=builder /workspace/bin/casing /bin/casing
ENTRYPOINT ["/bin/casing"]
