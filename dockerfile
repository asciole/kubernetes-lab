FROM golang:1.15

COPY . .

RUN go env -w GO111MODULE=auto

RUN go build -o server .

CMD ["./server"]