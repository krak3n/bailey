FROM golang:1.12-alpine
RUN apk update && apk add build-base git
WORKDIR /client.store.svc
COPY ./ /client.store.svc
RUN make client.store.svc

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /client.store.svc
COPY --from=0 /client.store.svc/bin/client.store.svc .
CMD ["./client.store.svc"]
