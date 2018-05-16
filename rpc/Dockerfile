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

# Reference cli
FROM sawtooth-seth-cli as cli

# Build RPC
FROM ubuntu:xenial AS rpc

COPY . /project/sawtooth-seth

ENV PATH=${PATH}:/project/sawtooth-seth/bin

RUN apt-get update \
 && apt-get install -y -q --allow-downgrades \
    curl \
    gcc \
    git \
    libssl-dev \
    libzmq3-dev \
    openssl \
    pkg-config \
    python3 \
    python3-pip \
    unzip \
 && pip3 install grpcio-tools \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

RUN curl -OLsS https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip \
 && unzip protoc-3.5.1-linux-x86_64.zip -d protoc3 \
 && rm protoc-3.5.1-linux-x86_64.zip

RUN curl https://sh.rustup.rs -sSf > /usr/bin/rustup-init \
 && chmod +x /usr/bin/rustup-init \
 && rustup-init -y --default-toolchain 1.24.1

ENV PATH=${PATH}:/root/.cargo/bin:/protoc3/bin:/project/sawtooth-seth/rpc/bin \
    CARGO_INCREMENTAL=0

WORKDIR /project/sawtooth-seth
RUN cargo install --version 1.4.4 --force protobuf \
 && seth-protogen rust

WORKDIR /project
RUN git clone https://github.com/blockchaintp/sawtooth-core \
 && cd sawtooth-core \
 && git fetch --all \
 && git checkout 1-0-staging-01 \
 && ./bin/protogen rust \
 && cd sdk/rust \
 && cargo build

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

COPY --from=cli /bin/seth /bin
COPY --from=rpc /project/sawtooth-seth/rpc/bin/seth-rpc /bin
