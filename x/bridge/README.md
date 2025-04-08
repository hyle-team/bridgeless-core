# `x/bridge`

## Abstract
Bridge module is developed to store such bridge data on-chain.
It is used by bridge to get up-to-date info about all chains, tokens and transfers.

---

## State
### Params

Bridge module params contains next fields:
- Module admin address - module admin is responsible for making changes in parties lists and thresholds, also only him
can add new info about chains, tokens and transfers
- Parties list - parties list is the list of active validators which are parts of TSS protocol.
- Newbies list - list of validators having enough delegation to become a part of TSS.
- Goodbye list - list of validators failed requirements and are going to be kicked removed from active parties list.
- Blacklist list - list of parties banned from being a part of TSS.
- Staking threshold - a value that validator has to stake to be eligible to be a part of TSS.
- Tss threshold - minimum required number of signers.

Definition:

```protobuf
// Params defines the parameters for the module.
message Params {
  string module_admin = 1;
  repeated Party parties =2;
  repeated Party newbies = 3;
  repeated Party goodbye_list = 4;
  repeated Party blacklist = 5;
  string stake_threshold = 6;
  uint32 tss_threshold = 7;
}
```

Example: 

```json
{
    "module_admin": "bridge1...",
    "parties": [{
      "address": "bridge1..."
    }
    ],
    "newbies": [
      {
        "address": "bridge1..."
      }
    ],
    "goodbye_list": [],
    "blacklist": [],
    "stake_threshold": "1000000000000",
    "tss_threshold": 5
  }
```

### Chain
**Chain** defines necessary chain properties used during bridging process.

Definition:

```protobuf
message Chain {
  string id = 1;
  ChainType type =2;
  // bridge_address is the address of the bridge contract on the chain
  string bridge_address = 3;
  // operator is the address of the operator of the bridge contract
  string operator = 4;
}
```

Example: 
```json
{
      "id": "0",
      "type": "BITCOIN",
      "bridge_address": "addr...",
      "operator": "addr...."
}
```

### Token
**Token** defines necessary token properties like id or decimals used during bridging process.

Definition:
```protobuf
message TokenInfo {
  string address = 1;
  uint64 decimals = 2;
  string chain_id = 3;;
  uint64 token_id = 4;
  bool is_wrapped = 5;
}


message TokenMetadata {
  string name = 1;
  string symbol = 2;
  string uri = 3;
}

message Token {
  uint64 id = 1;
  TokenMetadata metadata = 2 [(gogoproto.nullable) = false];
  // info is the token information on different chains
  repeated TokenInfo info = 3 [(gogoproto.nullable) = false];
}
```

Example:    
```json
{
      "id": "1",
      "metadata": {
        "name": "TESTNET",
        "symbol": "TEST",
        "uri": ""
      },
      "info": [
        {
          "address": "0x0000000000000000000000000000000000000000",
          "decimals": "18",
          "chain_id": "0",
          "token_id": "1",
          "is_wrapped": true
        },
        {
          "address": "0x0000000000000000000000000000000000000000",
          "decimals": "18",
          "chain_id": "0",
          "token_id": "0",
          "is_wrapped": false
        }
      ]
}
```

### Transaction
**Transaction** defines bridge transaction details.

Definition:

```protobuf
message Transaction {
  string deposit_chain_id = 1;
  string deposit_tx_hash = 2;
  uint64 deposit_tx_index = 3;
  uint64 deposit_block = 4;
  string deposit_token = 5;
  string deposit_amount = 6;
  string depositor = 7;
  string receiver = 8;
  string withdrawal_chain_id = 9;
  string withdrawal_tx_hash = 10;
  string withdrawal_token = 11;
  string signature = 12;
  bool is_wrapped = 13;
  string withdrawal_amount = 14;
}
```

Example:
```json
{
      "deposit_chain_id": "0",
      "deposit_tx_hash": "0x...",
      "deposit_tx_index": "0",
      "deposit_block": "0",
      "deposit_token": "0x0000000000000000000000000000000000000000",
      "deposit_amount": "1212",
      "depositor": "0x...",
      "receiver": "0x...",
      "withdrawal_chain_id": "1",
      "withdrawal_tx_hash": "",
      "withdrawal_token": "0x0000000000000000000000000000000000000000",
      "signature": "0x...",
      "is_wrapped": false,
      "withdrawal_amount": "1212"
    }
```

___

## Transactions
## RPC
___
### Chains
___
### InsertChain

**InsertChain** - adds new chain data to core.

Message example:
```protobuf
message MsgInsertChain {
  string creator = 1;
  Chain chain  = 2 [(gogoproto.nullable) = false];
}
```
___
### DeleteChain
**DeleteChain**-deletes chain data from core.

