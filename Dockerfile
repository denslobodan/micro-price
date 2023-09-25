FROM golang:1.21.1-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o /pricefetcher

EXPOSE 3030

CMD ["/pricefetcher"]

# docker build --tag fetcher:1 .
#  docker run -p 3030:3030 -d fetcher:1