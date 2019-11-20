# **简单博客网站API设计**  

## **前言**
这是中山大学数据科学与计算机学院2019年服务计算的作业项目。所有代码与博客将被上传至github当中。  
Github项目地址: [https://github.com/StarashZero/ServerComputing/tree/master/hw6](https://github.com/StarashZero/ServerComputing/tree/master/hw6)  
个人主页: [https://starashzero.github.io](https://starashzero.github.io)   

## **作业简介**  
模仿 [Github](https://developer.github.com/v3/)，设计一个博客网站的 API  

## **API**  
博客网站使用HTTPS协议，所有请求根地址: ```https://spblog.com```，数据默认以JSON形式传递  

### Current version  
网站可以存在多个版本，默认请求最新版本，可以自行选择其他版本  
```Accept: application/vnd.spBlog.v3+json```  

### Schema  
所有API通过HTTPS访问，数据以JSON形式的形式发送和接收  
```js  
curl -i https://spblog.com/users/articles/firstblog
HTTP/1.1 200 OK
Server: nginx
Date: Fri, 12 Oct 2012 23:33:14 GMT
Content-Type: application/json; charset=utf-8
Connection: keep-alive
Status: 200 OK
ETag: "a00049ba79152d03380c34652f2cb612"
X-SpBlog-Media-Type: spblog.v3
X-RateLimit-Limit: 5000
X-RateLimit-Remaining: 4987
X-RateLimit-Reset: 1350085394
Content-Length: 5
Cache-Control: max-age=0, private, must-revalidate
X-Content-Type-Options: nosniff
```   

### Authentication  
对博客进行上传、修改需要提供账户信息，以及部分私密博客需要认证才能获得  
* Basic authentication  
    这里提供一个基本认证方式  
    ```curl -u "username" https://spblog.com```  
    确认用户存在后会要求继续输入密码，密码正确则完成认证，获得对应权限(对该账户文章修改、上传以及访问私密文章等)，否则返回```404 Not Found```  

### Articles list  
可以查看用户的所有已发布博客，响应中还会附带每篇博客的概述(若为指定则默认文章第一段)  
```GET /user/articles```  
可包含的参数如下:   
  
| | |  
|---|---|---|  
|private|是否包含私有博客，若为true则需认证|   
|max_length|选取博客的最大数量(时间排序)，不设置或为null则表示无限制|    
示例:  
curl -i -uuser https://spblog.com/user/articles?private=true&max_length=10  
表示查看最近发布的10篇博客(包含私密博客)  
### Get article  
可以请求获得单篇博客，响应中会附带该博客的详细信息  
```GET /user/articles/firstBlog/```  
可附带private参数，使用方式与Articles list一致  
### Search article  
可以搜索具有某种特征的博客，响应中会附带所有满足条件的博客简要信息  
```GET /user/search/```  
可附带的参数如下:  
|||  
-|-|-|  
private|是否包含私有博客，若为true则需认证|  
max_length|选取博客的最大数量(时间排序)，不设置或为null则表示无限制|  
keyword|用于查找的关键词|  
in_content|是否搜索博客内容，若为false则只搜索标题|  
实例:  
```curl -uuser -i https://spblog.com/user/search?keyword=api&private=true&in_content=true```  
表示搜索包含"api"的所有博客  
### Upload article  
上传(发布)博客，用户必须已被认证通过  
```PUT /user/upload```  
可被包含的参数如下  
|||
-|-|-|  
private|是否为私密文章|  
title|文章标题|  
summary|文章概述，空置则默认为文章第一段|  
content|文章内容|  
可使用JSON格式上传较长的参数，示例:    
```js
curl -uuser -i -d '
{
    "title":"FirstBlog", 
    "summary":"summary of first blog",  
    "content":"some words"
}
'  
https://spblog.com/user/upload?private=true
```  

### Update article  
可以对已发布的文章进行更新，使用方式与upload相似，但是update必须提供已存在博客的ID，其他包含的参数会覆盖原有的参数，若不附带则不变  
```PUT /user/upload```  
可被包含的参数如下  
|||
-|-|-|  
id|文章id|
private|是否为私密文章|  
title|文章标题|  
summary|文章概述，空置则默认为文章第一段|  
content|文章内容|  
可使用JSON格式上传较长的参数，示例:    
```js
curl -uuser -i 
https://spblog.com/user/update?private=false&id=1
```  
意思是将id为1的博客修改为公开  

### Delete article  
删除博客需要用户已认证通过，并提供已发布文章的id号    
```Delete /user/delete```  
示例:  
```js
curl -uuser -i 
https://spblog.com/user/delete?id=1
```  
意为将id为1的博客删除  