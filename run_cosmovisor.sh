export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

export DAEMON_NAME=soarchaind  
export DAEMON_HOME=$HOME/.soarchain

make remove
make build
make reset
make init
make config
make keys
make parameter_token_denomination
make parameter_voting_period
make allocate_genesis_accounts
make sign_genesis_transaction
make collect_genesis_tx
make validate_genesis

mkdir -p $DAEMON_HOME/cosmovisor/genesis/bin  
cp $(which soarchaind) $DAEMON_HOME/cosmovisor/genesis/bin

cosmovisor run start