## About this [Linnen](https://github.com/linnenn)
自己使用的go项目,可以免费使用,复制,不过希望你遵守[MIT](https://opensource.org/licenses/MIT) 协议，让所有人受益.

## 使用的一些组件说明

- go 版本 1.14
- gin 框架
- redis 登录信息暂时存放在redis
- mysql 存储
用户登录,检测使用cookie/session,后面会支持jwt

其他的一些都在go.mod中,可以根据需要自己增删

## 目录说明
- [x] main 文件
项目的执行入口，完成所有的初始化工作，也是项目的执行入口
 1, 初始化配置
 2, 初始化数据库，redis
 3, 初始化路由

- [x] model 模型 
定义结构体，跟随数据库中模型定义的结构
- [x] routes 路由
定义路由，完成路由分组，路由处理
- [x] service 提供服务
提供对模型的封装，完成数据的增删改查的操作
- [x] storage 存储，日志
存储操作记录，完成日志记录
- [x] upload 上传目录
文件上传目录，存储文件
- [x] resource 静态资源
系统的静态资源，模版目录
- [x] config 配置文件
系统配置项目
- [x] api 提供请求接口服务的目录
对外提供接口，处理外部请求



## License
The app is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).