load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "91585017debb61982f7054c9688857a2ad1fd823fc3f9cb05048b0025c47d023",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.42.0/rules_go-v0.42.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.42.0/rules_go-v0.42.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "d3fa66a39028e97d76f9e2db8f1b0c11c099e8e01bf363a923074784e451f809",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.33.0/bazel-gazelle-v0.33.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.33.0/bazel-gazelle-v0.33.0.tar.gz",
    ],
)

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "b1e80761a8a8243d03ebca8845e9cc1ba6c82ce7c5179ce2b295cd36f7e394bf",
    urls = [
        "https://github.com/bazelbuild/rules_docker/releases/download/v0.25.0/rules_docker-v0.25.0.tar.gz",
    ],
)

load(
    "@io_bazel_rules_docker//toolchains/docker:toolchain.bzl",
    docker_toolchain_configure = "toolchain_configure",
)

docker_toolchain_configure(
    name = "docker_config",
    docker_flags = [
        "--log-level=info",
    ],
    docker_path = "${CONTAINER_CMD:-$(command -v podman||command -v docker)}",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_repository(
    name = "co_honnef_go_tools",
    importpath = "honnef.co/go/tools",
    sum = "h1:/hemPrYIhOhy8zYrNj+069zDB68us2sMGsfkFJO0iZs=",
    version = "v0.0.0-20190523083050-ea95bdfd59fc",
)

go_repository(
    name = "com_github_14rcole_gopopulate",
    importpath = "github.com/14rcole/gopopulate",
    sum = "h1:SCbEWT58NSt7d2mcFdvxC9uyrdcTfvBbPLThhkDmXzg=",
    version = "v0.0.0-20180821133914-b175b219e774",
)

go_repository(
    name = "com_github_acarl005_stripansi",
    importpath = "github.com/acarl005/stripansi",
    sum = "h1:licZJFw2RwpHMqeKTCYkitsPqHNxTmd4SNR5r94FGM8=",
    version = "v0.0.0-20180116102854-5a71ef0e047d",
)

go_repository(
    name = "com_github_afex_hystrix_go",
    importpath = "github.com/afex/hystrix-go",
    sum = "h1:rFw4nCn9iMW+Vajsk51NtYIcwSTkXr+JGrMd36kTDJw=",
    version = "v0.0.0-20180502004556-fa1af6a1f4f5",
)

go_repository(
    name = "com_github_agnivade_levenshtein",
    importpath = "github.com/agnivade/levenshtein",
    sum = "h1:3oJU7J3FGFmyhn8KHjmVaZCN5hxTr7GxgRue+sxIXdQ=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_alecthomas_template",
    importpath = "github.com/alecthomas/template",
    sum = "h1:JYp7IbQjafoB+tBA3gMyHYHrpOtNuDiK/uB5uXxq5wM=",
    version = "v0.0.0-20190718012654-fb15b899a751",
)

go_repository(
    name = "com_github_alecthomas_units",
    importpath = "github.com/alecthomas/units",
    sum = "h1:s6gZFSlWYmbqAuRjVTiNNhvNRfY2Wxp9nhfyel4rklc=",
    version = "v0.0.0-20211218093645-b94a6e3cc137",
)

go_repository(
    name = "com_github_andreyvit_diff",
    importpath = "github.com/andreyvit/diff",
    sum = "h1:bvNMNQO63//z+xNgfBlViaCIJKLlCJ6/fmUseuG0wVQ=",
    version = "v0.0.0-20170406064948-c7f18ee00883",
)

go_repository(
    name = "com_github_apache_thrift",
    importpath = "github.com/apache/thrift",
    sum = "h1:5hryIiq9gtn+MiLVn0wP37kb/uTeRZgN08WoCsAhIhI=",
    version = "v0.13.0",
)

go_repository(
    name = "com_github_appscode_jsonpatch",
    importpath = "github.com/appscode/jsonpatch",
    sum = "h1:e82Bj+rsBSnpsmjiIGlc9NiKSBpJONZkamk/F8GrCR0=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_armon_circbuf",
    importpath = "github.com/armon/circbuf",
    sum = "h1:QEF07wC0T1rKkctt1RINW/+RMTVmiwxETico2l3gxJA=",
    version = "v0.0.0-20150827004946-bbbad097214e",
)

go_repository(
    name = "com_github_armon_consul_api",
    importpath = "github.com/armon/consul-api",
    sum = "h1:G1bPvciwNyF7IUmKXNt9Ak3m6u9DE1rF+RmtIkBpVdA=",
    version = "v0.0.0-20180202201655-eb2c6b5be1b6",
)

go_repository(
    name = "com_github_armon_go_metrics",
    importpath = "github.com/armon/go-metrics",
    sum = "h1:8GUt8eRujhVEGZFFEjBj46YV4rDjvGrNxb0KMWYkL2I=",
    version = "v0.0.0-20180917152333-f0300d1749da",
)

go_repository(
    name = "com_github_armon_go_radix",
    importpath = "github.com/armon/go-radix",
    sum = "h1:BUAU3CGlLvorLI26FmByPp2eC2qla6E1Tw+scpcg/to=",
    version = "v0.0.0-20180808171621-7fddfc383310",
)

go_repository(
    name = "com_github_aryann_difflib",
    importpath = "github.com/aryann/difflib",
    sum = "h1:pv34s756C4pEXnjgPfGYgdhg/ZdajGhyOvzx8k+23nw=",
    version = "v0.0.0-20170710044230-e206f873d14a",
)

go_repository(
    name = "com_github_asaskevich_govalidator",
    importpath = "github.com/asaskevich/govalidator",
    sum = "h1:idn718Q4B6AGu/h5Sxe66HYVdqdGu2l9Iebqhi/AEoA=",
    version = "v0.0.0-20190424111038-f61b66f89f4a",
)

go_repository(
    name = "com_github_aws_aws_lambda_go",
    importpath = "github.com/aws/aws-lambda-go",
    sum = "h1:SuCy7H3NLyp+1Mrfp+m80jcbi9KYWAs9/BXwppwRDzY=",
    version = "v1.13.3",
)

go_repository(
    name = "com_github_aws_aws_sdk_go",
    importpath = "github.com/aws/aws-sdk-go",
    sum = "h1:0xphMHGMLBrPMfxR2AmVjZKcMEESEgWF8Kru94BNByk=",
    version = "v1.27.0",
)

go_repository(
    name = "com_github_aws_aws_sdk_go_v2",
    importpath = "github.com/aws/aws-sdk-go-v2",
    sum = "h1:qZ+woO4SamnH/eEbjM2IDLhRNwIwND/RQyVlBLp3Jqg=",
    version = "v0.18.0",
)

go_repository(
    name = "com_github_azure_go_ansiterm",
    importpath = "github.com/Azure/go-ansiterm",
    sum = "h1:UQHMgLO+TxOElx5B5HZ4hJQsoJ/PvUvKRhJHDQXO8P8=",
    version = "v0.0.0-20210617225240-d185dfc1b5a1",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest",
    importpath = "github.com/Azure/go-autorest/autorest",
    sum = "h1:5YWtOnckcudzIw8lPPBcWOnmIFWMtHci1ZWAZulMSx0=",
    version = "v0.9.6",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_adal",
    importpath = "github.com/Azure/go-autorest/autorest/adal",
    sum = "h1:O1X4oexUxnZCaEUGsvMnr8ZGj8HI37tNezwY4npRqA0=",
    version = "v0.8.2",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_date",
    importpath = "github.com/Azure/go-autorest/autorest/date",
    sum = "h1:yW+Zlqf26583pE43KhfnhFcdmSWlm5Ew6bxipnr/tbM=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_azure_go_autorest_autorest_mocks",
    importpath = "github.com/Azure/go-autorest/autorest/mocks",
    sum = "h1:qJumjCaCudz+OcqE9/XtEPfvtOjOmKaui4EOpFI6zZc=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_azure_go_autorest_logger",
    importpath = "github.com/Azure/go-autorest/logger",
    sum = "h1:ruG4BSDXONFRrZZJ2GUXDiUyVpayPmb1GnWeHDdaNKY=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_azure_go_autorest_tracing",
    importpath = "github.com/Azure/go-autorest/tracing",
    sum = "h1:TRn4WjSnkcSy5AEG3pnbtFSwNtwzjr4VYyQflFE619k=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_beorn7_perks",
    importpath = "github.com/beorn7/perks",
    sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_bgentry_speakeasy",
    importpath = "github.com/bgentry/speakeasy",
    sum = "h1:ByYyxL9InA1OWqxJqqp2A5pYHUrCiAL6K3J+LKSsQkY=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_bketelsen_crypt",
    importpath = "github.com/bketelsen/crypt",
    sum = "h1:+0HFd5KSZ/mm3JmhmrDukiId5iR6w4+BdFtfSy4yWIc=",
    version = "v0.0.3-0.20200106085610-5cbc8cc4026c",
)

go_repository(
    name = "com_github_blang_semver",
    importpath = "github.com/blang/semver",
    sum = "h1:cQNTCjp13qL8KC3Nbxr/y2Bqb63oX6wdnnjpJbkM4JQ=",
    version = "v3.5.1+incompatible",
)

go_repository(
    name = "com_github_boltdb_bolt",
    importpath = "github.com/boltdb/bolt",
    sum = "h1:JQmyP4ZBrce+ZQu0dY660FMfatumYDLun9hBCUVIkF4=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_brancz_gojsontoyaml",
    importpath = "github.com/brancz/gojsontoyaml",
    sum = "h1:DMb8SuAL9+demT8equqMMzD8C/uxqWmj4cgV7ufrpQo=",
    version = "v0.0.0-20190425155809-e8bd32d46b3d",
)

go_repository(
    name = "com_github_burntsushi_toml",
    importpath = "github.com/BurntSushi/toml",
    sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_burntsushi_xgb",
    importpath = "github.com/BurntSushi/xgb",
    sum = "h1:1BDTz0u9nC3//pOCMdNH+CiXJVYJh5UQNCOBG7jbELc=",
    version = "v0.0.0-20160522181843-27f122750802",
)

go_repository(
    name = "com_github_campoy_embedmd",
    importpath = "github.com/campoy/embedmd",
    sum = "h1:V4kI2qTJJLf4J29RzI/MAt2c3Bl4dQSYPuflzwFH2hY=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_casbin_casbin_v2",
    importpath = "github.com/casbin/casbin/v2",
    sum = "h1:bTwon/ECRx9dwBy2ewRVr5OiqjeXSGiTUY74sDPQi/g=",
    version = "v2.1.2",
)

go_repository(
    name = "com_github_cenkalti_backoff",
    importpath = "github.com/cenkalti/backoff",
    sum = "h1:tNowT99t7UNflLxfYYSlKYsBpXdEet03Pg2g16Swow4=",
    version = "v2.2.1+incompatible",
)

go_repository(
    name = "com_github_census_instrumentation_opencensus_proto",
    importpath = "github.com/census-instrumentation/opencensus-proto",
    sum = "h1:glEXhBS5PSLLv4IXzLA5yPRVX4bilULVyxxbrfOtDAk=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_certifi_gocertifi",
    importpath = "github.com/certifi/gocertifi",
    sum = "h1:MmeatFT1pTPSVb4nkPmBFN/LRZ97vPjsFKsZrU3KKTs=",
    version = "v0.0.0-20180905225744-ee1a9a0726d2",
)

go_repository(
    name = "com_github_cespare_xxhash",
    importpath = "github.com/cespare/xxhash",
    sum = "h1:a6HrQnmkObjyL+Gs60czilIUGqrzKutQD6XZog3p+ko=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_cespare_xxhash_v2",
    importpath = "github.com/cespare/xxhash/v2",
    sum = "h1:DC2CZ1Ep5Y4k3ZQ899DldepgrayRUGE6BBZ/cd9Cj44=",
    version = "v2.2.0",
)

go_repository(
    name = "com_github_chzyer_logex",
    importpath = "github.com/chzyer/logex",
    sum = "h1:Swpa1K6QvQznwJRcfTfQJmTE72DqScAa40E+fbHEXEE=",
    version = "v1.1.10",
)

go_repository(
    name = "com_github_chzyer_readline",
    importpath = "github.com/chzyer/readline",
    sum = "h1:fY5BOSpyZCqRo5OhCuC+XN+r/bBCmeuuJtjz+bCNIf8=",
    version = "v0.0.0-20180603132655-2972be24d48e",
)

go_repository(
    name = "com_github_chzyer_test",
    importpath = "github.com/chzyer/test",
    sum = "h1:q763qf9huN11kDQavWsoZXJNW3xEE4JJyHa5Q25/sd8=",
    version = "v0.0.0-20180213035817-a1ea475d72b1",
)

go_repository(
    name = "com_github_clbanning_x2j",
    importpath = "github.com/clbanning/x2j",
    sum = "h1:EdRZT3IeKQmfCSrgo8SZ8V3MEnskuJP0wCYNpe+aiXo=",
    version = "v0.0.0-20191024224557-825249438eec",
)

go_repository(
    name = "com_github_client9_misspell",
    importpath = "github.com/client9/misspell",
    sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
    version = "v0.3.4",
)

go_repository(
    name = "com_github_cncf_udpa_go",
    importpath = "github.com/cncf/udpa/go",
    sum = "h1:WBZRG4aNOuI15bLRrCgN8fCq8E5Xuty6jGbmSNEvSsU=",
    version = "v0.0.0-20191209042840-269d4d468f6f",
)

go_repository(
    name = "com_github_cockroachdb_datadriven",
    importpath = "github.com/cockroachdb/datadriven",
    sum = "h1:OaNxuTZr7kxeODyLWsRMC+OD03aFUH+mW6r2d+MWa5Y=",
    version = "v0.0.0-20190809214429-80d97fb3cbaa",
)

go_repository(
    name = "com_github_codahale_hdrhistogram",
    importpath = "github.com/codahale/hdrhistogram",
    sum = "h1:qMd81Ts1T2OTKmB4acZcyKaMtRnY5Y44NuXGX2GFJ1w=",
    version = "v0.0.0-20161010025455-3a0bb77429bd",
)

go_repository(
    name = "com_github_container_storage_interface_spec",
    importpath = "github.com/container-storage-interface/spec",
    sum = "h1:bD9KIVgaVKKkQ/UbVUY9kCaH/CJbhNxe0eeB4JeJV2s=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_containernetworking_cni",
    importpath = "github.com/containernetworking/cni",
    sum = "h1:9OIL/sZmMYDBe+G8svzILAlulUpaDTUjeAbtH/JNLBo=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_containers_image_v5",
    importpath = "github.com/containers/image/v5",
    sum = "h1:h1FCOXH6Ux9/p/E4rndsQOC4yAdRU0msRTfLVeQ7FDQ=",
    version = "v5.5.1",
)

go_repository(
    name = "com_github_containers_libtrust",
    importpath = "github.com/containers/libtrust",
    sum = "h1:Q8ePgVfHDplZ7U33NwHZkrVELsZP5fYj9pM5WBZB2GE=",
    version = "v0.0.0-20190913040956-14b96171aa3b",
)

go_repository(
    name = "com_github_containers_ocicrypt",
    importpath = "github.com/containers/ocicrypt",
    sum = "h1:Q0/IPs8ohfbXNxEfyJ2pFVmvJu5BhqJUAmc6ES9NKbo=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_containers_storage",
    importpath = "github.com/containers/storage",
    sum = "h1:tw/uKRPDnmVrluIzer3dawTFG/bTJLP8IEUyHFhltYk=",
    version = "v1.20.2",
)

go_repository(
    name = "com_github_coreos_bbolt",
    importpath = "github.com/coreos/bbolt",
    sum = "h1:wZwiHHUieZCquLkDL0B8UhzreNWsPHooDAG3q34zk0s=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_coreos_etcd",
    importpath = "github.com/coreos/etcd",
    sum = "h1:8F3hqu9fGYLBifCmRCJsicFqDx/D68Rt3q1JMazcgBQ=",
    version = "v3.3.13+incompatible",
)

go_repository(
    name = "com_github_coreos_go_etcd",
    importpath = "github.com/coreos/go-etcd",
    sum = "h1:bXhRBIXoTm9BYHS3gE0TtQuyNZyeEMux2sDi4oo5YOo=",
    version = "v2.0.0+incompatible",
)

go_repository(
    name = "com_github_coreos_go_oidc",
    importpath = "github.com/coreos/go-oidc",
    sum = "h1:mh48q/BqXqgjVHpy2ZY7WnWAbenxRjsz9N1i1YxjHAk=",
    version = "v2.2.1+incompatible",
)

go_repository(
    name = "com_github_coreos_go_semver",
    importpath = "github.com/coreos/go-semver",
    sum = "h1:yi21YpKnrx1gt5R+la8n5WgS0kCrsPp33dmEyHReZr4=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_coreos_go_systemd",
    importpath = "github.com/coreos/go-systemd",
    sum = "h1:W8b4lQ4tFF21aspRGoBuCNV6V2fFJBF+pm1J6OY8Lys=",
    version = "v0.0.0-20190620071333-e64a0ec8b42a",
)

go_repository(
    name = "com_github_coreos_pkg",
    importpath = "github.com/coreos/pkg",
    sum = "h1:lBNOc5arjvs8E5mO2tbpBpLoyyu8B6e44T7hJy6potg=",
    version = "v0.0.0-20180928190104-399ea9e2e55f",
)

go_repository(
    name = "com_github_coreos_prometheus_operator",
    importpath = "github.com/coreos/prometheus-operator",
    sum = "h1:kd7mysk8mCdwquBcPLyuRoRFNJCpgezXu8yUvIYE2nc=",
    version = "v0.35.0",
)

go_repository(
    name = "com_github_cpuguy83_go_md2man",
    importpath = "github.com/cpuguy83/go-md2man",
    sum = "h1:BSKMNlYxDvnunlTymqtgONjNnaRV1sTpcovwwjF22jk=",
    version = "v1.0.10",
)

go_repository(
    name = "com_github_cpuguy83_go_md2man_v2",
    importpath = "github.com/cpuguy83/go-md2man/v2",
    sum = "h1:wfIWP927BUkWJb2NmU/kNDYIBTh/ziUX91+lVfRxZq4=",
    version = "v2.0.4",
)

go_repository(
    name = "com_github_creack_pty",
    importpath = "github.com/creack/pty",
    sum = "h1:uDmaGzcdjhF4i/plgjmEsriH11Y0o7RKapEf/LDaM3w=",
    version = "v1.1.9",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:U9qPSI2PIWSS1VwoXQT9A3Wy9MM3WgvqSxFWenqJduM=",
    version = "v1.1.2-0.20180830191138-d8f796af33cc",
)

go_repository(
    name = "com_github_davecgh_go_xdr",
    importpath = "github.com/davecgh/go-xdr",
    sum = "h1:qg9VbHo1TlL0KDM0vYvBG9EY0X0Yku5WYIPoFWt8f6o=",
    version = "v0.0.0-20161123171359-e6a2ba005892",
)

go_repository(
    name = "com_github_dgrijalva_jwt_go",
    importpath = "github.com/dgrijalva/jwt-go",
    sum = "h1:7qlOGliEKZXTDg6OTjfoBKDXWrumCAMpl/TFQ4/5kLM=",
    version = "v3.2.0+incompatible",
)

go_repository(
    name = "com_github_dgryski_go_sip13",
    importpath = "github.com/dgryski/go-sip13",
    sum = "h1:RMLoZVzv4GliuWafOuPuQDKSm1SJph7uCRnnS61JAn4=",
    version = "v0.0.0-20181026042036-e10d5fee7954",
)

go_repository(
    name = "com_github_docker_distribution",
    importpath = "github.com/docker/distribution",
    sum = "h1:a5mlkVzth6W5A4fOsS3D2EO5BUmsJpcB+cRlLU7cSug=",
    version = "v2.7.1+incompatible",
)

go_repository(
    name = "com_github_docker_docker",
    importpath = "github.com/docker/docker",
    sum = "h1:Sm8iD2lifO31DwXfkGzq8VgA7rwxPjRsYmeo0K/dF9Y=",
    version = "v1.4.2-0.20191219165747-a9416c67da9f",
)

go_repository(
    name = "com_github_docker_docker_credential_helpers",
    importpath = "github.com/docker/docker-credential-helpers",
    sum = "h1:zI2p9+1NQYdnG6sMU26EX4aVGlqbInSQxQXLvzJ4RPQ=",
    version = "v0.6.3",
)

go_repository(
    name = "com_github_docker_go_connections",
    importpath = "github.com/docker/go-connections",
    sum = "h1:El9xVISelRB7BuFusrZozjnkIM5YnzCViNKohAFqRJQ=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_docker_go_metrics",
    importpath = "github.com/docker/go-metrics",
    sum = "h1:AgB/0SvBxihN0X8OR4SjsblXkbMvalQ8cjmtKQ2rQV8=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_docker_go_units",
    importpath = "github.com/docker/go-units",
    sum = "h1:3uh0PgVws3nIA0Q+MwDC8yjEPf9zjRfZZWXZYDct3Tw=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_docker_libnetwork",
    importpath = "github.com/docker/libnetwork",
    sum = "h1:rACxqwRsHD075Vb7qTimAgMSKWoM5zER0Dhws/SgCVo=",
    version = "v0.0.0-20190731215715-7f13a5c99f4b",
)

go_repository(
    name = "com_github_docker_libtrust",
    importpath = "github.com/docker/libtrust",
    sum = "h1:UhxFibDNY/bfvqU5CAUmr9zpesgbU6SWc8/B4mflAE4=",
    version = "v0.0.0-20160708172513-aabc10ec26b7",
)

go_repository(
    name = "com_github_docker_spdystream",
    importpath = "github.com/docker/spdystream",
    sum = "h1:ZfSZ3P3BedhKGUhzj7BQlPSU4OvT6tfOKe3DVHzOA7s=",
    version = "v0.0.0-20181023171402-6480d4af844c",
)

go_repository(
    name = "com_github_docopt_docopt_go",
    importpath = "github.com/docopt/docopt-go",
    sum = "h1:bWDMxwH3px2JBh6AyO7hdCn/PkvCZXii8TGj7sbtEbQ=",
    version = "v0.0.0-20180111231733-ee0de3bc6815",
)

go_repository(
    name = "com_github_dustin_go_humanize",
    importpath = "github.com/dustin/go-humanize",
    sum = "h1:GzkhY7T5VNhEkwH0PVJgjz+fX1rhBrR7pRT3mDkpeCY=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_eapache_go_resiliency",
    importpath = "github.com/eapache/go-resiliency",
    sum = "h1:1NtRmCAqadE2FN4ZcN6g90TP3uk8cg9rn9eNK2197aU=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_eapache_go_xerial_snappy",
    importpath = "github.com/eapache/go-xerial-snappy",
    sum = "h1:YEetp8/yCZMuEPMUDHG0CW/brkkEp8mzqk2+ODEitlw=",
    version = "v0.0.0-20180814174437-776d5712da21",
)

go_repository(
    name = "com_github_eapache_queue",
    importpath = "github.com/eapache/queue",
    sum = "h1:YOEu7KNc61ntiQlcEeUIoDTJ2o8mQznoNvUhiigpIqc=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_edsrzf_mmap_go",
    importpath = "github.com/edsrzf/mmap-go",
    sum = "h1:CEBF7HpRnUCSJgGUb5h1Gm7e3VkmVDrR8lvWVLtrOFw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_elazarl_go_bindata_assetfs",
    importpath = "github.com/elazarl/go-bindata-assetfs",
    sum = "h1:G/bYguwHIzWq9ZoyUQqrjTmJbbYn3j3CKKpKinvZLFk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_emicklei_go_restful",
    importpath = "github.com/emicklei/go-restful",
    sum = "h1:8KpYO/Xl/ZudZs5RNOEhWMBY4hmzlZhhRd9cu+jrZP4=",
    version = "v2.15.0+incompatible",
)

go_repository(
    name = "com_github_emicklei_go_restful_openapi",
    importpath = "github.com/emicklei/go-restful-openapi",
    sum = "h1:ohRZ1yEZERGzqaozBgxa3A0lt6c6KF14xhs3IL9ECwg=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_emicklei_go_restful_swagger12",
    importpath = "github.com/emicklei/go-restful-swagger12",
    sum = "h1:V94anc0ZG3Pa/cAMwP2m1aQW3+/FF8Qmw/GsFyTJAp4=",
    version = "v0.0.0-20170926063155-7524189396c6",
)

go_repository(
    name = "com_github_envoyproxy_go_control_plane",
    importpath = "github.com/envoyproxy/go-control-plane",
    sum = "h1:4cmBvAEBNJaGARUEs3/suWRyfyBfhf7I60WBZq+bv2w=",
    version = "v0.9.1-0.20191026205805-5f8ba28d4473",
)

go_repository(
    name = "com_github_envoyproxy_protoc_gen_validate",
    importpath = "github.com/envoyproxy/protoc-gen-validate",
    sum = "h1:EQciDnbrYxy13PgWoY8AqoxGiPrpgBZ1R8UNe3ddc+A=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_evanphx_json_patch",
    importpath = "github.com/evanphx/json-patch",
    sum = "h1:jBYDEEiFBPxA0v50tFdvOzQQTCvpL6mnFh5mB2/l16U=",
    version = "v5.6.0+incompatible",
)

go_repository(
    name = "com_github_fatih_color",
    importpath = "github.com/fatih/color",
    sum = "h1:8xPHl4/q1VyqGIPif1F+1V3Y3lSmrq01EabUW3CoW5s=",
    version = "v1.9.0",
)

go_repository(
    name = "com_github_fortytw2_leaktest",
    importpath = "github.com/fortytw2/leaktest",
    sum = "h1:u8491cBMTQ8ft8aeV+adlcytMZylmA5nnwwkRZjI8vw=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_franela_goblin",
    importpath = "github.com/franela/goblin",
    sum = "h1:gb2Z18BhTPJPpLQWj4T+rfKHYCHxRHCtRxhKKjRidVw=",
    version = "v0.0.0-20200105215937-c9ffbefa60db",
)

go_repository(
    name = "com_github_franela_goreq",
    importpath = "github.com/franela/goreq",
    sum = "h1:a9ENSRDFBUPkJ5lCgVZh26+ZbGyoVJG7yb5SSzF5H54=",
    version = "v0.0.0-20171204163338-bcd34c9993f8",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    sum = "h1:n+5WquG0fcWoWp6xPWfHdbskMCQaFnG6PfBrh1Ky4HY=",
    version = "v1.6.0",
)

go_repository(
    name = "com_github_fsouza_go_dockerclient",
    importpath = "github.com/fsouza/go-dockerclient",
    sum = "h1:B94x7idrPK5yFSMWliAvakUQlTDLgO7iJyHtpbYxGGM=",
    version = "v0.0.0-20171004212419-da3951ba2e9e",
)

go_repository(
    name = "com_github_fullsailor_pkcs7",
    importpath = "github.com/fullsailor/pkcs7",
    sum = "h1:RDBNVkRviHZtvDvId8XSGPu3rmpmSe+wKRcEWNgsfWU=",
    version = "v0.0.0-20190404230743-d7302db945fa",
)

go_repository(
    name = "com_github_getsentry_raven_go",
    importpath = "github.com/getsentry/raven-go",
    sum = "h1:F2m41rgyxoveZKD+Z6xwyAbtdNeVvhpi9BpQLvt5oRU=",
    version = "v0.0.0-20190513200303-c977f96e1095",
)

go_repository(
    name = "com_github_ghodss_yaml",
    importpath = "github.com/ghodss/yaml",
    sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_gin_contrib_cors",
    importpath = "github.com/gin-contrib/cors",
    sum = "h1:doAsuITavI4IOcd0Y19U4B+O0dNWihRyX//nn4sEmgA=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_gin_contrib_sse",
    importpath = "github.com/gin-contrib/sse",
    sum = "h1:Y/yl/+YNO8GZSjAhjMsSuLt29uWRFHdHYUb5lYOV9qE=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_gin_gonic_gin",
    importpath = "github.com/gin-gonic/gin",
    sum = "h1:4idEAncQnU5cB7BeOkPtxjfCSye0AAm1R0RVIqJ+Jmg=",
    version = "v1.9.1",
)

go_repository(
    name = "com_github_globalsign_mgo",
    importpath = "github.com/globalsign/mgo",
    sum = "h1:DujepqpGd1hyOd7aW59XpK7Qymp8iy83xq74fLr21is=",
    version = "v0.0.0-20181015135952-eeefdecb41b8",
)

go_repository(
    name = "com_github_go_bindata_go_bindata",
    importpath = "github.com/go-bindata/go-bindata",
    sum = "h1:5vjJMVhowQdPzjE1LdxyFF7YFTXg5IgGVW4gBr5IbvE=",
    version = "v3.1.2+incompatible",
)

go_repository(
    name = "com_github_go_gl_glfw",
    importpath = "github.com/go-gl/glfw",
    sum = "h1:QbL/5oDUmRBzO9/Z7Seo6zf912W/a6Sr4Eu0G/3Jho0=",
    version = "v0.0.0-20190409004039-e6da0acd62b1",
)

go_repository(
    name = "com_github_go_gl_glfw_v3_3_glfw",
    importpath = "github.com/go-gl/glfw/v3.3/glfw",
    sum = "h1:WtGNWLvXpe6ZudgnXrq0barxBImvnnJoMEhXAzcbM0I=",
    version = "v0.0.0-20200222043503-6f7a984d4dc4",
)

go_repository(
    name = "com_github_go_kit_kit",
    importpath = "github.com/go-kit/kit",
    sum = "h1:dXFJfIHVvUcpSgDOV+Ne6t7jXri8Tfv2uOLHUZ2XNuo=",
    version = "v0.10.0",
)

go_repository(
    name = "com_github_go_kit_log",
    importpath = "github.com/go-kit/log",
    sum = "h1:MRVx0/zhvdseW+Gza6N9rVzU/IVzaeE1SFI4raAhmBU=",
    version = "v0.2.1",
)

go_repository(
    name = "com_github_go_logfmt_logfmt",
    importpath = "github.com/go-logfmt/logfmt",
    sum = "h1:otpy5pqBCBZ1ng9RQ0dPu4PN7ba75Y/aA+UpowDyNVA=",
    version = "v0.5.1",
)

go_repository(
    name = "com_github_go_logr_logr",
    importpath = "github.com/go-logr/logr",
    sum = "h1:2y3SDp0ZXuc6/cjLSZ+Q3ir+QB9T/iG5yYRXqsagWSY=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_go_logr_zapr",
    importpath = "github.com/go-logr/zapr",
    sum = "h1:QHVo+6stLbfJmYGkQ7uGHUCu5hnAFAj6mDe6Ea0SeOo=",
    version = "v1.2.4",
)

go_repository(
    name = "com_github_go_openapi_analysis",
    importpath = "github.com/go-openapi/analysis",
    sum = "h1:8b2ZgKfKIUTVQpTb77MoRDIMEIwvDVw40o3aOXdfYzI=",
    version = "v0.19.5",
)

go_repository(
    name = "com_github_go_openapi_errors",
    importpath = "github.com/go-openapi/errors",
    sum = "h1:a2kIyV3w+OS3S97zxUndRVD46+FhGOUBDFY7nmu4CsY=",
    version = "v0.19.2",
)

go_repository(
    name = "com_github_go_openapi_jsonpointer",
    importpath = "github.com/go-openapi/jsonpointer",
    sum = "h1:eCs3fxoIi3Wh6vtgmLTOjdhSpiqphQ+DaPn38N2ZdrE=",
    version = "v0.19.6",
)

go_repository(
    name = "com_github_go_openapi_jsonreference",
    importpath = "github.com/go-openapi/jsonreference",
    sum = "h1:3sVjiK66+uXK/6oQ8xgcRKcFgQ5KXa2KvnJRumpMGbE=",
    version = "v0.20.2",
)

go_repository(
    name = "com_github_go_openapi_loads",
    importpath = "github.com/go-openapi/loads",
    sum = "h1:5I4CCSqoWzT+82bBkNIvmLc0UOsoKKQ4Fz+3VxOB7SY=",
    version = "v0.19.4",
)

go_repository(
    name = "com_github_go_openapi_runtime",
    importpath = "github.com/go-openapi/runtime",
    sum = "h1:csnOgcgAiuGoM/Po7PEpKDoNulCcF3FGbSnbHfxgjMI=",
    version = "v0.19.4",
)

go_repository(
    name = "com_github_go_openapi_spec",
    importpath = "github.com/go-openapi/spec",
    sum = "h1:ixzUSnHTd6hCemgtAJgluaTSGYpLNpJY4mA2DIkdOAo=",
    version = "v0.19.4",
)

go_repository(
    name = "com_github_go_openapi_strfmt",
    importpath = "github.com/go-openapi/strfmt",
    sum = "h1:eRfyY5SkaNJCAwmmMcADjY31ow9+N7MCLW7oRkbsINA=",
    version = "v0.19.3",
)

go_repository(
    name = "com_github_go_openapi_swag",
    importpath = "github.com/go-openapi/swag",
    sum = "h1:yMBqmnQ0gyZvEb/+KzuWZOXgllrXT4SADYbvDaXHv/g=",
    version = "v0.22.3",
)

go_repository(
    name = "com_github_go_openapi_validate",
    importpath = "github.com/go-openapi/validate",
    sum = "h1:QhCBKRYqZR+SKo4gl1lPhPahope8/RLt6EVgY8X80w0=",
    version = "v0.19.5",
)

go_repository(
    name = "com_github_go_playground_assert_v2",
    importpath = "github.com/go-playground/assert/v2",
    sum = "h1:JvknZsQTYeFEAhQwI4qEt9cyV5ONwRHC+lYKSsYSR8s=",
    version = "v2.2.0",
)

go_repository(
    name = "com_github_go_playground_locales",
    importpath = "github.com/go-playground/locales",
    sum = "h1:EWaQ/wswjilfKLTECiXz7Rh+3BjFhfDFKv/oXslEjJA=",
    version = "v0.14.1",
)

go_repository(
    name = "com_github_go_playground_universal_translator",
    importpath = "github.com/go-playground/universal-translator",
    sum = "h1:Bcnm0ZwsGyWbCzImXv+pAJnYK9S473LQFuzCbDbfSFY=",
    version = "v0.18.1",
)

go_repository(
    name = "com_github_go_playground_validator_v10",
    importpath = "github.com/go-playground/validator/v10",
    sum = "h1:vgvQWe3XCz3gIeFDm/HnTIbj6UGmg/+t63MyGU2n5js=",
    version = "v10.14.0",
)

go_repository(
    name = "com_github_go_sql_driver_mysql",
    importpath = "github.com/go-sql-driver/mysql",
    sum = "h1:7LxgVwFb2hIQtMm87NdgAVfXjnt4OePseqT1tKx+opk=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_go_stack_stack",
    importpath = "github.com/go-stack/stack",
    sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
    version = "v1.8.0",
)

go_repository(
    name = "com_github_gobuffalo_flect",
    importpath = "github.com/gobuffalo/flect",
    sum = "h1:PAVD7sp0KOdfswjAw9BpLCU9hXo7wFSzgpQ+zNeks/A=",
    version = "v0.2.2",
)

go_repository(
    name = "com_github_godbus_dbus",
    importpath = "github.com/godbus/dbus",
    sum = "h1:BWhy2j3IXJhjCbC68FptL43tDKIq8FladmaTs3Xs7Z8=",
    version = "v0.0.0-20190422162347-ade71ed3457e",
)

go_repository(
    name = "com_github_gogo_googleapis",
    importpath = "github.com/gogo/googleapis",
    sum = "h1:kFkMAZBNAn4j7K0GiZr8cRYzejq68VbheufiV3YuyFI=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_gogo_protobuf",
    importpath = "github.com/gogo/protobuf",
    sum = "h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_golang_glog",
    importpath = "github.com/golang/glog",
    sum = "h1:VKtxabqXZkF25pY9ekfRL6a582T4P37/31XEstQ5p58=",
    version = "v0.0.0-20160126235308-23def4e6c14b",
)

go_repository(
    name = "com_github_golang_groupcache",
    importpath = "github.com/golang/groupcache",
    sum = "h1:oI5xCqsCo564l8iNU+DwB5epxmsaqB+rhGL0m5jtYqE=",
    version = "v0.0.0-20210331224755-41bb18bfe9da",
)

go_repository(
    name = "com_github_golang_lint",
    importpath = "github.com/golang/lint",
    sum = "h1:2hRPrmiwPrp3fQX967rNJIhQPtiGXdlQWAxKbKw3VHA=",
    version = "v0.0.0-20180702182130-06c8688daad7",
)

go_repository(
    name = "com_github_golang_mock",
    importpath = "github.com/golang/mock",
    sum = "h1:G5FRp8JnTd7RQH5kemVNlMeyXQAztQ3mOWV95KxsXH8=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    sum = "h1:KhyjKVUg7Usr/dYsdSqoFveMYd5ko72D+zANwlG1mmg=",
    version = "v1.5.3",
)

go_repository(
    name = "com_github_golang_snappy",
    importpath = "github.com/golang/snappy",
    sum = "h1:aeE13tS0IiQgFjYdoL8qN3K1N2bXXtI6Vi51/y7BpMw=",
    version = "v0.0.2",
)

go_repository(
    name = "com_github_gonum_blas",
    importpath = "github.com/gonum/blas",
    sum = "h1:Q0Jsdxl5jbxouNs1TQYt0gxesYMU4VXRbsTlgDloZ50=",
    version = "v0.0.0-20181208220705-f22b278b28ac",
)

go_repository(
    name = "com_github_gonum_floats",
    importpath = "github.com/gonum/floats",
    sum = "h1:EvokxLQsaaQjcWVWSV38221VAK7qc2zhaO17bKys/18=",
    version = "v0.0.0-20181209220543-c233463c7e82",
)

go_repository(
    name = "com_github_gonum_graph",
    importpath = "github.com/gonum/graph",
    sum = "h1:NcVXNHJrvrcAv8SVYKzKT2zwtEXU1DK2J+azsK7oz2A=",
    version = "v0.0.0-20170401004347-50b27dea7ebb",
)

go_repository(
    name = "com_github_gonum_internal",
    importpath = "github.com/gonum/internal",
    sum = "h1:8jtTdc+Nfj9AR+0soOeia9UZSvYBvETVHZrugUowJ7M=",
    version = "v0.0.0-20181124074243-f884aa714029",
)

go_repository(
    name = "com_github_gonum_lapack",
    importpath = "github.com/gonum/lapack",
    sum = "h1:7qnwS9+oeSiOIsiUMajT+0R7HR6hw5NegnKPmn/94oI=",
    version = "v0.0.0-20181123203213-e4cdc5a0bff9",
)

go_repository(
    name = "com_github_gonum_matrix",
    importpath = "github.com/gonum/matrix",
    sum = "h1:V2IgdyerlBa/MxaEFRbV5juy/C3MGdj4ePi+g6ePIp4=",
    version = "v0.0.0-20181209220409-c518dec07be9",
)

go_repository(
    name = "com_github_google_btree",
    importpath = "github.com/google/btree",
    sum = "h1:gK4Kx5IaGY9CD5sPJ36FHiBJ6ZXl0kilRiiCj+jdYp4=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    sum = "h1:ofyhxvXcZhMsU5ulbFiLKl/XBFqE1GSq7atu8tAmTRI=",
    version = "v0.6.0",
)

go_repository(
    name = "com_github_google_gofuzz",
    importpath = "github.com/google/gofuzz",
    sum = "h1:xRy4A+RhZaiKjJ1bPfwQ8sedCA+YS2YcCHW6ec7JMi0=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_google_martian",
    importpath = "github.com/google/martian",
    sum = "h1:/CP5g8u/VJHijgedC/Legn3BAbAaWPgecwXBIDzw5no=",
    version = "v2.1.0+incompatible",
)

go_repository(
    name = "com_github_google_martian_v3",
    importpath = "github.com/google/martian/v3",
    sum = "h1:pMen7vLs8nvgEYhywH3KDWJIJTeEr2ULsVWHWYHQyBs=",
    version = "v3.0.0",
)

go_repository(
    name = "com_github_google_pprof",
    importpath = "github.com/google/pprof",
    sum = "h1:1FjCyPC+syAzJ5/2S8fqdZK1R22vvA0J7JZKcuOIQ7Y=",
    version = "v0.0.0-20211214055906-6f57359322fd",
)

go_repository(
    name = "com_github_google_renameio",
    importpath = "github.com/google/renameio",
    sum = "h1:GOZbcHa3HfsPKPlmyPyN2KEohoMXOhdMbHrvbpl2QaA=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_google_uuid",
    importpath = "github.com/google/uuid",
    sum = "h1:MtMxsa51/r9yyhkyLsVeVt0B+BGQZzpQiTQ4eHZ8bc4=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_googleapis_gax_go_v2",
    importpath = "github.com/googleapis/gax-go/v2",
    sum = "h1:sjZBwGj9Jlw33ImPtvFviGYvseOtDM7hkSKB7+Tv3SM=",
    version = "v2.0.5",
)

go_repository(
    name = "com_github_googleapis_gnostic",
    importpath = "github.com/googleapis/gnostic",
    sum = "h1:9fHAtK0uDfpveeqqo1hkEZJcFvYXAiCN3UutL8F9xHw=",
    version = "v0.5.5",
)

go_repository(
    name = "com_github_gopherjs_gopherjs",
    importpath = "github.com/gopherjs/gopherjs",
    sum = "h1:EGx4pi6eqNxGaHF6qqu48+N2wcFQ5qg5FXgOdqsJ5d8=",
    version = "v0.0.0-20181017120253-0766667cb4d1",
)

go_repository(
    name = "com_github_gorilla_context",
    importpath = "github.com/gorilla/context",
    sum = "h1:AWwleXJkX/nhcU9bZSnZoi3h/qGYqQAGhq6zZe/aQW8=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_gorilla_mux",
    importpath = "github.com/gorilla/mux",
    sum = "h1:i40aqfkR1h2SlN9hojwV5ZA91wcXFOvkdNIeFDP5koI=",
    version = "v1.8.0",
)

go_repository(
    name = "com_github_gorilla_websocket",
    importpath = "github.com/gorilla/websocket",
    sum = "h1:gmztn0JnHVt9JZquRuzLw3g4wouNVzKL15iLr/zn/QY=",
    version = "v1.5.1",
)

go_repository(
    name = "com_github_gregjones_httpcache",
    importpath = "github.com/gregjones/httpcache",
    sum = "h1:pdN6V1QBWetyv/0+wjACpqVH+eVULgEjkurDLq3goeM=",
    version = "v0.0.0-20180305231024-9cad4c3443a7",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_middleware",
    importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
    sum = "h1:+9834+KizmvFV7pXQGSXQTsaWhq2GjuNUt0aUU0YBYw=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_go_grpc_prometheus",
    importpath = "github.com/grpc-ecosystem/go-grpc-prometheus",
    sum = "h1:Ovs26xHkKqVztRpIrF/92BcuyuQ/YW4NSIpoGtfXNho=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
    sum = "h1:gmcG1KaJ57LophUzW0Hy8NmPhnMZb4M0+kPpLofRdBo=",
    version = "v1.16.0",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_health_probe",
    importpath = "github.com/grpc-ecosystem/grpc-health-probe",
    sum = "h1:UxmGBzaBcWDQuQh9E1iT1dWKQFbizZ+SpTd1EL4MSqs=",
    version = "v0.2.1-0.20181220223928-2bf0a5b182db",
)

go_repository(
    name = "com_github_hashicorp_consul_api",
    importpath = "github.com/hashicorp/consul/api",
    sum = "h1:HXNYlRkkM/t+Y/Yhxtwcy02dlYwIaoxzvxPnS+cqy78=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_hashicorp_consul_sdk",
    importpath = "github.com/hashicorp/consul/sdk",
    sum = "h1:UOxjlb4xVNF93jak1mzzoBatyFju9nrkxpVwIp/QqxQ=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_hashicorp_errwrap",
    importpath = "github.com/hashicorp/errwrap",
    sum = "h1:hLrqtEDnRye3+sgx6z4qVLNuviH3MR5aQ0ykNJa/UYA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_cleanhttp",
    importpath = "github.com/hashicorp/go-cleanhttp",
    sum = "h1:dH3aiDG9Jvb5r5+bYHsikaOUIpcM0xvgMXVoDkXMzJM=",
    version = "v0.5.1",
)

go_repository(
    name = "com_github_hashicorp_go_immutable_radix",
    importpath = "github.com/hashicorp/go-immutable-radix",
    sum = "h1:AKDB1HM5PWEA7i4nhcpwOrO2byshxBjXVn/J/3+z5/0=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_msgpack",
    importpath = "github.com/hashicorp/go-msgpack",
    sum = "h1:zKjpN5BK/P5lMYrLmBHdBULWbJ0XpYR+7NGzqkZzoD4=",
    version = "v0.5.3",
)

go_repository(
    name = "com_github_hashicorp_go_multierror",
    importpath = "github.com/hashicorp/go-multierror",
    sum = "h1:iVjPR7a6H0tWELX5NxNe7bYopibicUzc7uPribsnS6o=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_net",
    importpath = "github.com/hashicorp/go.net",
    sum = "h1:sNCoNyDEvN1xa+X0baata4RdcpKwcMS6DH+xwfqPgjw=",
    version = "v0.0.1",
)

go_repository(
    name = "com_github_hashicorp_go_rootcerts",
    importpath = "github.com/hashicorp/go-rootcerts",
    sum = "h1:Rqb66Oo1X/eSV1x66xbDccZjhJigjg0+e82kpwzSwCI=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_sockaddr",
    importpath = "github.com/hashicorp/go-sockaddr",
    sum = "h1:GeH6tui99pF4NJgfnhp+L6+FfobzVW3Ah46sLo0ICXs=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_syslog",
    importpath = "github.com/hashicorp/go-syslog",
    sum = "h1:KaodqZuhUoZereWVIYmpUgZysurB1kBLX2j0MwMrUAE=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_go_uuid",
    importpath = "github.com/hashicorp/go-uuid",
    sum = "h1:2gKiV6YVmrJ1i2CKKa9obLvRieoRGviZFL26PcT/Co8=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_hashicorp_go_version",
    importpath = "github.com/hashicorp/go-version",
    sum = "h1:3vNe/fWF5CBgRIguda1meWhsZHy3m8gCJ5wx+dIzX/E=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_hashicorp_golang_lru",
    importpath = "github.com/hashicorp/golang-lru",
    sum = "h1:YDjusn29QI/Das2iO9M0BHnIbxPeyuCHsjMW+lJfyTc=",
    version = "v0.5.4",
)

go_repository(
    name = "com_github_hashicorp_hcl",
    importpath = "github.com/hashicorp/hcl",
    sum = "h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_logutils",
    importpath = "github.com/hashicorp/logutils",
    sum = "h1:dLEQVugN8vlakKOUE3ihGLTZJRB4j+M2cdTm/ORI65Y=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_mdns",
    importpath = "github.com/hashicorp/mdns",
    sum = "h1:WhIgCr5a7AaVH6jPUwjtRuuE7/RDufnUvzIr48smyxs=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hashicorp_memberlist",
    importpath = "github.com/hashicorp/memberlist",
    sum = "h1:EmmoJme1matNzb+hMpDuR/0sbJSUisxyqBGG676r31M=",
    version = "v0.1.3",
)

go_repository(
    name = "com_github_hashicorp_serf",
    importpath = "github.com/hashicorp/serf",
    sum = "h1:YZ7UKsJv+hKjqGVUUbtE3HNj79Eln2oQ75tniF6iPt0=",
    version = "v0.8.2",
)

go_repository(
    name = "com_github_hpcloud_tail",
    importpath = "github.com/hpcloud/tail",
    sum = "h1:nfCOvKYfkgYP8hkirhJocXT2+zOD8yUNjXaWfTlyFKI=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_hudl_fargo",
    importpath = "github.com/hudl/fargo",
    sum = "h1:0U6+BtN6LhaYuTnIJq4Wyq5cpn6O2kWrxAtcqBmYY6w=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_ianlancetaylor_demangle",
    importpath = "github.com/ianlancetaylor/demangle",
    sum = "h1:uGg2frlt3IcT7kbV6LEp5ONv4vmoO2FW4qSO+my/aoM=",
    version = "v0.0.0-20210905161508-09a460cdf81d",
)

go_repository(
    name = "com_github_imdario_mergo",
    importpath = "github.com/imdario/mergo",
    sum = "h1:Y+UAYTZ7gDEuOfhxKWy+dvb5dRQ6rJjFSdX2HZY1/gI=",
    version = "v0.3.7",
)

go_repository(
    name = "com_github_improbable_eng_thanos",
    importpath = "github.com/improbable-eng/thanos",
    sum = "h1:iZfU7exq+RD5Lnb8n3Eh9MNYoRLeyeGO/85AvEkLg+8=",
    version = "v0.3.2",
)

go_repository(
    name = "com_github_inconshreveable_mousetrap",
    importpath = "github.com/inconshreveable/mousetrap",
    sum = "h1:wN+x4NVGpMsO7ErUn/mUI3vEoE6Jt13X2s0bqwp9tc8=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_influxdata_influxdb1_client",
    importpath = "github.com/influxdata/influxdb1-client",
    sum = "h1:/WZQPMZNsjZ7IlCpsLGdQBINg5bxKQ1K1sh6awxLtkA=",
    version = "v0.0.0-20191209144304-8bf82d3c094d",
)

go_repository(
    name = "com_github_jmespath_go_jmespath",
    importpath = "github.com/jmespath/go-jmespath",
    sum = "h1:pmfjZENx5imkbgOkpRUYLnmbU7UEFbjtDA2hxJ1ichM=",
    version = "v0.0.0-20180206201540-c2b33e8439af",
)

go_repository(
    name = "com_github_joefitzgerald_rainbow_reporter",
    importpath = "github.com/joefitzgerald/rainbow-reporter",
    sum = "h1:AuMG652zjdzI0YCCnXAqATtRBpGXMcAnrajcaTrSeuo=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_jonboulle_clockwork",
    importpath = "github.com/jonboulle/clockwork",
    sum = "h1:UOGuzwb1PwsrDAObMuhUnj0p5ULPj8V/xJ7Kx9qUBdQ=",
    version = "v0.2.2",
)

go_repository(
    name = "com_github_jpillora_backoff",
    importpath = "github.com/jpillora/backoff",
    sum = "h1:uvFg412JmmHBHw7iwprIxkPMI+sGQ4kzOWsMeHnm2EA=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_json_iterator_go",
    importpath = "github.com/json-iterator/go",
    sum = "h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=",
    version = "v1.1.12",
)

go_repository(
    name = "com_github_jsonnet_bundler_jsonnet_bundler",
    importpath = "github.com/jsonnet-bundler/jsonnet-bundler",
    sum = "h1:T/HtHFr+mYCRULrH1x/RnoB0prIs0rMkolJhFMXNC9A=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_jstemmer_go_junit_report",
    importpath = "github.com/jstemmer/go-junit-report",
    sum = "h1:6QPYqodiu3GuPL+7mfx+NwDdp2eTkp9IfEUpgAwUN0o=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_jtolds_gls",
    importpath = "github.com/jtolds/gls",
    sum = "h1:xdiiI2gbIgH/gLH7ADydsJ1uDOEzR8yvV7C0MuV77Wo=",
    version = "v4.20.0+incompatible",
)

go_repository(
    name = "com_github_julienschmidt_httprouter",
    importpath = "github.com/julienschmidt/httprouter",
    sum = "h1:U0609e9tgbseu3rBINet9P48AI/D3oJs4dN7jwJOQ1U=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_k8snetworkplumbingwg_network_attachment_definition_client",
    importpath = "github.com/k8snetworkplumbingwg/network-attachment-definition-client",
    sum = "h1:VzM3TYHDgqPkettiP6I6q2jOeQFL4nrJM+UcAc4f6Fs=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_kelseyhightower_envconfig",
    importpath = "github.com/kelseyhightower/envconfig",
    sum = "h1:Im6hONhd3pLkfDFsbRgu68RDNkGF1r3dvMUtDTo2cv8=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_kisielk_errcheck",
    importpath = "github.com/kisielk/errcheck",
    sum = "h1:e8esj/e4R+SAOwFwN+n3zr0nYeCyeweozKfO23MvHzY=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_kisielk_gotool",
    importpath = "github.com/kisielk/gotool",
    sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_klauspost_compress",
    importpath = "github.com/klauspost/compress",
    sum = "h1:eLeJ3dr/Y9+XRfJT4l+8ZjmtB5RPJhucH2HeCV5+IZY=",
    version = "v1.10.8",
)

go_repository(
    name = "com_github_klauspost_pgzip",
    importpath = "github.com/klauspost/pgzip",
    sum = "h1:TQ7CNpYKovDOmqzRHKxJh0BeaBI7UdQZYc6p7pMQh1A=",
    version = "v1.2.4",
)

go_repository(
    name = "com_github_knetic_govaluate",
    importpath = "github.com/Knetic/govaluate",
    sum = "h1:1G1pk05UrOh0NlF1oeaaix1x8XzrfjIDK47TY0Zehcw=",
    version = "v3.0.1-0.20171022003610-9aa49832a739+incompatible",
)

go_repository(
    name = "com_github_konsorten_go_windows_terminal_sequences",
    importpath = "github.com/konsorten/go-windows-terminal-sequences",
    sum = "h1:CE8S1cTafDpPvMhIxNJKvHsGVBgn1xWYf1NbHQhywc8=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_konveyor_controller",
    importpath = "github.com/konveyor/controller",
    sum = "h1:It0jfQEyth6AC5z8f4oTmd4RaRsTlKubKfLh/TlLYqk=",
    version = "v0.10.0",
)

go_repository(
    name = "com_github_kr_logfmt",
    importpath = "github.com/kr/logfmt",
    sum = "h1:T+h1c/A9Gawja4Y9mFVWj2vyii2bbUNDw3kt9VxK2EY=",
    version = "v0.0.0-20140226030751-b84e30acd515",
)

go_repository(
    name = "com_github_kr_pretty",
    importpath = "github.com/kr/pretty",
    sum = "h1:flRD4NNwYAUpkphVc1HcthR4KEIFJ65n8Mw5qdRn3LE=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_kr_pty",
    importpath = "github.com/kr/pty",
    sum = "h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_kr_text",
    importpath = "github.com/kr/text",
    sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_kubernetes_csi_csi_lib_utils",
    importpath = "github.com/kubernetes-csi/csi-lib-utils",
    sum = "h1:t1cS7HTD7z5D7h9iAdjWuHtMxJPb9s1fIv34rxytzqs=",
    version = "v0.7.0",
)

go_repository(
    name = "com_github_kubernetes_csi_csi_test",
    importpath = "github.com/kubernetes-csi/csi-test",
    sum = "h1:ia04uVFUM/J9n/v3LEMn3rEG6FmKV5BH9QLw7H68h44=",
    version = "v2.0.0+incompatible",
)

go_repository(
    name = "com_github_kubernetes_csi_external_snapshotter_v2",
    importpath = "github.com/kubernetes-csi/external-snapshotter/v2",
    sum = "h1:t5bmB3Y8nCaLA4aFrIpX0zjHEF/HUkJp6f5rm7BsVzM=",
    version = "v2.1.1",
)

go_repository(
    name = "com_github_kubernetes_sigs_kube_storage_version_migrator",
    importpath = "github.com/kubernetes-sigs/kube-storage-version-migrator",
    sum = "h1:XZfEWaxPMR/cRvQ/SFWOjok7YEnURVYrCiIltx/0HGY=",
    version = "v0.0.0-20191127225502-51849bc15f17",
)

go_repository(
    name = "com_github_kylelemons_godebug",
    importpath = "github.com/kylelemons/godebug",
    sum = "h1:MtvEpTB6LX3vkb4ax0b5D2DHbNAUsen0Gx5wZoq3lV4=",
    version = "v0.0.0-20170820004349-d65d576e9348",
)

go_repository(
    name = "com_github_leodido_go_urn",
    importpath = "github.com/leodido/go-urn",
    sum = "h1:XlAE/cm/ms7TE/VMVoduSpNBoyc2dOxHs5MZSwAN63Q=",
    version = "v1.2.4",
)

go_repository(
    name = "com_github_liggitt_tabwriter",
    importpath = "github.com/liggitt/tabwriter",
    sum = "h1:9TO3cAIGXtEhnIaL+V+BEER86oLrvS+kWobKpbJuye0=",
    version = "v0.0.0-20181228230101-89fcab3d43de",
)

go_repository(
    name = "com_github_lightstep_lightstep_tracer_common_golang_gogo",
    importpath = "github.com/lightstep/lightstep-tracer-common/golang/gogo",
    sum = "h1:143Bb8f8DuGWck/xpNUOckBVYfFbBTnLevfRZ1aVVqo=",
    version = "v0.0.0-20190605223551-bc2310a04743",
)

go_repository(
    name = "com_github_lightstep_lightstep_tracer_go",
    importpath = "github.com/lightstep/lightstep-tracer-go",
    sum = "h1:vi1F1IQ8N7hNWytK9DpJsUfQhGuNSc19z330K6vl4zk=",
    version = "v0.18.1",
)

go_repository(
    name = "com_github_lyft_protoc_gen_validate",
    importpath = "github.com/lyft/protoc-gen-validate",
    sum = "h1:KNt/RhmQTOLr7Aj8PsJ7mTronaFyx80mRTT9qF261dA=",
    version = "v0.0.13",
)

go_repository(
    name = "com_github_magiconair_properties",
    importpath = "github.com/magiconair/properties",
    sum = "h1:ZC2Vc7/ZFkGmsVC9KvOjumD+G5lXy2RtTKyzRKO2BQ4=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_mailru_easyjson",
    importpath = "github.com/mailru/easyjson",
    sum = "h1:UGYAvKxe3sBsEDzO8ZeWOSlIQfWFlxbzLZe7hwFURr0=",
    version = "v0.7.7",
)

go_repository(
    name = "com_github_mattn_go_colorable",
    importpath = "github.com/mattn/go-colorable",
    sum = "h1:c1ghPdyEDarC70ftn0y+A/Ee++9zz8ljHG1b13eJ0s8=",
    version = "v0.1.8",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    importpath = "github.com/mattn/go-isatty",
    sum = "h1:JITubQf0MOLdlGRuRq+jtsDlekdYPia9ZFsB8h/APPA=",
    version = "v0.0.19",
)

go_repository(
    name = "com_github_mattn_go_runewidth",
    importpath = "github.com/mattn/go-runewidth",
    sum = "h1:Lm995f3rfxdpd6TSmuVCHVb/QhupuXlYr8sCI/QdE+0=",
    version = "v0.0.9",
)

go_repository(
    name = "com_github_mattn_go_shellwords",
    importpath = "github.com/mattn/go-shellwords",
    sum = "h1:Y7Xqm8piKOO3v10Thp7Z36h4FYFjt5xB//6XvOrs2Gw=",
    version = "v1.0.10",
)

go_repository(
    name = "com_github_mattn_go_sqlite3",
    importpath = "github.com/mattn/go-sqlite3",
    sum = "h1:JL0eqdCOq6DJVNPSvArO/bIV9/P7fbGrV00LZHc+5aI=",
    version = "v1.14.18",
)

go_repository(
    name = "com_github_matttproud_golang_protobuf_extensions",
    importpath = "github.com/matttproud/golang_protobuf_extensions",
    sum = "h1:mmDVorXM7PCGKw94cs5zkfA9PSy5pEvNWRP0ET0TIVo=",
    version = "v1.0.4",
)

go_repository(
    name = "com_github_maxbrunsfeld_counterfeiter",
    importpath = "github.com/maxbrunsfeld/counterfeiter",
    sum = "h1:fJasMUaV/LYZvzK4bUOj13rNXc4fhVzU0Vu1OlcGUd4=",
    version = "v0.0.0-20181017030959-1aadac120687",
)

go_repository(
    name = "com_github_maxbrunsfeld_counterfeiter_v6",
    importpath = "github.com/maxbrunsfeld/counterfeiter/v6",
    sum = "h1:s0HwWQiNYF+YpoOncE8OxHVYG3YShNiRG8iuPDiSDWM=",
    version = "v6.2.1",
)

go_repository(
    name = "com_github_microsoft_go_winio",
    importpath = "github.com/Microsoft/go-winio",
    sum = "h1:ygIc8M6trr62pF5DucadTWGdEB4mEyvzi0e2nbcmcyA=",
    version = "v0.4.15-0.20190919025122-fc70bd9a86b5",
)

go_repository(
    name = "com_github_microsoft_hcsshim",
    importpath = "github.com/Microsoft/hcsshim",
    sum = "h1:VrfodqvztU8YSOvygU+DN1BGaSGxmrNfqOv5oOuX2Bk=",
    version = "v0.8.9",
)

go_repository(
    name = "com_github_miekg_dns",
    importpath = "github.com/miekg/dns",
    sum = "h1:9jZdLNd/P4+SfEJ0TNyxYpsK8N4GtfylBLqtbYN1sbA=",
    version = "v1.0.14",
)

go_repository(
    name = "com_github_mistifyio_go_zfs",
    importpath = "github.com/mistifyio/go-zfs",
    sum = "h1:gAMO1HM9xBRONLHHYnu5iFsOJUiJdNZo6oqSENd4eW8=",
    version = "v2.1.1+incompatible",
)

go_repository(
    name = "com_github_mitchellh_cli",
    importpath = "github.com/mitchellh/cli",
    sum = "h1:iGBIsUe3+HZ/AD/Vd7DErOt5sU9fa8Uj7A2s1aggv1Y=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mitchellh_go_homedir",
    importpath = "github.com/mitchellh/go-homedir",
    sum = "h1:lukF9ziXFxDFPkA1vsr5zpc1XuPDn/wFntq5mG+4E0Y=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_mitchellh_go_testing_interface",
    importpath = "github.com/mitchellh/go-testing-interface",
    sum = "h1:fzU/JVNcaqHQEcVFAKeR41fkiLdIPrefOvVG1VZ96U0=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mitchellh_gox",
    importpath = "github.com/mitchellh/gox",
    sum = "h1:lfGJxY7ToLJQjHHwi0EX6uYBdK78egf954SQl13PQJc=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_mitchellh_hashstructure",
    importpath = "github.com/mitchellh/hashstructure",
    sum = "h1:ZkRJX1CyOoTkar7p/mLS5TZU4nJ1Rn/F8u9dGS02Q3Y=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mitchellh_iochan",
    importpath = "github.com/mitchellh/iochan",
    sum = "h1:C+X3KsSTLFVBr/tK1eYN/vs4rJcvsiLU338UhYPJWeY=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    importpath = "github.com/mitchellh/mapstructure",
    sum = "h1:CpVNEelQCZBooIPDn+AR3NpivK/TIKU8bDxdASFVQag=",
    version = "v1.4.1",
)

go_repository(
    name = "com_github_moby_term",
    importpath = "github.com/moby/term",
    sum = "h1:HfkjXDfhgVaN5rmueG8cL8KKeFNecRCXFhaJ2qZ5SKA=",
    version = "v0.0.0-20221205130635-1aeaba878587",
)

go_repository(
    name = "com_github_modern_go_concurrent",
    importpath = "github.com/modern-go/concurrent",
    sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
    version = "v0.0.0-20180306012644-bacd9c7ef1dd",
)

go_repository(
    name = "com_github_modern_go_reflect2",
    importpath = "github.com/modern-go/reflect2",
    sum = "h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_morikuni_aec",
    importpath = "github.com/morikuni/aec",
    sum = "h1:nP9CBfwrvYnBRgY6qfDQkygYDmYwOilePFkwzv4dU8A=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mrnold_go_libnbd",
    importpath = "github.com/mrnold/go-libnbd",
    sum = "h1:bIszEpQZKre4DMqEIO5HCv/MA44ujD2w7BIQg/cOFts=",
    version = "v1.4.1-cdi",
)

go_repository(
    name = "com_github_mtrmac_gpgme",
    importpath = "github.com/mtrmac/gpgme",
    sum = "h1:dNOmvYmsrakgW7LcgiprD0yfRuQQe8/C8F6Z+zogO3s=",
    version = "v0.1.2",
)

go_repository(
    name = "com_github_munnerz_goautoneg",
    importpath = "github.com/munnerz/goautoneg",
    sum = "h1:C3w9PqII01/Oq1c1nUAm88MOHcQC9l5mIlSMApZMrHA=",
    version = "v0.0.0-20191010083416-a7dc8b61c822",
)

go_repository(
    name = "com_github_mwitkow_go_conntrack",
    importpath = "github.com/mwitkow/go-conntrack",
    sum = "h1:KUppIJq7/+SVif2QVs3tOP0zanoHgBEVAwHxUSIzRqU=",
    version = "v0.0.0-20190716064945-2f068394615f",
)

go_repository(
    name = "com_github_mxk_go_flowrate",
    importpath = "github.com/mxk/go-flowrate",
    sum = "h1:y5//uYreIhSUg3J1GEMiLbxo1LJaP8RfCpH6pymGZus=",
    version = "v0.0.0-20140419014527-cca7078d478f",
)

go_repository(
    name = "com_github_nats_io_jwt",
    importpath = "github.com/nats-io/jwt",
    sum = "h1:+RB5hMpXUUA2dfxuhBTEkMOrYmM+gKIZYS1KjSostMI=",
    version = "v0.3.2",
)

go_repository(
    name = "com_github_nats_io_nats_go",
    importpath = "github.com/nats-io/nats.go",
    sum = "h1:ik3HbLhZ0YABLto7iX80pZLPw/6dx3T+++MZJwLnMrQ=",
    version = "v1.9.1",
)

go_repository(
    name = "com_github_nats_io_nats_server_v2",
    importpath = "github.com/nats-io/nats-server/v2",
    sum = "h1:i2Ly0B+1+rzNZHHWtD4ZwKi+OU5l+uQo1iDHZ2PmiIc=",
    version = "v2.1.2",
)

go_repository(
    name = "com_github_nats_io_nkeys",
    importpath = "github.com/nats-io/nkeys",
    sum = "h1:6JrEfig+HzTH85yxzhSVbjHRJv9cn0p6n3IngIcM5/k=",
    version = "v0.1.3",
)

go_repository(
    name = "com_github_nats_io_nuid",
    importpath = "github.com/nats-io/nuid",
    sum = "h1:5iA8DT8V7q8WK2EScv2padNa/rTESc1KdnPw4TC2paw=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_nxadm_tail",
    importpath = "github.com/nxadm/tail",
    sum = "h1:8feyoE3OzPrcshW5/MJ4sGESc5cqmGkGCWlco4l0bqY=",
    version = "v1.4.11",
)

go_repository(
    name = "com_github_nytimes_gziphandler",
    importpath = "github.com/NYTimes/gziphandler",
    sum = "h1:ZUDjpQae29j0ryrS0u/B8HZfJBtBQHjqw2rQ2cqUQ3I=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_oklog_oklog",
    importpath = "github.com/oklog/oklog",
    sum = "h1:wVfs8F+in6nTBMkA7CbRw+zZMIB7nNM825cM1wuzoTk=",
    version = "v0.3.2",
)

go_repository(
    name = "com_github_oklog_run",
    importpath = "github.com/oklog/run",
    sum = "h1:Ru7dDtJNOyC66gQ5dQmaCa0qIsAUFY3sFpK1Xk8igrw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_oklog_ulid",
    importpath = "github.com/oklog/ulid",
    sum = "h1:EGfNDEx6MqHz8B3uNV6QAib1UR2Lm97sHi3ocA6ESJ4=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_olekukonko_tablewriter",
    importpath = "github.com/olekukonko/tablewriter",
    sum = "h1:58+kh9C6jJVXYjt8IE48G2eWl6BjwU5Gj0gqY84fy78=",
    version = "v0.0.0-20170122224234-a0225b3f23b5",
)

go_repository(
    name = "com_github_oneofone_xxhash",
    importpath = "github.com/OneOfOne/xxhash",
    sum = "h1:KMrpdQIwFcEqXDklaen+P1axHaj9BSKzvpUUfnHldSE=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_onsi_ginkgo",
    importpath = "github.com/onsi/ginkgo",
    sum = "h1:8xi0RTUf59SOSfEtZMvwTvXYMzG4gV23XVHOZiXNtnE=",
    version = "v1.16.5",
)

go_repository(
    name = "com_github_onsi_gomega",
    importpath = "github.com/onsi/gomega",
    sum = "h1:hvMK7xYz4D3HapigLTeGdId/NcfQx1VHMJc60ew99+8=",
    version = "v1.30.0",
)

go_repository(
    name = "com_github_op_go_logging",
    importpath = "github.com/op/go-logging",
    sum = "h1:lDH9UUVJtmYCjyT0CI4q8xvlXPxeZ0gYCVvWbmPlp88=",
    version = "v0.0.0-20160315200505-970db520ece7",
)

go_repository(
    name = "com_github_opencontainers_go_digest",
    importpath = "github.com/opencontainers/go-digest",
    sum = "h1:apOUWs51W5PlhuyGyz9FCeeBIOUDA/6nW8Oi/yOhh5U=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_opencontainers_image_spec",
    importpath = "github.com/opencontainers/image-spec",
    sum = "h1:9yCKha/T5XdGtO0q9Q9a6T5NUCsTn/DrBg0D7ufOcFM=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_opencontainers_runc",
    importpath = "github.com/opencontainers/runc",
    sum = "h1:4+xo8mtWixbHoEm451+WJNUrq12o2/tDsyK9Vgc/NcA=",
    version = "v1.0.0-rc90",
)

go_repository(
    name = "com_github_opencontainers_runtime_spec",
    importpath = "github.com/opencontainers/runtime-spec",
    sum = "h1:eNUVfm/RFLIi1G7flU5/ZRTHvd4kcVuzfRnL6OFlzCI=",
    version = "v0.1.2-0.20190507144316-5b71a03e2700",
)

go_repository(
    name = "com_github_opencontainers_selinux",
    importpath = "github.com/opencontainers/selinux",
    sum = "h1:F6DgIsjgBIcDksLW4D5RG9bXok6oqZ3nvMwj4ZoFu/Q=",
    version = "v1.5.2",
)

go_repository(
    name = "com_github_openshift_api",
    importpath = "github.com/openshift/api",
    sum = "h1:sgr89m3ejIIKhSbTtHq7HEZ80et4IAXDrJlk+u+rYX8=",
    version = "v0.0.0-20230613151523-ba04973d3ed1",
)

go_repository(
    name = "com_github_openshift_build_machinery_go",
    importpath = "github.com/openshift/build-machinery-go",
    sum = "h1:RR4ah7FfaPR1WePizm0jlrsbmPu91xQZnAsVVreQV1k=",
    version = "v0.0.0-20220913142420-e25cf57ea46d",
)

go_repository(
    name = "com_github_openshift_client_go",
    importpath = "github.com/openshift/client-go",
    sum = "h1:Nij5OnaECrkmcRQMAE9LMbQXPo95aqFnf+12B7SyFVI=",
    version = "v0.0.0-20230503144108-75015d2347cb",
)

go_repository(
    name = "com_github_openshift_custom_resource_status",
    importpath = "github.com/openshift/custom-resource-status",
    sum = "h1:C3DL44LEbvlbItfd8mT5jWrqPfHnSOQoQf/sypqA6A4=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_openshift_library_go",
    importpath = "github.com/openshift/library-go",
    sum = "h1:RXJqcNJPjUFlZA27J2fbL7KVMBUvhPxWiH3YkSjrR7w=",
    version = "v0.0.0-20230626162119-954ade536d6d",
)

go_repository(
    name = "com_github_openshift_prom_label_proxy",
    importpath = "github.com/openshift/prom-label-proxy",
    sum = "h1:GW8OxGwBbI2kCqjb5PQfVXRAuCJbYyX1RYs9R3ISjck=",
    version = "v0.1.1-0.20191016113035-b8153a7f39f1",
)

go_repository(
    name = "com_github_opentracing_basictracer_go",
    importpath = "github.com/opentracing/basictracer-go",
    sum = "h1:YyUAhaEfjoWXclZVJ9sGoNct7j4TVk7lZWlQw5UXuoo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_opentracing_contrib_go_observer",
    importpath = "github.com/opentracing-contrib/go-observer",
    sum = "h1:lM6RxxfUMrYL/f8bWEUqdXrANWtrL7Nndbm9iFN0DlU=",
    version = "v0.0.0-20170622124052-a52f23424492",
)

go_repository(
    name = "com_github_opentracing_opentracing_go",
    importpath = "github.com/opentracing/opentracing-go",
    sum = "h1:pWlfV3Bxv7k65HYwkikxat0+s3pV4bsqf19k25Ur8rU=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_openzipkin_contrib_zipkin_go_opentracing",
    importpath = "github.com/openzipkin-contrib/zipkin-go-opentracing",
    sum = "h1:ZCnq+JUrvXcDVhX/xRolRBZifmabN1HcS1wrPSvxhrU=",
    version = "v0.4.5",
)

go_repository(
    name = "com_github_openzipkin_zipkin_go",
    importpath = "github.com/openzipkin/zipkin-go",
    sum = "h1:nY8Hti+WKaP0cRsSeQ026wU03QsM762XBeCXBb9NAWI=",
    version = "v0.2.2",
)

go_repository(
    name = "com_github_operator_framework_go_appr",
    importpath = "github.com/operator-framework/go-appr",
    sum = "h1:c7gnBIMtxxenMKXZjeCuQaDfn7IGGmgh9laEGsOEeU4=",
    version = "v0.0.0-20180917210448-f2aef88446f2",
)

go_repository(
    name = "com_github_operator_framework_operator_lifecycle_manager",
    importpath = "github.com/operator-framework/operator-lifecycle-manager",
    sum = "h1:500FBy57oogxOEDRQpDtVD1E5ZzmWh2n2iliXZv169s=",
    version = "v0.0.0-20190725173916-b56e63a643cc",
)

go_repository(
    name = "com_github_operator_framework_operator_marketplace",
    importpath = "github.com/operator-framework/operator-marketplace",
    sum = "h1:47MQUQRBZqwyTPLEHoFlbGRv63p0OvxpPp5g6FUQXQs=",
    version = "v0.0.0-20190216021216-57300a3ef3ba",
)

go_repository(
    name = "com_github_operator_framework_operator_registry",
    importpath = "github.com/operator-framework/operator-registry",
    sum = "h1:oDIevJvKXFsp7BEb7iJHuLvuhPZYBtIx5oZQ7iSISAs=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_ostreedev_ostree_go",
    importpath = "github.com/ostreedev/ostree-go",
    sum = "h1:TnbXhKzrTOyuvWrjI8W6pcoI9XPbLHFXCdN2dtUw7Rw=",
    version = "v0.0.0-20190702140239-759a8c1ac913",
)

go_repository(
    name = "com_github_ovirt_go_ovirt",
    importpath = "github.com/ovirt/go-ovirt",
    sum = "h1:SGD34D2dmrTzqyLToETTuZnN+YxZF1QaXsvqrCOkIPo=",
    version = "v0.0.0-20230808190322-9fd1992199b2",
)

go_repository(
    name = "com_github_pact_foundation_pact_go",
    importpath = "github.com/pact-foundation/pact-go",
    sum = "h1:OYkFijGHoZAYbOIb1LWXrwKQbMMRUv1oQ89blD2Mh2Q=",
    version = "v1.0.4",
)

go_repository(
    name = "com_github_pascaldekloe_goe",
    importpath = "github.com/pascaldekloe/goe",
    sum = "h1:Lgl0gzECD8GnQ5QCWA8o6BtfL6mDH5rQgM4/fX3avOs=",
    version = "v0.0.0-20180627143212-57f6aae5913c",
)

go_repository(
    name = "com_github_pborman_uuid",
    importpath = "github.com/pborman/uuid",
    sum = "h1:+ZZIw58t/ozdjRaXh/3awHfmWRbzYxJoAdNJxe/3pvw=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_pelletier_go_toml",
    importpath = "github.com/pelletier/go-toml",
    sum = "h1:T5zMGML61Wp+FlcbWjRDT7yAxhJNAiPPLOFECq181zc=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_performancecopilot_speed",
    importpath = "github.com/performancecopilot/speed",
    sum = "h1:2WnRzIquHa5QxaJKShDkLM+sc0JPuwhXzK8OYOyt3Vg=",
    version = "v3.0.0+incompatible",
)

go_repository(
    name = "com_github_petar_gollrb",
    importpath = "github.com/petar/GoLLRB",
    sum = "h1:AwcgVYzW1T+QuJ2fc55ceOSCiVaOpdYUNpFj9t7+n9U=",
    version = "v0.0.0-20130427215148-53be0d36a84c",
)

go_repository(
    name = "com_github_peterbourgon_diskv",
    importpath = "github.com/peterbourgon/diskv",
    sum = "h1:UBdAOUP5p4RWqPBg048CAvpKN+vxiaj6gdUUzhl4XmI=",
    version = "v2.0.1+incompatible",
)

go_repository(
    name = "com_github_pierrec_lz4",
    importpath = "github.com/pierrec/lz4",
    sum = "h1:2xWsjqPFWcplujydGg4WmhC/6fZqK42wMM8aXeqhl0I=",
    version = "v2.0.5+incompatible",
)

go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_pkg_profile",
    importpath = "github.com/pkg/profile",
    sum = "h1:hnbDkaNWPCLMO9wGLdBFTIZvzDrDfBM2072E1S9gJkA=",
    version = "v1.7.0",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:Jamvg5psRIccs7FGNTlIRMkT8wgtp5eCXdBlqhYGL6U=",
    version = "v1.0.1-0.20181226105442-5d4384ee4fb2",
)

