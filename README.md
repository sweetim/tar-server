# Tar Server 

[![CircleCI](https://circleci.com/gh/sweetim/tar-server.svg?style=svg)](https://circleci.com/gh/sweetim/tar-server)
[![Maintainability](https://api.codeclimate.com/v1/badges/c982a5e7ccbc5f56a877/maintainability)](https://codeclimate.com/github/sweetim/tar-server/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/c982a5e7ccbc5f56a877/test_coverage)](https://codeclimate.com/github/sweetim/tar-server/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/sweetim/tar-server)](https://goreportcard.com/report/github.com/sweetim/tar-server)

This server will package the selected folder into a `.tar` and to be downloaded into the user PC

### Motivation

This application is to develop to solve the annoying problem of transferring files from the stored data PC to the host PC by plugging USB hard drive. This method is slow and time consuming, besides when transfering large files into the USB hard drives, it will usually slow down to few `MB/s`. This application will solve this annoying problem and will constantly stream the data at maximum ethernet speed.

Now you could enjoy fast transfer of large data to your host PC :)

### Deployment using Docker

    docker run -it --rm -v <DIR_TO_SHARE>:/data timx/tar-server

### Deployment using Docker compose

    version: "3"

    services:
        tarserver:
            image: timx/tar-server
            ports:
                - "3000:3000"
            environment:
                - DIR_PATH=<DIR_TO_SHARE>
            volumes:
                - <DIR_TO_SHARE>:<DIR_TO_SHARE>:ro

### Build from source

    go build && ./tar-server
