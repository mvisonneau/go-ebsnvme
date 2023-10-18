FROM busybox:1.36-glibc

WORKDIR /

COPY go-ebsnvme /usr/local/sbin/

ENTRYPOINT ["/usr/local/sbin/go-ebsnvme"]
CMD [""]
