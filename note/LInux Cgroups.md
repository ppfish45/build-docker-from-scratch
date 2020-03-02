## Linux Cgroups

### Motivation （source. Wikipedia)

On a single physical machine, we want to **limit**, **account for** and **isolate** the resource usage (CPU, memory, disk IO, network, etc.) of a collection of processes.

### Features （source. Wikipedia)

+ Resource limiting

    groups can be set to not exceed a configured memory limit, which also includes the file system cache

+ Prioritization

    some groups may get a larger share of CPU utilization or disk I/O throughput

+ Accounting

    measures a group's resource usage, which may be used, for example, for billing purposes

+ Control

    freezing groups of processes, their checkpointing and restarting

### Relations among {cgroup, subsystem, hierarchy}

+ The whole cgroup system is like a forest

+ A **hierarchy** is a tree in the forest

+ Multiple **subsystems** can be attached to the same **hierarchy**

+ One **subsystem** can only be attached to one **hierarchy**

+ Every tree node in **hierarchies** is a **cgroup**

### Relations between {cgroup, process}

+ A **process** can be a member of multiple **cgroups**. But those **cgroups** should belong to different **hierarchies**

+ When a subprocess is forked, it belongs to the same **cgroup** as its parent. However, it can be moved to other **cgroups**
