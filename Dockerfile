FROM img.paradise-soft.com.tw/pds-backend/golang:1.16.15-alpine3.15 AS go-builder
RUN apk add git
# 走 go mod vendor 模式
ENV CGO_ENABLED=0 GOFLAGS=-mod=vendor
WORKDIR /build
# cache 不常變動的檔案 cmd/ vendor/ go.mod go.sum 加速 build image 用
COPY cmd ./cmd
COPY vendor ./vendor
COPY go.mod .
COPY go.sum .
# 載入業務邏輯及設定
COPY service ./service
COPY conf.d ./conf.d
ARG buildVersion
ARG buildCommitID
RUN go build -ldflags \
    " \
    -X 'gitlab.geax.io/pdsbe/central-admin.BuildVersion=${buildVersion}' \
    -X 'gitlab.geax.io/pdsbe/central-admin.BuildCommitID=${buildCommitID}' \
    " \
    -o mall /build/cmd/

# 只複製執行時所需檔案，降低 image 大小
FROM reg.paradise-soft.com.tw:5000/alpine:3.12
RUN apk update && \
    apk upgrade && \
    apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=go-builder /build/mall /app/mall
COPY --from=go-builder /build/conf.d /app/conf.d
ENTRYPOINT ["./mall"]
