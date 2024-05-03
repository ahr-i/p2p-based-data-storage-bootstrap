# P2P Based Distributed-data-storage / Bootstrap
It is a distributed-data-storage that can be used in a peer-to-peer (P2P) environment.   
   
Use it in conjunction with distributed-data-storage-webui.   
(https://github.com/ahr-i/p2p-based-data-storage-webui)

## 1. Start
### 1.1 Clone
```
git clone https://github.com/ahr-i/p2p-based-data-storage-bootstrap.git
```

### 1.2 build
```
cd p2p-based-data-storage-bootstrap
docker build -t dds/bootstrap .
```

### 1.3 create network
```
docker network create --subnet 200.0.0.0/16 p2p
```

### 1.4 Run
```
docker run -d --name dds_bootstrap --network p2p --ip 200.0.0.20 -p 3000:3000 dds/bootstrap
```
It's better to set a static IP.   
   
Use it in conjunction with distributed-data-storage-webui.   
(https://github.com/ahr-i/p2p-based-data-storage-webui)
