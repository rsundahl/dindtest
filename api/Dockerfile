FROM alpine

RUN apk update && apk add bash curl

COPY ./api /

ENTRYPOINT ["./api"]

EXPOSE 8781
