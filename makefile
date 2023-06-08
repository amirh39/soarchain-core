remove:
	rm -rf ~/.soarchain

build:
	ignite chain build

reset:
	soarchaind tendermint unsafe-reset-all

init:
	soarchaind init hermes --chain-id soarchaindevnet

config:
	soarchaind config keyring-backend test
	soarchaind config chain-id soarchaindevnet

keys:
	soarchaind keys add hermes --keyring-backend test --algo secp256k1
	soarchaind keys add reserve --keyring-backend test --algo secp256k1
	soarchaind keys add soarMasterAccount --recover
	soarchaind keys add client --recover

parameter_token_denomination:
	cat ~/.soarchain/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="udmotus"' > ~/.soarchain/config/tmp_genesis.json && mv ~/.soarchain/config/tmp_genesis.json ~/.soarchain/config/genesis.json
	cat ~/.soarchain/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="udmotus"' > ~/.soarchain/config/tmp_genesis.json && mv ~/.soarchain/config/tmp_genesis.json ~/.soarchain/config/genesis.json
	cat ~/.soarchain/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="udmotus"' > ~/.soarchain/config/tmp_genesis.json && mv ~/.soarchain/config/tmp_genesis.json ~/.soarchain/config/genesis.json
	cat ~/.soarchain/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="udmotus"' > ~/.soarchain/config/tmp_genesis.json && mv ~/.soarchain/config/tmp_genesis.json ~/.soarchain/config/genesis.json
	sed -i 's/\("masterAccount": "\)[^"]*\(".*\)/\1'"$$(soarchaind keys show soarMasterAccount -a)"'\2/' ~/.soarchain/config/genesis.json

	cat ~/.soarchain/config/genesis.json | jq '.app_state["poa"]["masterKey"]["masterCertificate"]="-----BEGIN CERTIFICATE-----\nMIIB4TCCAYegAwIBAgIQTylBUpEkZd8CaYHSaLbBHzAKBggqhkjOPQQDAjBIMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMB4XDTIzMDMzMDA2NTUxNVoXDTQ4MDMzMDA2NTUxNVowSDEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEoMCYGA1UEAwwfU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLaCOXbFw/dRJXzXtvhSFWt92aUkdwRZPLmJWZFBFX55+XIDQsCGsQeMmU4pqsnXEB4/r842uYUinWsdzg4xUoqjUzBRMB0GA1UdDgQWBBRqxTRE6ZPuogp88TrNw1cwAYyPMjAfBgNVHSMEGDAWgBRqxTRE6ZPuogp88TrNw1cwAYyPMjAPBgNVHRMBAf8EBTADAQH/MAoGCCqGSM49BAMCA0gAMEUCIAHpI8Y6zPLaitMOGNAzzDAKb0PJw2r49vjzkFl5TIGPAiEArPJTReSmEnUJWFTcEIuYoWcRIBDI+GpianTVfX4uxNI=\n-----END CERTIFICATE-----"' > ~/.soarchain/config/tmp_genesis.json && mv ~/.soarchain/config/tmp_genesis.json ~/.soarchain/config/genesis.json

parameter_voting_period:
	cat <<< $$(jq '.app_state.gov.voting_params.voting_period = "600s"' ~/.soarchain/config/genesis.json) > ~/.soarchain/config/genesis.json

allocate_genesis_accounts:
	soarchaind add-genesis-account hermes 10000000udmotus --keyring-backend test
	soarchaind add-genesis-account soarMasterAccount 10000000udmotus --keyring-backend test
	soarchaind add-genesis-account client 4750000udmotus --keyring-backend test
	soarchaind add-genesis-account reserve 77425000000000udmotus 

sign_genesis_transaction:
	soarchaind gentx hermes 2000000udmotus --chain-id soarchaindevnet

collect_genesis_tx:
	soarchaind collect-gentxs

validate_genesis:
	soarchaind validate-genesis

ifdef start
	soarchaind start --log_level info --minimum-gas-prices=0.0001udmotus
endif

all: reset init config keys parameter_token_denomination allocate_genesis_accounts sign_genesis_transaction collect_genesis_tx validate_genesis