go_repository(
    name = "com_github_posener_complete",
    importpath = "github.com/posener/complete",
    sum = "h1:ccV59UEOTzVDnDUEFdT95ZzHVZ+5+158q8+SJb2QV5w=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_pquerna_cachecontrol",
    importpath = "github.com/pquerna/cachecontrol",
    sum = "h1:yJMy84ti9h/+OEWa752kBTKv4XC30OtVVHYv/8cTqKc=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_pquerna_ffjson",
    importpath = "github.com/pquerna/ffjson",
    sum = "h1:kyf9snWXHvQc+yxE9imhdI8YAm4oKeZISlaAR+x73zs=",
    version = "v0.0.0-20190813045741-dac163c6c0a9",
)

go_repository(
    name = "com_github_prometheus_client_golang",
    importpath = "github.com/prometheus/client_golang",
    sum = "h1:rl2sfwZMtSthVU752MqfjQozy7blglC+1SOtjMAMh+Q=",
    version = "v1.17.0",
)

go_repository(
    name = "com_github_prometheus_client_model",
    importpath = "github.com/prometheus/client_model",
    sum = "h1:VQw1hfvPvk3Uv6Qf29VrPF32JB6rtbgI6cYPYQjL0Qw=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_prometheus_common",
    importpath = "github.com/prometheus/common",
    sum = "h1:2BGz0eBc2hdMDLnO/8n0jeB3oPrt2D08CekT0lneoxM=",
    version = "v0.45.0",
)

