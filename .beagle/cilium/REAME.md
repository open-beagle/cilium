# cilium

## debug

```bash
# amd64
docker run -it --rm \
-v $PWD/:/go/src/github.com/cilium/cilium \
-v /go/pkg/mod:/go/pkg/mod \
-v /tmp/.cache:/root/.cache \
-w /go/src/github.com/cilium/cilium \
-e TARGETOS=linux \
-e TARGETARCH=amd64 \
registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-builder:v1.9-20210814-amd64 \
bash .beagle/cilium/build.sh

docker run -it --rm \
-v $PWD/:/go/src/github.com/cilium/cilium \
-v /go/pkg/mod:/go/pkg/mod \
-v /tmp/.cache:/root/.cache \
-w /go/src/github.com/cilium/cilium \
-e TARGETOS=linux \
-e TARGETARCH=arm64 \
-e GOARCH=arm64 \
-e CC=aarch64-linux-gnu-gcc \
registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-builder:v1.9-20220412-cross \
bash .beagle/cilium/build.sh

file .buildx/cilium/linux/arm64/usr/bin/cilium-agent
ldd .buildx/cilium/linux/arm64/usr/bin/cilium-agent

docker run -it --rm \
-v $PWD/:/go/src/github.com/cilium/cilium \
-v /go/pkg/mod:/go/pkg/mod \
-v /tmp/.cache:/root/.cache \
-w /go/src/github.com/cilium/cilium \
-e TARGETOS=linux \
-e TARGETARCH=ppc64le \
-e GOARCH=ppc64le \
-e CC=powerpc64le-linux-gnu-gcc \
registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-builder:v1.9-20220412-cross \
bash .beagle/cilium/build.sh

file .buildx/cilium/linux/ppc64le/usr/bin/cilium-agent
ldd .buildx/cilium/linux/ppc64le/usr/bin/cilium-agent

# arm64
docker run -it --rm \
-v $PWD/:/go/src/github.com/cilium/cilium \
-v /go/pkg/mod:/go/pkg/mod \
-v /tmp/.cache:/root/.cache \
-w /go/src/github.com/cilium/cilium \
-e TARGETOS=linux \
-e TARGETARCH=arm64 \
registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-builder:v1.9-20210814-arm64 \
bash .beagle/cilium/build.sh

# ppc64le
docker run -it --rm \
-v $PWD/:/go/src/github.com/cilium/cilium \
-v /go/pkg/mod:/go/pkg/mod \
-v /tmp/.cache:/root/.cache \
-w /go/src/github.com/cilium/cilium \
-e TARGETOS=linux \
-e TARGETARCH=ppc64le \
registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-builder:v1.9-20210814-ppc64le \
bash .beagle/cilium/build.sh
```
