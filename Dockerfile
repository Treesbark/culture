# To Build: docker build -f Dockerfile -t culture .
# To Run: docker run -p 8089:8089 -t culture

FROM golang:1.16

WORKDIR /
COPY . .
RUN go get -d github.com/gorilla/mux github.com/bitly/go-simplejson

EXPOSE 8089

CMD ["go","run","main.go"]