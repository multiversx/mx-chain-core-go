module github.com/multiversx/mx-chain-core-go

go 1.26

require (
	github.com/btcsuite/btcd/btcutil v1.2.0
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.3
	github.com/mr-tron/base58 v1.3.0
	github.com/pelletier/go-toml v1.9.5
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.11.1
	golang.org/x/crypto v0.54.0
	google.golang.org/grpc v1.67.1

)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

replace github.com/gogo/protobuf => github.com/multiversx/protobuf v1.3.2
