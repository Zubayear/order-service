module OrderService

go 1.16

require (
	entgo.io/ent v0.10.1
	github.com/Microsoft/go-winio v0.5.2 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20220113124808-70ae35bab23f // indirect
	github.com/asim/go-micro/plugins/config/encoder/yaml/v4 v4.0.0-20220401061051-1f086a300299
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/google/uuid v1.3.0
	github.com/google/wire v0.5.0 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/miekg/dns v1.1.48 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/urfave/cli/v2 v2.4.0 // indirect
	github.com/xanzy/ssh-agent v0.3.1 // indirect
	go-micro.dev/v4 v4.6.0
	golang.org/x/crypto v0.0.0-20220331220935-ae2d96664a29 // indirect
	golang.org/x/net v0.0.0-20220401154927-543a649e0bdd // indirect
	golang.org/x/sys v0.0.0-20220330033206-e17cdc41300f // indirect
	golang.org/x/tools v0.1.10 // indirect
	google.golang.org/protobuf v1.28.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