Message example:
```protobuf
message MsgDeleteChain {
  string creator = 1;
  string chain_id = 2;
}
```
___

### Tokens

___

### InsertToken

**InsertToken** - inserts token data to core.

Message example:

```protobuf
message MsgInsertToken {
  string creator = 1;
  Token token  = 2 [(gogoproto.nullable) = false];
}

```
___
### UpdateToken 
**UpdateToken** - updates queried token data on core.

Message example:

```protobuf
message MsgUpdateToken {
  string creator = 1;
  uint64 token_id = 2;
  TokenMetadata metadata = 3 [(gogoproto.nullable) = false];
}
```
___

### DeleteToken
**DeleteToken** - deletes token data from core.

Message example:

```protobuf
message MsgDeleteToken {
  string creator = 1;
  uint64 token_id = 2;
}
```

___
### Transactions
___

### SubmitTransactions

**SubmitTransactions** - stores bridge transactions data to core.

Message example:

```protobuf
message MsgSubmitTransactions {
  option (cosmos.msg.v1.signer) = "submitter";

  string submitter = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  repeated core.bridge.Transaction transactions = 2 [(gogoproto.nullable) = false];
}
```
___
### TokenInfo
___

### AddTokenInfo
**AddTokenInfo** - stores token info to core.

Message example:

```protobuf
message MsgAddTokenInfo {
  string creator = 1;
  TokenInfo info = 3 [(gogoproto.nullable) = false];
}
```
___

### RemoveTokenInfo

Message example:

```protobuf
message MsgRemoveTokenInfo {
  string creator = 1;
  uint64 token_id = 2;
  string chain_id = 3;
}
```
___
### Parties
___

### SetParties

Message example:

```protobuf
message MsgSetPartiesList {
  string creator =1;
  repeated Party parties = 2;
}
```
___

### SetNewbieParties

Message example:

```protobuf
message MsgSetNewbiesList {
  string creator =1;
  repeated Party parties = 2;
}
```
___

### SetGoodbyeParties

Message example:

```protobuf
message MsgSetGoodbyePartiesList {
  string creator =1;
  repeated Party parties = 2;
}
```
___

### SetBlacklistParties

Message example:

```protobuf
message MsgSetBlacklistPartiesList {
  string creator =1;
  repeated Party parties = 2;
}
```
___
### Threshold
___

### SetStakeThreshold

Message example:
```protobuf
message MsgSetStakeThreshold {
  string creator =1;
  string amount = 2;
}
```
___

### SetTssThreshold

Message example:
```protobuf
message MsgSetTssThreshold {
  string creator =1;
  uint32  amount = 2;
}
```
___

## CLI
___

### Transactions
___
### Chains
___
### InsertChain

**InsertChain** - adds new chain data to core.

```
bridgeless-cored tx bridge chains insert bridge1... chain.json
```
___
### DeleteChain
**DeleteChain**-deletes chain data from core.

```
 bridgeless-cored tx bridge chains remove bridge1... 1
```
___

### Tokens

___

### InsertToken

**InsertToken** - inserts token data to core.


```
 bridgeless-cored tx bridge tokens insert bridge1... token.json
```
___
### UpdateToken
**UpdateToken** - updates queried token data on core.

```
bridgeless-cored tx bridge tokens update bridge1... token.json
```
___

### DeleteToken
**DeleteToken** - deletes token data from core.

```
 bridgeless-cored tx bridge tokens remove bridge1... 1
```

___
### Transactions
___

### SubmitTransactions

**SubmitTransactions** - stores bridge transactions data to core.

```
bridgeless-cored tx bridge transactions submit bridge1... tx.json
```
___
### TokenInfo
___

### AddTokenInfo
**AddTokenInfo** - stores token info to core.

```
bridgeless-cored tx bridge tokens add-info bridge1... info.json
```
___

### RemoveTokenInfo

```
bridgeless-cored tx bridge tokens remove-info bridge1... 1 1
```
___
### Parties
___

### SetParties

```
 bridgeless-cored tx bridge parties set default bridge1... bridge1...,bridge1...
```
___

### SetNewbieParties

```
 bridgeless-cored tx bridge parties set newbies bridge1... bridge1...,bridge1...
```
___

### SetGoodbyeParties

```
 bridgeless-cored tx bridge parties set goodbye bridge1... bridge1...,bridge1...
```
___

### SetBlacklistParties

```
 bridgeless-cored tx bridge parties set blacklist bridge1... bridge1...,bridge1...
```
___
### Threshold
___

### SetStakeThreshold

```
bridgeless-cored tx bridge threshold set-stake-threshold bridge1... "1000"
```
___

