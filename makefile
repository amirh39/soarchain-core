include common.mk

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
	soarchaind keys add apollo --keyring-backend test --algo secp256k1
	soarchaind keys add client --recover --keyring-backend test --algo secp256k1
	soarchaind keys add soarMasterAccount --recover --keyring-backend test --algo secp256k1
	soarchaind keys add challenger --recover --keyring-backend test --algo secp256k1
	soarchaind keys add runner --recover --keyring-backend test --algo secp256k1

parameter_token_denomination:
	cat ~/.soarchain/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="utmotus"' > ~/.soarchain/config/tmp_genesis.json && mv ~/.soarchain/config/tmp_genesis.json ~/.soarchain/config/genesis.json
	cat ~/.soarchain/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="utmotus"' > ~/.soarchain/config/tmp_genesis.json && mv ~/.soarchain/config/tmp_genesis.json ~/.soarchain/config/genesis.json
	cat ~/.soarchain/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="utmotus"' > ~/.soarchain/config/tmp_genesis.json && mv ~/.soarchain/config/tmp_genesis.json ~/.soarchain/config/genesis.json
	cat ~/.soarchain/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="utmotus"' > ~/.soarchain/config/tmp_genesis.json && mv ~/.soarchain/config/tmp_genesis.json ~/.soarchain/config/genesis.json
	sed -i 's/\("masterAccount": "\)[^"]*\(".*\)/\1'"$$(soarchaind keys show soarMasterAccount -a)"'\2/' ~/.soarchain/config/genesis.json

	cat ~/.soarchain/config/genesis.json | jq '.app_state["poa"]["masterKey"]["masterCertificate"]="-----BEGIN CERTIFICATE-----\nMIIB4TCCAYegAwIBAgIQTylBUpEkZd8CaYHSaLbBHzAKBggqhkjOPQQDAjBIMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMB4XDTIzMDMzMDA2NTUxNVoXDTQ4MDMzMDA2NTUxNVowSDEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEoMCYGA1UEAwwfU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLaCOXbFw/dRJXzXtvhSFWt92aUkdwRZPLmJWZFBFX55+XIDQsCGsQeMmU4pqsnXEB4/r842uYUinWsdzg4xUoqjUzBRMB0GA1UdDgQWBBRqxTRE6ZPuogp88TrNw1cwAYyPMjAfBgNVHSMEGDAWgBRqxTRE6ZPuogp88TrNw1cwAYyPMjAPBgNVHRMBAf8EBTADAQH/MAoGCCqGSM49BAMCA0gAMEUCIAHpI8Y6zPLaitMOGNAzzDAKb0PJw2r49vjzkFl5TIGPAiEArPJTReSmEnUJWFTcEIuYoWcRIBDI+GpianTVfX4uxNI=\n-----END CERTIFICATE-----"' > ~/.soarchain/config/tmp_genesis.json && mv ~/.soarchain/config/tmp_genesis.json ~/.soarchain/config/genesis.json

parameter_voting_period:
	cat <<< $$(jq '.app_state.gov.voting_params.voting_period = "600s"' ~/.soarchain/config/genesis.json) > ~/.soarchain/config/genesis.json

allocate_genesis_accounts:

	soarchaind add-genesis-account soarMasterAccount 10000000000utmotus --keyring-backend test
	soarchaind add-genesis-account client 10000000000utmotus --keyring-backend test
	soarchaind add-genesis-account apollo 10000000000utmotus --keyring-backend test
	soarchaind add-genesis-account challenger 10000000000utmotus --keyring-backend test
	soarchaind add-genesis-account runner 10000000000utmotus --keyring-backend test
	soarchaind add-genesis-account soar1qyhtcgw54973l3tz7fag27480q5qzt7cmsv9th 47500000utmotus #airdrop
	soarchaind add-genesis-account soar1wfly3s05fvtuqs7lpesr8nas6rydm96jh88m9v 100700000utmotus #coummunityPool
	soarchaind add-genesis-account soar1743qv44dgty0zv4t7vnnzdfhd7ftjfsmpreggg 77425000utmotus #inverstorSeed
	soarchaind add-genesis-account soar1c9k0cjhq0sma2mskl6re9mx93lxkavzzm6xdj4 36100000utmotus #StrategicReserve0
	soarchaind add-genesis-account soar1paaxlh6luwlxvv9smf935nj53hz0yk7wna2hu2 285000000utmotus #team
	soarchaind add-genesis-account soar1m6u0zxu4hkg4ycqawgrvmnlqhudqr32ydgal0m 36100000utmotus #StrategicReserve1
	soarchaind add-genesis-account soar1hmj5fccg6nuns2scxz0pvwqgyy44ntp2knvnfa 36100000utmotus #StrategicReserve2
	soarchaind add-genesis-account soar1s6n8jr600zhnefzdgv0d5z5mdhkx7au8k0glh0 36100000utmotus #StrategicReserve3
	soarchaind add-genesis-account soar1hcr2r7v54gdus8ud7u7vm0wggt9y7pp2qfxssc 36100000utmotus #StrategicReserve4

sign_genesis_transaction:
	soarchaind gentx apollo 2000000utmotus --chain-id soarchaindevnet

collect_genesis_tx:
	soarchaind collect-gentxs

validate_genesis:
	soarchaind validate-genesis

ifdef start
	soarchaind start --log_level info --minimum-gas-prices=0.0001utmotus
endif

all: reset init config keys parameter_token_denomination allocate_genesis_accounts sign_genesis_transaction collect_genesis_tx validate_genesis