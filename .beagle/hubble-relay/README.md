# cilium

## hubble-relay

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
bash .beagle/hubble-relay/build.sh

# arm64
docker run -it --rm \
-v $PWD/:/go/src/github.com/cilium/cilium \
-v /go/pkg/mod:/go/pkg/mod \
-v /tmp/.cache:/root/.cache \
-w /go/src/github.com/cilium/cilium \
-e TARGETOS=linux \
-e TARGETARCH=arm64 \
registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-builder:v1.9-20210814-arm64 \
bash .beagle/hubble-relay/build.sh

# ppc64le
docker run -it --rm \
-v $PWD/:/go/src/github.com/cilium/cilium \
-v /go/pkg/mod:/go/pkg/mod \
-v /tmp/.cache:/root/.cache \
-w /go/src/github.com/cilium/cilium \
-e TARGETOS=linux \
-e TARGETARCH=ppc64le \
registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-builder:v1.9-20210814-ppc64le \
bash .beagle/hubble-relay/build.sh
```
