FROM --platform=${TARGETPLATFORM:-linux/amd64} golang:1.21 as build-env

WORKDIR /go/src/app
ADD . /go/src/app

RUN go test -mod=vendor -cover ./...
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-w -s" -mod=vendor -o /go/bin/app


FROM --platform=${TARGETPLATFORM:-linux/amd64} gcr.io/distroless/static:966f4bd97f611354c4ad829f1ed298df9386c2ec
# latest-amd64 -> 966f4bd97f611354c4ad829f1ed298df9386c2ec
# https://github.com/GoogleContainerTools/distroless/tree/master/base

LABEL name="..."
LABEL repository="http://github.com/.../..."
LABEL homepage="http://github.com/.../..."

LABEL maintainer="First Last <email@domain.com>"
LABEL com.github.actions.name="..."
LABEL com.github.actions.description="..."
LABEL com.github.actions.icon="file-text"
LABEL com.github.actions.color="purple"

COPY --from=build-env /go/bin/app /app

CMD ["/app"]
