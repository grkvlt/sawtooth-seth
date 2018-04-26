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

# Description:
#   Builds the environment needed to build the Sawtooth Seth docs
#   Running the image will put the Sawtooth Seth docs in
#   sawtooth-seth/docs/build on your local machine.
#
# Build:
#   $ cd sawtooth-seth
#   $ docker build . -f docs/Dockerfile -t seth-build-docs
#
# Run:
#   $ cd sawtooth-seth
#   $ docker run -v $(pwd)/docs:/project/sawtooth-seth/docs seth-build-docs

FROM sawtooth-seth-tp AS tp
FROM sawtooth-seth-rpc AS rpc
FROM sawtooth-seth-cli AS cli

FROM ubuntu:xenial

COPY --from=cli /bin/seth /bin
COPY --from=tp /bin/seth-tp /bin
COPY --from=rpc /bin/seth-rpc /bin

ENV PATH=$PATH:/go/bin:/project/sawtooth-seth/bin:/project/sawtooth-seth/cli/bin:/project/sawtooth-seth/common/bin:/project/sawtooth-seth/processor/bin:/project/sawtooth-seth/rpc/bin:/protoc3/bin:/root/.cargo/bin:/usr/lib/go-1.9/bin

RUN apt-get update \
 && apt-get install -y -q \
    curl \
    dvipng \
    gcc \
    git \
    libssl-dev \
    libzmq3-dev \
    make \
    openssl \
    python3 \
    software-properties-common \
    sudo \
    texlive-fonts-recommended \
    texlive-latex-base \
    texlive-latex-extra \
    texlive-latex-recommended \
    unzip \
    zip \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/* \
 && pip3 install \
    sphinx==1.5.6 \
    sphinxcontrib-httpdomain \
    sphinxcontrib-openapi \
    sphinx_rtd_theme

COPY . /project/sawtooth-seth

WORKDIR /project/sawtooth-seth/docs
CMD make html latexpdf