### SetTssThreshold

```
 bridgeless-cored tx bridge threshold set-tss-threshold bridge1... 5
```
___
### Query
___

### Params

```
bridgeless-cored query bridge params
```

Response example:

```
params:
  blacklist: []
  goodbye_list: []
  module_admin: bridge1...
  newbies:
  - address: bridge1..., bridge1...
  parties: []
  stakeThreshold: "100000000"
  tssThreshold: 1
```
___
### Chains
___

### QueryChains

```
bridgeless-cored query bridge chains
```

Response example:
```
chains:
- bridge_address: m4...
  id: "1"
  operator: m4...
  type: BITCOIN
- bridge_address: Zx...
  id: "2"
  operator: Zx...
  type: ZANO
- bridge_address: 0x...
  id: "3"
  operator: 0x...
  type: EVM
- bridge_address: 0x...
  id: "4"
  operator: 0x...
  type: EVM
```
___

### QueryChainById

```
bridgeless-cored query bridge chain 1
```

Response example:
```
chain:
  bridge_address: m4...
  id: "1"
  operator: m4...
  type: BITCOIN
```
___
### Tokens
___
### QueryAllTokens

```
bridgeless-cored query bridge tokens
```

Response example:

```
tokens:
- id: "1"
  info:
  - address: 0x...
    chain_id: "3"
    decimals: "18"
    is_wrapped: true
    token_id: "1"
  - address: "0x..."
    chain_id: "3"
    decimals: "18"
    is_wrapped: false
    token_id: "1"
  metadata:
    name: TESTNET
    symbol: TEST
    uri: url
- id: "2"
  info:
  - address: "0x0000000000000000000000000000000000000000"
    chain_id: "4"
    decimals: "18"
    is_wrapped: false
    token_id: "2"
  - address: 0x0000000000000000000000000000000000000000
    chain_id: "5"
    decimals: "18"
    is_wrapped: true
    token_id: "2"
```

___
### QueryTokenById

```
 bridgeless-cored query bridge token 1
```

Response example:
```
token:
  id: "1"
  info:
  - address: 0x0000000000000000000000000000000000000000
    chain_id: "2"
    decimals: "18"
    is_wrapped: true
    token_id: "1"
  - address: "0x0000000000000000000000000000000000000000"
    chain_id: "3"
    decimals: "18"
    is_wrapped: false
    token_id: "1"
  metadata:
    name: TESTNET
    symbol: TEST
    uri: url
```
___

### QueryTokenPairs

```
bridgeless-cored query bridge token-pairs 1 0x0000000000000000000000000000000000000000 2
```

Response example:

```
    address: 0x0000000000000000000000000000000000000000
    chain_id: "2"
    decimals: "18"
    is_wrapped: true
    token_id: "1"
```
___

### QueryTokenInfo

```
bridgeless-core % bridgeless-cored query bridge token-info 1 0x...
```

Response example:
```
info:
  address: 0xcA14eE670F7E827cd558E14d6AE68cB1a6d2e269
  chain_id: "35442"
  decimals: "18"
  is_wrapped: true
  token_id: "1"
```
___
### Transactions
___

### QueryAllTransactions

```
bridgeless-cored query bridge transactions
```

Response example:

```
transactions:
- deposit_amount: "2"
  deposit_block: "1"
  deposit_chain_id: "1"
  deposit_token: "0x0000000000000000000000000000000000000000"
  deposit_tx_hash: 0x0000000000000000000000000000000000000000
  deposit_tx_index: "0"
  depositor: 0x0000000000000000000000000000000000000000
  is_wrapped: false
  receiver: 0x0000000000000000000000000000000000000000
  signature: 0x0000000000000000000000000000000000000000
  withdrawal_amount: "1212"
  withdrawal_chain_id: "1"
  withdrawal_token: 0x0000000000000000000000000000000000000000
  withdrawal_tx_hash: ""
```
___

### QueryTransactionById
```
bridgeless-cored query bridge transaction 0x/2/0x
```

Response example:

```
transactions:
- deposit_amount: "2"
  deposit_block: "1"
  deposit_chain_id: "1"
  deposit_token: "0x0000000000000000000000000000000000000000"
  deposit_tx_hash: 0x0000000000000000000000000000000000000000
  deposit_tx_index: "0"
  depositor: 0x0000000000000000000000000000000000000000
  is_wrapped: false
  receiver: 0x0000000000000000000000000000000000000000
  signature: 0x0000000000000000000000000000000000000000
  withdrawal_amount: "1212"
  withdrawal_chain_id: "1"
  withdrawal_token: 0x0000000000000000000000000000000000000000
  withdrawal_tx_hash: ""
```



