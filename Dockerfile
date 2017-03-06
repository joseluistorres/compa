FROM golang:1.6-onbuild
COPY . /Users/joseluistorres/golang/src/github.com/compa
RUN go get -d -v
RUN go install -v
ENV PORT=8000
EXPOSE 8000
ENV RETHINKDB_URL_PORT="172.17.0.2:28015"