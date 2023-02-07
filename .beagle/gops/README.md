# gops

```bash
# git
docker run --rm \
-v $PWD/.buildx/gops/src:/go/src/github.com/google/gops \
-v $PWD/.beagle/gops/:/.beagle/gops/ \
-w /.beagle/gops/ \
registry.cn-qingdao.aliyuncs.com/wod/golang:1.16.6-alpine \
./gops.sh

# devops-go-arch
docker run -it --rm \
-w /go/src/github.com/google/gops/ \
-v $PWD/.buildx/gops/src/:/go/src/github.com/google/gops/ \
-v $PWD/.buildx/gops/bin/:/go/src/github.com/google/gops/dist/ \
-e CI_WORKSPACE=/go/src/github.com/google/gops \
-e PLUGIN_BINARY=gops \
-e PLUGIN_GOARGS='-mod vendor' \
registry.cn-qingdao.aliyuncs.com/wod/devops-go-arch:1.16.6-alpine

# gops
docker build --rm \
--tag registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-gops:v0.3.10 \
--file ./.beagle/gops/Dockerfile ./.buildx/gops/bin/

docker push registry.cn-qingdao.aliyuncs.com/wod-arm/cilium-gops:v0.3.10
```
