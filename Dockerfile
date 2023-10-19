FROM golang:1.18-alpine as buildStage
WORKDIR /app
COPY . .
RUN go build -o /rest ./cmd/rest
RUN go build -p /grpc ./cmd/grpc

FROM alpine
COPY --from=buildStage ./app/rest /rest
CMD [ "rest" ]

FROM alpine
COPY --from=buildStage ./app/grpc /grpc
CMD [ "grpc" ]