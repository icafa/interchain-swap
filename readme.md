# Interchain swap

## Overview

This codebase is the minimal demo to test interchain query and ica txs with osmosis chain.

### Developer Documentation

Interchain Queries developer docs can be found here
https://github.com/strangelove-ventures/async-icq/blob/main/README.md

Medium blog -
https://link.medium.com/a70uOK1cFwb

## Demo

### Start the instance of `interchain-swapd`

```bash
ignite chain serve -c sender.yml --reset-once
```

### Start the instance of osmosisd

```bash
sh start_osmosisd.sh
```

### Setup osmosis pool

Create a new liquidity pool and provide initial liquidity to it.

```sh
osmosisd tx gamm create-pool --pool-file="./pool.json"
```

```json
{
  "weights": "5stake,5uosmo",
  "initial-deposit": "100000000stake,1000000uosmo",
  "swap-fee": "0.002",
  "exit-fee": "0.0",
  "future-governor": "168h"
}
```

### Initialize connection and start hermes for ICA relayer

```bash
sh hermes_init.sh
sh network/hermes/start.sh
```

### Configure and start the icq relayer

```bash
rm -rf ~/.ignite/relayer
```

```bash
ignite relayer configure -a \
--source-rpc "http://localhost:26659" \
--source-faucet "http://localhost:4500" \
--source-port "interquery" \
--source-gasprice "0.0stake" \
--source-gaslimit 5000000 \
--source-prefix "cosmos" \
--source-version "icq-1" \
--target-rpc "http://localhost:26559" \
--target-faucet "http://localhost:4501" \
--target-port "icqhost" \
--target-gasprice "0.0stake" \
--target-gaslimit 300000 \
--target-prefix "cosmos"  \
--target-version "icq-1"
```

```bash
ignite relayer connect
```

### Send the query to "receiver" chain

```bash
interchain-swapd tx interquery send-query-osmosis-price channel-0 cosmos1ez43ye5qn3q2zwh8uvswppvducwnkq6w6mthgl --chain-id=sender --node=tcp://localhost:26659 --home ~/.sender --from alice
```

```bash
interchain-swapd query interquery query-price-state 1 --chain-id=sender --node=tcp://localhost:26659
```

### Create interchain account

```bash
interchain-swapd tx interquery register-ica channel-1 ""
```

### Query interchain account address

```bash
interchain-swapd query interquery query-ica [connection-id] [owner]
```

### Send tokens to interchain-account

```
osmosisd tx bank send bob $ICA_ADDRESS 10000000uosmo --keyring-backend=test
```

### Send ICA swap tx

```bash
interchain-swapd tx interquery swap-exact-amount-in [connection-id] [token-in] [token-out-min-amount] --pool-id --from --chain-id
```

Example

```bash
interchain-swapd tx interquery swap-exact-amount-in connection-0 100000stake 1400 --swap-route-pool-ids 1 --swap-route-denoms uosmo --from WALLET_NAME --chain-id osmosis-1
```
