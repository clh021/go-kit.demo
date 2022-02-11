# go-kit-demo

## 代码结构

```
├─common # 公共目录
│  ├─errors
│  └─utils # 封装的通用工具类
│      └─consul
├─user
│  ├─conf # 配置定义
│  ├─dao # 数据库增删改
│  ├─endpoint # 集成中间件（go-kit）
│  ├─global 
│  ├─initialize # 初始化定义
│  ├─model # 数据模型层，定义入参、出参
│  ├─pb # proto 文件
│  ├─server # 启动、路由
│  │  ├─grpc
│  │  └─http
│  ├─service # 编写主要逻辑（go-kit）
│  ├─test
│  └─transport # 网络协议层，转码编码（go-kit）
│      ├─grpc
│      └─http
└─vendor # go 包
```

## 简介

基于 go-kit 实现的同时支持 http 和 grpc 的微服务。

按业务模块拆分，每个业务模块拆分一个目录；模块内含go-kit 框架 transport、endpoint、service 三层模型，和开发常用的目录。

common 为公共目录，目前封装了 consul，可被各个业务模块调用，后续会持续集成。



