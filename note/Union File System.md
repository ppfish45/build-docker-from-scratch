## Union File System

### Overview

UnionFS implements a **union mount** for other file systems. It allows files and directories of separate file systems, known as branches, to be transparently overlaid, forming a single coherent file system. Contents of directories which have the same path within the merged branches **will be seen together in a single merged directory**, within the new, virtual filesystem.

### Copy-on-write

https://zhuanlan.zhihu.com/p/47683490

The different branches may be either *read-only* or *read-write* file systems, so that writes to the virtual, merged copy are **directed** to a specific real file system.

This allows a file system to **appear as writable**, but without actually allowing writes to change the file system, also known as **copy-on-write**.

## AUFS (Advanced Multi-Layered Unification Filesystem)

### Image Layer and AUFS

A docker image consists of a series of **read-only layers**. Image layers are stored in `/var/lib/docker/aufs/diff` of Docker hosts filesystem. While metadata of how to stack those layers is stored in `/var/lib/docker/aufs/layers`.

All the changes in a running container will be applied on *read-write* layer of AUFS, which will not affect the base layers.

### Example

Suppose now we have such file structure

```
.
├── container-layer
│   └── container-layer.txt
├── image-layer1
│   └── image-layer1.txt
├── image-layer2
│   └── image-layer2.txt
├── image-layer3
│   └── image-layer3.txt
├── image-layer4
│   └── image-layer4.txt
└── mnt
```

and we want to mount `container-layer` as well as `image-layer${n}` to location `mnt` in a AUFS mannar. We run the following command:

```
sudo mount -t aufs -o dirs=./container-layer:./image-layer1:./image-layer2:./image-layer3:./image-layer4 none ./mnt
```

By default, the first directory after `dirs` argument is *read-write* and the remaining is *read-only*. Run `cat /sys/fs/aufs/si_60321d6a3564259d/*`, we can check this conclusion

```
# cat /sys/fs/aufs/si_60321d6a3564259d/*
/home/ppfish/Repos/build-docker-from-scratch/aufs/container-layer=rw
/home/ppfish/Repos/build-docker-from-scratch/aufs/image-layer1=ro
/home/ppfish/Repos/build-docker-from-scratch/aufs/image-layer2=ro
/home/ppfish/Repos/build-docker-from-scratch/aufs/image-layer3=ro
/home/ppfish/Repos/build-docker-from-scratch/aufs/image-layer4=ro
64
65
66
67
68
/home/ppfish/Repos/build-docker-from-scratch/aufs/container-layer/.aufs.xino
```

Let's try out Copy-on-write mechanism. We append some text to `mnt/image-layer1.txt`.

```
# cat image-layer1.txt
123I am image layer 1
```

and then check out where the changes goes.

```
tree
.
├── container-layer
│   ├── container-layer.txt
│   └── image-layer1.txt
├── image-layer1
│   └── image-layer1.txt
├── image-layer2
│   └── image-layer2.txt
├── image-layer3
│   └── image-layer3.txt
├── image-layer4
│   └── image-layer4.txt
└── mnt
    ├── container-layer.txt
    ├── image-layer1.txt
    ├── image-layer2.txt
    ├── image-layer3.txt
    └── image-layer4.txt
```

A copy of `image-layer1.txt` with changes are created in `container-layer` folder, namely, the only *read-write* folder. This is because `image-layer1` folder is *read-only*, any changes in AUFS will be redirected to the *read-write* folder.

**A question is, how to deal with deletions in `mnt`?**

The answer is intuitive. Every time when `file1` is deleted in `mnt`, a `.wh.file1` will be generated in the *read-write* layer, namely, `container-layer` in this example to hide all `file1` in *read-only* layers.