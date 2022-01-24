FROM golang:latest as builder
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o food-delivery .

FROM alpine:latest
WORKDIR /app/
COPY --from=builder /app .
CMD ["/app/food-delivery"]