go_repository(
    name = "com_github_prometheus_procfs",
    importpath = "github.com/prometheus/procfs",
    sum = "h1:xRC8Iq1yyca5ypa9n1EZnWZkt7dwcoRPQwX/5gwaUuI=",
    version = "v0.11.1",
)

go_repository(
    name = "com_github_prometheus_prometheus",
    importpath = "github.com/prometheus/prometheus",
    sum = "h1:EekL1S9WPoPtJL2NZvL+xo38iMpraOnyEHOiyZygMDY=",
    version = "v2.3.2+incompatible",
)

go_repository(
    name = "com_github_prometheus_tsdb",
    importpath = "github.com/prometheus/tsdb",
    sum = "h1:w1tAGxsBMLkuGrFMhqgcCeBkM5d1YI24udArs+aASuQ=",
    version = "v0.8.0",
)

go_repository(
    name = "com_github_puerkitobio_purell",
    importpath = "github.com/PuerkitoBio/purell",
    sum = "h1:WEQqlqaGbrPkxLJWfBwQmfEAE1Z7ONdDLqrN38tNFfI=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_puerkitobio_urlesc",
    importpath = "github.com/PuerkitoBio/urlesc",
    sum = "h1:d+Bc7a5rLufV/sSk/8dngufqelfh6jnri85riMAaF/M=",
    version = "v0.0.0-20170810143723-de5bf2ad4578",
)

