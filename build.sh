export REGISTRY_ORG=arturshadnik
export REGISTRY=quay.io
export REGISTRY_TAG=latest-dev

CONTROLLER_IMAGE=quay.io/${REGISTRY_ORG}/forklift-controller:${REGISTRY_TAG}
OPERATOR_IMAGE=quay.io/${REGISTRY_ORG}/forklift-operator:${REGISTRY_TAG}
VIRT_V2V_IMAGE=quay.io/${REGISTRY_ORG}/forklift-virt-v2v:${REGISTRY_TAG}

bazel run --package_path=virt-v2v push-forklift-virt-v2v
# bazel run push-forklift-controller --@io_bazel_rules_docker//transitions:enable=false

# bazel run push-forklift-operator --@io_bazel_rules_docker//transitions:enable=false
# bazel run push-forklift-operator-bundle --action_env OPERATOR_IMAGE=${OPERATOR_IMAGE} --action_env CONTROLLER_IMAGE=${CONTROLLER_IMAGE} --action_env VIRT_V2V_IMAGE=${VIRT_V2V_IMAGE} --@io_bazel_rules_docker//transitions:enable=false
# # The build of the catalog requires already pushed bundle
# # For http registry add --action_env OPM_OPTS="--use-http"
# bazel run push-forklift-operator-index --action_env REGISTRY=${REGISTRY} --action_env REGISTRY_ORG=${REGISTRY_ORG} --action_env REGISTRY_TAG=${REGISTRY_TAG} --@io_bazel_rules_docker//transitions:enable=false
