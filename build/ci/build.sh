set -euxo pipefail

docker build -t $CI_REGISTRY_IMAGE --file ./build/.Dockerfile .