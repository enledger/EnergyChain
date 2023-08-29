# EnergyChain
PoA Private Chain of EnLedger and it's Corporate Partners - Tracks shares and performs energy-utility functionalities

# energychain
**energychain** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).


# to run a TESTNET node:

git clone https://www.github.com/enledger/energychain

make sure you have installed the dependencies: go, ignite, node

a Dockerfile-ubuntu is provided for running the testnet node under docker, customize top line ubuntu version

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your energychain testnet node (in development)

to use the EnergyChain self-hosted Vue.js based web app wallet (to connect Keplr wallet in the browser)

Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).





