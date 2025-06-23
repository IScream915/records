## 本教程将逐步介绍如何在一台ubuntu虚拟机中部署一个包含zookeeper的kafka服务

### 一. 部署环境
#### 1. 示例系统: ubuntu 24.04 LTS
#### 2. 示例docker版本 Docker version 28.0.1

### 二. 创建一个docker network
#### 在不同的docker container A, B中, A和B是无法相互通讯的, 由于kafka的部署通常需要建立在zookeeper的基础上(也有启用KRaft而无需zookeeper的kafka, 但是本教程仅针对zookeeper-kafka架构的kafka部署), 因此需要创建一个网络使得不同的docker之间能够相互通讯.

这里我们创建一个名为zookeeper-kafka的桥接网络用于zookeeper和kafka通讯

sudo docker network create zookeeper-kafka --driver bridge 

### 三. 部署zookeeper container
#### 本教程仅针对不开启KRaft功能的kafka, 如果想要部署一个启用KRaft功能的kafka则不需要部署zookeeper container

sudo docker run -d --name zookeeper &#92; \
--network zookeeper-kafka &#92; \
-e ALLOW_ANONYMOUS_LOGIN=yes &#92; \
bitnami/zookeeper:latest

#### PS: 这里的 -e ALLOW_ANONYMOUS_LOGIN=yes 
向container内注入环境变量 ALLOW_ANONYMOUS_LOGIN 并设值为 yes, 目的是允许匿名客户端绕过认证直接连接, 这仅仅是便于快速开发或测试场景，但不建议在生产环境开启, 本教程这样设置是以个人开发作为主要应用场景, 如果是用作正常的产线使用则需要按照个人需求更改配置.

#### 在部署完成zookeeper container之后可以通过以下指令来验证zookeeper服务是否已经启用

sudo docker logs -f zookeeper

如果执行后没有发现有报错则说明已经启用成功

### 四. 部署kakfa container
#### 再次声明!!! 本教程是针对不启用KRaft模式的kafka
#### 由于kafka的更新迭代, 在 3.7 及以上版本的kafka都会默认启用KRaft模式, 而且无法关闭, 但 3.4 系列仍保留对 zookeeper 的完整支持, 因此这里我们需要部署 3.4 版本的kafka container才能正常的运行

sudo docker run -d --name kafka &#92; \
--network zookeeper-kakfa &#92; \
-p 9092:9092 &#92; \
-e ALLOW_PLAINTEXT_LISTENER=yes &#92; \
-e KAFKA_ENABLE_KRAFT=no &#92; \
-e KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181 &#92; \
-e KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://你虚拟机的IP地址:9092 &#92; \
bitnami/kafka:3.4

#### 在部署完成kafka container之后可以通过以下指令来验证kafka服务是否已经启用

sudo docker logs -f kafka

如果执行后没有发现有报错则说明已经启用成功

### 五. 部署kafka-map可视化管理工具(可选)
#### 这个管理工具可以用图形化界面来管理kafka, 让调试, 观测, 操作kafka变得更加的简单, 非常的建议安装

sudo docker run -d --name kafka-map &#92; \
--network zookeeper-kafka &#92; \
-p 9001:8080 &#92; \
-v /opt/kafka-map/data:/usr/local/kafka-map/data &#92; \
-e DEFAULT_USERNAME=你自定义的用户名 &#92; \
-e DEFAULT_PASSWORD=你自定义的密码 &#92; \
--restart always dushixiang/kafka-map:latest


部署完成之后按照以下顺序执行代码, 启用kafka服务:

sudo docker start zookeeper \
sudo docker start kafka \
sudo docker start kafka-map 

之后你可以在你的本地通过

http://你的虚拟机IP:9001/ 

来访问kafka-map进行kafka的图形化管理
PS: 访问kafka-map的时候记得关闭科学上网工具, 不然可能会出现访问不到的情况
