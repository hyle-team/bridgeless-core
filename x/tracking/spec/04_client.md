# Client

## Query


### CLI

A user can query and interact with the `tracking` module using the CLI.

The `query` commands allow users to query `tracking` state.

```sh
simd query tracking --help
```
#### positions

The `positions` command allow users to query all positions

```sh
simd query tracking positions [flags]
```
Example Output:
```ssh
pagination:
  next_key: null
  total: "2"
positions:
- address: 0x078c1353f521c4172cF45703F50557f976de6881
  amount: "800000000000000000"
  recipient: 0x078c1353f521c4172cF45703F50557f976de6881
- address: 0x45337D9891E0fB7Ab2C14451a80C9ea0D40dbC50
  amount: "800000000000000000"
  recipient: 0x45337D9891E0fB7Ab2C14451a80C9ea0D40dbC50
```

The `position [address]` command allow users to query all positions for address

```sh
simd query tracking positions [address] [flags]
```
Example Output:
```ssh
position:
  address: 0x078c1353f521c4172cF45703F50557f976de6881
  amount: "800000000000000000"
  recipient: 0x078c1353f521c4172cF45703F50557f976de6881
```

#### params

The `params` command allow users to query tracking module params

```ssh
simd query tracking params [flags]
```

Example Output: 

```ssh
params:
  borrow_event_name: Borrowed
  contract_address: 0x09e240ee85Af8A051eBBAA35a16e2D795Bbc78a6
  liquidation_event_name: Liquidated
  liquidator_address: bridge103n4cmjt2je8nqcxg9y9desyhy6m57u52kkuc4
  oracle_address: oracle_address
  sender_address: 0x7C675c6e4b54b2798306414856e604B935bA7b94
```

### HTTP

A user can query and interact with the `tracking` module using the HTTP.

#### positions

The `positions` endpoint allow users to query all nfts
```
/cosmos/tracking/positions
```
Example:

```
curl -X GET "http://localhost:1317/cosmos/tracking/positions" -H  "accept: application/json"
```

Example Output:
```
{
  "positions": [
    {
      "address": "0x078c1353f521c4172cF45703F50557f976de6881",
      "recipient": "0x078c1353f521c4172cF45703F50557f976de6881",
      "amount": "800000000000000000"
    },
    {
      "address": "0x45337D9891E0fB7Ab2C14451a80C9ea0D40dbC50",
      "recipient": "0x45337D9891E0fB7Ab2C14451a80C9ea0D40dbC50",
      "amount": "800000000000000000"
    }
  ],
  "pagination": {
    "next_key": null,
    "total": "2"
  }
}
```

The `position/{address}` endpoint allow users to query all positions for address
```
/cosmos/tracking/positions/{address}
```
Example:

```
curl -X GET "http://localhost:1317/cosmos/tracking/positions/{address}" -H  "accept: application/json"
```

Example Output:
```
{
  "position": {
    "address": "0x078c1353f521c4172cF45703F50557f976de6881",
    "recipient": "0x078c1353f521c4172cF45703F50557f976de6881",
    "amount": "800000000000000000"
  }
}
```

#### params

The `params` command allow users to query tracking module params

Example:
```
curl -X GET "http://localhost:1317/cosmos/tracking/params" -H  "accept: application/json"
```

Example Output:

```ssh
{
  "params": {
    "borrow_event_name": "Borrowed",
    "contract_address": "0x09e240ee85Af8A051eBBAA35a16e2D795Bbc78a6",
    "sender_address": "0x7C675c6e4b54b2798306414856e604B935bA7b94",
    "liquidator_address": "bridge103n4cmjt2je8nqcxg9y9desyhy6m57u52kkuc4",
    "liquidation_event_name": "Liquidated",
    "oracle_address": "oracle_address"
  }
}
```