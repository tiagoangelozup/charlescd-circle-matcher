FROM istio/proxyv2:1.13.2

ENV GO_VERSION=1.17.8

RUN envoy --version
RUN apt update && apt install make tar golang-go -y
RUN mkdir /src

WORKDIR /src

ENTRYPOINT ["make", "test"]
