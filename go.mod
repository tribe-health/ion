module github.com/pion/ion

go 1.13

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
<<<<<<< HEAD
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/at-wat/ebml-go v0.11.0
	github.com/cloudwebrtc/go-protoo v0.0.0-20200602160428-0a199e23f7e0
	github.com/cloudwebrtc/nats-protoo v0.0.0-20200604135451-87b43396e8de
	github.com/coreos/etcd v3.3.22+incompatible // indirect
=======
	github.com/cloudwebrtc/go-protoo v1.0.0
	github.com/coreos/etcd v3.3.25+incompatible // indirect
>>>>>>> upstream/v1.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis/v7 v7.4.0
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/google/go-cmp v0.5.2 // indirect
	github.com/google/uuid v1.1.2
	github.com/gorilla/websocket v1.4.2
	github.com/nats-io/nats.go v1.10.0
	github.com/notedit/sdp v0.0.4
<<<<<<< HEAD
	github.com/pion/ion-avp v1.0.12
	github.com/pion/ion-sfu v1.0.24
	github.com/pion/rtcp v1.2.4
	github.com/pion/rtp v1.6.1
	github.com/pion/stun v0.3.5
	github.com/pion/transport v0.10.1
	github.com/pion/webrtc/v2 v2.2.26
	github.com/pion/webrtc/v3 v3.0.0-beta.7
=======
	github.com/pion/ion-avp v1.1.1
	github.com/pion/ion-log v1.0.0
	github.com/pion/ion-sfu v1.2.0
	github.com/pion/webrtc/v3 v3.0.0-beta.12
>>>>>>> upstream/v1.0.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	go.etcd.io/etcd v3.3.25+incompatible
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)