go_repository(
    name = "com_github_rcrowley_go_metrics",
    importpath = "github.com/rcrowley/go-metrics",
    sum = "h1:9ZKAASQSHhDYGoxY8uLVpewe1GDZ2vu2Tr/vTdVAkFQ=",
    version = "v0.0.0-20181016184325-3113b8401b8a",
)

go_repository(
    name = "com_github_robfig_cron",
    importpath = "github.com/robfig/cron",
    sum = "h1:ZjScXvvxeQ63Dbyxy76Fj3AT3Ut0aKsyd2/tl3DTMuQ=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_rogpeppe_fastuuid",
    importpath = "github.com/rogpeppe/fastuuid",
    sum = "h1:gu+uRPtBe88sKxUCEXRoeCvVG90TJmwhiqRpvdhQFng=",
    version = "v0.0.0-20150106093220-6724a57986af",
)

go_repository(
    name = "com_github_rogpeppe_go_charset",
    importpath = "github.com/rogpeppe/go-charset",
    sum = "h1:BN/Nyn2nWMoqGRA7G7paDNDqTXE30mXGqzzybrfo05w=",
    version = "v0.0.0-20180617210344-2471d30d28b4",
)

go_repository(
    name = "com_github_rogpeppe_go_internal",
    importpath = "github.com/rogpeppe/go-internal",
    sum = "h1:TMyTOH3F/DB16zRVcYyreMH6GnZZrwQVAoYjRBZyWFQ=",
    version = "v1.10.0",
)

go_repository(
    name = "com_github_rs_cors",
    importpath = "github.com/rs/cors",
    sum = "h1:+88SsELBHx5r+hZ8TCkggzSstaWNbDvThkVK8H6f9ik=",
    version = "v1.7.0",
)

go_repository(
    name = "com_github_russross_blackfriday",
    importpath = "github.com/russross/blackfriday",
    sum = "h1:HyvC0ARfnZBqnXwABFeSZHpKvJHJJfPz81GNueLj0oo=",
    version = "v1.5.2",
)

go_repository(
    name = "com_github_russross_blackfriday_v2",
    importpath = "github.com/russross/blackfriday/v2",
    sum = "h1:JIOH55/0cWyOuilr9/qlrm0BSXldqnqwMsf35Ld67mk=",
    version = "v2.1.0",
)

go_repository(
    name = "com_github_ryanuber_columnize",
    importpath = "github.com/ryanuber/columnize",
    sum = "h1:UFr9zpz4xgTnIE5yIMtWAMngCdZ9p/+q6lTbgelo80M=",
    version = "v0.0.0-20160712163229-9b3edd62028f",
)

go_repository(
    name = "com_github_samuel_go_zookeeper",
    importpath = "github.com/samuel/go-zookeeper",
    sum = "h1:p3Vo3i64TCLY7gIfzeQaUJ+kppEO5WQG3cL8iE8tGHU=",
    version = "v0.0.0-20190923202752-2cc03de413da",
)

go_repository(
    name = "com_github_sclevine_spec",
    importpath = "github.com/sclevine/spec",
    sum = "h1:1Jwdf9jSfDl9NVmt8ndHqbTZ7XCCPbh1jI3hkDBHVYA=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_sean_seed",
    importpath = "github.com/sean-/seed",
    sum = "h1:nn5Wsu0esKSJiIVhscUtVbo7ada43DJhG55ua/hjS5I=",
    version = "v0.0.0-20170313163322-e2103e2c3529",
)

go_repository(
    name = "com_github_sergi_go_diff",
    importpath = "github.com/sergi/go-diff",
    sum = "h1:Kpca3qRNrduNnOQeazBd0ysaKrUJiIuISHxogkT9RPQ=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_shopify_sarama",
    importpath = "github.com/Shopify/sarama",
    sum = "h1:9oksLxC6uxVPHPVYUmq6xhr1BOF/hHobWH2UzO67z1s=",
    version = "v1.19.0",
)

go_repository(
    name = "com_github_shopify_toxiproxy",
    importpath = "github.com/Shopify/toxiproxy",
    sum = "h1:TKdv8HiTLgE5wdJuEML90aBgNWsokNbMijUGhmcoBJc=",
    version = "v2.1.4+incompatible",
)

go_repository(
    name = "com_github_shurcool_sanitized_anchor_name",
    importpath = "github.com/shurcooL/sanitized_anchor_name",
    sum = "h1:PdmoCO6wvbs+7yrJyMORt4/BmY5IYyJwS/kOiWx8mHo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    sum = "h1:trlNQbNUG3OdDrDil03MCb1H2o9nJ1x4/5LYw7byDE0=",
    version = "v1.9.0",
)

go_repository(
    name = "com_github_smartystreets_assertions",
    importpath = "github.com/smartystreets/assertions",
    sum = "h1:zE9ykElWQ6/NYmHa3jpm/yHnI4xSofP+UP6SpjHcSeM=",
    version = "v0.0.0-20180927180507-b2de0cb4f26d",
)

go_repository(
    name = "com_github_smartystreets_goconvey",
    importpath = "github.com/smartystreets/goconvey",
    sum = "h1:fv0U8FUIMPNf1L9lnHLvLhgicrIVChEkdzIKYqbNC9s=",
    version = "v1.6.4",
)

go_repository(
    name = "com_github_soheilhy_cmux",
    importpath = "github.com/soheilhy/cmux",
    sum = "h1:jjzc5WVemNEDTLwv9tlmemhC73tI08BNOIGwBOo10Js=",
    version = "v0.1.5",
)

go_repository(
    name = "com_github_sony_gobreaker",
    importpath = "github.com/sony/gobreaker",
    sum = "h1:oMnRNZXX5j85zso6xCPRNPtmAycat+WcoKbklScLDgQ=",
    version = "v0.4.1",
)

go_repository(
    name = "com_github_spaolacci_murmur3",
    importpath = "github.com/spaolacci/murmur3",
    sum = "h1:qLC7fQah7D6K1B0ujays3HV9gkFtllcxhzImRR7ArPQ=",
    version = "v0.0.0-20180118202830-f09979ecbc72",
)

go_repository(
    name = "com_github_spf13_afero",
    importpath = "github.com/spf13/afero",
    sum = "h1:5jhuqJyZCZf2JRofRvN/nIFgIWNzPa3/Vz8mYylgbWc=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_spf13_cast",
    importpath = "github.com/spf13/cast",
    sum = "h1:oget//CVOEoFewqQxwr0Ej5yjygnqGkvggSE/gB35Q8=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_spf13_cobra",
    importpath = "github.com/spf13/cobra",
    sum = "h1:e5/vxKd/rZsfSJMUX1agtjeTDf+qv1/JdBF8gg5k9ZM=",
    version = "v1.8.1",
)

go_repository(
    name = "com_github_spf13_jwalterweatherman",
    importpath = "github.com/spf13/jwalterweatherman",
    sum = "h1:XHEdyB+EcvlqZamSM4ZOMGlc93t6AcsBEu9Gc1vn7yk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_spf13_pflag",
    importpath = "github.com/spf13/pflag",
    sum = "h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=",
    version = "v1.0.5",
)

go_repository(
    name = "com_github_spf13_viper",
    importpath = "github.com/spf13/viper",
    sum = "h1:xVKxvI7ouOI5I+U9s2eeiUfMaWBVoXA3AWskkrqK0VM=",
    version = "v1.7.0",
)

go_repository(
    name = "com_github_stevvooe_resumable",
    importpath = "github.com/stevvooe/resumable",
    sum = "h1:4bT0pPowCpQImewr+BjzfUKcuFW+KVyB8d1OF3b6oTI=",
    version = "v0.0.0-20180830230917-22b14a53ba50",
)

