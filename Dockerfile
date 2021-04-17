FROM golang:alpine
WORKDIR /src
COPY ./src .
COPY ./static /static
RUN go build -o /bin/sswc .
ENTRYPOINT [ "/bin/sswc" ]
EXPOSE 8080
