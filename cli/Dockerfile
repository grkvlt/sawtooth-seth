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

ENV GOPATH=/project/sawtooth-seth/cli
ENV PATH=$PATH:/project/sawtooth-seth/cli/bin:/project/sawtooth-seth/bin:/usr/lib/go-1.9/bin

COPY . /project/sawtooth-seth

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
    github.com/btcsuite/btcd/btcec \
    github.com/golang/mock/gomock \
    github.com/golang/mock/mockgen \
    github.com/golang/protobuf/proto \
    github.com/golang/protobuf/protoc-gen-go \
    github.com/btcsuite/btcd/btcec \
    github.com/grkvlt/sawtooth-sdk-go/... \
    github.com/jessevdk/go-flags \
    github.com/pebbe/zmq4 \
    github.com/pelletier/go-toml \
    github.com/satori/go.uuid \
    golang.org/x/crypto/ssh/terminal

RUN go get github.com/hyperledger/sawtooth-sdk-go \
 && cd $GOPATH/src/github.com/hyperledger/sawtooth-sdk-go \
 && go generate

RUN seth-protogen go

WORKDIR ${GOPATH}/src/seth_cli/cli
ENV GOPATH=${GOPATH}:/project/sawtooth-seth/burrow:/project/sawtooth-seth/common
RUN go build -o /project/sawtooth-seth/cli/bin/seth

FROM ubuntu:xenial

RUN npm install \
    ethereumjs-abi \
    web3

COPY --from=cli /project/sawtooth-seth/cli/bin/seth /bin/seth
