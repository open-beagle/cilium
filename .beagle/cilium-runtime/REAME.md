# cilium-runtime

images/runtime/Dockerfile

```bash
# amd64
docker build \
  --build-arg UBUNTU_IMAGE=registry.cn-qingdao.aliyuncs.com/wod/ubuntu:20.04-amd64 \
  --build-arg GOPS_IMAGE=registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-gops:v0.3.10 \
  --build-arg CNI_IMAGE=registry.cn-qingdao.aliyuncs.com/wod/kube-cni-plugins:v0.9.1-amd64 \
  --build-arg LLVM_IMAGE=registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-llvm:20210814 \
  --build-arg BPF_IMAGE=registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-bpftool:20210814 \
  --build-arg IPROUTE_IMAGE=registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-iproute2:20211011 \
  --build-arg ARCH=amd64 \
  --tag registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-runtime:v1.9-20211011-amd64 \
  --file ./.beagle/cilium-runtime/dockerfile ./images/runtime

docker push registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-runtime:v1.9-20211011-amd64

# arm64
docker build \
  --build-arg UBUNTU_IMAGE=registry.cn-qingdao.aliyuncs.com/wod/ubuntu:20.04-arm64 \
  --build-arg GOPS_IMAGE=registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-gops:v0.3.10 \
  --build-arg CNI_IMAGE=registry.cn-qingdao.aliyuncs.com/wod/kube-cni-plugins:v0.9.1-arm64 \
  --build-arg LLVM_IMAGE=registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-llvm:20210814 \
  --build-arg BPF_IMAGE=registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-bpftool:20210814 \
  --build-arg IPROUTE_IMAGE=registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-iproute2:20211011 \
  --build-arg ARCH=arm64 \
  --tag registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-runtime:v1.9-20211011-arm64 \
  --file ./.beagle/cilium-runtime/dockerfile ./images/runtime

docker push registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-runtime:v1.9-20211011-arm64

# ppc64le
docker build \
  --build-arg UBUNTU_IMAGE=registry.cn-qingdao.aliyuncs.com/wod/ubuntu:20.04-ppc64le \
  --build-arg GOPS_IMAGE=registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-gops:v0.3.10 \
  --build-arg CNI_IMAGE=registry.cn-qingdao.aliyuncs.com/wod/kube-cni-plugins:v0.9.1-ppc64le \
  --build-arg LLVM_IMAGE=registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-llvm:20210814 \
  --build-arg BPF_IMAGE=registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-bpftool:20210814 \
  --build-arg IPROUTE_IMAGE=registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-iproute2:20211011 \
  --build-arg ARCH=ppc64le \
  --tag registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-runtime:v1.9-20211011-ppc64le \
  --file ./.beagle/cilium-runtime/dockerfile ./images/runtime

docker push registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-runtime:v1.9-20211011-ppc64le
```
