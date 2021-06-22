FROM golang:1.16-alpine as build
WORKDIR /work
COPY . /work
RUN go build main.go

FROM alpine
COPY --from=build /work/main /
CMD ./main
