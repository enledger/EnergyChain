# EnergyChain - A Blockhain for Energy & Utilities Data
PoA Private Chain of EnLedger and it's Corporate Partners - Tracks shares and performs energy-utility functionalities

**EnergyChain** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

# EnergyChain Consensus: Proof-of-Authority (PoA) Ruleset

EneryChain is a PRIVATE blockchain meant for industrial tracking of data related to the energy industry including but not limited to:

Utilities and Appliance Usage, Home and Building Performance Data, Electrical Grid & Microgrid Monitoring / Optimization / Control

To become a **Validator** you must be a corporate partner of energychain, nodes will mainly be hosted at EnLedger-sponsored server locations or microgrid control sites, but may be hosted by computing partners.

There is NO BLOCK REWARD and NO STAKING TOKEN INFLATION. EnergyChain runs on a **proof-of-authority** consensus protocol SOME transactions are processed by the validators for FREE such as voting and security-token transactions.

# EnergyChain Tokenomics

**Token: energystake**, <u>total token supply: 100</u>

> all energystake token are owned by EnLedger Corporation, and DELEGATED to validators to effect PoA consensus under the standard Cosmos-SDK Delegated Proof-of-Stake (DPoS) consensus ruleset.

**Token: eecoin**, total token supply (unlimited, minted 10k at a time as needed)

> this token is a GAMING token which is NOT SOLD OR REDEEMED BY ENLEDGER CORP., it is given as a reward for in-game activities, achievements, and online tournament prizes within the **EnergyLand Metaverse**.

> you can get some EECoin by accessing the **faucet**, at the EnLedger *Virtual Watering Hole* fountain (todo: add link), or by accessing the testnet / mainnet faucet (todo: add instructions)

**Token: EnergyAsset_A1 (EA_A1)**, <u>token supply 100</u> (supply increases when shares issued)

> a placeholder for now, EA_A1 token is a SECURITY TOKEN, linked to CleanGrid Investment Trust's FIRST portfolio offering, a development fund for GREEN ENERGY ASSET & MICROGRID development, deployment, maintenance, and requisitioning (operated and managed by CleanGrid Holdings, LLC)

**Token: Multifamily_A1 (M_A1)**, <u>total token supply 100</u> (supply increases when shares issued)

> a placeholder for now, M_A1 token is a SECURITY TOKEN, linked to CleanGrid Investment Trust's SECOND (planned) portfolio offering, a development fund for MULTIFAMILY real estate development, maintenance, aquisitions and sales (operated and managed by CleanGrid Holdings, LLC)

**Token: Commercial_A1 (C_A1)**, <u>total token supply 100</u> (supply increases when shares issued)

> a placeholder for now, C_A1 token is a SECURITY TOKEN, linked to CleanGrid Investment Trust's THIRD (planned) portfolio offering, a development fund for COMMERCIAL real estate development, maintenance, aquisitions and sales (operated and managed by CleanGrid Holdings, LLC)

**Token: VacationLux_A1 (VL_A1)**, <u>total token supply 100</u> (supply increases when shares issued)

> a placeholder for now, VL_A1 token is a SECURITY TOKEN, linked to CleanGrid Investment Trust's THIRD (planned) portfolio offering, a development fund for VACATION & LUXURY real estate development, maintenance, aquisitions and sales (operated and managed by CleanGrid Holdings, LLC)

# NFTS on EnergyChain

NFTs on EnergyChain are use as registry objects notarizing ownership and info related to utilities & appliances, digital-twin homes and businesses, virtual objects in the EnergyLand metaverse.

Access to PRIVATE data linked to publicly-registered assets represented as NFTs on EnergyChain is affected via **token-controlled-access** to off-network resources stored on the PRIVATE EnLedger dataserver network or complementary private databases or "data lake" repositories run by EnLedger corporate partners.

EnLedger's data network provides a *privacy-shield** to prevent your data from being sold or monetized without your consent, and OPTIONALLY allow you to monetize your own datastreams for your own benefit, and control which commercial entities get access to it, how much access and what data access they get, and for what purposes (and for what price).

NFTs on EnergyChain may also represent energy assets such as solar / battery installations, community microgrids, generators, phase controllers, boilers and hot-water / underground heating systems installed at CleanGrid-sponsored multifamily real estate developments, and be used to requisition, dispatch, monitor, optimize, and control these assets and networks.

NFTs on EnergyChain use the standard Cosmos SDK [NFT module](https://docs.cosmos.network/v0.47/modules/nft) (note: do we need to use a different custom module, maybe from the irisnet project?)

# to run an EnergyChain TESTNET node:

git clone https://www.github.com/enledger/energychain

make sure you have installed the dependencies: go, ignite, node

a Dockerfile-ubuntu is provided for running the testnet node under docker, customize top line ubuntu version

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your energychain testnet node (in development)

# Self-Hosted Wallet Functionality

to use the EnergyChain self-hosted Vue.js based web app wallet (to connect Keplr wallet in the browser)

Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).





