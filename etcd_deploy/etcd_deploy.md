## 部署一个单节点的etcd

    sudo docker run -d --name etcd \
    -p 2379:2379 \
    -e ALLOW_NONE_AUTHENTICATION=yes \
    -e ETCD_LISTEN_CLIENT_URLS=http://0.0.0.0:2379 \
    -e ETCD_ADVERTISE_CLIENT_URLS=http://0.0.0.0:2379 \
    bitnami/etcd:latest

