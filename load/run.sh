   #!/bin/bash

IP=$1
PORT=$2
PREFIX=$3

if [ -z "${IP}" ]; then
    IP=127.0.0.1
fi

if [ -z "${PORT}" ]; then
    PORT=8002
fi

if [ -z "${PREFIX}" ]; then
    PREFIX=dloan-account
fi

echo "----------------------------------"
echo "Environment"
echo "----------------------------------"
echo "Host: ${IP}"
echo "Port: ${PORT}"
echo "PREFIX: ${PREFIX}"

echo "----------------------------------"
echo "Pre success case01"
echo "----------------------------------"

# TS_CreditTransfer_ISavings-CBS_001_Pre
hey -c 2 -n 200 -m POST \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-D case/01.json \
http://${IP}:${PORT}/${PREFIX}/v1/accounts


# hey -c 4 -n 100 -m GET http://localhost:8002
# hey -c 4 -n 4 -m POST http://localhost:8002