FROM golang:1.17
WORKDIR /Backend
COPY . .
RUN go env -w GO111MODULE=off
RUN go get github.com/gorilla/mux
RUN go get github.com/rs/cors
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go get go.mongodb.org/mongo-driver/bson
EXPOSE 3000
CMD ["go","run","Servidor.go"]
