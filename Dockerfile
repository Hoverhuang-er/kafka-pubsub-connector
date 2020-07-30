# Build Stage
FROM golang:1.15rc1-buster AS Build
COPY . .
RUN  make prepare && make release
# Runtime
FROM debian:10.4
COPY --from=Build connector /usr/bin
CMD ["./usr/bin/connector"]