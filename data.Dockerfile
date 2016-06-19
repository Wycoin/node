FROM alpine:3.4

RUN apk update
RUN apk add openssl

RUN apk update && \
    apk add bash rsync && \
    rm -rf /var/cache/apk/*


RUN adduser -D wycoin
USER wycoin

COPY bin/data.entrypoint /entry.sh
WORKDIR /bin

ENTRYPOINT ["/entry.sh"]
CMD ["true"]

