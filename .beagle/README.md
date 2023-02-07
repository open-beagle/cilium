# cilium

<https://github.com/cilium/cilium>

## repo

```bash
git remote add upstream git@github.com:cilium/cilium.git

git fetch upstream

git merge v1.11.13
```

## build

完整的 cilium 项目运行需要五个镜像

- cilium/cilium
- cilium/operator
- cilium/hubble-relay
- cilium/hubble-ui
- cilium/hubble-ui-backend

本项目编译产生其中三个镜像

- cilium/cilium
- cilium/operator
- cilium/hubble-relay

cilium , 编译难度“地狱级”。

- 基础镜像 cilium-runtime，这个镜像由 llvm，iproute2, bpftool 组成。
- 构建镜像 cilium-builder，本身需要由 cilium/cilium-builder 来进行源码构建。
- 依赖镜像 cilium-envoy，envoy 编译难度大，还好提供了 arm 版，负责深度处理过滤流量的事务。
- 依赖镜像 hubble，golang 镜像简单，负责观察流量。

依赖外部项目

- cloud/cilium-envoy
- cloud/cilium-hubble
- cloud/cilium-image-tools
- cloud/cilium-hubble-ui

## cross-build

由于 proxylib 库，在编译过程中使用了 buildmod=c-shared 模式进行 go build，这个模式暂时不支持 mips64le 架构下编译，
所以当前的 cilium 镜像无法支持 mips64le 架构.

## hub.docker.com

```bash
docker pull cilium/cilium:v1.11.13 && \
docker tag cilium/cilium:v1.11.13 registry.cn-qingdao.aliyuncs.com/wod/cilium:v1.11.13 && \
docker push registry.cn-qingdao.aliyuncs.com/wod/cilium:v1.11.13

docker pull cilium/operator-generic:v1.11.13 && \
docker tag cilium/operator-generic:v1.11.13 registry.cn-qingdao.aliyuncs.com/wod/cilium-operator-generic:v1.11.13 && \
docker push registry.cn-qingdao.aliyuncs.com/wod/cilium-operator-generic:v1.11.13

docker pull cilium/hubble-relay:v1.11.13 && \
docker tag cilium/hubble-relay:v1.11.13 registry.cn-qingdao.aliyuncs.com/wod/cilium-hubble-relay:v1.11.13 && \
docker push registry.cn-qingdao.aliyuncs.com/wod/cilium-hubble-relay:v1.11.13
```

## cache

```bash
# 构建缓存-->推送缓存至服务器
docker run --rm \
  -e PLUGIN_REBUILD=true \
  -e PLUGIN_ENDPOINT=$PLUGIN_ENDPOINT \
  -e PLUGIN_ACCESS_KEY=$PLUGIN_ACCESS_KEY \
  -e PLUGIN_SECRET_KEY=$PLUGIN_SECRET_KEY \
  -e DRONE_REPO_OWNER="open-beagle" \
  -e DRONE_REPO_NAME="cilium" \
  -e PLUGIN_MOUNT="./.git" \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  registry.cn-qingdao.aliyuncs.com/wod/devops-s3-cache:1.0

# 读取缓存-->将缓存从服务器拉取到本地
docker run --rm \
  -e PLUGIN_RESTORE=true \
  -e PLUGIN_ENDPOINT=$PLUGIN_ENDPOINT \
  -e PLUGIN_ACCESS_KEY=$PLUGIN_ACCESS_KEY \
  -e PLUGIN_SECRET_KEY=$PLUGIN_SECRET_KEY \
  -e DRONE_REPO_OWNER="open-beagle" \
  -e DRONE_REPO_NAME="cilium" \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  registry.cn-qingdao.aliyuncs.com/wod/devops-s3-cache:1.0
```