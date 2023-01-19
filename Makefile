remove:
	rm -rf ~/.soarchain

reset:
	soarchaind tendermint unsafe-reset-all

init:
	soarchaind init devnet-validator-zeus --chain-id soarchaindevnet

config:
	soarchaind config keyring-backend test
	soarchaind config chain-id soarchaindevnet

keys:
	soarchaind keys add investorWallet --keyring-backend test --algo secp256k1
	soarchaind keys add airdropWallet --keyring-backend test --algo secp256k1
	soarchaind keys add strategicWallet --keyring-backend test --algo secp256k1
	soarchaind keys add communityWallet --keyring-backend test --algo secp256k1
	soarchaind keys add zeus --keyring-backend test --algo secp256k1

parameter_token_denomination:
	cat ~/.soarchain/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="soar"' > ~/.soarchain/config/tmp_genesis.json && mv ~/.soarchain/config/tmp_genesis.json ~/.soarchain/config/genesis.json
	cat ~/.soarchain/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="soar"' > ~/.soarchain/config/tmp_genesis.json && mv ~/.soarchain/config/tmp_genesis.json ~/.soarchain/config/genesis.json
	cat ~/.soarchain/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="soar"' > ~/.soarchain/config/tmp_genesis.json && mv ~/.soarchain/config/tmp_genesis.json ~/.soarchain/config/genesis.json
	cat ~/.soarchain/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="soar"' > ~/.soarchain/config/tmp_genesis.json && mv ~/.soarchain/config/tmp_genesis.json ~/.soarchain/config/genesis.json

allocate_genesis_accounts:
	soarchaind add-genesis-account investorWallet 77425000000000soar --keyring-backend test
	soarchaind add-genesis-account airdropWallet 47500000000000soar --keyring-backend test
	soarchaind add-genesis-account strategicWallet 180500000000000soar --keyring-backend test
	soarchaind add-genesis-account communityWallet 100700000000000soar --keyring-backend test
	soarchaind add-genesis-account zeus 10000000soar --keyring-backend test

sign_genesis_transaction:
	soarchaind gentx zeus 10000000soar --keyring-backend test --chain-id soarchaindevnet

collect_genesis_tx:
	soarchaind collect-gentxs

validate_genesis:
	soarchaind validate-genesis

ifdef start
	soarchaind start --log_level info --minimum-gas-prices=0.0001soar
endif

all: reset init config keys parameter_token_denomination allocate_genesis_accounts sign_genesis_transaction collect_genesis_tx validate_genesis