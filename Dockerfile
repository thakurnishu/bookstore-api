FROM golang:1.22rc1-alpine3.19 AS build
WORKDIR /app
COPY . /app/
RUN go build -o bookstore .

FROM alpine:latest as security_provider
RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot


FROM scratch
COPY --from=security_provider /etc/passwd /etc/passwd
USER nonroot
COPY --from=build /app/bookstore /app/bookstore
WORKDIR /app
EXPOSE 3000
CMD [ "./bookstore" ] 