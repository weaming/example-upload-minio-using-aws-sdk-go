#!/bin/bash

datapath="${1:-$HOME/data}"
echo $datapath
KEY="ZSCR0B4UJ2ZRQWAWUX0Z"
SECRET="I9PDoAzKuTAQsHCpRvsC4jEId8jwsnmvBfIWVMVq"
docker run -d -e "MINIO_ACCESS_KEY=$KEY" -e "MINIO_SECRET_KEY=$SECRET" \
    -p 9000:9000 \
    -p 9001:9001 \
    --name minio \
    -v "$datapath"/data:/data \
    -v "$datapath"/config:/root/.minio \
    minio/minio server /data --console-address ":9001"
echo MINIO_ACCESS_KEY=$KEY
echo MINIO_SECRET_KEY=$SECRET