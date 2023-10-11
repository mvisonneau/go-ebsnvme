FROM busybox:1.36-glibc

WORKDIR /

COPY go-ebsnvme /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/go-ebsnvme"]
CMD [""]
