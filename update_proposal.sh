export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

export DAEMON_NAME=soarchaind  
export DAEMON_HOME=$HOME/.soarchain

make build

mkdir -p $DAEMON_HOME/cosmovisor/upgrades/test1/bin  
cp $(which soarchaind) $DAEMON_HOME/cosmovisor/upgrades/test1/bin

make upgrade_proposal