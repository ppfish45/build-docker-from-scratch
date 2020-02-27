## Linux Namespace

### An example

A company has a great server and it wants to sell tomcat instances running on this server. But each user needs to have `root` privilege to execute certain commands. It is hard to achieve such an isolation status.

### Solution

`Linux Namespace` makes it possible to partion the kernal resources such that one set of processes only sees one set of corresponding resources. A user with UID of `n` will have `root` privilege in its own namespace, while on the real machine, it is only the user with UID of `n`.

PID (Process ID) can be virtualized as well. The PID should be numbered as a normal Linux machine in every single namespace - an `init` process with PID of 1, and others come in an ascending order.

### Categories of Namespaces

+ Mount Namespace

    `mount` is used to establish mappings between physical storage and file system. For example, to mount an external storage, we can use `mount -t vfat /dev/sdd1 /mnt/usb`. After that, by accessing `/mnt/usb` we can access the files we want.

    If we assign a new mount namespace to a process, a copy of all mounts will be copied to this namespace. Any subsequent modifies on mount points will only take effect on this copy instead of the mount namespace of the host. Namely, we can create a brand new (virtual) file system for a process.

    https://www.zhihu.com/question/263885160


+ UTS Namespace

    UTS namespaces allow a single system to appear to have different `hostname` and `NIS domain name` to different processes.

    https://www.cnblogs.com/sparkdev/p/9377072.html

+ IPC Namespace

    To isolate `System V IPC` (IPC - Inter-Process Communication) and `POSIX message queue`.

+ PID Namespace

    One process may possess different `PID` s in different PID namespaces.

    **Why the PID of clone process in `ps -a` is not `1` ?**

    Because results of `ps` depends on content in `/proc`. We should run `mount -t proc proc /proc` to ensure only processes of this new PID namespace will be presented in `/proc`.

+ Network Namespace

+ User Namespace

    To isolate users and groups. Namely, a process may have different user ID and group ID inside and outside the user namespace.

    A non-root user can be given `root` privilege in a user namespace. 

### Syscalls

+ `clone` creates a new process. Its flags is used to specify which **new** namespace it should be migrated to.

+ `unshare` disassociates one process from one namespace

+ `setns` enters a specified namespace
