# syntax=docker/dockerfile:experimental

####################################
#  Setup env for build and checks  #
####################################
FROM golang:1.19 AS build

# Enable access to private repos
WORKDIR /app

COPY . .
RUN if [ ! -d "./vendor" ]; then make build.vendor; fi

ARG build_args
RUN GOOS=linux GOARCH=amd64 make build.local BUILD_ARGS="${build_args}"


################
#   Run step   #
################
FROM gcr.io/distroless/base

COPY --from=build /app/target/hvac /usr/bin/run

ENTRYPOINT ["/usr/bin/run"]
