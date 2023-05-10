# build the chain

FROM ignitehq/cli:0.25.2 AS soarchain-full-node

COPY . /soarchain-core

USER root

RUN ls -a && cd /soarchain-core && ignite chain build -y

# init the chain

FROM soarchain-full-node

USER root

RUN soarchaind init docker_node --chain-id soarchaindevnet

RUN rm -rf /root/.soarchain/config/genesis.json && cp -r /soarchain-core/genesis.json /root/.soarchain/config/

# RUN rm -rf /root/.soarchain/config/config.toml && cp -r /soarchain-core/config.toml /root/.soarchain/config/

RUN sed -i 's/persistent_peers = ""/persistent_peers = "a76be618efb2d0f585691b1f7a8ec414a5c75a2f@34.171.22.26:26656"/g' /root/.soarchain/config/config.toml

RUN soarchaind config keyring-backend test

RUN soarchaind keys add docker_key --keyring-backend test --algo secp256k1

CMD ["/root/bin/soarchaind", "start"]
