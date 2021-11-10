# etcd

## docker install

1、拉取最新镜像
docker pull bitnami/etcd:latest

2、建立网络
docker network create app-tier --driver bridge

3、启动
docker run -d --name etcd-server \
--network app-tier \
--publish 2379:2379 \
--publish 2380:2380 \
--env ALLOW_NONE_AUTHENTICATION=yes \
--env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379 \
bitnami/etcd:latest

4、查看etcd版本
curl -L http://127.0.0.1:2379/version
