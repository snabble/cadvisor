module github.com/google/cadvisor

go 1.13

require (
	cloud.google.com/go v0.34.0
	github.com/Microsoft/go-winio v0.3.2
	github.com/Rican7/retry v0.1.1-0.20160712041035-272ad122d6e5
	github.com/SeanDolphin/bqschema v0.0.0-20150424181127-f92a08f515e1
	github.com/Shopify/sarama v1.19.0
	github.com/abbot/go-http-auth v0.0.0-20140618235127-c0ef4539dfab
	github.com/aws/aws-sdk-go v1.19.6
	github.com/beorn7/perks v0.0.0-20180321164747-3a771d992973
	github.com/blang/semver v3.1.0+incompatible
	github.com/containerd/console v0.0.0-20170925154832-84eeaae905fa
	github.com/containerd/containerd v1.0.2
	github.com/containerd/typeurl v0.0.0-20170912152501-f6943554a7e7
	github.com/coreos/go-systemd v0.0.0-20160527140244-4484981625c1
	github.com/coreos/pkg v0.0.0-20160530111557-7f080b6c11ac
	github.com/coreos/rkt v1.25.0
	github.com/cyphar/filepath-securejoin v0.2.2-0.20170720062807-ae69057f2299
	github.com/davecgh/go-spew v1.1.1
	github.com/docker/distribution v2.6.0-rc.1.0.20170726174610-edc3ab29cdff+incompatible
	github.com/docker/docker v1.4.2-0.20180612054059-a9fbbdc8dd87
	github.com/docker/go-connections v0.3.0
	github.com/docker/go-units v0.2.1-0.20151230175859-0bbddae09c5a
	github.com/eapache/go-resiliency v1.1.0
	github.com/eapache/queue v1.1.0
	github.com/euank/go-kmsg-parser v2.0.0+incompatible
	github.com/garyburd/redigo v0.0.0-20150301180006-535138d7bcd7
	github.com/go-ini/ini v1.9.0
	github.com/godbus/dbus v0.0.0-20151105175453-c7fdd8b5cd55
	github.com/gogo/protobuf v1.2.0
	github.com/golang/protobuf v1.2.0
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/influxdb/influxdb v0.9.6-0.20151125225445-9eab56311373
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af
	github.com/jteeuwen/go-bindata v3.0.7+incompatible // indirect
	github.com/karrick/godirwalk v1.7.5
	github.com/klauspost/crc32 v0.0.0-20151223135126-a3b15ae34567
	github.com/kr/pretty v0.0.0-20140723054909-088c856450c0
	github.com/kr/text v0.0.0-20130911015532-6807e777504f
	github.com/mattn/go-shellwords v1.0.4-0.20180201004752-39dbbfa24bbc
	github.com/matttproud/golang_protobuf_extensions v1.0.1
	github.com/mesos/mesos-go v0.0.7-0.20180413204204-29de6ff97b48
	github.com/mindprince/gonvml v0.0.0-20171110221305-fee913ce8fb2
	github.com/mistifyio/go-zfs v2.1.2-0.20170901132433-166dd29edf05+incompatible
	github.com/mrunalp/fileutils v0.0.0-20160930181131-4ee1cc9a8058
	github.com/olivere/elastic v6.2.23+incompatible
	github.com/olivere/elastic/v7 v7.0.7
	github.com/opencontainers/go-digest v1.0.0-rc1
	github.com/opencontainers/image-spec v1.0.0-rc6.0.20170604055404-372ad780f634
	github.com/opencontainers/runc v1.0.0-rc4.0.20180108143726-d5b4a3eddbe4
	github.com/opencontainers/runtime-spec v1.0.0
	github.com/opencontainers/selinux v1.0.0-rc1.0.20170621221121-4a2974bf1ee9
	github.com/pborman/uuid v0.0.0-20150824212802-cccd189d45f7
	github.com/pkg/errors v0.8.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/pquerna/ffjson v0.0.0-20171002144729-d49c2bc1aa13
	github.com/prometheus/client_golang v0.9.3-0.20190127221311-3c4408c8b829
	github.com/prometheus/client_model v0.0.0-20190115171406-56726106282f
	github.com/prometheus/common v0.2.0
	github.com/prometheus/procfs v0.0.0-20190117184657-bf6a532e95b1
	github.com/seccomp/libseccomp-golang v0.0.0-20150813023252-1b506fc7c24e
	github.com/sirupsen/logrus v1.2.0
	github.com/stretchr/objx v0.1.1
	github.com/stretchr/testify v1.3.0
	github.com/syndtr/gocapability v0.0.0-20160928074757-e7cb7fa329f4
	github.com/vishvananda/netlink v0.0.0-20150820014904-1e2e08e8a2dc
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	golang.org/x/oauth2 v0.0.0-20190226205417-e64efc72b421
	golang.org/x/sys v0.0.0-20190509141414-a5b02f93d862
	golang.org/x/text v0.3.0
	google.golang.org/api v0.3.1
	google.golang.org/genproto v0.0.0-20190307195333-5fe7a883aa19
	google.golang.org/grpc v1.19.0
	gopkg.in/olivere/elastic.v2 v2.0.12
	k8s.io/klog v0.3.0
	k8s.io/utils v0.0.0-20190712204705-3dccf664f023
)