go_repository(
    name = "com_github_streadway_amqp",
    importpath = "github.com/streadway/amqp",
    sum = "h1:WhxRHzgeVGETMlmVfqhRn8RIeeNoPr2Czh33I4Zdccw=",
    version = "v0.0.0-20190827072141-edfb9018d271",
)

go_repository(
    name = "com_github_streadway_handy",
    importpath = "github.com/streadway/handy",
    sum = "h1:AhmOdSHeswKHBjhsLs/7+1voOxT+LLrSk/Nxvk35fug=",
    version = "v0.0.0-20190108123426-d5acb3125c2a",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    sum = "h1:1zr/of2m5FGMsad5YfcqgdqdWrIhu+EBEJRhR1U7z/c=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:CcVxjf3Q8PM0mHUKJCdn+eZZtm5yQwehR5yeSVQQcUk=",
    version = "v1.8.4",
)

go_repository(
    name = "com_github_subosito_gotenv",
    importpath = "github.com/subosito/gotenv",
    sum = "h1:Slr1R9HxAlEKefgq5jn9U+DnETlIUa6HfgEzj0g5d7s=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_syndtr_gocapability",
    importpath = "github.com/syndtr/gocapability",
    sum = "h1:b6uOv7YOFK0TYG7HtkIgExQo+2RdLuwRft63jn2HWj8=",
    version = "v0.0.0-20180916011248-d98352740cb2",
)

go_repository(
    name = "com_github_tchap_go_patricia",
    importpath = "github.com/tchap/go-patricia",
    sum = "h1:GkY4dP3cEfEASBPPkWd+AmjYxhmDkqO9/zg7R0lSQRs=",
    version = "v2.3.0+incompatible",
)

go_repository(
    name = "com_github_tidwall_pretty",
    importpath = "github.com/tidwall/pretty",
    sum = "h1:HsD+QiTn7sK6flMKIvNmpqz1qrpP3Ps6jOKIKMooyg4=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_tmc_grpc_websocket_proxy",
    importpath = "github.com/tmc/grpc-websocket-proxy",
    sum = "h1:6fotK7otjonDflCTK0BCfls4SPy3NcCVb5dqqmbRknE=",
    version = "v0.0.0-20220101234140-673ab2c3ae75",
)

go_repository(
    name = "com_github_ugorji_go",
    importpath = "github.com/ugorji/go",
    sum = "h1:/68gy2h+1mWMrwZFeD1kQialdSzAb432dtpeJ42ovdo=",
    version = "v1.1.7",
)

go_repository(
    name = "com_github_ugorji_go_codec",
    importpath = "github.com/ugorji/go/codec",
    sum = "h1:BMaWp1Bb6fHwEtbplGBGJ498wD+LKlNSl25MjdZY4dU=",
    version = "v1.2.11",
)

go_repository(
    name = "com_github_ulikunitz_xz",
    importpath = "github.com/ulikunitz/xz",
    sum = "h1:YvTNdFzX6+W5m9msiYg/zpkSURPPtOlzbqYjrFn7Yt4=",
    version = "v0.5.7",
)

go_repository(
    name = "com_github_xi2_xz",
    importpath = "github.com/xi2/xz",
    sum = "h1:nIPpBwaJSVYIxUFsDv3M8ofmx9yWTog9BfvIu0q41lo=",
    version = "v0.0.0-20171230120015-48954b6210f8",
)

go_repository(
    name = "com_github_urfave_cli",
    importpath = "github.com/urfave/cli",
    sum = "h1:+mkCCcOFKPnCmVYVcURKps1Xe+3zP90gSYGNfRkjoIY=",
    version = "v1.22.1",
)

go_repository(
    name = "com_github_vbatts_tar_split",
    importpath = "github.com/vbatts/tar-split",
    sum = "h1:0Odu65rhcZ3JZaPHxl7tCI3V/C/Q9Zf82UFravl02dE=",
    version = "v0.11.1",
)

go_repository(
    name = "com_github_vbauerster_mpb_v5",
    importpath = "github.com/vbauerster/mpb/v5",
    sum = "h1:zIICVOm+XD+uV6crpSORaL6I0Q1WqOdvxZTp+r3L9cw=",
    version = "v5.2.2",
)

go_repository(
    name = "com_github_vektah_gqlparser",
    importpath = "github.com/vektah/gqlparser",
    sum = "h1:ZsyLGn7/7jDNI+y4SEhI4yAxRChlv15pUHMjijT+e68=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_vishvananda_netlink",
    importpath = "github.com/vishvananda/netlink",
    sum = "h1:bqNY2lgheFIu1meHUFSH3d7vG93AFyqg3oGbJCOJgSM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_vishvananda_netns",
    importpath = "github.com/vishvananda/netns",
    sum = "h1:OviZH7qLw/7ZovXvuNyL3XQl8UFofeikI1NW1Gypu7k=",
    version = "v0.0.0-20191106174202-0a2b9b5464df",
)

go_repository(
    name = "com_github_vividcortex_ewma",
    importpath = "github.com/VividCortex/ewma",
    sum = "h1:MnEK4VOv6n0RSY4vtRe3h11qjxL3+t0B8yOL8iMXdcM=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_vividcortex_gohistogram",
    importpath = "github.com/VividCortex/gohistogram",
    sum = "h1:6+hBz+qvs0JOrrNhhmR7lFxo5sINxBCGXrdtl/UvroE=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_vmware_govmomi",
    importpath = "github.com/vmware/govmomi",
    sum = "h1:Hqu2Uke2itC+cNoIcFQBLEZvX9wBRTTOP04J7V1fqRw=",
    version = "v0.34.1",
)

go_repository(
    name = "com_github_vmware_vmw_guestinfo",
    importpath = "github.com/vmware/vmw-guestinfo",
    sum = "h1:sH9mEk+flyDxiUa5BuPiuhDETMbzrt9A20I2wktMvRQ=",
    version = "v0.0.0-20170707015358-25eff159a728",
)

go_repository(
    name = "com_github_xeipuuv_gojsonpointer",
    importpath = "github.com/xeipuuv/gojsonpointer",
    sum = "h1:6cLsL+2FW6dRAdl5iMtHgRogVCff0QpRi9653YmdcJA=",
    version = "v0.0.0-20190809123943-df4f5c81cb3b",
)

go_repository(
    name = "com_github_xeipuuv_gojsonreference",
    importpath = "github.com/xeipuuv/gojsonreference",
    sum = "h1:EzJWgHovont7NscjpAxXsDA8S8BMYve8Y5+7cuRE7R0=",
    version = "v0.0.0-20180127040603-bd5ef7bd5415",
)

go_repository(
    name = "com_github_xeipuuv_gojsonschema",
    importpath = "github.com/xeipuuv/gojsonschema",
    sum = "h1:LhYJRs+L4fBtjZUfuSZIKGeVu0QRy8e5Xi7D17UxZ74=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_xiang90_probing",
    importpath = "github.com/xiang90/probing",
    sum = "h1:eY9dn8+vbi4tKz5Qo6v2eYzo7kUS51QINcR5jNpbZS8=",
    version = "v0.0.0-20190116061207-43a291ad63a2",
)

go_repository(
    name = "com_github_xlab_handysort",
    importpath = "github.com/xlab/handysort",
    sum = "h1:j2hhcujLRHAg872RWAV5yaUrEjHEObwDv3aImCaNLek=",
    version = "v0.0.0-20150421192137-fb3537ed64a1",
)

go_repository(
    name = "com_github_xordataexchange_crypt",
    importpath = "github.com/xordataexchange/crypt",
    sum = "h1:ESFSdwYZvkeru3RtdrYueztKhOBCSAAzS4Gf+k0tEow=",
    version = "v0.0.3-0.20170626215501-b2862e3d0a77",
)

go_repository(
    name = "com_github_yuin_goldmark",
    importpath = "github.com/yuin/goldmark",
    sum = "h1:fVcFKWvrslecOb/tg+Cc05dkeYx540o0FuFt3nUVDoE=",
    version = "v1.4.13",
)

go_repository(
    name = "com_google_cloud_go",
    importpath = "cloud.google.com/go",
    sum = "h1:e0WKqKTd5BnrG8aKH3J3h+QvEIQtSUcf2n5UZ5ZgLtQ=",
    version = "v0.26.0",
)

go_repository(
    name = "com_google_cloud_go_bigquery",
    importpath = "cloud.google.com/go/bigquery",
    sum = "h1:PQcPefKFdaIzjQFbiyOgAqyx8q5djaE7x9Sqe712DPA=",
    version = "v1.8.0",
)

go_repository(
    name = "com_google_cloud_go_datastore",
    importpath = "cloud.google.com/go/datastore",
    sum = "h1:/May9ojXjRkPBNVrq+oWLqmWCkr4OU5uRY29bu0mRyQ=",
    version = "v1.1.0",
)

go_repository(
    name = "com_google_cloud_go_firestore",
    importpath = "cloud.google.com/go/firestore",
    sum = "h1:9x7Bx0A9R5/M9jibeJeZWqjeVEIxYW9fZYqB9a70/bY=",
    version = "v1.1.0",
)

go_repository(
    name = "com_google_cloud_go_pubsub",
    importpath = "cloud.google.com/go/pubsub",
    sum = "h1:ukjixP1wl0LpnZ6LWtZJ0mX5tBmjp1f8Sqer8Z2OMUU=",
    version = "v1.3.1",
)

go_repository(
    name = "com_google_cloud_go_storage",
    importpath = "cloud.google.com/go/storage",
    sum = "h1:STgFzyU5/8miMl0//zKh2aQeTyeaUH3WN9bSUiJ09bA=",
    version = "v1.10.0",
)

go_repository(
    name = "com_shuralyov_dmitri_gpu_mtl",
    importpath = "dmitri.shuralyov.com/gpu/mtl",
    sum = "h1:VpgP7xuJadIUuKccphEpTJnWhS2jkQyMt6Y7pJCD7fY=",
    version = "v0.0.0-20190408044501-666a987793e9",
)

go_repository(
    name = "com_sourcegraph_sourcegraph_appdash",
    importpath = "sourcegraph.com/sourcegraph/appdash",
    sum = "h1:ucqkfpjg9WzSUubAO62csmucvxl4/JeW3F4I4909XkM=",
    version = "v0.0.0-20190731080439-ebfcffb1b5c0",
)

go_repository(
    name = "in_gopkg_alecthomas_kingpin_v2",
    importpath = "gopkg.in/alecthomas/kingpin.v2",
    sum = "h1:jMFz6MfLP0/4fUyZle81rXUoxOBFi19VUFKVDOQfozc=",
    version = "v2.2.6",
)

go_repository(
    name = "in_gopkg_asn1_ber_v1",
    importpath = "gopkg.in/asn1-ber.v1",
    sum = "h1:TxyelI5cVkbREznMhfzycHdkp5cLA7DpE+GKjSslYhM=",
    version = "v1.0.0-20181015200546-f715ec2f112d",
)

go_repository(
    name = "in_gopkg_check_v1",
    importpath = "gopkg.in/check.v1",
    sum = "h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=",
    version = "v1.0.0-20201130134442-10cb98267c6c",
)

go_repository(
    name = "in_gopkg_cheggaaa_pb_v1",
    importpath = "gopkg.in/cheggaaa/pb.v1",
    sum = "h1:Ev7yu1/f6+d+b3pi5vPdRPc6nNtP1umSfcWiEfRqv6I=",
    version = "v1.0.25",
)

go_repository(
    name = "in_gopkg_errgo_v2",
    importpath = "gopkg.in/errgo.v2",
    sum = "h1:0vLT13EuvQ0hNvakwLuFZ/jYrLp5F3kcWHXdRggjCE8=",
    version = "v2.1.0",
)

go_repository(
    name = "in_gopkg_fsnotify_v1",
    importpath = "gopkg.in/fsnotify.v1",
    sum = "h1:xOHLXZwVvI9hhs+cLKq5+I5onOuwQLhQwiu63xxlHs4=",
    version = "v1.4.7",
)

go_repository(
    name = "in_gopkg_gcfg_v1",
    importpath = "gopkg.in/gcfg.v1",
    sum = "h1:m8OOJ4ccYHnx2f4gQwpno8nAX5OGOh7RLaaz0pj3Ogs=",
    version = "v1.2.3",
)

go_repository(
    name = "in_gopkg_go_playground_assert_v1",
    importpath = "gopkg.in/go-playground/assert.v1",
    sum = "h1:xoYuJVE7KT85PYWrN730RguIQO0ePzVRfFMXadIrXTM=",
    version = "v1.2.1",
)

go_repository(
    name = "in_gopkg_go_playground_validator_v9",
    importpath = "gopkg.in/go-playground/validator.v9",
    sum = "h1:SvGtYmN60a5CVKTOzMSyfzWDeZRxRuGvRQyEAKbw1xc=",
    version = "v9.29.1",
)

go_repository(
    name = "in_gopkg_inf_v0",
    importpath = "gopkg.in/inf.v0",
    sum = "h1:73M5CoZyi3ZLMOyDlQh031Cx6N9NDJ2Vvfl76EDAgDc=",
    version = "v0.9.1",
)

go_repository(
    name = "in_gopkg_ini_v1",
    importpath = "gopkg.in/ini.v1",
    sum = "h1:AQvPpx3LzTDM0AjnIRlVFwFFGC+npRopjZxLJj6gdno=",
    version = "v1.51.0",
)

go_repository(
    name = "in_gopkg_ldap_v2",
    importpath = "gopkg.in/ldap.v2",
    sum = "h1:wiu0okdNfjlBzg6UWvd1Hn8Y+Ux17/u/4nlk4CQr6tU=",
    version = "v2.5.1",
)

go_repository(
    name = "in_gopkg_natefinch_lumberjack_v2",
    importpath = "gopkg.in/natefinch/lumberjack.v2",
    sum = "h1:bBRl1b0OH9s/DuPhuXpNl+VtCaJXFZ5/uEFST95x9zc=",
    version = "v2.2.1",
)

go_repository(
    name = "in_gopkg_resty_v1",
    importpath = "gopkg.in/resty.v1",
    sum = "h1:CuXP0Pjfw9rOuY6EP+UvtNvt5DSqHpIxILZKT/quCZI=",
    version = "v1.12.0",
)

go_repository(
    name = "in_gopkg_square_go_jose_v2",
    importpath = "gopkg.in/square/go-jose.v2",
    sum = "h1:NGk74WTnPKBNUhNzQX7PYcTLUjoq7mzKk2OKbvwk2iI=",
    version = "v2.6.0",
)

go_repository(
    name = "in_gopkg_tomb_v1",
    importpath = "gopkg.in/tomb.v1",
    sum = "h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
    version = "v1.0.0-20141024135613-dd632973f1e7",
)

go_repository(
    name = "in_gopkg_warnings_v0",
    importpath = "gopkg.in/warnings.v0",
    sum = "h1:wFXVbFY8DY5/xOe1ECiWdKCzZlxgshcYVNkBHstARME=",
    version = "v0.1.2",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
    version = "v2.4.0",
)

go_repository(
    name = "in_gopkg_yaml_v3",
    importpath = "gopkg.in/yaml.v3",
    sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
    version = "v3.0.1",
)

go_repository(
    name = "io_etcd_go_bbolt",
    importpath = "go.etcd.io/bbolt",
    sum = "h1:j+zJOnnEjF/kyHlDDgGnVL/AIqIJPq8UoB2GSNfkUfQ=",
    version = "v1.3.7",
)

go_repository(
    name = "io_etcd_go_etcd",
    importpath = "go.etcd.io/etcd",
    sum = "h1:Gqga3zA9tdAcfqobUGjSoCob5L3f8Dt5EuOp3ihNZko=",
    version = "v0.5.0-alpha.5.0.20200819165624-17cef6e3e9d5",
)

go_repository(
    name = "io_k8s_api",
    importpath = "k8s.io/api",
    sum = "h1:Gj1HtbSdB4P08C8rs9AR94MfSGpRhJgsS+GF9V26xMM=",
    version = "v0.28.3",
)

go_repository(
    name = "io_k8s_apiextensions_apiserver",
    importpath = "k8s.io/apiextensions-apiserver",
    sum = "h1:Od7DEnhXHnHPZG+W9I97/fSQkVpVPQx2diy+2EtmY08=",
    version = "v0.28.3",
)

go_repository(
    name = "io_k8s_apimachinery",
    importpath = "k8s.io/apimachinery",
    sum = "h1:B1wYx8txOaCQG0HmYF6nbpU8dg6HvA06x5tEffvOe7A=",
    version = "v0.28.3",
)

go_repository(
    name = "io_k8s_apiserver",
    importpath = "k8s.io/apiserver",
    sum = "h1:8Ov47O1cMyeDzTXz0rwcfIIGAP/dP7L8rWbEljRcg5w=",
    version = "v0.28.3",
)

go_repository(
    name = "io_k8s_client_go",
    importpath = "k8s.io/client-go",
    sum = "h1:2OqNb72ZuTZPKCl+4gTKvqao0AMOl9f3o2ijbAj3LI4=",
    version = "v0.28.3",
)

go_repository(
    name = "io_k8s_cluster_bootstrap",
    importpath = "k8s.io/cluster-bootstrap",
    replace = "k8s.io/cluster-bootstrap",
    sum = "h1:QY1qpB83m+mNlDcXLT+EpTGiUNVSNV0ybmgd/yZR5mA=",
    version = "v0.16.4",
)

go_repository(
    name = "io_k8s_code_generator",
    importpath = "k8s.io/code-generator",
    sum = "h1:I847QvdpYx7xKiG2KVQeCSyNF/xU9TowaDAg601mvlw=",
    version = "v0.28.3",
)

go_repository(
    name = "io_k8s_component_base",
    importpath = "k8s.io/component-base",
    sum = "h1:rDy68eHKxq/80RiMb2Ld/tbH8uAE75JdCqJyi6lXMzI=",
    version = "v0.28.3",
)

go_repository(
    name = "io_k8s_gengo",
    importpath = "k8s.io/gengo",
    sum = "h1:pWEwq4Asjm4vjW7vcsmijwBhOr1/shsbSYiWXmNGlks=",
    version = "v0.0.0-20230829151522-9cce18d56c01",
)

go_repository(
    name = "io_k8s_klog",
    importpath = "k8s.io/klog",
    sum = "h1:Pt+yjF5aB1xDSVbau4VsWe+dQNzA0qv1LlXdC2dF6Q8=",
    version = "v1.0.0",
)

go_repository(
    name = "io_k8s_klog_v2",
    importpath = "k8s.io/klog/v2",
    sum = "h1:U/Af64HJf7FcwMcXyKm2RPM22WZzyR7OSpYj5tg3cL0=",
    version = "v2.110.1",
)

go_repository(
    name = "io_k8s_kube_aggregator",
    importpath = "k8s.io/kube-aggregator",
    sum = "h1:NYgl5PDV/oX1yqAZIkRnb+KtW+eLykzc6hHg81ECgiI=",
    version = "v0.27.1",
)

go_repository(
    name = "io_k8s_kube_openapi",
    importpath = "k8s.io/kube-openapi",
    sum = "h1:aVUu9fTY98ivBPKR9Y5w/AuzbMm96cd3YHRTU83I780=",
    version = "v0.0.0-20231010175941-2dd684a91f00",
)

go_repository(
    name = "io_k8s_kubernetes",
    importpath = "k8s.io/kubernetes",
    replace = "k8s.io/kubernetes",
    sum = "h1:x6Q6M9nNBm9thoKj+PJr3HDPfHlz7FR35Ri61xpNfgg=",
    version = "v0.19.3",
)

go_repository(
    name = "io_k8s_sigs_apiserver_network_proxy_konnectivity_client",
    importpath = "sigs.k8s.io/apiserver-network-proxy/konnectivity-client",
    sum = "h1:trsWhjU5jZrx6UvFu4WzQDrN7Pga4a7Qg+zcfcj64PA=",
    version = "v0.1.2",
)

go_repository(
    name = "io_k8s_sigs_controller_runtime",
    importpath = "sigs.k8s.io/controller-runtime",
    sum = "h1:2TuvuokmfXvDUamSx1SuAOO3eTyye+47mJCigwG62c4=",
    version = "v0.16.3",
)

go_repository(
    name = "io_k8s_sigs_controller_tools",
    importpath = "sigs.k8s.io/controller-tools",
    sum = "h1:3u2RCwOlp0cjCALAigpOcbAf50pE+kHSdueUosrC/AE=",
    version = "v0.5.0",
)

go_repository(
    name = "io_k8s_sigs_structured_merge_diff",
    importpath = "sigs.k8s.io/structured-merge-diff",
    replace = "sigs.k8s.io/structured-merge-diff",
    sum = "h1:zD2IemQ4LmOcAumeiyDWXKUI2SO0NYDe3H6QGvPOVgU=",
    version = "v1.0.1-0.20191108220359-b1b620dd3f06",
)

go_repository(
    name = "io_k8s_sigs_structured_merge_diff_v3",
    importpath = "sigs.k8s.io/structured-merge-diff/v3",
    sum = "h1:0KsuGbLhWdIxv5DA1OnbFz5hI/Co9kuxMfMUa5YsAHY=",
    version = "v3.0.0-20200116222232-67a7b8c61874",
)

go_repository(
    name = "io_k8s_sigs_structured_merge_diff_v4",
    importpath = "sigs.k8s.io/structured-merge-diff/v4",
    sum = "h1:150L+0vs/8DA78h1u02ooW1/fFq/Lwr+sGiqlzvrtq4=",
    version = "v4.4.1",
)

go_repository(
    name = "io_k8s_sigs_testing_frameworks",
    importpath = "sigs.k8s.io/testing_frameworks",
    sum = "h1:vK0+tvjF0BZ/RYFeZ1E6BYBwHJJXhjuZ3TdsEKH+UQM=",
    version = "v0.1.2",
)

go_repository(
    name = "io_k8s_sigs_yaml",
    importpath = "sigs.k8s.io/yaml",
    sum = "h1:a2VclLzOGrwOHDiV8EfBGhvjHvP46CtW5j6POvhYGGo=",
    version = "v1.3.0",
)

