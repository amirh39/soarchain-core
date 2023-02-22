remove:
	rm -rf ~/.soarchain

build:
	ignite chain build

reset:
	soarchaind tendermint unsafe-reset-all

init:
	soarchaind init devnet-validator-zeus --chain-id soarchaindevnet

config:
	soarchaind config keyring-backend test
	soarchaind config chain-id soarchaindevnet

keys:
	soarchaind keys add soarMasterAccount --keyring-backend test --algo secp256k1
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
	echo $$(soarchaind keys show soarMasterAccount -a)
	sed -i '' 's/\("masterAccount": "\)[^"]*\(".*\)/\1'"$$(soarchaind keys show soarMasterAccount -a)"'\2/' ~/.soarchain/config/genesis.json

parameter_voting_period:
	cat <<< $$(jq '.app_state.gov.voting_params.voting_period = "20s"' ~/.soarchain/config/genesis.json) > ~/.soarchain/config/genesis.json

allocate_genesis_accounts:
	soarchaind add-genesis-account soarMasterAccount 100000000soar --keyring-backend test
	soarchaind add-genesis-account investorWallet 77425000000000soar --keyring-backend test
	soarchaind add-genesis-account airdropWallet 47500000000000soar --keyring-backend test
	soarchaind add-genesis-account strategicWallet 180500000000000soar --keyring-backend test
	soarchaind add-genesis-account communityWallet 100700000000000soar --keyring-backend test
	soarchaind add-genesis-account zeus 100000000soar --keyring-backend test

sign_genesis_transaction:
	soarchaind gentx zeus 10000000soar --keyring-backend test --chain-id soarchaindevnet

collect_genesis_tx:
	soarchaind collect-gentxs

validate_genesis:
	soarchaind validate-genesis

upgrade_proposal:
	soarchaind tx gov submit-proposal software-upgrade test1 --title upgrade --description upgrade --upgrade-height 35 --chain-id soarchaindevnet --from zeus --yes

vote_proposal:	
	soarchaind tx gov deposit 1 10000000soar --from zeus --chain-id soarchaindevnet --yes
	soarchaind tx gov vote 1 yes --from zeus --chain-id soarchaindevnet --yes

ifdef start
	soarchaind start --log_level info --minimum-gas-prices=0.0001soar
endif

all: reset init config keys parameter_token_denomination allocate_genesis_accounts sign_genesis_transaction collect_genesis_tx validate_genesis