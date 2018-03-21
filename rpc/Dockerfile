# Copyright 2018 Intel Corporation
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

# Build cli
FROM ubuntu:xenial AS cli

COPY . /project/sawtooth-seth

ENV GOPATH=/project/sawtooth-seth/cli \
    PATH=${PATH}:/project/sawtooth-seth/cli/bin:/project/sawtooth-seth/bin:/usr/lib/go-1.9/bin

RUN apt-get update \
 && apt-get install -y -q \
    git \
    libssl-dev \
    libzmq3-dev \
    openssl \
    pkg-config \
    python3 \
    unzip \
    python3-pip \
 && pip3 install grpcio-tools \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

RUN go get -u \
    github.com/golang/protobuf/proto \
    github.com/golang/protobuf/protoc-gen-go \
    github.com/btcsuite/btcd/btcec \
    github.com/grkvlt/sawtooth-sdk-go/... \
    github.com/jessevdk/go-flags \
    github.com/pelletier/go-toml \
    golang.org/x/crypto/ssh/terminal

RUN seth-protogen go

WORKDIR ${GOPATH}/src/seth_cli/cli
ENV GOPATH=${GOPATH}:/project/sawtooth-seth/burrow:/project/sawtooth-seth/common
RUN go build -o /project/sawtooth-seth/cli/bin/seth

# Build RPC
FROM ubuntu:xenial AS rpc

COPY . /project/sawtooth-seth

ENV PATH=${PATH}:/project/sawtooth-seth/bin

RUN apt-get update \
 && apt-get install -y -q --allow-downgrades \
    pkg-config \
    curl \
    gcc \
    unzip \
    openssl \
    libssl-dev \
    libzmq3-dev \
    python3 \
    python3-pip \
 && pip3 install grpcio-tools \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

RUN curl -OLsS https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip \
 && unzip protoc-3.5.1-linux-x86_64.zip -d protoc3 \
 && rm protoc-3.5.1-linux-x86_64.zip

RUN curl https://sh.rustup.rs -sSf > /usr/bin/rustup-init \
 && chmod +x /usr/bin/rustup-init \
 && rustup-init -y

ENV PATH=${PATH}:/root/.cargo/bin:/protoc3/bin:/project/sawtooth-seth/rpc/bin \
    CARGO_INCREMENTAL=0

WORKDIR /project/sawtooth-seth/rpc

RUN mkdir /project/sawtooth-seth/rpc/bin \
 && cargo build \
 && cp ./target/debug/rpc /project/sawtooth-seth/rpc/bin/seth-rpc

FROM ubuntu:xenial

RUN apt-get update \
 && apt-get install -y -q \
    libssl-dev \
    libzmq3-dev \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /

COPY --from=cli /project/sawtooth-seth/cli/bin/seth /bin
COPY --from=rpc /project/sawtooth-seth/rpc/bin/seth-rpc /bin
