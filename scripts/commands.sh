#!/bin/bash

RED="\e[31m"
GREEN="\e[32m"
ENDCOLOR="\e[0m"

export ALICE_TX="--from alice --gas 200000 --fees 5000000tkyve --yes --chain-id kyve-local"
export JSON="--output json"

check_tx() {
  echo "🔍 Checking transaction..."
#  local cmd_code=$(jq -r '.code' | tr -d '.')
#  if [[ "$cmd_code" -ne 0 ]]; then
##    TODO: print raw_log
#    echo -e "${RED}✗ Command failed with code $cmd_code.${ENDCOLOR}"
#    return
#  fi

  local tx_hash=$(jq -r '.txhash' | tr -d '.')
  local code=$(jq -r '.code' | tr -d '.')
  if [[ -z "$tx_hash" || "$code" -ne 0 ]]; then
    echo "✏️ raw_log:"
    jq '.raw_log'
    echo -e "$RED✗ Transaction failed to broadcast.$ENDCOLOR"
    return
  fi

  local sleep_time=0.5
  local elapsed_seconds=0
  local progress_bar=""

  # Run a loop to check for tx_output
  while true; do
    local tx_output=$(kyved q tx "$tx_hash" --output json 2>/dev/null)
#    local tx_output=$(kyved q tx "$tx_hash" --output json)
    local code=$(echo "$tx_output" | jq -r '.code')

    if [[ -z "$code" ]]; then
      printf "\r⏳ Fetching transaction: %02d seconds | Progress: [%-30s]" "$elapsed_seconds" "$progress_bar"
      sleep $sleep_time
      ((elapsed_seconds++))
      progress_bar+="="
    else
      break
    fi
  done

  if [[ "$code" -eq 0 ]]; then
    echo "✏️ logs:"
    echo "$tx_output" | jq '.logs'
    echo ""
    echo -e "${GREEN}✅ Transaction $tx_hash successful!${ENDCOLOR}"
  else
    echo "✏️ raw_log:"
    echo "$tx_output" | jq '.raw_log'
    echo ""
    echo -e "${RED}✗ Transaction $tx_hash failed with code $code.${ENDCOLOR}"
  fi
}
export -f check_tx

echo "👋 Welcome to the KYVE CLI!"
echo "Examples:"
echo "kyved tx bank send alice kyve1jve70gkvgyvdnrxw4q7ry7vq2asu25ac0m48vk 1000tkyve \$ALICE_TX \$JSON | check_tx"
echo ""
echo "Funders:"
echo "kyved tx funders create-funder Alice \$ALICE_TX \$JSON | check_tx"
echo "kyved tx funders update-funder Alice \$ALICE_TX \$JSON | check_tx"
