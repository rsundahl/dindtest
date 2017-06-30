FROM alpine

RUN apk update && apk add bash curl

COPY ./server /

ENTRYPOINT ["./server"]

EXPOSE 8780
