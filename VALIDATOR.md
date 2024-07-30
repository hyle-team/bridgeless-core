# Validators’ Guide


## Node (Bridgeless-core)
* Chain id: `bridge_2607-1`
* Bridgeless-core binary (linux/amd64): https://github.com/hyle-team/bridgeless-core/releases/tag/v12.1.7
* Genesis file: https://github.com/hyle-team/bridgeless-core/blob/develop/config/genesis.json

## Local installation


Also, you can build core from sources by yourself: use “https://github.com/hyle-team/bridgeless-core/releases/tag/v12.1.7” release information.

If you are using Ubuntu linux, please install musl-dev using `sudo apt install musl-dev` command to be able to use Alpine binary on your machine

To generate configs file download the bridgeless-core binary file and rename it to `bridgeless-core`. 

Set envs:

    export MONIKER_NAME=YOU_VALIDATOR_NAME
    export BRIDGELESS_HOME=YOU_CORE_HOME_PATH
    export BRIDGELESS_NODE=tcp://167.99.26.8:26657

Then to init node struct and generate configs u should execute this 

    bridgeless-cored init $MONIKER_NAME main --chain-id bridge_2607-1  --home=BRIDGELESS_HOME --keyring-backend test

Move the genesis file to `$BRIDGELESS_HOME/config`

Create validator private key:

    bridgeless-core keys add <key-name> --keyring-backend test --home=$BRIDGELESS_HOME


Dont forget to save the mnemonic and address. That address will be used for your validator staking. 
Send you address (bridge…) to `vl@distributedlab.com`.

Please, backup the following files and folders:

    $BRIDGELESS_HOME/config/priv_validator_key.json
    $BRIDGELESS_HOME/config/node_key.json
    $BRIDGELESS_HOME/keyring-test


Get the node id: 

    bridgeless-core tendermint show-node-id --home=$BRIDGELESS_HOME

The next step is a setting peers into config.toml. Find a `persistent_peers` field and set here at least your node info(`node_id@node_ip:26656`) and one or two the others.
In our case, past should past `0b872fe5d809863a94fe0e666a334ff06cd8d2cf@167.99.26.8:26656`. It's good to have at least 3 peers

To run use env variables:

    - name: DAEMON_NAME
      value: "bridgeless-core"

    - name: DAEMON_HOME
      value: $BRIDGELESS_HOME

    - name: DAEMON_ALLOW_DOWNLOAD_BINARIES
      value: "true"


To the simplest updates, run a chain using a `cosmovisor` tool.

    mkdir -p $DAEMON_HOME/cosmovisor/genesis/bin && cp YOU_STORE_CORE_BIN(name bridgeless-core) $DAEMON_HOME/cosmovisor/genesis/bin && cosmovisor run start --home=$BRIDGELESS_HOME --rpc.laddr tcp://0.0.0.0:26657

### Finishing with the validator
After receiving confirmation from ourside about token transfer, execute command to stake tokens
and become a validator. You need to stake exactly `100000000000000000000bridge` (10^18) that is equal to 1 ABRIDGE.

    bridgeless-core tx staking create-validator --amount 100000000000000000000bridge --commission-max-change-rate "0.01" --commission-max-rate "0.2" --commission-rate "0.1" --min-self-delegation "1" --details "Meet new bridgeless validator" --pubkey $(bridgless-core tendermint show-validator --home=$BRIDGELESS_HOME) --moniker $MONIKER_NAME --chain-id bridge_2607-1 --fees 0bridge --from $LOCAL_VALIDATOR_ADDRESS --home=$BRIDGELESS_HOME --node=$BRIDGELESS_NODE --keyring-backend=test



