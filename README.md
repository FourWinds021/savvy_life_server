# savvy_life

# savvy_life

## 项目说明
系统为服务简单框架。

## 项目结构
项目重要目录逻辑
1. admin: 服务管理脚本
2. client: 服务客户端
3. common: 服务基础逻辑封装
4. config: 服务配置文件，及https相关证书文件
5. extern: 外部依赖服务封装
6. logs: 服务日志记录目录
7. middlewares: 服务中间件
8. proto: 服务主要接口及结构体定义，其中运行 go run generate.go 命令自动生成proto代码
9. server: 服务主要接口及业务逻辑封装
10. timer: 服务定时器，涉及标注题目批阅定时任务和标注超时监控定时任务
11. tool: 服务维护相关工具，数据库表更、服务打包部署、标注题目批量生产和日志清理工具等

## 项目打包部署
一、项目打包方式
make clean; make debug # debug bin文件打包命令，执行文件目录 build/
make clean; make release # release正式包打包命令，包目录 build/
二、项目部署
1. savvy_life是软链目录，指向最新执行版本服务；
2. 初次部署时，可手动创建软链；
3. 新版本服务部署基本步骤：
   待更新包 savvy_life_release_vxxx.tar.gz
   3.1 ./tool/deploy.sh -v savvy_life_release_vxxx # 新版本服务更新
   3.2 cd ./tool; ./init.sh -u "mysql user" -p "mysql passwd" -f ./sql/vxxx.sql # DDL, DML sql变更
   3.3 配置文件变更
   3.4 ./savvy_life/admin/savvy_life.sh start # 启动服务
   三、 项目接口在线文档
   配置文件可设置在线接口文档是否打开，以及具体服务地址等