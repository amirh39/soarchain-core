export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

make remove
make build
make reset
make init
make config
make keys
make parameter_token_denomination
make allocate_genesis_accounts
make sign_genesis_transaction
make collect_genesis_tx
make validate_genesis
if [ "$1" = "start" ]; then
   soarchaind start --log_level info --minimum-gas-prices=0.0001soar
fi