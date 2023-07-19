module github.com/celo-org/rosetta

go 1.15

require (
	github.com/allegro/bigcache v1.2.1 // indirect
	// TODO EN: update once gingerbread release is available
	github.com/celo-org/celo-blockchain v0.0.0-20230713152423-4a1e3fc3b943
	github.com/celo-org/celo-bls-go v0.6.4 // indirect
	github.com/celo-org/celo-bls-go-android v0.6.4 // indirect
	github.com/celo-org/celo-bls-go-ios v0.6.4 // indirect
	github.com/celo-org/celo-bls-go-linux v0.6.4 // indirect
	github.com/celo-org/celo-bls-go-macos v0.6.4 // indirect
	github.com/celo-org/celo-bls-go-other v0.6.4 // indirect
	github.com/celo-org/celo-bls-go-windows v0.6.4 // indirect
	// TODO EN: update once kliento is fully updated
	// points to head of eelanagaraj/gingerbread-updates
	github.com/celo-org/kliento v0.2.1-0.20230719154442-54f3f19facdc
	github.com/coinbase/rosetta-sdk-go v0.5.9
	github.com/felixge/httpsnoop v1.0.1
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/gorilla/handlers v1.4.2
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/onsi/gomega v1.10.1
	github.com/rjeczalik/notify v0.9.2 // indirect
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.6.3
	golang.org/x/crypto v0.1.0 // indirect
	golang.org/x/net v0.1.0
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4
)

// TODO EN: undo once kliento dep is published
// replace github.com/celo-org/kliento => ../kliento
