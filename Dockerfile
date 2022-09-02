From golang:latest

MAINTAINER RobKing_Goroutine_ChatRoom

WORKDIR /app/demo

COPY . .

RUN GOPROXY="https://goproxy.io" 
GO111MODULE=on
 go build main.go

EXPOSE 9999

ENTRYPOINT ["./main"]