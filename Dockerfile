FROM golang:alpine
RUN apk add --update --no-cache git
RUN apk add --no-cache gcc musl-dev
RUN go get -u github.com/gorilla/mux
RUN go get github.com/mattn/go-sqlite3
RUN mkdir -p /go/src/github.com/prashantkamdar
WORKDIR /go/src/github.com/prashantkamdar
RUN  git clone https://github.com/prashantkamdar/todolist.git
WORKDIR /go/src/github.com/prashantkamdar/todolist/src
RUN go build -o /go/bin/todolistapp .
# 546 MB

FROM alpine
COPY --from=0 /go/bin/todolistapp .
CMD ["./todolistapp"]
# 17.6 MB