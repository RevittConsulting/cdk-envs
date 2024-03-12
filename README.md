# cdk-envs

![CI](https://github.com/RevittConsulting/cdk-envs/actions/workflows/CI.yml/badge.svg)

cdk-envs is a dockerized app, in development.

***

## Pre-requisites / Dependencies

- [Make](https://www.gnu.org/software/make/)
- [Docker](https://www.docker.com/) + [Docker Compose](https://docs.docker.com/compose/)
- [Node.js](https://nodejs.org/en/)
- [npm](https://www.npmjs.com/)
- [Go](https://golang.org/)

***

## Config

Set up a config file in the `/api` directory of the project called chains.yaml. There is an example of this config in the root. This must be put in the api folder.

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

## Run the app

```bash
git clone https://github.com/RevittConsulting/cdk-envs
```

You can mount your data directory that contains your mdbx.dat data files. Do this by placing your data in `/data` in the project root.

To build and run the containers:

```bash
make cdk-envs
```

***