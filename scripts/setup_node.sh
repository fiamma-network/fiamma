#!/bin/bash
set -e

if [ $# -lt 1 ]; then
	echo "Usage: $0 <moniker>"
	exit 1
else
    MONIKER=$1
fi

NODE_HOME=$HOME/.fiamma
CHAIN_BINARY=fiammad
: ${CHAIN_ID:="fiamma-testnet-1"}

: ${PEER_ADDR="13.112.47.16,13.113.190.239"}

PEER_ARRAY=(${PEER_ADDR//,/ })

# You should ensure that the fiammad binary files have been correctly installed.
$CHAIN_BINARY comet unsafe-reset-all
$CHAIN_BINARY init $MONIKER \
    --chain-id $CHAIN_ID --overwrite

for ADDR in "${PEER_ARRAY[@]}"; do
    GENESIS=$(curl -f "$ADDR:26657/genesis" | jq '.result.genesis')
    if [ -n "$GENESIS" ]; then
        echo "$GENESIS" > $NODE_HOME/config/genesis.json;
        break;
    fi
done

PERSISTENT_PEERS=()

for ADDR in "${PEER_ARRAY[@]}"; do
    PEER_ID=$(curl -s "$ADDR:26657/status" | jq -r '.result.node_info.id')
    if [ -n "$PEER_ID" ]; then
        PERSISTENT_PEERS+=("$PEER_ID@$ADDR:26656")
    fi
done

CONFIG_STRING=$(IFS=,; echo "${PERSISTENT_PEERS[*]}")

$CHAIN_BINARY config set config p2p.persistent_peers "$CONFIG_STRING" --skip-validate