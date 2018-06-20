From busybox:glibc

RUN mkdir -p /root/service
COPY ./helloworld /root/service/helloworld
RUN chmod +x  /root/service/helloworld

WORKDIR /root/service
ENTRYPOINT ["/root/service/helloworld"]
