# Currently only C++ development environment
# Includes emscripten SDK for WebAssembly compilation
FROM fedora:28

RUN dnf install -y gcc-c++ make cmake ragel boost-devel libubsan libasan
RUN dnf install -y git nodejs jre llvm
RUN dnf install -y xz

RUN git clone https://github.com/juj/emsdk.git /emsdk
WORKDIR /emsdk
RUN ./emsdk install latest
RUN source ./emsdk_env.sh

RUN curl --output /tmp/fastcomp.tar.gz https://codeload.github.com/kripken/emscripten-fastcomp/tar.gz/1.38.15
RUN tar -C / -xf /tmp/fastcomp.tar.gz && rm -f /tmp/fastcomp.tar.gz
RUN mv /emscripten-fastcomp-1.38.15 /fastcomp
RUN curl --output /tmp/fastcomp-clang.tar.gz https://codeload.github.com/kripken/emscripten-fastcomp-clang/tar.gz/1.38.15
RUN tar -C /fastcomp/tools -xf /tmp/fastcomp-clang.tar.gz && rm -f /tmp/fastcomp-clang.tar.gz
RUN mv /fastcomp/tools/emscripten-fastcomp-clang-1.38.15 /fastcomp/tools/clang
RUN curl --output /tmp/extra.zip https://codeload.github.com/llvm-mirror/clang-tools-extra/zip/master
RUN unzip /tmp/extra.zip -d /fastcomp/tools/clang/tools
RUN mv /fastcomp/tools/clang/tools/clang-tools-extra-master/ /fastcomp/tools/clang/tools/extra
RUN mkdir /fastcomp/build

WORKDIR /fastcomp/build

RUN cmake ..
RUN make

WORKDIR /

ENV LLVM_ROOT /fastcomp/build/bin
ENV EMSCRIPTEN /emsdk/emscripten/1.38.14

VOLUME /src
WORKDIR /src
