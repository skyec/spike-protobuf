# Dockerfile for dev environment

FROM ubuntu:trusty
MAINTAINER Skye Cove <skye.cove@gmail.com>
RUN apt-get update && apt-get install -y --no-install-recommends \
 bzr \
 cmake \
 curl \
 g++ \
 git \
 make \
 man-db \
 mercurial \
 ncurses-dev \
 procps \
 python-dev \
 python-pip \
 ssh \
 sudo \
 tmux \
 unzip \
 vim \
 xz-utils \
 && rm -rf /var/lib/apt/lists/* \
 && pip install ipython \
 && git clone https://github.com/gmarik/Vundle.vim.git /root/.vim/bundle/Vundle.vim \
 && git clone https://github.com/Valloric/YouCompleteMe.git /root/.vim/bundle/YouCompleteMe \
 && cd /root/.vim/bundle/YouCompleteMe && git submodule update --init --recursive \
 && ./install.sh --clang-completer

COPY bashrc /root/.bashrc

RUN curl -O https://storage.googleapis.com/golang/go1.7.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.7.linux-amd64.tar.gz

RUN git config --global credential.helper 'cache --timeout=86400'

ENV GOPATH /go
ENV GOBIN /go/bin
ENV PATH /usr/local/go/bin:/go/bin:/installs/go_appengine:$PATH
ENV HOME /root
WORKDIR /go/src

RUN go version | grep go1.7

COPY vimrc /root/.vimrc
RUN vim +PluginInstall  +qall
RUN vim +GoInstallBinaries +qall

RUN mkdir /pbuff \
  && curl -LO https://github.com/google/protobuf/releases/download/v3.0.0/protoc-3.0.0-linux-x86_64.zip \
  && unzip -d /pbuff protoc-3.0.0-linux-x86_64.zip \
  && go get -u github.com/golang/protobuf/proto \
  && go get -u github.com/golang/protobuf/protoc-gen-go

ENV PATH /pbuff/bin:$PATH

CMD ["tmux", "-u2"]

# Run Docker with
# docker run -itv ~/yourworkspace/src:/go/src yourimage 

