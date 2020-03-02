## Union File System

### Overview

UnionFS implements a **union mount** for other file systems. It allows files and directories of separate file systems, known as branches, to be transparently overlaid, forming a single coherent file system. Contents of directories which have the same path within the merged branches **will be seen together in a single merged directory**, within the new, virtual filesystem.

### Copy-on-write

https://zhuanlan.zhihu.com/p/47683490

The different branches may be either *read-only* or _read-write_ file systems, so that writes to the virtual, merged copy are **directed** to a specific real file system.

This allows a file system to **appear as writable**, but without actually allowing writes to change the file system, also known as **copy-on-write**.

## AUFS (Advanced Multi-Layered Unification Filesystem)

### Image Layer and AUFS

A docker image consists of a series of **read-only layers**. Image layers are stored in `/var/lib/docker/aufs/diff` of Docker hosts filesystem. While metadata of how to stack those layers is stored in `/var/lib/docker/aufs/layers`.