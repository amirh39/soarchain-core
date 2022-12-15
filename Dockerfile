# build the chain

FROM ignitehq/cli:latest AS soarchain-full-node

COPY . /soarchain-core

USER root

RUN ls -a && cd /soarchain-core && ignite chain build -y

# init the chain

FROM soarchain-full-node

USER root

RUN soarchaind init docker_node --chain-id soarchaindevnet

RUN rm -rf /root/.soarchain/config/genesis.json && cp -r /soarchain-core/genesis.json /root/.soarchain/config/

RUN sed -i 's/persistent_peers = ""/persistent_peers = "66efd221eceb6fd5f0ad9f60fbd09e7177b7feb4@104.248.142.45:26656"/g' /root/.soarchain/config/config.toml

RUN soarchaind config keyring-backend test

RUN soarchaind keys add docker_key --keyring-backend test --algo secp256k1

CMD ["/root/bin/soarchaind", "start"]


