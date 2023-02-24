#!/bin/sh

rm -rf ~/.osmosisd

osmosisd init --chain-id=receiver test
osmosisd keys add alice --keyring-backend=test
osmosisd keys add bob --keyring-backend=test
osmosisd add-genesis-account $(osmosisd keys show alice -a --keyring-backend=test) 100000000000000uosmo,100000000000000stake
osmosisd add-genesis-account $(osmosisd keys show bob -a --keyring-backend=test) 100000000000000uosmo,100000000000000stake
osmosisd gentx alice 100000000stake --chain-id=receiver --keyring-backend=test
osmosisd collect-gentxs

sed -i -e 's#"tcp://0.0.0.0:1317"#"tcp://0.0.0.0:1312"#g' $HOME/.osmosisd/config/app.toml
sed -i -e 's#":8080"#":9096"#g' $HOME/.osmosisd/config/app.toml
sed -i -e 's#":26657"#":26559"#g' $HOME/.osmosisd/config/config.toml
sed -i -e 's#":26656"#":26662"#g' $HOME/.osmosisd/config/config.toml

osmosisd start --minimum-gas-prices=0stake
