FROM golang:1.15.2-alpine3.12
COPY main.go .
RUN go build main.go
FROM scratch
COPY --from=0 /go/main .
CMD ["./main"]