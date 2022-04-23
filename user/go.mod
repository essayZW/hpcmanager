module github.com/essayZW/hpcmanager/user

go 1.17

replace github.com/essayZW/hpcmanager => ../

require (
	github.com/asim/go-micro/plugins/broker/rabbitmq/v4 v4.0.0-20220311080335-e5a35d38f931
	github.com/asim/go-micro/plugins/registry/etcd/v4 v4.0.0-20220419144745-367771923c07
	github.com/asim/go-micro/plugins/sync/etcd/v3 v3.7.1-0.20211012115553-1cd7cfaa6cbb
	github.com/essayZW/hpcmanager v0.0.0-00010101000000-000000000000
	github.com/essayZW/hpcmanager/gateway v0.0.0-00010101000000-000000000000
	github.com/essayZW/hpcmanager/hpc v0.0.0-00010101000000-000000000000
	github.com/essayZW/hpcmanager/permission v0.0.0-00010101000000-000000000000
	github.com/go-redis/redis/v8 v8.11.5
	github.com/google/uuid v1.3.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/mozillazg/go-pinyin v0.19.0
	go-micro.dev/v4 v4.6.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/guregu/null.v4 v4.0.0
)

require (
	github.com/Microsoft/go-winio v0.5.0 // indirect
	github.com/ProtonMail/go-crypto v0.0.0-20210428141323-04723f9f07d7 // indirect
	github.com/acomagu/bufpipe v1.0.3 // indirect
	github.com/asim/go-micro/plugins/config/encoder/yaml/v4 v4.0.0-20220118152736-9e0be6c85d75 // indirect
	github.com/asim/go-micro/plugins/config/source/etcd/v4 v4.0.0-20220118152736-9e0be6c85d75 // indirect
	github.com/asim/go-micro/plugins/logger/zap/v4 v4.0.0-20220118152736-9e0be6c85d75 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-git/gcfg v1.5.0 // indirect
	github.com/go-git/go-billy/v5 v5.3.1 // indirect
	github.com/go-git/go-git/v5 v5.4.2 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/kevinburke/ssh_config v0.0.0-20201106050909-4977a11b4351 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/streadway/amqp v1.0.0 // indirect
	github.com/urfave/cli/v2 v2.3.0 // indirect
	github.com/xanzy/ssh-agent v0.3.0 // indirect
	go.etcd.io/etcd/api/v3 v3.5.0 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.0 // indirect
	go.etcd.io/etcd/client/v3 v3.5.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.20.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220319134239-a9b59b0215f8 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c // indirect
	google.golang.org/grpc v1.45.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/essayZW/hpcmanager/gateway => ../gateway

replace github.com/essayZW/hpcmanager/user => ./

replace github.com/essayZW/hpcmanager/hpc => ../hpc

replace github.com/essayZW/hpcmanager/permission => ../permission

replace github.com/essayZW/hpcmanager/project => ../project

replace github.com/essayZW/hpcmanager/node => ../node

replace github.com/essayZW/hpcmanager/fee => ../fee

replace github.com/essayZW/hpcmanager/fss => ../fss

replace github.com/essayZW/hpcmanager/award => ../award
