# Chain Development Utilities

![CI-CDU](https://github.com/RevittConsulting/chain-dev-utils/actions/workflows/CI-CDU.yml/badge.svg)
![CI-Frontend](https://github.com/RevittConsulting/chain-dev-utils/actions/workflows/CI-Frontend.yml/badge.svg)

Chain dev utils is a dockerized app with a set of tools for developers working in the crypto ecosystem.

The app currently has these services:

- Chains RPC API
    - This takes all the chains in the chains.yaml file and displays them as a card on the front end. You can then run RPC calls to the chains, and it will update via websockets.
- MDBX Viewer
    - This is a DB viewer for the MDBX database files that are stored in your mounted data directory. You can view the data and search through it.
- TxSender [WIP]
    - This is a tool that will send a transaction on a chain. It will then display the transaction hash and the receipt.
- Datastreamer [WIP]
    - This is under development.

***

## Pre-requisites / Dependencies

- [Make](https://www.gnu.org/software/make/)
- [Docker](https://www.docker.com/) + [Docker Compose](https://docs.docker.com/compose/)
- [Node.js](https://nodejs.org/en/)
- [npm](https://www.npmjs.com/)
- [Go](https://golang.org/)

***

## Config for chains

Set up a config file in the `/cdu` directory of the project called chains.yaml. There is an example of this config in the root. This must be put in the cdu folder. Add any amount of chains to this file.

```yaml
Chains:
  # Add any amount of chains here
  Cardona:
    NetworkName: "Polygon zkEVM Cardona Testnet"

    Etherscan: "https://sepolia.etherscan.io"
    Blockscout: "https://eth-sepolia.blockscout.com"
    Polygonscan: "https://cardona-zkevm.polygonscan.com"
    CurrencySymbol: "ETH"

    L1ChainId: 11155111
    L1RpcUrl: "https://rpc.sepolia.org"

    L2ChainId: 2442
    L2RpcUrl: "https://rpc.cardona.zkevm-rpc.com/"
    L2DataStreamUrl: "datastream.cardona.zkevm-rpc.com:6900"

    RollupManagerAddress: "0x32d33D5137a7cFFb54c5Bf8371172bcEc5f310ff"
    RollupAddress: "0xA13Ddb14437A8F34897131367ad3ca78416d6bCa"

    TopicsVerification: "0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966"
    TopicsSequence: "0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766"
```

***

## MDBX Viewer

MDBX viewer is a frontend tool that allows you to perform various operations on the MDBX database files that are stored in your data directory that you can mount through the argument `data`. You can view the data and search through it.

To set up your data, you must build CDU with the argument `data=/Path/to/your/data`.

***

## Run the app

```bash
git clone https://github.com/RevittConsulting/chain-dev-utils
```

To build and run the containers with a path to your data directory:

```bash
make cdu data=/Path/to/your/data
```

***