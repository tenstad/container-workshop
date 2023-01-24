FROM golang:latest as build

WORKDIR /app

COPY src .

RUN CGO_ENABLED=0 go build

FROM scratch

COPY --from=build /app/app /app

CMD ["/app"]
