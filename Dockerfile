FROM golang:1.19.1-alpine
WORKDIR /app
COPY *.mod .
COPY *.sum .
RUN go mod download
COPY . .
EXPOSE 3030
RUN go build -o main .
CMD ["/app/main"]



