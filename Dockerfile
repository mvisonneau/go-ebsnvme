FROM busybox:1.36-glibc

WORKDIR /

COPY go-ebsvnme /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/go-ebsvnme"]
CMD [""]
