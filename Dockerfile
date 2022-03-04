FROM golang:1.17.4
WORKDIR /go/src/app
RUN mkdir HtmlBase
COPY . ./HtmlBase
WORKDIR ./HtmlBase/src
RUN go mod download
RUN go build
# RUN ls
EXPOSE 8080
CMD ["./HtmlBase"]