FROM golang:1.19-alpine

WORKDIR /Users/Shanaka/Documents/programing/GO/go-simple-rest-application

COPY . .

RUN go get -u github.com/go-sql-driver/mysql

RUN go build -o main main.go

CMD ["./main"]