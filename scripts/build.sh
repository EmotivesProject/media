#!/bin/bash

docker build -f docker/media/Dockerfile -t ghcr.io/tombowyerresearchproject/media_api .
docker build -f docker/media/Dockerfile.server -t ghcr.io/tombowyerresearchproject/media_server .