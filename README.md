Jenkins Docker permission denied error

There’s an easy fix to the “Permission denied while trying to connect to the Docker daemon socket” error you encounter when you run a Jenkins build or a pipeline’s Jenkinsfile that accesses a Docker image. It’s just a single terminal command and then a reboot:

*********************************************************************************************************************************************
Summery:
> sudo usermod -a -G docker jenkins OR > sudo usermod -a -G docker leader[$USER]

> sudo chmod 666 /var/run/docker.sock
> ls -l /var/run/docker.sock
srw-rw-rw- 1 root docker 0 Mar 13 06:05 /var/run/docker.sock
*********************************************************************************************************************************************

https://www.baeldung.com/linux/docker-permission-denied-daemon-socket-error
Run this command and Jenkins will be able to invoke a Docker run command and the Docker daemon socket issues will go away.

1. Introduction

Owing to their numerous benefits, containers have become the IT industry’s buzzword regarding technology. The Docker platform is the most well-known and commonly used container platform, with a sizable developer community.

The permission error is the most common issue dealt with in the context of software while accessing it.

In this tutorial, we’ll explore the causes of this error and provide solutions to fix it, including using sudo and adding users to the Docker group.

Without any further ado, let’s get into the nitty-gritty details of it.
2. Troubleshooting the Docker Socket Connection Issue

Docker Engine is a complete container runtime environment that includes the Docker CLI, Docker API, and Docker daemon. Further, the Docker daemon is a background process that runs on a host machine; it receives the request from the Docker CLI or client and connects to the Docker API for managing Docker containers, images, volumes, networks, and other Docker objects.

Also, it uses the Unix Network Socket for communicating between the Docker client and daemon:

$ docker run ubuntu:latest /bin/bash
docker: Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock: Post "http://%2Fvar%2Frun%2Fdocker.sock/v1.24/containers/create": dial unix /var/run/docker.sock: connect: permission denied.
See 'docker run --help'.
$

Clients always experience this “Permission denied while trying to connect to the Docker daemon socket” error whenever the user doesn’t have the appropriate access rights to the Docker daemon socket.

Next, let’s see how we can fix this problem.
3. Changing the Permission of /var/run/docker.sock to ‘666‘ Is Dangerous

By default, only root and the docker user-group members can access the Docker socket:

$ ls -l /var/run/docker.sock
srw-rw---- 1 root docker 0 May 13 06:05 /var/run/docker.sock

Some of us may come up with a straightforward “solution” – using the chmod command to modify the permissions of the sock file so that any user can read and write to the socket:

$ sudo chmod 666 /var/run/docker.sock

$ ls -l /var/run/docker.sock
srw-rw-rw- 1 root docker 0 Mar 13 06:05 /var/run/docker.sock

After restarting the Docker service, any user can manage Docker images and start containers. This indeed evades the “permission denied” issue:

$ systemctl restart docker.service

$ docker run -it ubuntu:latest /bin/bash
root@fa679b3875fe:/#

We may think it’s convenient for users to start a Docker container. However, this permission change is potentially dangerous. An example can explain it quickly.

Let’s say we have a user baeldunguser in our system. Now, let’s start the Ubuntu container and mount the host machine’s root (/) directory as the /host volume:

baeldunguser$ docker run -it -v /:/host ubuntu:latest /bin/bash
root@fa679b3875fe:/#

As the output above shows, we’ve entered the container’s shell successfully. Now, let’s make some changes to the /host directory, for example, creating a file:

root@fa679b3875fe:/# touch /host/changed-by-docker-container.hello

Then, if we check the host machine, we can see the file is created in the root (/) directory:

baeldunguser$ ls -l /changed-by-docker-container.hello
-rw-r--r-- 1 root root 0 Jun  3 20:38 /changed-by-docker-container.hello

However, only the root user is allowed to write to the root (/) directory:

baeldunguser$ ls -ld /
drwxr-xr-x 19 root root 4096 Jun  3 20:38 /

Of course, we don’t want to break the system on the host machine. So, we created a file as an example. But we’ve come to the realization that the baeldunguser user has the potential to cause significant damage to the host system through a Docker container.

Therefore, we shouldn’t allow anyone to read and write to the Docker socket daemon.

Next, let’s see the proper solutions to the permission problem.
4. Resolving Permission Errors With Docker and User Privileges

There are multiple ways to resolve this permission issue. Let’s start with a temporary solution to solve this problem. As a quick fix, we can use the sudo (superuser do) command in Linux to elevate the privileges or permissions of the user. It will further help us run the Docker commands without any issues:

$ sudo docker run -it ubuntu:latest /bin/bash
[sudo] password for baeldunguser:
root@5a909c4c6138:/#

Now, follow the steps below to permanently resolve the permission issue on accessing the Docker daemon socket.

First, check whether the Docker user group is available on the host machine. If not, we can add it using groupadd command:

$ sudo groupadd docker

$ getent group docker
docker:x:999:

$ awk -F':' '/docker/{print $4}' /etc/group

Next, let’s add the user to the Docker group, which grants the necessary access rights to the Docker daemon socket:

$ sudo usermod -aG docker baeldunguser

To make the changes effective, we must log out of the previous session and re-enter the new session after adding the user to the docker group:

$ getent group docker
docker:x:999:baeldunguser

$ awk -F':' '/docker/{print $4}' /etc/group
baeldunguser

Lastly, apply the configuration changes by restarting the Docker daemon service using the systemctl restart docker.service command:

$ sudo systemctl restart docker.service
$ sudo systemctl status docker.service
● docker.service - Docker Application Container Engine
Loaded: loaded (/lib/systemd/system/docker.service; enabled; vendor preset: enabled)
Active: active (running) since Mon 2023-03-13 07:23:16 IST; 2s ago
Docs: https://docs.docker.com
Main PID: 10196 (dockerd)
Tasks: 13
CGroup: /system.slice/docker.service
└─10196 /usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock

Now, we’ve successfully executed the Docker commands from the baeldunguser userspace with appropriate privileges:

$ docker run -it ubuntu:latest /bin/bash
root@fa679b3875fe:/#

5. Conclusion

In this article, we’ve explored how to temporarily fix the Docker permission issues using the sudo command and permanently add the user to the appropriate group.

Overall, by adhering to best deployment practices, users can avoid such common errors and ensure the smooth operation of their Docker environment.