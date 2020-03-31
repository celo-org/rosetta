
# USAGE INSTRUCTIONS:
#
# Docker Container Label: gcr.io/celo-testnet/rosetta
#
# Local Build For testing: 
#    docker build -t gcr.io/celo-testnet/rosetta:$USER .
#
# Run rosetta
#    docker rm rosetta    # Removes rosetta named container
#    docker run --name rosetta gcr.io/celo-testnet/rosetta:$USER
#    # or to create & delete
#    docker run -rm gcr.io/celo-testnet/rosetta:$USER
#
# Build & Deploy:
#    export COMMIT_SHA=$(git rev-parse HEAD)
#    docker build --build-arg COMMIT_SHA=$COMMIT_SHA -t gcr.io/celo-testnet/rosetta:$COMMIT_SHA .
#    docker push gcr.io/celo-testnet/rosetta:$COMMIT_SHA


#---------------------------------------------------------------------
# Stage 0: Build Rust library
# Outputs: rust build output @ /bls-zexe/target/
#---------------------------------------------------------------------
FROM ubuntu:16.04 as rustbuilder

RUN apt update && apt install -y curl muRUN curl https://sh.rustup.rs -sSf | sh -s -- -y
ENV PATH=$PATH:~/.cargo/bin
RUN $HOME/.cargo/bin/rustup install 1.41.0 && $HOME/.cargo/bin/rustup default 1.41.0 && $HOME/.cargo/bin/rustup target add x86_64-unknown-linux-musl
COPY ./external/bls-zexe /bls-zexe
RUN cd /bls-zexe && $HOME/.cargo/bin/cargo build --target x86_64-unknown-linux-musl --release


#---------------------------------------------------------------------
# Stage 1: Build Rosetta
# Outputs: binary @ /rosetta/rosetta 
#---------------------------------------------------------------------
FROM golang:1.13-alpine as builder
WORKDIR /rosetta
RUN apk add --no-cache make gcc musl-dev linux-headers git



# Copy BLS library + rust build output
COPY external ./external
RUN mkdir -p external/bls-zexe/target/release/
COPY --from=rustbuilder /bls-zexe/target/x86_64-unknown-linux-musl/release/libepoch_snark.a external/bls-zexe/target/release/

# Downnload dependencies
# We use strategy depicted here: https://github.com/golang/go/wiki/Modules#how-do-i-download-modules-needed-to-build-specific-packages-or-tests
COPY go.mod .
COPY go.sum .
RUN go list ./...

# Build project
#  (this saves to redownload everything when go.mod/sum didn't change)
COPY . .
RUN go build -o rosetta .


#---------------------------------------------------------------------
# Stage 1: Build Final container
# Integrates celo-blockchain & rosetta builds into a single container
# Outputs: rosetta & geth binaries on /usr/loca/bin
#---------------------------------------------------------------------
FROM gcr.io/celo-testnet/geth:rosetta
ARG COMMIT_SHA

RUN apk add --no-cache ca-certificates
COPY --from=builder /rosetta/rosetta /usr/local/bin/
RUN echo $COMMIT_SHA > /version.txt
RUN mkdir /data
EXPOSE 8080/tcp
ENV ROSETTA_DATADIR="/data"
ENV ROSETTA_GETH="/usr/local/bin/geth"
ENV ROSETTA_GENESIS="/data/genesis.json"
ENTRYPOINT ["/usr/local/bin/rosetta"]
CMD ["serve", "local"]


