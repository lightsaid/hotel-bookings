# Hotel Bookings 

你好，欢迎来到我写的的酒店预定管理系统服务端代码仓库。 👏👏👏 


## 环境安装

开发环境基于 Mac 或 Linux。

1. 执行 `make docker/up` 命令构建环境，使用下面命令

2. 使用DBMS工具（如：navicat、SQLyog、DBeaver...）连接到mysql，设置变数据库编码为 utf8mb4、utf8mb4_general_ci, 或者执行下面SQL也行。
```sql
ALTER DATABASE `db_hotel_bookings` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;
```


3. 执行 `make migrate/up` 命令迁移所有SQL。

更多 `make` 命令参考 `Makefile` 命令定义。