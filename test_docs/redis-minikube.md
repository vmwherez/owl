# Redis Enterprise and `minikube` for Kubernetes Development

### Learn about the basics steps needed to setup Minikube and run a three-node Redis Enterprise cluster on your local laptop.

by [Vick Kelkar](https://dzone.com/users/3190689/vickkelkar.html)

Kubernetes is an open-source container orchestration system used to deploy, scale, and manage containerized applications. Kubernetes is a project hosted by the Cloud Native Computing Foundation ([CNCF](https://www.cncf.io/)). At a very high level, it contains two types of resources: a master node (which is the cluster coordinator) and nodes, which are the workers that run containerized applications.

Minikube is a tool used to run a Kubernetes cluster on a local machine. Minikube is a single-node Kubernetes cluster inside a VM on your laptop. Minikube can be used to try out Kubernetes and or develop with it day-to-day.

The [Redis Enterprise](https://redislabs.com/) offering extends Redis, the most popular database used with Docker containers. Redis Enterprise delivers high performance, low latency, and high availability to organizations. This blog post will show you the basics steps needed to setup Minikube and run a three-node Redis Enterprise cluster on your local laptop.)

## Installation

We are going to use the [homebrew](https://brew.sh/) package manager to install Minikube and the Kubernetes command line tool on your local laptop.

```
$ brew cask install minikube
$ brew install kubernetes-cli
```

## Starting Minikube ![img](https://redislabs.com/wp-content/uploads/2018/01/20171213_minikube.png)

Minikube offers the ability to change the Virtual Machine (VM) driver. For this blog post, we used the vmwarefusion driver.

```
$ minikube start --memory 12288 --disk-size 30g --vm-driver vmwarefusion
```

The output of the command should look like:

![Image title](https://redislabs.com/wp-content/uploads/2018/01/20171212_minikube_startup2.png)

## Verify Minikube installation

We will use the Kubernetes command-line tool `kubectl` to verify the Minikube installation. You can verify the Minikube install using:

```
$ kubectl cluster-info
```

The output should look like:

```
Kubernetes master is running at https://192.168.21.147:8443
```

## Redis Enterprise on Minikube

### Installation

Minikube is now running on your laptop and `kubectl` cli was able to successfully query the master node to get the status of the cluster. We will now deploy the Redis Enterprise service in Minikube with three replica sets. The YAML configuration for the Redis Enterprise deployment and service can be found [here](https://gist.github.com/v-kelkar/6dd938ec8aa9e26cb39edcb7420683e9).

We will use the YAML to create the deployment and service in the Kubernetes cluster.

```
$ kubectl apply -f redis-enterprise5.yaml
deployment "redis5" created
service "redis5" created
```

We can verify that three pods were created by issuing the command:

```
Kubectl get pods -o wide
```

### Prerequisite for Cluster Configuration

We need to change the binding of the CCS from local loopback address 127.0.0.1 to 0.0.0.0.

```
$ kubectl exec -it <name-of-your-pod> bash
$ sed 's/bind 127.0.0.1/bind 0.0.0.0/g' -i /opt/redislabs/config/ccs-redis.conf
$ /opt/redislabs/bin/supervisorctl restart ccs
```

### Configuration of Cluster

We need to setup the Redis Enterprise cluster as a master node by logging into the pod and issuing a create cluster command.

```
$ kubectl exec -it <name-of-your-pod> bash
```

We will use the rladmin utility to create the new cluster:

```
root@redis5-7749c97f4d-vk6rn:~# /opt/redislabs/bin/rladmin cluster create name cluster.local username <your_email_addr> password <your_passwd>
```

We will use the rladmin utility to join the two nodes to cluster:

```
root@redis5-7749c97f4d-hcggk:~#/opt/redislabs/bin/rladmin cluster join username <your_email_addr> password <master_password> nodes <master_IP_addr>
```

And:

```
root@redis5-7749c97f4d-p74nx:~#/opt/redislabs/bin/rladmin cluster join username <your_email_addr> password <master_password> nodes <master_IP_addr>
```

### Database Creation on Cluster

We will use the REST API of Redis Enterprise to create a database on the master node:

```
curl -k -u "vick@redislabs.com:<password>" --request POST --url "https://localhost:9443/v1/bdbs" --header 'content-type: application/json' --data '{"name":"demo-db","type":"redis","memory_size":536870912,"port":11000,"data_persistence":"aof","replication":true,"shards_placement":"sparse","proxy_policy":"all-master-shards"}'
```

We can look at the status of the cluster and database using the rladmin utility included in Redis Enterprise download:

- You can read more about Redis Enterprise's shard placement policy and proxy policy on our [documentation page](https://redislabs.com/redis-enterprise-documentation/administering/designing-production/networking/multiple-active-proxy/).
- Redis Enterprise includes replication and persistence by default. You can read about persistence on our Redis Enterprise [database documentation page](https://redislabs.com/redis-enterprise-documentation/concepts-architecture/data-access/persistence/).

### Redis Enterprise Dashboard Access

You can access the Redis Enterprise dashboard, which is running on port 8443, by setting up a secure tunnel between local port and the pod port. Once the tunnel is established, you can reach the Redis Enterprise dashboard at https://127.0.0.1:8443:

```
kubectl port-forward <my-pod-name> <localport>:<pod-port>

Example: kubectl port-forward redis5-58dc568c56-7qk22 8443:8443
```

### Local Access to Redis Enterprise

You can connect to your database using the IP of the node and the port specified during database creation:

```
$ redis-cli -h <ip_addr_of master_node> -p <11000>
```

## What's Next

We are working on a Kubernetes native Redis Enterprise container that will take advantage of the new primitives introduced in Kubernetes 1.8 and above. We are working on releasing a new version of the Redis Enterprise container image that will leverage both the new Persistent Sets and the Storage class primitives while providing a better cluster bootstrapping experience. In the meantime, [learn more about the Redis Enterprise](https://redislabs.com/why-redis/redis-enterprise/) offering.