# Copyright 2017 Intel Corporation
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# ------------------------------------------------------------------------------

FROM ubuntu:xenial AS cli

COPY . /project/sawtooth-seth

ENV GOPATH=/project/sawtooth-seth/cli \
    PATH=${PATH}:/project/sawtooth-seth/cli/bin:/project/sawtooth-seth/bin:/usr/lib/go-1.9/bin

RUN apt-get update \
 && apt-get install -y -q \
    curl \
    git \
    golang-1.9-go \
    libssl-dev \
    libzmq3-dev \
    openssl \
    python3 \
    python3-pip \
 && pip3 install grpcio-tools \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

RUN go get -u \
    github.com/golang/protobuf/proto \
    github.com/golang/protobuf/protoc-gen-go \
    github.com/btcsuite/btcd/btcec \
    github.com/grkvlt/sawtooth-go-sdk/... \
    github.com/jessevdk/go-flags \
    github.com/pelletier/go-toml \
    golang.org/x/crypto/ssh/terminal

RUN seth-protogen go

WORKDIR ${GOPATH}/src/seth_cli/cli
ENV GOPATH=${GOPATH}:/project/sawtooth-seth/burrow:/project/sawtooth-seth/common
RUN go build -o /project/sawtooth-seth/cli/bin/seth

FROM ubuntu:xenial

RUN apt-get update \
 && apt-get install -y -q \
    software-properties-common \
    curl \
    git \
    build-essential \
 && add-apt-repository ppa:ethereum/ethereum \
 && apt-get update \
 && apt-get install -y -q \
    solc \
 && curl -s -S -o /tmp/setup-node.sh https://deb.nodesource.com/setup_6.x \
 && chmod 755 /tmp/setup-node.sh \
 && /tmp/setup-node.sh \
 && apt-get install nodejs -y -q \
 && rm /tmp/setup-node.sh \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

RUN npm install \
    ethereumjs-abi \
    web3

COPY --from=cli /project/sawtooth-seth/cli/bin/seth /bin/seth
