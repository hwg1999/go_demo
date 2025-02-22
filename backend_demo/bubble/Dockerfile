FROM golang:alpine AS builder
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o bubble .

FROM scratch
COPY ./templates /templates
COPY ./static /static
COPY ./conf /conf
COPY --from=builder /build/bubble /
ENTRYPOINT ["/bubble", "conf/config.ini"]