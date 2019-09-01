# 基于centos的私有云的学习与搭建 

## **前言**
这是中山大学数据科学与计算机学院2019年服务计算的作业项目。所有代码与博客将被上传至github当中。  
项目地址: [https://github.com/StarashZero/ServerComputing](https://github.com/StarashZero/ServerComputing)  
个人主页: [https://starashzero.github.io](https://starashzero.github.io)   
实验要求: [https://pmlpml.github.io/ServiceComputingOnCloud/ex-install-cloud](https://pmlpml.github.io/ServiceComputingOnCloud/ex-install-cloud)  

## **安装VirtualBox**  
1. VirtualBox与git安装过程省略  
2. 创建虚拟机内部虚拟网络  
    点击VirtualBox管理->主机网络管理器，创建一块虚拟网卡，配置如下图：  
    ![](picture/1.png)  
    配置完成后可以在cmd中通过ipconfig指令查看网卡信息  
    ![](picture/2.png)  
## **创建Linux虚拟机**  
1. 下载[Centos](https://www.centos.org/download/), minimalISO即可。  
2. 创建虚拟机，配置如下图：  
![](picture/3.png)  
**注:** 设置网卡信息，第一块网卡必须是 NAT；第二块网卡连接方式： Host-Only，接口就是前面创建的虚拟网卡  
3. 按照提示安装虚拟机:    
**注:** 安装过程如果出现无法使用鼠标的情况，需要将虚拟机显示项中的显卡控制器设置为VBoxVGA, 如下图： 
![](picture/4.png)  
等待虚拟机安装完成  
![](picture/5.png) 
安装完毕后进入虚拟机，来到命令行界面，输入之前设置的用户名和密码   
![](picture/6.png)  
输入sudo nmtui，选择第二项激活网卡  
![](picture/7.png)  
网卡激活以后应当就可以上网了，使用sudo yum install wget，安装wget
![](picture/8.png)  
配置镜像可以暂时略过，除非出现网络问题。  
使用yum update升级内核  
![](picture/9.png)  
退出虚拟机，当前虚拟机作为源本，方便以后创建虚拟机。  
  
    接下来复制一个新的虚拟机，右击之前的虚拟机，选择复制。  
**注**: 选择重新初始化所有网卡的 MAC 地址  
点击链接复制，得到一个新虚拟机
![](picture/12.png)  
打开新虚拟机，输入sudo nmtui，同样激活网卡，同时选择第三个选项修改主机名。  
完成后虚拟机就已经连上网了。  
使用service sshd restart开启虚拟机的ssh功能，则在主机上可以通过ssh -p 2222 YOURNAME@127.0.0.1 连接到虚拟机当中  
![](picture/10.png)  
为了增强虚拟机可用性，可以按照一个GNOME桌面    
    1. 安装桌面 yum groupinstall "GNOME Desktop"
    2. 设置启动目标为桌面 ln -sf /lib/systemd/system/runlevel5.target /etc/systemd/system/default.target
    3. 重启  
  
    安装完成后如下图  
![](picture/11.png)  

## **配置用远程桌面访问你的虚拟机**  
1. 首先需要下载[VirtualBox扩展包](https://download.virtualbox.org/virtualbox/6.0.10/Oracle_VM_VirtualBox_Extension_Pack-6.0.10.vbox-extpack)  
2. 在VirtualBox全局设定->拓展中添加下载好的扩展包  
![](picture/13.png)  
3. 在虚拟机设置->显示->远程桌面中，打开服务器并将端口号设为5000  
4. 通过远程桌面连接访问虚拟机  
![](picture/14.png)
![](picture/15.png)  
**实验完成**