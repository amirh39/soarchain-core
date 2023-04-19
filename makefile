remove:
	rm -rf ~/.soarchain

build:
	ignite chain build

reset:
	soarchaind tendermint unsafe-reset-all

init:
	soarchaind init devnet-validator-apollo --chain-id soarchaindevnet

config:
	soarchaind config keyring-backend test
	soarchaind config chain-id soarchaindevnet

keys:
	soarchaind keys add client --recover
	soarchaind keys add challenger-runner --keyring-backend test --algo secp256k1
	soarchaind keys add soarMasterAccount --recover
	soarchaind keys add investorWallet --keyring-backend test --algo secp256k1
	soarchaind keys add airdropWallet --keyring-backend test --algo secp256k1
	soarchaind keys add strategicWallet --keyring-backend test --algo secp256k1
	soarchaind keys add communityWallet --keyring-backend test --algo secp256k1
	soarchaind keys add apollo --keyring-backend test --algo secp256k1

parameter_token_denomination:
	cat ~/.soarchain/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="umotus"' > ~/.soarchain/config/tmp_genesis.json && mv ~/.soarchain/config/tmp_genesis.json ~/.soarchain/config/genesis.json
	cat ~/.soarchain/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="umotus"' > ~/.soarchain/config/tmp_genesis.json && mv ~/.soarchain/config/tmp_genesis.json ~/.soarchain/config/genesis.json
	cat ~/.soarchain/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="umotus"' > ~/.soarchain/config/tmp_genesis.json && mv ~/.soarchain/config/tmp_genesis.json ~/.soarchain/config/genesis.json
	cat ~/.soarchain/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="umotus"' > ~/.soarchain/config/tmp_genesis.json && mv ~/.soarchain/config/tmp_genesis.json ~/.soarchain/config/genesis.json
	sed -i 's/\("masterAccount": "\)[^"]*\(".*\)/\1'"$$(soarchaind keys show soarMasterAccount -a)"'\2/' ~/.soarchain/config/genesis.json

	cat ~/.soarchain/config/genesis.json | jq '.app_state["poa"]["masterKey"]["masterCertificate"]="-----BEGIN CERTIFICATE-----\nMIIB4TCCAYegAwIBAgIQTylBUpEkZd8CaYHSaLbBHzAKBggqhkjOPQQDAjBIMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMB4XDTIzMDMzMDA2NTUxNVoXDTQ4MDMzMDA2NTUxNVowSDEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEoMCYGA1UEAwwfU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLaCOXbFw/dRJXzXtvhSFWt92aUkdwRZPLmJWZFBFX55+XIDQsCGsQeMmU4pqsnXEB4/r842uYUinWsdzg4xUoqjUzBRMB0GA1UdDgQWBBRqxTRE6ZPuogp88TrNw1cwAYyPMjAfBgNVHSMEGDAWgBRqxTRE6ZPuogp88TrNw1cwAYyPMjAPBgNVHRMBAf8EBTADAQH/MAoGCCqGSM49BAMCA0gAMEUCIAHpI8Y6zPLaitMOGNAzzDAKb0PJw2r49vjzkFl5TIGPAiEArPJTReSmEnUJWFTcEIuYoWcRIBDI+GpianTVfX4uxNI=\n-----END CERTIFICATE-----"' > ~/.soarchain/config/tmp_genesis.json && mv ~/.soarchain/config/tmp_genesis.json ~/.soarchain/config/genesis.json

parameter_voting_period:
	cat <<< $$(jq '.app_state.gov.voting_params.voting_period = "20s"' ~/.soarchain/config/genesis.json) > ~/.soarchain/config/genesis.json

allocate_genesis_accounts:
	soarchaind add-genesis-account challenger-runner 10000000000000umotus --keyring-backend test
	soarchaind add-genesis-account client 10000000000000umotus --keyring-backend test
	soarchaind add-genesis-account soarMasterAccount 10000000000000umotus --keyring-backend test
	soarchaind add-genesis-account investorWallet 77425000000000umotus --keyring-backend test
	soarchaind add-genesis-account airdropWallet 47500000000000umotus --keyring-backend test
	soarchaind add-genesis-account strategicWallet 180500000000000umotus --keyring-backend test
	soarchaind add-genesis-account communityWallet 100700000000000umotus --keyring-backend test
	soarchaind add-genesis-account apollo 100000000000000umotus --keyring-backend test

sign_genesis_transaction:
	soarchaind gentx apollo 1000000umotus --keyring-backend test --chain-id soarchaindevnet

collect_genesis_tx:
	soarchaind collect-gentxs

validate_genesis:
	soarchaind validate-genesis

upgrade_proposal:
	soarchaind tx gov submit-proposal software-upgrade test1 --title upgrade --description upgrade --upgrade-height 35 --chain-id soarchaindevnet --from apollo --yes

vote_proposal:	
	soarchaind tx gov deposit 1 1000000umotus --from apollo --chain-id soarchaindevnet --yes
	soarchaind tx gov vote 1 yes --from apollo --chain-id soarchaindevnet --yes

ifdef start
	soarchaind start --log_level info --minimum-gas-prices=0.0001umotus
endif

all: reset init config keys parameter_token_denomination allocate_genesis_accounts sign_genesis_transaction collect_genesis_tx validate_genesis