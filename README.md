# 角色认证管理系统
基于以下技术栈：

- web框架：gin
- orm：gorm
- 日志: logrus
- 角色认证：casbin
- 热加载：fresh
- 配置解析：viper

## 代码热加载

基于fresh
使用：在项目的根目录下使用fresh命令即可

## 项目结构
- api：提供的api接口
- common：通用的基础结构
  - global: 全局变量，如：db、jwt、log
  - model：通用的结构，如：回复类、错误码、配置类
  - util：工具包
- config：配置文件
- docs：api文档、sql
- middleware：gin中间件
- model：实体类
- router：gin配置路由
- service：服务类