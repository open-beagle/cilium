# cilium-builder

images/builder/Dockerfile

```bash
# cross
docker build \
  --build-arg GOLANG_IMAGE=registry.cn-qingdao.aliyuncs.com/wod/golang:1.17.8-bullseye-amd64 \
  --build-arg COMPILERS_IMAGE=registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-image-compilers:20210814-bazel \
  --tag registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-builder:v1.9-20220412-cross \
  --file ./.beagle/cilium-builder/dockerfile ./.beagle/cilium-builder/

docker push registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-builder:v1.9-20220412-cross

# amd64
docker build \
  --build-arg GOLANG_IMAGE=registry.cn-qingdao.aliyuncs.com/wod/golang:1.17.8-bullseye-amd64 \
  --build-arg COMPILERS_IMAGE=registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-image-compilers:20210814-amd64 \
  --tag registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-builder:v1.9-20210814-amd64 \
  --file ./.beagle/cilium-builder/dockerfile ./.beagle/cilium-builder/

docker push registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-builder:v1.9-20210814-amd64

# arm64
docker build \
  --build-arg GOLANG_IMAGE=registry.cn-qingdao.aliyuncs.com/wod/golang:1.17.8-bullseye-arm64 \
  --build-arg COMPILERS_IMAGE=registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-image-compilers:20210814-arm64 \
  --tag registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-builder:v1.9-20210814-arm64 \
  --file ./.beagle/cilium-builder/dockerfile ./.beagle/cilium-builder/

docker push registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-builder:v1.9-20210814-arm64

# ppc64le
docker build \
  --build-arg GOLANG_IMAGE=registry.cn-qingdao.aliyuncs.com/wod/golang:1.17.8-bullseye-ppc64le \
  --build-arg COMPILERS_IMAGE=registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-image-compilers:20210814-ppc64le \
  --tag registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-builder:v1.9-20210814-ppc64le \
  --file ./.beagle/cilium-builder/dockerfile ./.beagle/cilium-builder/

docker push registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-builder:v1.9-20210814-ppc64le
```
