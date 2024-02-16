<h1 align="center">SZSK—每日一言API</h1>
<p align="center">简单的每日一言API程序</p>

<p align="center">
<a href="./README.md">中文介绍</a> |
<a href="./README.en.md">English Description</a> 
</p>

<p align="center"> 
<img src="https://img.shields.io/badge/Author-孙子烧烤-orange.svg" title="Author: 孙子烧烤">
<img src="https://img.shields.io/badge/Go-1.21.6-brightgreen.svg" title="Go" />
<img src="https://img.shields.io/badge/version-v1.1-brightgreen.svg" title="Version: v1.0">
<img src="https://img.shields.io/badge/GPL-3.0-brightgreen.svg" title="GPL-3.0">
<img src="https://gitee.com/szsk/yiyan/badge/star.svg?theme=dark" title="Star Count">  
<img src="https://gitee.com/szsk/yiyan/badge/fork.svg?theme=dark" title="Fork Count">  

<p align="center">
<a href="https://www.sunzishaokao.com/">官方网址</a> 
</p>

<p align="center">源码地址：<a href="https://gitee.com/szsk/yiyan">Gitee</a> | 
<a href="https://github.com/szsk2022/yiyan">GitHub</a>
</p>

#### 介绍
简单的每日一言API  
![](https://www.sunzishaokao.com/wp-content/uploads/2024/01/20240131005421418-C3974D21-A954-4516-8015-4C463337E78E.png)
>查看[Demo](https://sunzishaokao.com/plugin/quote/ "Demo")

#### 使用说明
1. 根据系统，下载我们最新编译的程序[Gitee-Releases](https://gitee.com/szsk/kms/releases "Releases")
2. 根据实际情况编辑config.yaml配置文件
3. ./yiyan即可启动！

请求参数：  
1. `/?lang=cn` 请求中文一言  
2. `/?lang=en` 请求英文一言

>Tips:  
>1. 如果没有请求参数即直接访问，默认请求中文一言  
>2. 程序内英文一言和中文一言各内置100条，您可在Public目录按需添加或删减一言！

如何嵌入到网页？→<a href="./HTML.md">HTML调用SZSK一言API</a> 

#### 编译说明
1. 将源代码克隆到本地  
```
git clone https://gitee.com/szsk/yiyan.git
````
2. cd到项目根目录，编译源代码  
```
go build
```

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

