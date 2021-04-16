FROM golang:alpine
WORKDIR /src
COPY ./src .
RUN go build -o /bin/sswc .
ENTRYPOINT [ "/bin/sswc" ]
EXPOSE 8080
