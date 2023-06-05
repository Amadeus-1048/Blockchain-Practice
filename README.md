# Blockchain-Practice
A notebook for learning Blockchain and Hyperledger Fabric.



# 搭建环境

## Docker

在官网安装docker，并配置镜像拉取加速器

![](https://tva1.sinaimg.cn/large/008vxvgGgy1h99a28xxscj31j70u077i.jpg)

 

配置完成点击Apply&Restart，等待docker重启完成

 

## Dependencies

以下命令仅为查看已安装的依赖的版本，安装过程自己搜一搜～

```Bash
# docker
docker --version
docker-compose --version

# go
go version

# node 和 npm
node -v
npm -v

# python
python --version
```

 

## Fabric

在 $gopath/src/github.com/hyperledger 目录下拉取两个仓库

```Bash
# 从git上克隆Hyperledger的一个Demo源码
git clone https://github.com/hyperledger/fabric-samples

# 从git上克隆fabric项目
git clone https://github.com/hyperledger/fabric.git
```

 

先打开如下网址（可能需要翻墙，注意这个是新版脚本，而非旧版，旧版不适用于mac）： [https://bit.ly/2ysbOFE](https://goo.gl/byy2Qj)

在fabric-samples根目录下创建"init.sh"脚本文件（终端执行"vi init.sh"），将网站的内容复制到"init.sh"文件中。

```Bash
# 设置"init.sh"文件最高权限777，执行"./init.sh"命令之前确保已经启动Docker。
vi init.sh
chmod 777 init.sh
./init.sh
```

 

执行了"./init.sh"命令之后，会下载一些镜像文件

![](https://tva1.sinaimg.cn/large/008vxvgGgy1h99a2ypsixj31iw0u042u.jpg)

 

在`fabric-samples`代码库的`test-network`目录中找到启动网络的脚本。 使用以下命令导航至测试网络目录（注意，旧版教程中进入的文件夹可能是first-network，但在新版中已经没有该文件夹了）：

```Bash
cd fabric-samples/test-network
```

 

在此目录中，您可以找到带注释的脚本`network.sh`，该脚本在本地计算机上使用Docker镜像建立Fabric网络。  可以运行`./network.sh -h`以打印脚本帮助文本

![](https://tva1.sinaimg.cn/large/008vxvgGgy1h99a3kjqupj325w0kuafd.jpg)

 

在`test-network`目录中，运行以下命令删除先前运行的所有容器或工程：

```Bash
./network.sh down
```

 

然后通过执行以下命令来启动网络

```Bash
./network.sh up
```

此命令创建一个由两个对等节点和一个排序节点组成的Fabric网络。  如果命令执行成功，您将看到已创建的节点的日志（可以使用 `docker ps` 重复查看）：

![](https://tva1.sinaimg.cn/large/008vxvgGgy1h99a3zp6ecj32f60io47q.jpg)

 

## 参考

[Hyperledger Fabric开发环境搭建（MacOS系统）](https://www.jianshu.com/p/3696da2584ff)

[MacOS下搭建Fabric开发环境_dongZhenSong的博客-CSDN博客](https://blog.csdn.net/dongzhensong/article/details/95500596)

[Mac环境安装Hyperledger Fabric](https://www.jianshu.com/p/a59ff954d3b2)

[官方手册](https://hyperledger-fabric.readthedocs.io/zh_CN/latest/test_network.html)


