module github.com/essayZW/hpcmanager

go 1.17

require (
	github.com/asim/go-micro/plugins/logger/zap/v4 v4.0.0-20220118152736-9e0be6c85d75
	github.com/essayZW/hpcmanager/gateway v0.0.0-00010101000000-000000000000
	go-micro.dev/v4 v4.5.0
	go.uber.org/zap v1.20.0
)

require (
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/net v0.0.0-20210510120150-4163338589ed // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20210603081109-ebe580a85c40 // indirect
	golang.org/x/text v0.3.6 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

require (
	github.com/asim/go-micro/plugins/config/encoder/yaml/v4 v4.0.0-20220118152736-9e0be6c85d75
	github.com/google/uuid v1.3.0 // indirect
	google.golang.org/protobuf v1.26.0
)

replace github.com/essayZW/hpcmanager/gateway => ./gateway

replace github.com/essayZW/hpcmanager/user => ./user
