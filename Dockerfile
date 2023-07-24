
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
# Stage 1: Build Rosetta
# Outputs: binary @ /rosetta/rosetta 
#---------------------------------------------------------------------
FROM golang:1.17-alpine as builder
WORKDIR /rosetta
RUN apk add --no-cache make gcc musl-dev linux-headers git

# Downnload dependencies & cache them in docker layer
COPY go.mod .
COPY go.sum .
RUN go mod download

# Build project
#  (this saves to redownload everything when go.mod/sum didn't change)
COPY . .
RUN go build --tags musl -o rosetta .


#---------------------------------------------------------------------
# Stage 2: Build Final container
# Integrates celo-blockchain & rosetta builds into a single container
# Outputs: rosetta & geth binaries on /usr/loca/bin
#---------------------------------------------------------------------
# geth mainnet
# FROM us.gcr.io/celo-org/geth:1.6.1
# TODO EN: for early testing
FROM us-west1-docker.pkg.dev/devopsre/celo-blockchain-public/geth:4a1e3fc3b9437cc45497a5da42aa797c0a5bbacf

ARG COMMIT_SHA

RUN apk add --no-cache ca-certificates
COPY --from=builder /rosetta/rosetta /usr/local/bin/
RUN echo $COMMIT_SHA > /version.txt
RUN mkdir /data
RUN mkdir /logs
RUN mkdir -p /var

EXPOSE 8080/tcp
EXPOSE 30503/tcp
EXPOSE 30503/udp

ENV ROSETTA_DATADIR="/data"
ENV ROSETTA_RPC_PORT="8080"
ENV ROSETTA_GETH_LOGFILE="/logs/celo.log"
ENV ROSETTA_GETH_IPCPATH="/var/geth.ipc"
ENV ROSETTA_GETH_BINARY="/usr/local/bin/geth"

ENTRYPOINT ["/usr/local/bin/rosetta"]
CMD ["run"]


