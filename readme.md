# Hotel Bookings 

你好，欢迎来到我写的的酒店预定管理系统服务端代码仓库。 👏👏👏 


## 环境安装

开发环境基于 Mac 或 Linux。

1. 执行 `make docker/up` 命令构建环境

2. 使用DBMS工具（如：navicat、SQLyog、DBeaver...）连接到mysql，设置变数据库编码为 utf8mb4、utf8mb4_general_ci, 或者执行下面SQL也行。
```sql
ALTER DATABASE `db_hotel_bookings` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;
```

3. 执行 `make migrate/up` 命令迁移所有SQL。

更多 `make` 命令参考 `Makefile` 命令定义。

### 启动服务

启动后台api服务命令: `make back/start`

启动前台api服务命令: `make front/start`

启动前端用户端命令(本地 Node:v16.20.2): 安装依赖 `cd webapps/hotel-app && npm i`, 启动项目 `npm run dev`

启动前端管理端命令(本地 Node:v16.20.2): 安装依赖 `cd webapps/hotel-mgt && npm i`, 启动项目 `npm run dev`


## 项目介绍

这部分主要介绍项目的结构和业务的描述。

### 后台 API 项目层次设计

项目核心分为4层：DB层(MySQL、Redis)、Service 层、Router 路由层、API 层 。

层与层调用链：router -> api -> service -> db。

错误处理：db -> service -> api -> router。其中db、service 层直接将错误抛出，不做任何处理，错误最终在api层记录。

系统日志：~~因为错误在 api 层处理的，所以日志也是由 api 打印，也足够调试，系统保持相对干净。~~ 
错误由service层处理，api 层使用并打印日志。

入参设计：service 层入参和 api 层入参使用共同的结构体，独属于service层的参数取消 json tag，避免参数定义过多。

### 前端用户端项目介绍

/webapps/hotel-app

技术栈：Vite、React、TypeScript、TailwindCSS


### 前端后台管理端项目介绍

/webapps/hotel-mgt

技术栈：Vite、Vue3、TypeScript、Element Plus