##
# BUILD CONTAINER
##

FROM goreleaser/goreleaser:v0.142.0 as builder

WORKDIR /build

COPY . .
RUN \
apk add --no-cache make ca-certificates ;\
make build-linux-amd64

##
# RELEASE CONTAINER
##

FROM scratch

WORKDIR /

COPY --from=builder /build/dist/go-ebsnvme_linux_amd64/go-ebsnvme /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/go-ebsnvme"]
CMD [""]
