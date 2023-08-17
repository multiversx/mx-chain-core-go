module github.com/multiversx/mx-chain-core-go

go 1.20

require (
	github.com/btcsuite/btcd/btcutil v1.1.3
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.2
	github.com/hashicorp/golang-lru v0.5.4
	github.com/mr-tron/base58 v1.2.0
	github.com/pelletier/go-toml v1.9.3
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.3.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.2.0 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

replace github.com/gogo/protobuf => github.com/multiversx/protobuf v1.3.2
