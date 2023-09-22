# Container From Scratch

## Lets start with creating containers using docker

```sh
docker run -it ubuntu /bin/bash
```

Lets understand the isolation given by containers
- Process isolation
- File System isolation
- HostName isolation
- Network isolation
- User isolation

## Lets build our own contianers from Scratch

### Lets create a process and see isolation provided by default
```sh
/bin/bash
```
> Basically, no isolation

### Lets try to provide isolation using linux features

#### Linux Namespaces
It allows us to create restricted views of systems like the process tree, network interfaces, and mounts and users.
- UTS(Unix Time Sharing) namespace: hostname and domain name
- PID namespace: process number
- Mounts namespace: mount points
- IPC namespace: Inter Process Communication resources
- Network namespace: network resources
- User namespace: User and Group ID numbers

```sh
man unshare
```


```sh
# lets create a process with user isolation
id -u
unshare -U
```

```sh
# lets create a process with uts (unix time-sharing system) isolation
# uts helps to unshare the domain and hostname
unshare -u
```

```sh
# lets create a process with process isolation
ps -aef --forest
unshare -fp --mount-proc 
```

```sh
# lets create a process with network isolation
unshare -n
```

```sh
# lets create a process with mount isolation
unshare -m
mkdir mount-dir
mount -n -o size=10m -t tmpfs tmpfs mount-dir
df mount-dir
```

### File system is still shared, lets bring some isolation there

```sh
# Lets create a new dir and say that is the root of the new process
mkdir tempdir
chroot tempdir/

# it missing the libs
sudo chroot bundle/rootfs
# bundle/rootfs has basic

ps aux # to list of processes
```

### What to do if few process takes more memory and cpu than others?
cgroups is to support resource limiting, prioritization, accounting and controlling.

- memory
- cpu
- network
- disk i/o
- device permissions

```sh
# if it is cgroup -> /sys/fs/cgroup/memory/fridaytechbytes
# if it is cgroupv2 -> /sys/fs/cgroup/fridaytechbytes
mount | grep cgroup # to test cgroup
mkdir -p /sys/fs/cgroup/fridaytechbytes
ls /sys/fs/cgroup/fridaytechbytes
echo 100000000 > /sys/fs/cgroup/fridaytechbytes/memory.high
echo 0 > /sys/fs/cgroup/fridaytechbytes/memory.swap.high
echo $$ > /sys/fs/cgroup/fridaytechbytes/cgroup.procs
```

## Simple golang based program
Refer [./main.go](./main.go)

## Conclusion
Containers are controlled and isolated group of processes



### Refs:
- https://www.youtube.com/watch?v=Utf-A4rODH8
- https://itnext.io/container-from-scratch-348838574160
- https://medium.com/@saschagrunert/demystifying-containers-part-i-kernel-space-2c53d6979504
- https://www.ianlewis.org/en/almighty-pause-container