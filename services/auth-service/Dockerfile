FROM golang:1.24.1-alpine as builder
WORKDIR /opt/app
COPY . .
RUN go mod tidy
RUN go build -o application
FROM golang:1.24.1-alpine
WORKDIR /opt/app
COPY --from=builder /opt/app/application /opt/app/
EXPOSE 3000
RUN chmod +x application
CMD [ "/opt/app/application" ]