go_repository(
    name = "io_k8s_utils",
    importpath = "k8s.io/utils",
    sum = "h1:eQ/4ljkx21sObifjzXwlPKpdGLrCfRziVtos3ofG/sQ=",
    version = "v0.0.0-20240102154912-e7106e64919e",
)

go_repository(
    name = "io_kubevirt_client_go",
    importpath = "kubevirt.io/client-go",
    sum = "h1:Twh3JdWUptj1TIs5sI+jwDvcUXaCL2lsfF9PLZoCw7s=",
    version = "v0.42.1",
)

go_repository(
    name = "io_kubevirt_containerized_data_importer",
    importpath = "kubevirt.io/containerized-data-importer",
    sum = "h1:BrvtyeHdRWINSi+dQM46hxduUr4UwdEdwdYdX6okR14=",
    version = "v1.34.0",
)

go_repository(
    name = "io_kubevirt_containerized_data_importer_api",
    importpath = "kubevirt.io/containerized-data-importer-api",
    sum = "h1:GdDt9BlR0qHejpMaPfASbsG8JWDmBf1s7xZBj5W9qn0=",
    version = "v1.59.0",
)

go_repository(
    name = "io_kubevirt_controller_lifecycle_operator_sdk",
    importpath = "kubevirt.io/controller-lifecycle-operator-sdk",
    sum = "h1:auv8LrA7gnLfQREnlGVPwgJpTxOEgnw4+mzXlUqKTxY=",
    version = "v0.2.3",
)

go_repository(
    name = "io_kubevirt_qe_tools",
    importpath = "kubevirt.io/qe-tools",
    sum = "h1:S6z9CATmgV2/z9CWetij++Rhu7l/Z4ObZqerLdNMo0Y=",
    version = "v0.1.6",
)

go_repository(
    name = "io_opencensus_go",
    importpath = "go.opencensus.io",
    sum = "h1:LYy1Hy3MJdrCdMwwzxA/dRok4ejH+RwNGbuoD9fCjto=",
    version = "v0.22.4",
)

go_repository(
    name = "io_rsc_binaryregexp",
    importpath = "rsc.io/binaryregexp",
    sum = "h1:HfqmD5MEmC0zvwBuF187nq9mdnXjXsSivRiXN7SmRkE=",
    version = "v0.2.0",
)

go_repository(
    name = "io_rsc_quote_v3",
    importpath = "rsc.io/quote/v3",
    sum = "h1:9JKUTTIUgS6kzR9mK1YuGKv6Nl+DijDNIc0ghT58FaY=",
    version = "v3.1.0",
)

go_repository(
    name = "io_rsc_sampler",
    importpath = "rsc.io/sampler",
    sum = "h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=",
    version = "v1.3.0",
)

go_repository(
    name = "ml_vbom_util",
    importpath = "vbom.ml/util",
    sum = "h1:O69FD9pJA4WUZlEwYatBEEkRWKQ5cKodWpdKTrCS/iQ=",
    version = "v0.0.0-20180919145318-efcd4e0f9787",
)

go_repository(
    name = "org_bitbucket_ww_goautoneg",
    importpath = "bitbucket.org/ww/goautoneg",
    replace = "github.com/markusthoemmes/goautoneg",
    sum = "h1:Qhv4Ni88zV+8TY65yr2ak8xU4sblgs6aRT9RuGM5SNU=",
    version = "v0.0.0-20190713162725-c6008fefa5b1",
)

go_repository(
    name = "org_golang_google_api",
    importpath = "google.golang.org/api",
    sum = "h1:yfrXXP61wVuLb0vBcG6qaOoIoqYEzOQS8jum51jkv2w=",
    version = "v0.30.0",
)

go_repository(
    name = "org_golang_google_appengine",
    importpath = "google.golang.org/appengine",
    sum = "h1:FZR1q0exgwxzPzp/aF+VccGrSfxfPpkBqjIIEq3ru6c=",
    version = "v1.6.7",
)

go_repository(
    name = "org_golang_google_genproto",
    importpath = "google.golang.org/genproto",
    sum = "h1:9NWlQfY2ePejTmfwUH1OWwmznFa+0kKcHGPDvcPza9M=",
    version = "v0.0.0-20230526161137-0005af68ea54",
)

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    sum = "h1:EhTqbhiYeixwWQtAEZAxmV9MGqcjEU2mFx52xCzNyag=",
    version = "v1.54.0",
)

go_repository(
    name = "org_golang_google_protobuf",
    importpath = "google.golang.org/protobuf",
    sum = "h1:g0LDEJHgrBl9N9r17Ru3sqWhkIx2NB67okBHPwC7hs8=",
    version = "v1.31.0",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:g1v0xeRhjcugydODzvb3mEM9SQ0HGp9s/nh3COQ/C30=",
    version = "v0.22.0",
)

go_repository(
    name = "org_golang_x_exp",
    importpath = "golang.org/x/exp",
    sum = "h1:+WEEuIdZHnUeJJmEUjyYC2gfUMj69yZXw17EnHg/otA=",
    version = "v0.0.0-20220722155223-a9213eeb770e",
)

go_repository(
    name = "org_golang_x_image",
    importpath = "golang.org/x/image",
    sum = "h1:+qEpEAPhDZ1o0x3tHzZTQDArnOixOzGD9HUJfcg0mb4=",
    version = "v0.0.0-20190802002840-cff245a6509b",
)

go_repository(
    name = "org_golang_x_lint",
    importpath = "golang.org/x/lint",
    sum = "h1:5hukYrvBGR8/eNkX5mdUezrA6JiaEZDtJb9Ei+1LlBs=",
    version = "v0.0.0-20190930215403-16217165b5de",
)

go_repository(
    name = "org_golang_x_mobile",
    importpath = "golang.org/x/mobile",
    sum = "h1:4+4C/Iv2U4fMZBiMCc98MG1In4gJY5YRhtpDNeDeHWs=",
    version = "v0.0.0-20190719004257-d2bd2a29d028",
)

go_repository(
    name = "org_golang_x_mod",
    importpath = "golang.org/x/mod",
    sum = "h1:dGoOF9QVLYng8IHTm7BAyWqCqSheQ5pYWGhzW00YJr0=",
    version = "v0.14.0",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:1PcaxkF854Fu3+lvBIx5SYn9wRlBzzcnHZSiaFFAb0w=",
    version = "v0.24.0",
)

go_repository(
    name = "org_golang_x_oauth2",
    importpath = "golang.org/x/oauth2",
    sum = "h1:smVPGxink+n1ZI5pkQa8y6fZT0RW0MgCO5bFpepy4B4=",
    version = "v0.12.0",
)

go_repository(
    name = "org_golang_x_sync",
    importpath = "golang.org/x/sync",
    sum = "h1:60k92dhOjHxJkrqnwsfl8KuaHbn/5dl0lUPUklKo3qE=",
    version = "v0.5.0",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:q5f1RH2jigJ1MoAWp2KTp3gm5zAGFUTarQZ5U386+4o=",
    version = "v0.19.0",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:ScX5w1eTa3QqT8oi6+ziP7dTV1S2+ALU0bI+0zXKWiQ=",
    version = "v0.14.0",
)

go_repository(
    name = "org_golang_x_time",
    importpath = "golang.org/x/time",
    sum = "h1:rg5rLMjNzMS1RkNLzCG38eapWhnYLFYXDXj2gOlr8j4=",
    version = "v0.3.0",
)

go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    sum = "h1:TLyB3WofjdOEepBHAU20JdNC1Zbg87elYofWYAY5oZA=",
    version = "v0.16.1",
)

go_repository(
    name = "org_golang_x_xerrors",
    importpath = "golang.org/x/xerrors",
    sum = "h1:H2TDz8ibqkAF6YGhCdN3jS9O0/s90v0rJh3X/OLHEUk=",
    version = "v0.0.0-20220907171357-04be3eba64a2",
)

go_repository(
    name = "org_gonum_v1_gonum",
    importpath = "gonum.org/v1/gonum",
    sum = "h1:2qZ38BsejXrhuetzb8UxucqrWDZKjypFSZA82hLCpZ4=",
    version = "v0.0.0-20190710053202-4340aa3071a0",
)

go_repository(
    name = "org_gonum_v1_netlib",
    importpath = "gonum.org/v1/netlib",
    sum = "h1:OE9mWmgKkjJyEmDAAtGMPjXu+YNeGvK9VTSHY6+Qihc=",
    version = "v0.0.0-20190313105609-8cb42192e0e0",
)

go_repository(
    name = "org_libvirt_libvirt_go_xml",
    importpath = "libvirt.org/libvirt-go-xml",
    sum = "h1:NaCRjbtz//xuTZOp1nDHbe0eu5BQlhIy5PPuc09EWtU=",
    version = "v7.4.0+incompatible",
)

go_repository(
    name = "org_uber_go_atomic",
    importpath = "go.uber.org/atomic",
    sum = "h1:9qC72Qh0+3MqyJbAn8YU5xVq1frD8bn3JtD2oXtafVQ=",
    version = "v1.10.0",
)

go_repository(
    name = "org_uber_go_multierr",
    importpath = "go.uber.org/multierr",
    sum = "h1:blXXJkSxSSfBVBlC76pxqeO+LN3aDfLQo+309xJstO0=",
    version = "v1.11.0",
)

go_repository(
    name = "org_uber_go_tools",
    importpath = "go.uber.org/tools",
    sum = "h1:0mgffUl7nfd+FpvXMVz4IDEaUSmT1ysygQC7qYo7sG4=",
    version = "v0.0.0-20190618225709-2cfd321de3ee",
)

go_repository(
    name = "org_uber_go_zap",
    importpath = "go.uber.org/zap",
    sum = "h1:sI7k6L95XOKS281NhVKOFCUNIvv9e0w4BF8N3u+tCRo=",
    version = "v1.26.0",
)

go_repository(
    name = "tools_gotest",
    importpath = "gotest.tools",
    sum = "h1:VsBPFP1AI068pPrMxtb/S8Zkgf9xEmTLJjfM+P5UIEo=",
    version = "v2.2.0+incompatible",
)

go_repository(
    name = "tools_gotest_v3",
    importpath = "gotest.tools/v3",
    sum = "h1:kG1BFyqVHuQoVQiR1bWGnfz/fmHvvuiSPIV7rvl360E=",
    version = "v3.0.2",
)

go_repository(
    name = "xyz_gomodules_jsonpatch_v2",
    importpath = "gomodules.xyz/jsonpatch/v2",
    sum = "h1:Ci3iUJyx9UeRx7CeFN8ARgGbkESwJK+KB9lLcWxY/Zw=",
    version = "v2.4.0",
)

go_rules_dependencies()

# NOTE: Keep the version in sync with Go toolchain in GitHub action.
go_register_toolchains(version = "1.21.9")

# override rules_docker issue with this dependency
# rules_docker 0.16 uses 0.1.4, bit since there the checksum changed, which is very weird, going with 0.1.4.1 to
go_repository(
    name = "com_github_google_go_containerregistry",
    importpath = "github.com/google/go-containerregistry",
    sha256 = "bc0136a33f9c1e4578a700f7afcdaa1241cfff997d6bba695c710d24c5ae26bd",
    strip_prefix = "google-go-containerregistry-efb2d62",
    type = "tar.gz",
    urls = ["https://api.github.com/repos/google/go-containerregistry/tarball/efb2d62d93a7705315b841d0544cb5b13565ff2a"],  # v0.1.4.1
)

# All dependencies defined with go_repository should be above
# gazelle_dependencies. Otherwise dependencies defined by Gazelle would be used
# because the first definitions of external repository wins.
gazelle_dependencies()

load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
)
load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

container_pull(
    name = "centos-stream-9",
    registry = "quay.io",
    repository = "centos/centos",
    tag = "stream9",
)

container_pull(
    name = "ubi9-minimal",
    registry = "registry.access.redhat.com",
    repository = "ubi9/ubi-minimal",
    tag = "latest",
)

container_pull(
    name = "ansible-operator",
    registry = "quay.io",
    repository = "operator-framework/ansible-operator",
    tag = "main",
)

container_pull(
    name = "opm-image",
    # v1.43.1
    digest = "sha256:dbeedc2489e6a971bcc2640b582482059938f6c1e14446db5fdb6f6c02aa01d5",
    registry = "quay.io",
    repository = "operator-framework/opm",
)

http_file(
    name = "opa",
    downloaded_file_path = "opa",
    executable = True,
    sha256 = "cd6b0b2d762571a746f0261890b155e6dd71cca90dad6b42b6fcf6dd7f619f08",
    urls = ["https://openpolicyagent.org/downloads/v0.65.0/opa_linux_amd64_static"],
)

http_file(
    name = "kustomize",
    downloaded_file_path = "kustomize.tar.gz",
    sha256 = "3ab32f92360d752a2a53e56be073b649abc1e7351b912c0fb32b960d1def854c",
    urls = [
        "https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize/v5.3.0/kustomize_v5.3.0_linux_amd64.tar.gz",
    ],
)

http_file(
    name = "operator-sdk",
    downloaded_file_path = "operator-sdk",
    executable = True,
    sha256 = "90ac1ed961585f16ff60d5187fd29592e5af6db8d074e1db3fde02d4145ff3fd",
    urls = ["https://github.com/operator-framework/operator-sdk/releases/download/v1.34.2/operator-sdk_linux_amd64"],
)

http_file(
    name = "opm",
    downloaded_file_path = "opm",
    executable = True,
    sha256 = "2afbcd0728f61d07502c67b0bd36ff6f7521f1a629ff4e23ff13a59dec93eeb0",
    urls = ["https://github.com/operator-framework/operator-registry/releases/download/v1.43.1/linux-amd64-opm"],
)

http_archive(
    name = "bazeldnf",
    sha256 = "de8ac0519db6c02bed2a976885379d97dc52f46d52b0338e03963def5c2b65fb",
    urls = [
        "https://github.com/rmohr/bazeldnf/releases/download/v0.5.7-rc1/bazeldnf-v0.5.7-rc1.tar.gz",
    ],
)

load(
    "@io_bazel_rules_go//go:deps.bzl",
    "go_register_toolchains",
    "go_rules_dependencies",
)
load("@bazeldnf//:deps.bzl", "bazeldnf_dependencies", "rpm")

bazeldnf_dependencies()

