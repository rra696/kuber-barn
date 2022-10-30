FROM ubuntu

WORKDIR /server

RUN apt-get update && apt-get install -y wget containerd

COPY bin/server .

EXPOSE 6050

ENTRYPOINT ["./server"]