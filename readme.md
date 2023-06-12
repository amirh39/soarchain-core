# Soarchain
**Soarchain** is multi-purpose Layer 1 network that enables a token-incentivized data and connectivity infrastructure for mobility.
​Soarchain is a global, distributed network that create a public data infrasturcture and connectivity layer for mobility members. 
Car miners - MOTUS - produce and are compensated in $MOTUS, the native cryptocurrency of the Soarchain blockchain. The Soarchain blockchain is a new, open source, public blockchain created entirely to incentivize the creation of physical, decentralized mobility network.

*Soarchain* is:
* Appchain, a single purpose blockchain designed for mobility.
* Interchain for Mobility, interoperable with other chains due to the Cosmos SDK.
* Gateway of Web3 for Mobility that brings existing Web3 applications to vehicles and mobility space.
* Infrastructure Independent with direct Vehicle-to-Vehicle communication.

### Explorer
- [Soarchain Explorer](https://explorer.soarchain.com/soar)

## Get started
cd soarchain-core <br />
Ignite chain build <br />
./run_makefile.sh  <br />
soarchaind start 

if soarchand command is not available; <br />
export GOPATH=$HOME/go <br />
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

### Configure

Your local soarchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/soarchain@latest! | sudo bash
```
`username/soarchain` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more
- [Light Paper](https://www.soarchain.com/lightpaper)
- [Official Documentation](https://docs.soarchain.com/)


---
## Development Tutorial

Key objects over chain:

<details><summary>Click to expand</summary>

- [VRF](https://github.com/soar-robotics/soarchain-core/blob/dev/x/poa/keeper/createVRF.go)
  - [Concept Introduction](#basic-introduction)
  - [Multiplier](#multiplier)
  - [VRF Data](#basic-introduction)
- [Runner](https://github.com/soar-robotics/soarchain-core/blob/dev/x/poa/keeper/msg_server_gen_runner.go)
  - [Concept Introduction](#basic-introduction)
- [Challenger](https://github.com/soar-robotics/soarchain-core/blob/dev/x/poa/keeper/msg_server_gen_challenger.go)
  - [Concept Introduction](#basic-introduction)
- [Runner Challenge](https://github.com/soar-robotics/soarchain-core/blob/dev/x/poa/keeper/msg_server_runner_challenge.go)
  - [Concept Introduction](#basic-introduction)

</details>

## VRF

#### Concept Introduction

A: A Verifiable Random Function (VRF) is the public-key version of a keyed cryptographic hash. Only the holder of the secret key can compute the hash, but anyone with the public key can verify the correctness of the hash. VRFs are useful for preventing enumeration of hash-based data structures.

#### Multiplier

When we were designing the formula to randomly generate a cool down time between challenges, we said we need a constant value too. The constant value will store by the Multiplier property.

#### VRF Data

Implementing VRF module on cosmos SDK needs to store some VRF values. 
- Index/Messages
- Proof/Publickey
- VRV/FinalVRV
- Multiplier

## Runner

#### Concept Introduction

A Runners are responsible for receiving V2N messages from the Motus and then broadcast them to the Motus devices nearby.

## Challenger

#### Concept Introduction

Challengers are responsible for challenging this receive & broadcast operation and verify if it was done correctly or not.

## Runner Challenge

#### Concept Introduction

Runner challenge is a service we use to reward/punish v2n challenges based on challenge results. In the case a malicious activity breaks our V2N network and prevents messages to be broadcasted or just publish wrong data to the network we would punish otherwise we would reward.



