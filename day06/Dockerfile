FROM golang:1.17-alpine as Build
WORKDIR /src
ADD go/ ./

RUN go build -ldflags="-s -w" main.go
RUN #go build main.go

FROM scratch
WORKDIR /bin
COPY --from=Build /src/main /bin/main

CMD ["./main"]
