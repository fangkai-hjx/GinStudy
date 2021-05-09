### Go语言的学习小项目
> 主要功能
> - 统一报错信息
> - 统一rest输出
> - validator
> - 日志功能
> - 安全
> - 接口文档
### 项目结构
--Project
-----cache

-----config

-----controller:控制器，负责接收参数、验证参数，调用service,统一输出

-----dao:到数据库的访问

-----doc:文档

-----global: 全局用到的变量,主要是配置、数据库连接、日志功能等

-----model:数据模型

-----pkg

-----request:

-----router:路由

-----service:主要的业务逻辑

main.go[入口函数]

------
