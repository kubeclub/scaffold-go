![github](https://img.shields.io/badge/scaffold-golang-brightgreen)
1. 工程结构
    - 模块划分/服务划分：为服务扩展性服务
    - 包结构划分：影响到代码的可阅读性、代码逻辑的高内聚低耦合等
2. 基础功能
    - 完善的日志框架
    - 配置读取
    - 鉴权
    - 基础工具
    - 中间件调用封装
      - http-restful
      - [http-graphgl](./docs/graphql.md)
      - mysql
      - redis
      - es
    - 单元测试

# golang 脚手架说明

仓库目录放了2个独立的模块，后续可以根据需求拆解到不同 git 仓库，在项目前期放一起可以方便修改。

```
// 将引用改成本地目录
replace exexm.com/golang_common => ./../golang_common //当前出于修改方便把 common 工程放一起，后续需要多服务共有可以独立出去
```
## 技术栈
- gin
- viper
- gorm
- jwt
- graphql
- ...

## golang_common
存放通用到代码，像操作各个中间件还有业务无关的各项基础服务。

```
├── conf     配置文件夹
├── lib      日志、配置、redis操作、es 操作、mysql 操作
├── test     测试代码
```
## golang_web
```
├── README.md
├── conf            配置文件夹
│   └── dev
│       ├── base.toml
│       ├── mysql_map.toml
│       └── redis_map.toml
├── controller      控制器
│   └── demo.go
├── dao             DB数据层
│   └── demo.go
├── docs            swagger文件层
├── dto             输入输出结构层
│   └── demo.go
├── go.mod
├── go.sum
├── main.go         入口文件
├── middleware      中间件层
│   ├── panic.go
│   ├── response.go
│   ├── token_auth.go
│   └── translation.go
├── public          公共文件
│   ├── log.go
│   ├── mysql.go
│   └── validate.go
└── router          路由层
│   ├── httpserver.go
│   └── route.go
└── services        逻辑处理层
```

## 快速启用
1. git clone 代码
2. 准备好 golang 1.16 环境,开启 GO111MODULE="on" ,必要时也设置下 GOPROXY=https://goproxy.cn , 我这边用的是 goland idea 工具
3. 2个独立的模块分别执行 go mod tidy
4. 初始化好测试需要的数据库
```
CREATE TABLE `area` (
 `id` bigint(20) NOT NULL AUTO_INCREMENT,
 `area_name` varchar(255) NOT NULL,
 `city_id` int(11) NOT NULL,
 `user_id` int(11) NOT NULL,
 `update_at` datetime NOT NULL,
 `create_at` datetime NOT NULL,
 `delete_at` datetime NOT NULL,
 PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='area';
INSERT INTO `area` (`id`, `area_name`, `city_id`, `user_id`, `update_at`, `create_at`, `delete_at`) VALUES (NULL, 'area_name', '1', '2', '2019-06-15 00:00:00', '2019-06-15 00:00:00', '2019-06-15 00:00:00');
```

5. 启动 golang_web 服务 go run main.go -config=./conf/exe-dev/

测试：http://127.0.0.1:8880/demo/dao?id=1