FROM golang:1.13-alpine as builder
RUN apk add --no-cache --virtual .build-deps \
bash \
gcc \
git \
musl-dev
RUN mkdir build
COPY . /build
WORKDIR /build
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o ReportHealCheck main.go
RUN adduser -S -D -H -h /build ReportHealCheck
USER ReportHealCheck
FROM scratch
COPY --from=builder /build/ReportHealCheck /app/
WORKDIR /app
ENTRYPOINT ["./ReportHealCheck"]