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

RUN sed -i 's|^\(persistent_peers = \).*|\1"e0ee4d996f87d0f2b7706e35e162531c7fd76493@34.165.238.45:26656"|' /root/.soarchain/config/config.toml
RUN sed -i 's|^\(timeout_commit = \).*|\1"15s"|' /root/.soarchain/config/config.toml

RUN soarchaind config keyring-backend test

RUN soarchaind keys add docker_key --keyring-backend test --algo secp256k1

CMD ["/root/bin/soarchaind", "start"]
