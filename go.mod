module github.com/regen-network/regen-ledger

go 1.14

require (
	github.com/cosmos/cosmos-sdk v0.34.4-0.20201005082218-6c1c2cce0461
	github.com/enigmampc/btcutil v1.0.3-0.20200723161021-e2fb6adb2a25
	github.com/gorilla/mux v1.8.0
	github.com/rakyll/statik v0.1.7
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.34.0-rc4.0.20201005104435-7e27e9b85222
	github.com/tendermint/tm-db v0.6.2
	google.golang.org/grpc v1.32.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

replace github.com/cosmos/cosmos-sdk => github.com/cosmos/cosmos-sdk v0.34.4-0.20201005082218-6c1c2cce0461