http_archive(
    name = "rules_proto",
    sha256 = "dc3fb206a2cb3441b485eb1e423165b231235a1ea9b031b4433cf7bc1fa460dd",
    strip_prefix = "rules_proto-5.3.0-21.7",
    urls = [
        "https://github.com/bazelbuild/rules_proto/archive/refs/tags/5.3.0-21.7.tar.gz",
    ],
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

go_repository(
    name = "com_github_a8m_tree",
    importpath = "github.com/a8m/tree",
    sum = "h1:4E8RufAN3UQ/weB6AnQ4y5miZCO0Yco8ZdGId41WuQs=",
    version = "v0.0.0-20210115125333-10a5fd5b637d",
)

go_repository(
    name = "com_github_alecthomas_kingpin_v2",
    importpath = "github.com/alecthomas/kingpin/v2",
    sum = "h1:H0aULhgmSzN8xQ3nX1uxtdlTHYoPLu5AhHxWrKI6ocU=",
    version = "v2.3.2",
)

go_repository(
    name = "com_github_antlr_antlr4_runtime_go_antlr",
    importpath = "github.com/antlr/antlr4/runtime/Go/antlr",
    sum = "h1:yL7+Jz0jTC6yykIK/Wh74gnTJnrGr5AyrNMXuA0gves=",
    version = "v1.4.10",
)

go_repository(
    name = "com_github_antlr_antlr4_runtime_go_antlr_v4",
    importpath = "github.com/antlr/antlr4/runtime/Go/antlr/v4",
    sum = "h1:7RFfzj4SSt6nnvCPbCqijJi1nWCd+TqAT3bYCStRC18=",
    version = "v4.0.0-20230305170008-8188dc5388df",
)

go_repository(
    name = "com_github_armon_go_socks5",
    importpath = "github.com/armon/go-socks5",
    sum = "h1:0CwZNZbxp69SHPdPJAN/hZIm0C4OItdklCFmMRWYpio=",
    version = "v0.0.0-20160902184237-e75332964ef5",
)

go_repository(
    name = "com_github_azure_go_ntlmssp",
    importpath = "github.com/Azure/go-ntlmssp",
    sum = "h1:ZU22z/2YRFLyf/P4ZwUYSdNCWsMEI0VeyrFoI2rAhJQ=",
    version = "v0.0.0-20211209120228-48547f28849e",
)

go_repository(
    name = "com_github_benbjohnson_clock",
    importpath = "github.com/benbjohnson/clock",
    sum = "h1:Q92kusRqC1XV2MjkWETPvjJVqKetz1OzxZB7mHJLju8=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_blang_semver_v4",
    importpath = "github.com/blang/semver/v4",
    sum = "h1:1PFHFE6yCCTv8C1TeyNNarDzntLi7wMI5i/pzqYIsAM=",
    version = "v4.0.0",
)

go_repository(
    name = "com_github_bytedance_sonic",
    importpath = "github.com/bytedance/sonic",
    sum = "h1:6iJ6NqdoxCDr6mbY8h18oSO+cShGSMRGCEo7F2h0x8s=",
    version = "v1.9.1",
)

go_repository(
    name = "com_github_cenkalti_backoff_v4",
    importpath = "github.com/cenkalti/backoff/v4",
    sum = "h1:y4OZtCnogmCPw98Zjyt5a6+QwPLGkiQsYW5oUqylYbM=",
    version = "v4.2.1",
)

go_repository(
    name = "com_github_chenzhuoyu_base64x",
    importpath = "github.com/chenzhuoyu/base64x",
    sum = "h1:qSGYFH7+jGhDF8vLC+iwCD4WpbV1EBDSzWkJODFLams=",
    version = "v0.0.0-20221115062448-fe3a3abad311",
)

go_repository(
    name = "com_github_coreos_go_systemd_v22",
    importpath = "github.com/coreos/go-systemd/v22",
    sum = "h1:RrqgGjYQKalulkV8NGVIfkXQf6YYmOyiJKk8iXXhfZs=",
    version = "v22.5.0",
)

go_repository(
    name = "com_github_distribution_distribution_v3",
    importpath = "github.com/distribution/distribution/v3",
    sum = "h1:yGaIDzPWkgU+yRvI2x/rGdOU1hl6bLZzm0mETEUSHwk=",
    version = "v3.0.0-20230511163743-f7717b7855ca",
)

go_repository(
    name = "com_github_dougm_pretty",
    importpath = "github.com/dougm/pretty",
    sum = "h1:tR3jsKPiO/mb6ntzk/dJlHZtm37CPfVp1C9KIo534+4=",
    version = "v0.0.0-20171025230240-2ee9d7453c02",
)

go_repository(
    name = "com_github_elazarl_goproxy",
    importpath = "github.com/elazarl/goproxy",
    sum = "h1:yUdfgN0XgIJw7foRItutHYUIhlcKzcSf5vDpdhQAKTc=",
    version = "v0.0.0-20180725130230-947c36da3153",
)

go_repository(
    name = "com_github_emicklei_go_restful_v3",
    importpath = "github.com/emicklei/go-restful/v3",
    sum = "h1:rAQeMHw1c7zTmncogyy8VvRZwtkmkZ4FxERmMY4rD+g=",
    version = "v3.11.0",
)

go_repository(
    name = "com_github_evanphx_json_patch_v5",
    importpath = "github.com/evanphx/json-patch/v5",
    sum = "h1:b91NhWfaz02IuVxO9faSllyAtNXHMPkC5J8sJCLunww=",
    version = "v5.6.0",
)

go_repository(
    name = "com_github_felixge_fgprof",
    importpath = "github.com/felixge/fgprof",
    sum = "h1:VvyZxILNuCiUCSXtPtYmmtGvb65nqXh2QFWc0Wpf2/g=",
    version = "v0.9.3",
)

go_repository(
    name = "com_github_felixge_httpsnoop",
    importpath = "github.com/felixge/httpsnoop",
    sum = "h1:s/nj+GCswXYzN5v2DpNMuMQYe+0DDwt5WVCU6CWBdXk=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_gabriel_vasile_mimetype",
    importpath = "github.com/gabriel-vasile/mimetype",
    sum = "h1:w5qFW6JKBz9Y393Y4q372O9A7cUSequkh1Q7OhCmWKU=",
    version = "v1.4.2",
)

go_repository(
    name = "com_github_getkin_kin_openapi",
    importpath = "github.com/getkin/kin-openapi",
    sum = "h1:j77zg3Ec+k+r+GA3d8hBoXpAc6KX9TbBPrwQGBIy2sY=",
    version = "v0.76.0",
)

go_repository(
    name = "com_github_go_asn1_ber_asn1_ber",
    importpath = "github.com/go-asn1-ber/asn1-ber",
    sum = "h1:vXT6d/FNDiELJnLb6hGNa309LMsrCoYFvpwHDF0+Y1A=",
    version = "v1.5.4",
)

go_repository(
    name = "com_github_go_ldap_ldap_v3",
    importpath = "github.com/go-ldap/ldap/v3",
    sum = "h1:JCKUtJPIcyOuG7ctGabLKMgIlKnGumD/iGjuWeEruDI=",
    version = "v3.4.3",
)

go_repository(
    name = "com_github_go_logr_stdr",
    importpath = "github.com/go-logr/stdr",
    sum = "h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=",
    version = "v1.2.2",
)

go_repository(
    name = "com_github_go_task_slim_sprig",
    importpath = "github.com/go-task/slim-sprig",
    sum = "h1:tfuBGBXKqDEevZMzYi5KSi8KkcZtzBcTgAUUtapy0OI=",
    version = "v0.0.0-20230315185526-52ccab3ef572",
)

go_repository(
    name = "com_github_goccy_go_json",
    importpath = "github.com/goccy/go-json",
    sum = "h1:CrxCmQqYDkv1z7lO7Wbh2HN93uovUHgrECaO5ZrCXAU=",
    version = "v0.10.2",
)

go_repository(
    name = "com_github_golang_jwt_jwt_v4",
    importpath = "github.com/golang-jwt/jwt/v4",
    sum = "h1:7cYmW1XlMY7h7ii7UhUyChSgS5wUJEnm9uZVTGqOWzg=",
    version = "v4.5.0",
)

go_repository(
    name = "com_github_google_cel_go",
    importpath = "github.com/google/cel-go",
    sum = "h1:3hZfSNiAU3KOiNtxuFXVp5WFy4hf/Ly3Sa4/7F8SXNo=",
    version = "v0.16.1",
)

go_repository(
    name = "com_github_google_gnostic",
    importpath = "github.com/google/gnostic",
    sum = "h1:FhTMOKj2VhjpouxvWJAV1TL304uMlb9zcDqkl6cEI54=",
    version = "v0.5.7-v3refs",
)

go_repository(
    name = "com_github_google_gnostic_models",
    importpath = "github.com/google/gnostic-models",
    sum = "h1:yo/ABAfM5IMRsS1VnXjTBvUb61tFIHozhlYvRgGre9I=",
    version = "v0.6.8",
)

go_repository(
    name = "com_github_gophercloud_gophercloud",
    importpath = "github.com/gophercloud/gophercloud",
    replace = "github.com/kubev2v/gophercloud",
    sum = "h1:2+joXldqFBM5YnNAINDK8/8866ML9W1kaa34voAxwhQ=",
    version = "v0.0.0-20230629135522-9d701a75c760",
)

go_repository(
    name = "com_github_gophercloud_utils",
    importpath = "github.com/gophercloud/utils",
    sum = "h1:vJyXd9+MB5vAKxpOo4z/PDSiPgKmEyJwHIDOdV4Y0KY=",
    version = "v0.0.0-20230418172808-6eab72e966e1",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway_v2",
    importpath = "github.com/grpc-ecosystem/grpc-gateway/v2",
    sum = "h1:BZHcxBETFHIdVyhyEfOvn/RdU/QGdLI4y34qQGjGWO0=",
    version = "v2.7.0",
)

go_repository(
    name = "com_github_jessevdk_go_flags",
    importpath = "github.com/jessevdk/go-flags",
    sum = "h1:4IU2WS7AumrZ/40jfhf4QVDMsQwqA7VEHozFRrGARJA=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_josharian_intern",
    importpath = "github.com/josharian/intern",
    sum = "h1:vlS4z54oSdjm0bgjRigI+G1HpF+tI+9rE5LLzOg8HmY=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_klauspost_cpuid_v2",
    importpath = "github.com/klauspost/cpuid/v2",
    sum = "h1:acbojRNwl3o09bUq+yDCtZFc1aiwaAAxtcn8YkZXnvk=",
    version = "v2.2.4",
)

go_repository(
    name = "com_github_matttproud_golang_protobuf_extensions_v2",
    importpath = "github.com/matttproud/golang_protobuf_extensions/v2",
    sum = "h1:jWpvCLoY8Z/e3VKvlsiIGKtc+UG6U5vzxaoagmhXfyg=",
    version = "v2.0.0",
)

go_repository(
    name = "com_github_moby_spdystream",
    importpath = "github.com/moby/spdystream",
    sum = "h1:cjW1zVyyoiM0T7b6UoySUFqzXMoqRckQtXwGPiBhOM8=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_niemeyer_pretty",
    importpath = "github.com/niemeyer/pretty",
    sum = "h1:fD57ERR4JtEqsWbfPhv4DMiApHyliiK5xCTNVSPiaAs=",
    version = "v0.0.0-20200227124842-a10e7caefd8e",
)

go_repository(
    name = "com_github_onsi_ginkgo_v2",
    importpath = "github.com/onsi/ginkgo/v2",
    sum = "h1:Bi2gGVkfn6gQcjNjZJVO8Gf0FHzMPf2phUei9tejVMs=",
    version = "v2.13.2",
)

go_repository(
    name = "com_github_pelletier_go_toml_v2",
    importpath = "github.com/pelletier/go-toml/v2",
    sum = "h1:0ctb6s9mE31h0/lhu+J6OPmVeDxJn+kYnJc2jZR9tGQ=",
    version = "v2.0.8",
)

go_repository(
    name = "com_github_rangelreale_osincli",
    importpath = "github.com/RangelReale/osincli",
    sum = "h1:x8Brv0YNEe6jY3V/hQglIG2nd8g5E2Zj5ubGKkPQctQ=",
    version = "v0.0.0-20160924135400-fababb0555f2",
)

go_repository(
    name = "com_github_rasky_go_xdr",
    importpath = "github.com/rasky/go-xdr",
    sum = "h1:lbe6PJ3nOQAUvpx9P3GtsQ/jyNBOHLV+cj2++uZrpa4=",
    version = "v0.0.0-20170217172119-4930550ba2e2",
)

go_repository(
    name = "com_github_stoewer_go_strcase",
    importpath = "github.com/stoewer/go-strcase",
    sum = "h1:Z2iHWqGXH00XYgqDmNgQbIBxf3wrNq0F3feEy0ainaU=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_twitchyliquid64_golang_asm",
    importpath = "github.com/twitchyliquid64/golang-asm",
    sum = "h1:SU5vSMR7hnwNxj24w34ZyCi/FmDZTkS4MhqMhdFk5YI=",
    version = "v0.15.1",
)

go_repository(
    name = "com_github_xhit_go_str2duration_v2",
    importpath = "github.com/xhit/go-str2duration/v2",
    sum = "h1:lxklc02Drh6ynqX+DdPyp5pCKLUQpRT8bp8Ydu2Bstc=",
    version = "v2.1.0",
)

go_repository(
    name = "com_github_xlab_treeprint",
    importpath = "github.com/xlab/treeprint",
    sum = "h1:HzHnuAF1plUN2zGlAFHbSQP2qJ0ZAD3XF5XD7OesXRQ=",
    version = "v1.2.0",
)

go_repository(
    name = "com_google_cloud_go_compute",
    importpath = "cloud.google.com/go/compute",
    sum = "h1:6aKEtlUiwEpJzM001l0yFkpXmUVXaN8W+fbkb2AZNbg=",
    version = "v1.20.1",
)

go_repository(
    name = "com_google_cloud_go_compute_metadata",
    importpath = "cloud.google.com/go/compute/metadata",
    sum = "h1:mg4jlk7mCAj6xXp9UJ4fjI9VUI5rubuGBW5aJ7UnBMY=",
    version = "v0.2.3",
)

go_repository(
    name = "io_etcd_go_etcd_api_v3",
    importpath = "go.etcd.io/etcd/api/v3",
    sum = "h1:4wSsluwyTbGGmyjJktOf3wFQoTBIURXHnq9n/G/JQHs=",
    version = "v3.5.9",
)

go_repository(
    name = "io_etcd_go_etcd_client_pkg_v3",
    importpath = "go.etcd.io/etcd/client/pkg/v3",
    sum = "h1:oidDC4+YEuSIQbsR94rY9gur91UPL6DnxDCIYd2IGsE=",
    version = "v3.5.9",
)

go_repository(
    name = "io_etcd_go_etcd_client_v2",
    importpath = "go.etcd.io/etcd/client/v2",
    sum = "h1:YZ2OLi0OvR0H75AcgSUajjd5uqKDKocQUqROTG11jIo=",
    version = "v2.305.9",
)

go_repository(
    name = "io_etcd_go_etcd_client_v3",
    importpath = "go.etcd.io/etcd/client/v3",
    sum = "h1:r5xghnU7CwbUxD/fbUtRyJGaYNfDun8sp/gTr1hew6E=",
    version = "v3.5.9",
)

go_repository(
    name = "io_etcd_go_etcd_pkg_v3",
    importpath = "go.etcd.io/etcd/pkg/v3",
    sum = "h1:6R2jg/aWd/zB9+9JxmijDKStGJAPFsX3e6BeJkMi6eQ=",
    version = "v3.5.9",
)

go_repository(
    name = "io_etcd_go_etcd_raft_v3",
    importpath = "go.etcd.io/etcd/raft/v3",
    sum = "h1:ZZ1GIHoUlHsn0QVqiRysAm3/81Xx7+i2d7nSdWxlOiI=",
    version = "v3.5.9",
)

go_repository(
    name = "io_etcd_go_etcd_server_v3",
    importpath = "go.etcd.io/etcd/server/v3",
    sum = "h1:vomEmmxeztLtS5OEH7d0hBAg4cjVIu9wXuNzUZx2ZA0=",
    version = "v3.5.9",
)

go_repository(
    name = "io_k8s_component_helpers",
    importpath = "k8s.io/component-helpers",
    sum = "h1:te9ieTGzcztVktUs92X53P6BamAoP73MK0qQP0WmDqc=",
    version = "v0.28.3",
)

go_repository(
    name = "io_k8s_kms",
    importpath = "k8s.io/kms",
    sum = "h1:jYwwAe96XELNjYWv1G4kNzizcFoZ50OOElvPansbw70=",
    version = "v0.28.3",
)

go_repository(
    name = "io_k8s_sigs_json",
    importpath = "sigs.k8s.io/json",
    sum = "h1:EDPBXCAspyGV4jQlpZSudPeMmr1bNJefnuqLsRAsHZo=",
    version = "v0.0.0-20221116044647-bc3834ca7abd",
)

go_repository(
    name = "io_k8s_sigs_kube_storage_version_migrator",
    importpath = "sigs.k8s.io/kube-storage-version-migrator",
    sum = "h1:qsCecgZHgdismlTt8xCmS/3numvpxrj58RWJeIg76wc=",
    version = "v0.0.4",
)

go_repository(
    name = "io_kubevirt_api",
    importpath = "kubevirt.io/api",
    sum = "h1:vt5bOpACArNFIudx1bcE1VeejQdh5wCd7Oz/uFBIkH8=",
    version = "v1.1.1",
)

go_repository(
    name = "io_kubevirt_controller_lifecycle_operator_sdk_api",
    importpath = "kubevirt.io/controller-lifecycle-operator-sdk/api",
    sum = "h1:QMrd0nKP0BGbnxTqakhDZAUhGKxPiPiN5gSDqKUmGGc=",
    version = "v0.0.0-20220329064328-f3cc58c6ed90",
)

go_repository(
    name = "io_opentelemetry_go_contrib_instrumentation_google_golang_org_grpc_otelgrpc",
    importpath = "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc",
    sum = "h1:xFSRQBbXF6VvYRf2lqMJXxoB72XI1K/azav8TekHHSw=",
    version = "v0.35.0",
)

go_repository(
    name = "io_opentelemetry_go_contrib_instrumentation_net_http_otelhttp",
    importpath = "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp",
    sum = "h1:sxoY9kG1s1WpSYNyzm24rlwH4lnRYFXUVVBmKMBfRgw=",
    version = "v0.35.1",
)

go_repository(
    name = "io_opentelemetry_go_otel",
    importpath = "go.opentelemetry.io/otel",
    sum = "h1:Y7DTJMR6zs1xkS/upamJYk0SxxN4C9AqRd77jmZnyY4=",
    version = "v1.10.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_exporters_otlp_internal_retry",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/internal/retry",
    sum = "h1:TaB+1rQhddO1sF71MpZOZAuSPW1klK2M8XxfrBMfK7Y=",
    version = "v1.10.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace",
    sum = "h1:pDDYmo0QadUPal5fwXoY1pmMpFcdyhXOmL5drCrI3vU=",
    version = "v1.10.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace_otlptracegrpc",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc",
    sum = "h1:KtiUEhQmj/Pa874bVYKGNVdq8NPKiacPbaRRtgXi+t4=",
    version = "v1.10.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_metric",
    importpath = "go.opentelemetry.io/otel/metric",
    sum = "h1:6SiklT+gfWAwWUR0meEMxQBtihpiEs4c+vL9spDTqUs=",
    version = "v0.31.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_sdk",
    importpath = "go.opentelemetry.io/otel/sdk",
    sum = "h1:jZ6K7sVn04kk/3DNUdJ4mqRlGDiXAVuIG+MMENpTNdY=",
    version = "v1.10.0",
)

go_repository(
    name = "io_opentelemetry_go_otel_trace",
    importpath = "go.opentelemetry.io/otel/trace",
    sum = "h1:npQMbR8o7mum8uF95yFbOEJffhs1sbCOfDh8zAJiH5E=",
    version = "v1.10.0",
)

go_repository(
    name = "io_opentelemetry_go_proto_otlp",
    importpath = "go.opentelemetry.io/proto/otlp",
    sum = "h1:IVN6GR+mhC4s5yfcTbmzHYODqvWAp3ZedA2SJPI1Nnw=",
    version = "v0.19.0",
)

go_repository(
    name = "io_rsc_pdf",
    importpath = "rsc.io/pdf",
    sum = "h1:k1MczvYDUvJBe93bYd7wrZLLUEcLZAuF824/I4e5Xr4=",
    version = "v0.1.1",
)

go_repository(
    name = "org_golang_google_genproto_googleapis_api",
    importpath = "google.golang.org/genproto/googleapis/api",
    sum = "h1:m8v1xLLLzMe1m5P+gCTF8nJB9epwZQUBERm20Oy1poQ=",
    version = "v0.0.0-20230525234035-dd9d682886f9",
)

go_repository(
    name = "org_golang_google_genproto_googleapis_rpc",
    importpath = "google.golang.org/genproto/googleapis/rpc",
    sum = "h1:0nDDozoAU19Qb2HwhXadU8OcsiO/09cnTqhUtq2MEOM=",
    version = "v0.0.0-20230525234030-28d5490b6b19",
)

go_repository(
    name = "org_golang_x_arch",
    importpath = "golang.org/x/arch",
    sum = "h1:02VY4/ZcO/gBOH6PUaoiptASxtXU10jazRCP865E97k=",
    version = "v0.3.0",
)

go_repository(
    name = "org_golang_x_term",
    importpath = "golang.org/x/term",
    sum = "h1:+ThwsDv+tYfnJFhF4L8jITxu1tdTWRTZpdsWgEgjL6Q=",
    version = "v0.19.0",
)

go_repository(
    name = "org_uber_go_goleak",
    importpath = "go.uber.org/goleak",
    sum = "h1:NBol2c7O1ZokfZ0LEU9K6Whx/KnwvepVetCUhtKja4A=",
    version = "v1.2.1",
)

rules_proto_dependencies()

rules_proto_toolchains()

rpm(
    name = "alternatives-0__1.24-1.el9.x86_64",
    sha256 = "ecc6f9b8b8bf2dcbe8f70452262c275e9c530165379b025f1f8370f656e42bcb",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/alternatives-1.24-1.el9.x86_64.rpm",
    ],
)

rpm(
    name = "audit-libs-0__3.1.5-1.el9.x86_64",
    sha256 = "c152e73767630f4781480c2866935dfc0991bd3da475e9507d016cc7f26da85c",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/audit-libs-3.1.5-1.el9.x86_64.rpm",
    ],
)

rpm(
    name = "basesystem-0__11-13.el9.x86_64",
    sha256 = "2ea5b1c117cf96a8f0cdba2cf67776c2934c286e2107f8e6ba0357a63f947a03",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/basesystem-11-13.el9.noarch.rpm",
    ],
)

rpm(
    name = "bash-0__5.1.8-9.el9.x86_64",
    sha256 = "ef7898e1ad43e638d3cb336bb2fbfc4160203568c30a7e955a0317b9a35406a6",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/bash-5.1.8-9.el9.x86_64.rpm",
    ],
)

rpm(
    name = "binutils-0__2.35.2-54.el9.x86_64",
    sha256 = "a9a664b5bee30e522a6546ac5f886590920e1751662ffe3ac5b56d2b73c33837",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/binutils-2.35.2-54.el9.x86_64.rpm",
    ],
)

rpm(
    name = "binutils-gold-0__2.35.2-54.el9.x86_64",
    sha256 = "16e44629b1adf2beeb1478b486228203d7a3c6c4af3437d5314f23137db16257",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/binutils-gold-2.35.2-54.el9.x86_64.rpm",
    ],
)

rpm(
    name = "bzip2-libs-0__1.0.8-8.el9.x86_64",
    sha256 = "339ff9bae82e635de5039bc6ff8be3130514d711f9b8e5a162e8e870fbfdfae2",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/bzip2-libs-1.0.8-8.el9.x86_64.rpm",
    ],
)

rpm(
    name = "ca-certificates-0__2024.2.69_v8.0.303-91.4.el9.x86_64",
    sha256 = "345cfd71983f45555db1f5c23e155f8f6720d640b68f54538c512cdc7b791432",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/ca-certificates-2024.2.69_v8.0.303-91.4.el9.noarch.rpm",
    ],
)

rpm(
    name = "centos-gpg-keys-0__9.0-11.el9.x86_64",
    sha256 = "72cba4981776b788dde10451dfd538df1651311ea30281108f6ac290f1d78ea5",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/centos-gpg-keys-9.0-11.el9.noarch.rpm",
    ],
)

rpm(
    name = "centos-stream-release-0__9.0-11.el9.x86_64",
    sha256 = "997fecdc6da163649356bdff0f33e4f00c114e09609d1d4288fc72d10bc188bf",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/centos-stream-release-9.0-11.el9.noarch.rpm",
    ],
)

rpm(
    name = "centos-stream-repos-0__9.0-11.el9.x86_64",
    sha256 = "bf73192c94f403df472a26c69ae20a27a5320df2e487ab38c207036cac536205",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/centos-stream-repos-9.0-11.el9.noarch.rpm",
    ],
)

rpm(
    name = "cmake-filesystem-0__3.26.5-2.el9.x86_64",
    sha256 = "8e398dd9ccda66f1b64951593600553a753dd554295869c6bcc9b42a5d9ca5a3",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/AppStream/x86_64/os/Packages/cmake-filesystem-3.26.5-2.el9.x86_64.rpm",
    ],
)

rpm(
    name = "coreutils-single-0__8.32-36.el9.x86_64",
    sha256 = "e157ec6ce92a0aa08f97b0de6908da195088a87f7e61366038e34a52a5c17577",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/coreutils-single-8.32-36.el9.x86_64.rpm",
    ],
)

rpm(
    name = "cpp-0__11.5.0-2.el9.x86_64",
    sha256 = "6d19aff72241ceefe83a3d029321b380b80f7d0267e1f8d28291b94b1daf7880",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/AppStream/x86_64/os/Packages/cpp-11.5.0-2.el9.x86_64.rpm",
    ],
)

rpm(
    name = "crypto-policies-0__20240822-1.gitbaf3e06.el9.x86_64",
    sha256 = "f01a7b6cd9d00c2954ef513925c5b8d5856f8b30910baccd877e54e343419fb1",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/crypto-policies-20240822-1.gitbaf3e06.el9.noarch.rpm",
    ],
)

rpm(
    name = "cyrus-sasl-lib-0__2.1.27-21.el9.x86_64",
    sha256 = "54c23bdba8e22e5dec8a9dc63036940c0fd41e31e12d54331e66ef65e2abe1ad",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/cyrus-sasl-lib-2.1.27-21.el9.x86_64.rpm",
    ],
)

rpm(
    name = "elfutils-debuginfod-client-0__0.191-4.el9.x86_64",
    sha256 = "9346d3d6224aababd3ebf109952052e83cc394f4dae21acda249ff38d6cfab81",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/elfutils-debuginfod-client-0.191-4.el9.x86_64.rpm",
    ],
)

rpm(
    name = "elfutils-default-yama-scope-0__0.191-4.el9.x86_64",
    sha256 = "a475deea8055449614718a57a74b54edf82fb2a13a858144e1392fd1e7ed74f2",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/elfutils-default-yama-scope-0.191-4.el9.noarch.rpm",
    ],
)

rpm(
    name = "elfutils-libelf-0__0.191-4.el9.x86_64",
    sha256 = "24768dafed7fb8f1d64c4cdb9cca6ff5495f09011c0cb01bdaa0ad67d6a4dc01",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/elfutils-libelf-0.191-4.el9.x86_64.rpm",
    ],
)

rpm(
    name = "elfutils-libs-0__0.191-4.el9.x86_64",
    sha256 = "4b3badb89d83bad0640c38899c9f13d6206713377803cde605d23612e47dd740",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/elfutils-libs-0.191-4.el9.x86_64.rpm",
    ],
)

rpm(
    name = "expat-0__2.5.0-2.el9.x86_64",
    sha256 = "80ad6b136853ccef91903dc4652edff6d617b1db50093a6b5601ba5e9e7265c5",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/expat-2.5.0-2.el9.x86_64.rpm",
    ],
)

rpm(
    name = "filesystem-0__3.16-5.el9.x86_64",
    sha256 = "e3389fe64db8e7527873791f3d57b7ba5db47bc0a9661f187667c43e8dbb7b7d",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/filesystem-3.16-5.el9.x86_64.rpm",
    ],
)

rpm(
    name = "gawk-0__5.1.0-6.el9.x86_64",
    sha256 = "0119e036237ba549b66404b6107c02023896eaae4e08fda4b296dbe9fb59c8f3",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/gawk-5.1.0-6.el9.x86_64.rpm",
    ],
)

rpm(
    name = "gcc-0__11.5.0-2.el9.x86_64",
    sha256 = "ce12a9ca76e71fb4eecf4f5dd6b8f16528887b92e75405159cc7a0e93b6787c7",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/AppStream/x86_64/os/Packages/gcc-11.5.0-2.el9.x86_64.rpm",
    ],
)

rpm(
    name = "gdbm-libs-1__1.23-1.el9.x86_64",
    sha256 = "adbdf6d81052d0cdc95fbdd5e5d456fe9a268cf1a35edf347ffea1271b11c86f",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/gdbm-libs-1.23-1.el9.x86_64.rpm",
    ],
)

rpm(
    name = "glib2-0__2.68.4-15.el9.x86_64",
    sha256 = "27fd0529efa3b405d8ef06fdeb29245a1163a077a70ce86b259ed5f105f5a237",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/glib2-2.68.4-15.el9.x86_64.rpm",
    ],
)

rpm(
    name = "glibc-0__2.34-120.el9.x86_64",
    sha256 = "d10f1b9ff7ce56fb114bda9bcd90aebd16826c348a7e5c74eda6186630daa252",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/glibc-2.34-120.el9.x86_64.rpm",
    ],
)

