# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM alpine:3.12.4
RUN sed -i 's!http://dl-cdn.alpinelinux.org/!https://mirrors.ustc.edu.cn/!g' /etc/apk/repositories && \
    apk --no-cache add tzdata curl bash vim && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk add --no-cache ca-certificates && \
    update-ca-certificates
WORKDIR /
COPY ./bin/kube-scheduler .
ENTRYPOINT ["/kube-scheduler"]
