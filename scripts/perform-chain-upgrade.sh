#!/bin/bash

RED="\e[31m"
ENDCOLOR="\e[0m"

# Check if upgrade name is provided
if [ "$#" -ne 1 ]; then
    echo -e "$RED✗ Usage: perform-chain-upgrade.sh [upgrade_name] [wait_blocks (default=30)]$ENDCOLOR"
    exit 1
fi

# Check if wait_blocks is provided
if [ "$#" -eq 2 ]; then
    WAIT_BLOCKS=$2
else
    WAIT_BLOCKS=30
fi

# Check if binary env is set
if [ -z "$BINARY" ]; then
    echo -e "$RED✗ BINARY env not set. Please source commands.sh before running this script$ENDCOLOR"
    exit 1
fi

PROPOSAL='{
  "messages": [
    {
      "@type": "/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade",
      "authority": "kyve10d07y265gmmuvt4z0w9aw880jnsr700jdv7nah",
      "plan": {
        "name": "MY_UPGRADE_NAME",
        "height": "MY_UPGRADE_HEIGHT",
        "time": "MY_TIME",
        "info": ""
      }
    }
  ],
  "metadata": "ipfs://QmSbXwSgAnhhkDk5LtvgEwKUf6iUjww5FbMngYs86uGxdG",
  "deposit": "20000000000000MY_DENOM",
  "title": "Upgrade to MY_UPGRADE_NAME",
  "summary": "This is an upgrade to MY_UPGRADE_NAME",
  "expedited": false
}'

NAME=$1
HEIGHT=$($BINARY status | jq '.sync_info.latest_block_height' | tr -d '\"')
HEIGHT=`expr "$HEIGHT" + $WAIT_BLOCKS`

DENOM=$($BINARY q staking params $JSON | jq '.params.bond_denom' | tr -d '\"')

TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")

PROPOSAL="${PROPOSAL/MY_UPGRADE_HEIGHT/$HEIGHT}"
PROPOSAL="${PROPOSAL/MY_UPGRADE_NAME/$NAME}"
PROPOSAL="${PROPOSAL/MY_DENOM/$DENOM}"
PROPOSAL="${PROPOSAL/MY_TIME/$TIME}"
printf "$PROPOSAL" > /tmp/proposal.json

# Submit proposal
$BINARY tx gov submit-proposal /tmp/proposal.json $ALICE_TX $CHAINHOME $TESTBACKEND $TX $JSON | check_tx
echo "Submitted proposal"
echo ""

# Get proposal ID
ID=$($BINARY q gov proposals --page-limit 1 --page-reverse --proposal-status voting-period $JSON | jq '.proposals[0].id | tonumber')

# Vote on proposal
$BINARY tx gov vote $ID yes $ALICE_TX $CHAINHOME $TESTBACKEND $JSON | check_tx
echo "Voted on proposal: $ID"
echo "Scheduled upgrade for $HEIGHT"