rpm(
    name = "glibc-common-0__2.34-120.el9.x86_64",
    sha256 = "562593e8052f32bb08217686d55624c7176d9994f248bb42843a1ca491a50a11",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/glibc-common-2.34-120.el9.x86_64.rpm",
    ],
)

rpm(
    name = "glibc-devel-0__2.34-120.el9.x86_64",
    sha256 = "c414b5bb681059c2dcee010dc74e01b4af04356dffc2218d12d89b2be1cfe9c8",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/AppStream/x86_64/os/Packages/glibc-devel-2.34-120.el9.x86_64.rpm",
    ],
)

rpm(
    name = "glibc-headers-0__2.34-120.el9.x86_64",
    sha256 = "78c19f1867d820a428e1f1265f1d0a815221a3f93f5bcfc74facd9f77c3d131d",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/AppStream/x86_64/os/Packages/glibc-headers-2.34-120.el9.x86_64.rpm",
    ],
)

rpm(
    name = "glibc-langpack-xh-0__2.34-120.el9.x86_64",
    sha256 = "46a556b5c1d233d2b9b575c271177bf8be5a2741b55b79b81b53c6ee6f57b74b",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/glibc-langpack-xh-2.34-120.el9.x86_64.rpm",
    ],
)

rpm(
    name = "gmp-1__6.2.0-13.el9.x86_64",
    sha256 = "932d8eb40b6b66bf6f54d95b8aa70c9f749951f719591f0b242753a60aef4b2b",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/gmp-6.2.0-13.el9.x86_64.rpm",
    ],
)

rpm(
    name = "gnutls-0__3.8.3-4.el9.x86_64",
    sha256 = "feaaa4dcb80f6433fedb61c5fc3870134faaf40656e028006731ea421b754328",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/gnutls-3.8.3-4.el9.x86_64.rpm",
    ],
)

rpm(
    name = "grep-0__3.6-5.el9.x86_64",
    sha256 = "1b77e1fc2072e009f6d2eb6216192e487cdae4d18dcd7063882319dbcc03861e",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/grep-3.6-5.el9.x86_64.rpm",
    ],
)

rpm(
    name = "kernel-headers-0__5.14.0-500.el9.x86_64",
    sha256 = "57726f60303e27888de119347fe2db18c26061d9db04e5b184a80841d480629b",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/AppStream/x86_64/os/Packages/kernel-headers-5.14.0-500.el9.x86_64.rpm",
    ],
)

rpm(
    name = "keyutils-libs-0__1.6.3-1.el9.x86_64",
    sha256 = "7e6b55523a8e7cec7171e5f71db634a88d3c37a04e798bed743b35ddc146164e",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/keyutils-libs-1.6.3-1.el9.x86_64.rpm",
    ],
)

rpm(
    name = "krb5-libs-0__1.21.1-3.el9.x86_64",
    sha256 = "8bebd20c005268e56e10cc3f8e17a6604215aeeb8d2f58accff62c8f19e62bbf",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/krb5-libs-1.21.1-3.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libacl-0__2.3.1-4.el9.x86_64",
    sha256 = "e00201fc534d05ad3d3243ab400e6c15983aa5c0a3996d344dc016f7ec937abc",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libacl-2.3.1-4.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libaio-0__0.3.111-13.el9.x86_64",
    sha256 = "96b8c1b764cc383db3661f1a9544bab7e59ee914da400857054f0ad1f6eb1742",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libaio-0.3.111-13.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libattr-0__2.5.1-3.el9.x86_64",
    sha256 = "b51006462a9feccf792430bab9dbc5e5933638bc91af7c0ac35c6529c9ea4a55",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libattr-2.5.1-3.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libblkid-0__2.37.4-20.el9.x86_64",
    sha256 = "b980b5b7115f03d7fca51a87084911afba3af759e3caa77838a813ea4813ad35",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libblkid-2.37.4-20.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libbrotli-0__1.0.9-6.el9.x86_64",
    sha256 = "a768d7fab68a758acaea3d114f0bd710cc147a50aa688dbac4b964a52673bc76",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libbrotli-1.0.9-6.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libcap-0__2.48-9.el9.x86_64",
    sha256 = "efb63957e8566a66e2b8b0b9b37e9d6e339b586887c16d878113613d6e3c792b",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libcap-2.48-9.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libcap-ng-0__0.8.2-7.el9.x86_64",
    sha256 = "84a21c19d507bbbb6049f771b5a11e28522c9522e414e6323cd92c1f6c12cbf5",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libcap-ng-0.8.2-7.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libcom_err-0__1.46.5-5.el9.x86_64",
    sha256 = "7d4ceee5318b213e4ee90514f00abb58ff06c6a1d5d7bc00802e16d0eda999cf",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libcom_err-1.46.5-5.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libcurl-0__7.76.1-31.el9.x86_64",
    sha256 = "555b75b749d4feaf74b0e6d8e754353ac2feb99888d622705a90aa0448f070b3",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libcurl-7.76.1-31.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libcurl-devel-0__7.76.1-31.el9.x86_64",
    sha256 = "801b52f1a69bab444ff53cab2f5694768f0eef8bb2ba1e5615057e4637dc2b2b",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/AppStream/x86_64/os/Packages/libcurl-devel-7.76.1-31.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libevent-0__2.1.12-6.el9.x86_64",
    sha256 = "45d1aecbf5cbbdd6c6e819d9088f956751e1c0d9b93362fd677f6e03f4998763",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libevent-2.1.12-6.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libffi-0__3.4.2-8.el9.x86_64",
    sha256 = "7107f69df5163e58e1c3757f6c7229bb68bafe50ea9a5d0bc3e84506190c4794",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libffi-3.4.2-8.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libgcc-0__11.5.0-2.el9.x86_64",
    sha256 = "72c142d95964627ba7d9119341b3764a71249da644e62fe6070798dc72c6b27e",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libgcc-11.5.0-2.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libgomp-0__11.5.0-2.el9.x86_64",
    sha256 = "bb7423f44ec4cd27b9d76105a426146c988caf22b62cfdf2b3f18597e8f6233b",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libgomp-11.5.0-2.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libidn2-0__2.3.0-7.el9.x86_64",
    sha256 = "e2183f1db43a2d9cb6c4b350d7e488da09c823d33177ab9f94464e9685b684d6",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libidn2-2.3.0-7.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libmount-0__2.37.4-20.el9.x86_64",
    sha256 = "562bb6596fc1558cc5d2a8143b9fca71728f0a728f402842981d6a6c4e19846b",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libmount-2.37.4-20.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libmpc-0__1.2.1-4.el9.x86_64",
    sha256 = "39e1496f4b848437043224fa9bc223a03733a7be3f4c7c9cc74eb1d8b81d9b06",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/AppStream/x86_64/os/Packages/libmpc-1.2.1-4.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libnghttp2-0__1.43.0-6.el9.x86_64",
    sha256 = "f081b41a0063262e4f05e63c4dbc317c63abc8c5dc8c568ad13dfb3c3d24f140",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libnghttp2-1.43.0-6.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libpkgconf-0__1.7.3-10.el9.x86_64",
    sha256 = "482e1e6f3bb6765c83ff8c813c3cc7fc06ae40d58759f52b2a4b1d41f63011cc",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libpkgconf-1.7.3-10.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libpsl-0__0.21.1-5.el9.x86_64",
    sha256 = "d907915a6f01f78919e385ce4e3f4ddc3a644f48049a35721d864954d5eb9680",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libpsl-0.21.1-5.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libselinux-0__3.6-1.el9.x86_64",
    sha256 = "f81fcfdf9c30fcf8dfd54c1906976c0390e6472b6cb37fe95d80d63a1b0ab7e3",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libselinux-3.6-1.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libsemanage-0__3.6-1.el9.x86_64",
    sha256 = "656aed89ad772da9264d6f1bce065d744896c8d7cb31c6fc199bbcdc02460f30",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libsemanage-3.6-1.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libsepol-0__3.6-1.el9.x86_64",
    sha256 = "0f1498b526009a2e6da5438717fe981131a087058595d094671e8425ecf7cf50",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libsepol-3.6-1.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libsigsegv-0__2.13-4.el9.x86_64",
    sha256 = "f3ee877e3dd76f61587d2e4687a033bc7b0d7138b2f84cf0f1af62efb748777b",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libsigsegv-2.13-4.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libssh-0__0.10.4-13.el9.x86_64",
    sha256 = "9b6e56105ea9077d005d400dab609b5eb866521d76bb8d1b240c08a50a23e0ce",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libssh-0.10.4-13.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libssh-config-0__0.10.4-13.el9.x86_64",
    sha256 = "8164abb5201d80698301d48e9c284c8ed01c6599b6f37c59faadb6482f6db76f",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libssh-config-0.10.4-13.el9.noarch.rpm",
    ],
)

rpm(
    name = "libstdc__plus____plus__-0__11.5.0-2.el9.x86_64",
    sha256 = "65e3a8674b62fade64c6d293caef93d694bace34d9e3e50014a1a11ce813dc00",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libstdc++-11.5.0-2.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libtasn1-0__4.16.0-8.el9.x86_64",
    sha256 = "926cb611745095f3c1a75acb7b090ab906d4c5e22d2d8e8c34fc15398f177e55",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libtasn1-4.16.0-8.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libunistring-0__0.9.10-15.el9.x86_64",
    sha256 = "42f4771c9226b8c14f83801deb383f41b3675174026ca66e14b0502d5b4ac949",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libunistring-0.9.10-15.el9.x86_64.rpm",
    ],
)

rpm(
    name = "liburing-0__2.5-1.el9.x86_64",
    sha256 = "dbeb30f4333754b3d7f16a984a3e663c3e20f2c2fc34717037aa9b875a4f2de6",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/AppStream/x86_64/os/Packages/liburing-2.5-1.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libuuid-0__2.37.4-20.el9.x86_64",
    sha256 = "880975f7b3ab5244889357ee328e7befd0b9f8231c60f277039346ccc54c99ce",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libuuid-2.37.4-20.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libverto-0__0.3.2-3.el9.x86_64",
    sha256 = "83d6e258d8d4e021e339ad8b7a03c68e136285f740939446ec2795edbdc5e922",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libverto-0.3.2-3.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libxcrypt-0__4.4.18-3.el9.x86_64",
    sha256 = "b07e135ce7cea03dd1ba3205606e0d4d42a51d1249a708b804521cfe3bc3454d",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libxcrypt-4.4.18-3.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libxcrypt-devel-0__4.4.18-3.el9.x86_64",
    sha256 = "4289873ed55d6a000381ce1c904fdb0d4a0c7566d6119d950ed299c2626e1893",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/AppStream/x86_64/os/Packages/libxcrypt-devel-4.4.18-3.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libxml2-0__2.9.13-6.el9.x86_64",
    sha256 = "f7c7e3a76a247f4106fc1cbbb5d170d993ad6fe7b1d27c7cee30d4048c3cb2d2",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libxml2-2.9.13-6.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libxml2-devel-0__2.9.13-6.el9.x86_64",
    sha256 = "e29caaa4ab1851eb05739895fa9706fa7d33dee3bae2dd03781257f7ffb98b74",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/AppStream/x86_64/os/Packages/libxml2-devel-2.9.13-6.el9.x86_64.rpm",
    ],
)

rpm(
    name = "libzstd-0__1.5.1-2.el9.x86_64",
    sha256 = "26edf97c452888a4b29dedec289f142bb0028367473e17311114f3e7cdba7f0b",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/libzstd-1.5.1-2.el9.x86_64.rpm",
    ],
)

rpm(
    name = "make-1__4.3-8.el9.x86_64",
    sha256 = "cf9c9f6912190a2849ed7a3474ef06e20b572a2180ec3176c46d256a81bc5bb3",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/make-4.3-8.el9.x86_64.rpm",
    ],
)

rpm(
    name = "mpfr-0__4.1.0-7.el9.x86_64",
    sha256 = "ad4526486c80739698dee8604a62c4d3953a6c030c23e30b368f44674a257dfb",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/mpfr-4.1.0-7.el9.x86_64.rpm",
    ],
)

rpm(
    name = "ncurses-base-0__6.2-10.20210508.el9.x86_64",
    sha256 = "3ccbaa39db3cc8ae78f9da42cab3859d0cdb302dd947cdebdb83a702e89074cf",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/ncurses-base-6.2-10.20210508.el9.noarch.rpm",
    ],
)

rpm(
    name = "ncurses-libs-0__6.2-10.20210508.el9.x86_64",
    sha256 = "c130660ede7c78b7b759d9752a68c5a1ebab8cc16b361450f8b3ab2003d63320",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/ncurses-libs-6.2-10.20210508.el9.x86_64.rpm",
    ],
)

rpm(
    name = "nettle-0__3.9.1-1.el9.x86_64",
    sha256 = "ecde91b76480ddcf0e617f99d41865e80ff91d345bb051a92dcd1fe9fb373ba9",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/nettle-3.9.1-1.el9.x86_64.rpm",
    ],
)

rpm(
    name = "numactl-libs-0__2.0.18-2.el9.x86_64",
    sha256 = "5a3c2f0c20c67b6b1b3ff45347307debc4429e2cc50ec0aef0eaff749d95195a",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/numactl-libs-2.0.18-2.el9.x86_64.rpm",
    ],
)

rpm(
    name = "openldap-0__2.6.6-3.el9.x86_64",
    sha256 = "961c79091639f6f82bf5fc172854d71eae4cb98f83494074857bdfe10687af41",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/openldap-2.6.6-3.el9.x86_64.rpm",
    ],
)

rpm(
    name = "openssl-devel-1__3.2.2-4.el9.x86_64",
    sha256 = "a323d1e6369355969abe061d3989a755b0dce3d612d03700557a1a99d9936cd3",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/AppStream/x86_64/os/Packages/openssl-devel-3.2.2-4.el9.x86_64.rpm",
    ],
)

rpm(
    name = "openssl-libs-1__3.2.2-4.el9.x86_64",
    sha256 = "f618b202712317772ee17dbd7fe7e20a31a7a3fc2a4a6838147a81d3daea197d",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/openssl-libs-3.2.2-4.el9.x86_64.rpm",
    ],
)

rpm(
    name = "p11-kit-0__0.25.3-2.el9.x86_64",
    sha256 = "2da9531864a099ca5a262361d1c1ff51095ef834d68ecd01c98a4bede61ecae9",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/p11-kit-0.25.3-2.el9.x86_64.rpm",
    ],
)

rpm(
    name = "p11-kit-trust-0__0.25.3-2.el9.x86_64",
    sha256 = "42a4f210469584d6171f912ea5c217cf59e5dddcc73a3a3bc29be7d86aafaa41",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/p11-kit-trust-0.25.3-2.el9.x86_64.rpm",
    ],
)

rpm(
    name = "pcre-0__8.44-4.el9.x86_64",
    sha256 = "8de7ca63e65ae5b063c37e5a1de1e7024e5b74024df170362e807b05f505eadc",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/pcre-8.44-4.el9.x86_64.rpm",
    ],
)

rpm(
    name = "pcre2-0__10.40-6.el9.x86_64",
    sha256 = "89ddfdc4ef6df4e623ae20be1d9d3407ac56ba77aaf590e211d5790ce307fe66",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/pcre2-10.40-6.el9.x86_64.rpm",
    ],
)

rpm(
    name = "pcre2-syntax-0__10.40-6.el9.x86_64",
    sha256 = "63813c993604e0625dede1c0f6dce0ac49e68a2f0c229c5492dd6cb8e059a9df",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/pcre2-syntax-10.40-6.el9.noarch.rpm",
    ],
)

rpm(
    name = "pkgconf-0__1.7.3-10.el9.x86_64",
    sha256 = "14c0ce0c3c7be16732e16ea0122a2f5eed652a9ec915f739d83fd5c76872cb60",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/pkgconf-1.7.3-10.el9.x86_64.rpm",
    ],
)

rpm(
    name = "pkgconf-m4-0__1.7.3-10.el9.x86_64",
    sha256 = "f670513f9304c282c08672407a80874b15307e5d1f23dfbd344fddd462e1b3b8",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/pkgconf-m4-1.7.3-10.el9.noarch.rpm",
    ],
)

rpm(
    name = "pkgconf-pkg-config-0__1.7.3-10.el9.x86_64",
    sha256 = "4cbc37f3776574da72a96d836bcdc7974b92ea3705d920ab1d976074d5b689af",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/pkgconf-pkg-config-1.7.3-10.el9.x86_64.rpm",
    ],
)

rpm(
    name = "publicsuffix-list-dafsa-0__20210518-3.el9.x86_64",
    sha256 = "cf13c91ecf9f734a7d2d3ae757c041ebdb7c4e2c7298733d74d4d87c3fd995a6",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/publicsuffix-list-dafsa-20210518-3.el9.noarch.rpm",
    ],
)

rpm(
    name = "python3-0__3.9.19-8.el9.x86_64",
    sha256 = "c3dd036263e18982d57b3cece244696d1dd51bb3a11a03e6ef30bc3bdef8b0b7",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/python3-3.9.19-8.el9.x86_64.rpm",
    ],
)

rpm(
    name = "python3-devel-0__3.9.19-8.el9.x86_64",
    sha256 = "71bb89e0d2eb4cb66b44cc9e5ce5e5fce834fc3cabec7e198178d9552c5149fb",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/AppStream/x86_64/os/Packages/python3-devel-3.9.19-8.el9.x86_64.rpm",
    ],
)

rpm(
    name = "python3-libs-0__3.9.19-8.el9.x86_64",
    sha256 = "a3d7ed5494de35bb34404309131f4b8c4a40ebabe5582767d0a72f8f964f0669",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/python3-libs-3.9.19-8.el9.x86_64.rpm",
    ],
)

rpm(
    name = "python3-pip-0__21.3.1-1.el9.x86_64",
    sha256 = "5c552404c6812cf6a2cdb130da36231a91b3924d5771d854f2489d429b1d04f9",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/AppStream/x86_64/os/Packages/python3-pip-21.3.1-1.el9.noarch.rpm",
    ],
)

rpm(
    name = "python3-pip-wheel-0__21.3.1-1.el9.x86_64",
    sha256 = "51c88553dff297396e6bf3494474cc88af1bceebc0e78fb45f02a4167777319e",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/python3-pip-wheel-21.3.1-1.el9.noarch.rpm",
    ],
)

rpm(
    name = "python3-setuptools-wheel-0__53.0.0-13.el9.x86_64",
    sha256 = "ec13fb51bd6b970771abdd5d3cd11c55e64f9cf7a55f8e158d1e2157ecd04ffa",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/python3-setuptools-wheel-53.0.0-13.el9.noarch.rpm",
    ],
)

rpm(
    name = "qemu-img-17__9.0.0-9.el9.x86_64",
    sha256 = "2bf775ba6b143e12ced2aa33a62da5e0a3c4800263b8a918f3ffb483da4b8399",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/AppStream/x86_64/os/Packages/qemu-img-9.0.0-9.el9.x86_64.rpm",
    ],
)

rpm(
    name = "readline-0__8.1-4.el9.x86_64",
    sha256 = "7e43f8266c9bd4fe4df2798d23d5c6e529f93030b7ef27c0dfbb35fe1bf55f06",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/readline-8.1-4.el9.x86_64.rpm",
    ],
)

rpm(
    name = "sed-0__4.8-9.el9.x86_64",
    sha256 = "110db587738ad9dbf8aa5a9dd2f442b12d6e760697e0f94c0b7de57741cbbc0d",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/sed-4.8-9.el9.x86_64.rpm",
    ],
)

rpm(
    name = "setup-0__2.13.7-10.el9.x86_64",
    sha256 = "1afafe063300e6bbe84c953faf0082ac89c54cf95cf622b4db0b2ef68d013f02",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/setup-2.13.7-10.el9.noarch.rpm",
    ],
)

rpm(
    name = "shadow-utils-2__4.9-9.el9.x86_64",
    sha256 = "5900b308b91db106efaf56b8cd37b61c9aafb15f7a90e7b014a9dc1804780d2c",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/shadow-utils-4.9-9.el9.x86_64.rpm",
    ],
)

rpm(
    name = "sqlite-libs-0__3.34.1-7.el9.x86_64",
    sha256 = "5d1f3688c827ce01e84c06892a325b31a76de65613d7c0e867e321b7dcc45c68",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/sqlite-libs-3.34.1-7.el9.x86_64.rpm",
    ],
)

rpm(
    name = "tzdata-0__2024a-2.el9.x86_64",
    sha256 = "d9647b27e5d4df583bf3acba85e70e4b89c90834110d18a0b08c5406e24cffc8",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/tzdata-2024a-2.el9.noarch.rpm",
    ],
)

rpm(
    name = "xz-devel-0__5.2.5-8.el9.x86_64",
    sha256 = "c9ba742f76b51daab90eb581de4c2a01623686fbbaaf87e59644b622d7e10fd7",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/AppStream/x86_64/os/Packages/xz-devel-5.2.5-8.el9.x86_64.rpm",
    ],
)

rpm(
    name = "xz-libs-0__5.2.5-8.el9.x86_64",
    sha256 = "2ad419ac36ffe4ba7ebc55372621db5a3ee497a96801a06ba8f7f020270ae851",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/xz-libs-5.2.5-8.el9.x86_64.rpm",
    ],
)

rpm(
    name = "zlib-0__1.2.11-40.el9.x86_64",
    sha256 = "4ed20865bc82692e597f350bb0cacfdcc3ebd347df5117ad3afcf8fdc5c43d79",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/BaseOS/x86_64/os/Packages/zlib-1.2.11-40.el9.x86_64.rpm",
    ],
)

rpm(
    name = "zlib-devel-0__1.2.11-40.el9.x86_64",
    sha256 = "862c0a619524e03926ae6778dfc06f802327035ad72a2dc2b58517d009f5b234",
    urls = [
        "https://composes.stream.centos.org/development/latest-CentOS-Stream/compose/AppStream/x86_64/os/Packages/zlib-devel-1.2.11-40.el9.x86_64.rpm",
    ],
)
