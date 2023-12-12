FROM ignitehq/cli:0.25.2 AS soarchain-full-node

COPY . /soarchain-core

USER root

RUN ls -a && cd /soarchain-core && ignite chain build -y

# init the chain
FROM soarchain-full-node

USER root

RUN soarchaind init docker_node --chain-id soarchaindevnet

# Install jq for JSON processing
RUN apt-get update && apt-get install -y jq

# Set chain parameters
RUN	cat /root/.soarchain/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="udmotus"' > /root/.soarchain/config/tmp_genesis.json && mv /root/.soarchain/config/tmp_genesis.json /root/.soarchain/config/genesis.json
RUN	cat /root/.soarchain/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="udmotus"' > /root/.soarchain/config/tmp_genesis.json && mv /root/.soarchain/config/tmp_genesis.json /root/.soarchain/config/genesis.json
RUN	cat /root/.soarchain/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="udmotus"' > /root/.soarchain/config/tmp_genesis.json && mv /root/.soarchain/config/tmp_genesis.json /root/.soarchain/config/genesis.json
RUN	cat /root/.soarchain/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="udmotus"' > /root/.soarchain/config/tmp_genesis.json && mv /root/.soarchain/config/tmp_genesis.json /root/.soarchain/config/genesis.json
RUN	cat /root/.soarchain/config/genesis.json | jq '.app_state["poa"]["masterKey"]["masterCertificate"]="-----BEGIN CERTIFICATE-----\nMIIB4TCCAYegAwIBAgIQTylBUpEkZd8CaYHSaLbBHzAKBggqhkjOPQQDAjBIMRwwGgYDVQQKDBNTb2FyIFJvYm90aWNzLCBJbmMuMSgwJgYDVQQDDB9Tb2FyIFJvYm90aWNzIFNlY3VyZSBFbGVtZW50IENBMB4XDTIzMDMzMDA2NTUxNVoXDTQ4MDMzMDA2NTUxNVowSDEcMBoGA1UECgwTU29hciBSb2JvdGljcywgSW5jLjEoMCYGA1UEAwwfU29hciBSb2JvdGljcyBTZWN1cmUgRWxlbWVudCBDQTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLaCOXbFw/dRJXzXtvhSFWt92aUkdwRZPLmJWZFBFX55+XIDQsCGsQeMmU4pqsnXEB4/r842uYUinWsdzg4xUoqjUzBRMB0GA1UdDgQWBBRqxTRE6ZPuogp88TrNw1cwAYyPMjAfBgNVHSMEGDAWgBRqxTRE6ZPuogp88TrNw1cwAYyPMjAPBgNVHRMBAf8EBTADAQH/MAoGCCqGSM49BAMCA0gAMEUCIAHpI8Y6zPLaitMOGNAzzDAKb0PJw2r49vjzkFl5TIGPAiEArPJTReSmEnUJWFTcEIuYoWcRIBDI+GpianTVfX4uxNI=\n-----END CERTIFICATE-----"' > /root/.soarchain/config/tmp_genesis.json && mv /root/.soarchain/config/tmp_genesis.json /root/.soarchain/config/genesis.json

RUN sed -i 's|^\(timeout_commit = \).*|\1"15s"|' /root/.soarchain/config/config.toml
RUN sed -i 's|laddr = "tcp://127.0.0.1:26657"|laddr = "tcp://0.0.0.0:26657"|' /root/.soarchain/config/config.toml
RUN sed -i '/# Enable defines if the API server should be enabled./!b;n;s/enable = false/enable = true/' /root/.soarchain/config/app.toml
RUN sed -i 's|\("masterAccount": "\)[^"]*|\1soar1qt8myp9424ng6rv4fwf65u9a0ttfschw5j4sp8|' /root/.soarchain/config/genesis.json

# Config keyring
RUN soarchaind config keyring-backend test

# Add local key
RUN soarchaind keys add docker_key --keyring-backend test --algo secp256k1

# Add genesis keys
RUN	soarchaind add-genesis-account docker_key 100000000000udmotus --keyring-backend test
RUN	soarchaind add-genesis-account soar1ghfnkjlc5gxpldat7hm50tgggwc6l5h7ydwy2a 36100000000udmotus
RUN	soarchaind add-genesis-account soar1ctvwnwl30getglarz2pp0gmnpggx4zznmlkv7j 36100000000udmotus
RUN	soarchaind add-genesis-account soar1y0kxkxk6prnx3ym569uacc82fqtvztsq3e78g3 36100000000udmotus
# Soarchain Master account
RUN	soarchaind add-genesis-account soar1qt8myp9424ng6rv4fwf65u9a0ttfschw5j4sp8 10000000000000udmotus

# Become a validator
RUN soarchaind gentx docker_key 2000000udmotus --chain-id soarchaindevnet

# Collect gentxs
RUN soarchaind collect-gentxs

# Validate the genesis file
RUN soarchaind validate-genesis

CMD ["/root/bin/soarchaind", "start"]