FROM golang:1.12.6 as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go get -d -v github.com/nsf/termbox-go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o hello  .
FROM scratch
COPY --from=builder /build/hello /app/
WORKDIR /app
CMD ["./hello"]
