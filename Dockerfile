FROM golang:1.18-bullseye

WORKDIR /app

COPY . . 
RUN go mod download 

RUN go build -o /goapp

EXPOSE 9000

CMD [ "/goapp" ]