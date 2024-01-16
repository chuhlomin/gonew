FROM --platform=${TARGETPLATFORM:-linux/amd64} golang:1.21 as build-env

WORKDIR /go/src/app
ADD . /go/src/app

RUN go test -mod=vendor -cover ./...
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -mod=vendor -o /go/bin/app


FROM --platform=${TARGETPLATFORM:-linux/amd64} gcr.io/distroless/static:966f4bd97f611354c4ad829f1ed298df9386c2ec
# latest-amd64 -> 966f4bd97f611354c4ad829f1ed298df9386c2ec
# https://github.com/GoogleContainerTools/distroless/tree/master/base

COPY --from=build-env /go/bin/app /app

CMD ["/app"]
