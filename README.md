# RobKing_Goroutine_ChatRoom
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

## 项目介绍

基于golang高并发特性开发的聊天室。

项目是如何设计与实现的？

### 客户端

客户端启动的时候进行初始化，最重要的就是两个管道，一个管道用于接受客户端请求信息，一个管道用于接受服务器发送的信息，同时开启三个协程，主协程用于处理用户的登录和注册，一个子协程用于接受连接中的数据，另一个子协程用于通过连接发送数据，以下是具体实现细节：

- 注册

  用户输入用户名和密码，通过子协程发送登录请求到 客户端接受信息的管道，接着通过子协程发送登录请求到服务器，服务器接收请求验证成功就会发送“success”，客户端接收协程收到之后注册成功，进行用户登录

- 登录

  还是一样的输入用户名和密码，通过子协程发送请求，服务器验证成功之后登录成功，之后继续开启两个协程，一个协程用于接受服务器的信息并显示在终端，另一个协程用于客户端信息的输入，可以直接发送消息，也可以通过to发送私聊消息，发送quit可以退出聊天室，发送clientlist可以查看当前在线用户。

- 接受服务器信息的子协程

  将服务器发送的信息放入到管道中

- 向服务器发送信息的子协程

  从客户端输入信息的管道中拿出数据发送到服务器

- 消息显示协程

  从管道中拿出数据，打印到终端

- 信息输入的协程

  将输入的信息放入到管道，通过发送信息的子协程发送到服务器

### 服务端

服务端启动的时候进行初始化，信息管道，还有一个哈希表用于缓存所有连接对象的信息，开启两个协程，子协程用于广播消息，主协程用于持续监听等待客户端的连接，当用客户端连接的时候，首先对这个客户端进行初始化，接发信息的管道，连接对象，同时开启三个协程，一个协程处理用户登录注册的认证，一个协程用于向该客户端发送信息，还有一个用于接受该用户的信息，下面是具体实现细节：

- 处理用户登录注册的协程

  注册主要是判断数据库中是否存在相同的用户名，没有的话发送“success”到该客户端 发送信息的管道，通过子协程发送给该客户端。登录主要是判断用户名和密码是否匹配，成功还是一样的处理，向所有的连接对象发送该用户进入聊天室的提示信息，同时将该连接对象放入到哈希表中，开启接受该用户信息的协程

- 接收该用户信息的协程

  通过从接受信息的管道中拿取数据，并将数据放入到服务器的信息管道中进行处理

- 向该客户端发送信息的协程

  通过将管道中的信息发送给该客户端

- 接收该用户的信息协程

  主要是通过连接接收信息，放入到接收信息的管道

- 广播信息的协程

  将服务器管道中的信息取出，根据信息内容不同的处理，如果是quit，从哈希表中删除该用户，如果是userlist，遍历哈希表，将所有在线用户名发送给该客户端，如果有to标志，将信息发送给指定的用户，普通信息则遍历所有的连接对象，将信息放入到每个客户端的管道中然后发送。

## **安装步骤**

1. Get a free API Key at [https://example.com](https://example.com)
2. Clone the repo

```sh
git clone https://github.com/RobKing9/RobKing_Goroutine_ChatRoom.git
```

## 文件目录说明

```
.
|-- README.md
|-- client
|   |-- client.go	主程序，初始客户端，连接服务器
|   |-- data.go	接收服务器信息放入管道，拿取管道信息发送给服务器
|   |-- login.go	用户登录
|   |-- message.go	从管道中拿出信息显示到终端，用户发送数据到管道中
|   `-- register.go	用户注册
|-- server
|   |-- broadcast.go	广播消息
|   |-- clientAuth.go	用户认证
|   |-- data.go	数据接收和发送
|   |-- mysql.go	连接数据库及相关的操作
|   `-- server.go	服务器主协程，等待客户端连接

```


## 部署

通过docker部署在服务器中，用户可以通过不同的网络连接交流


## 贡献者

请阅读**CONTRIBUTING.md** 查阅为该项目做出贡献的开发者。

## 如何参与开源项目

贡献使开源社区成为一个学习、激励和创造的绝佳场所。你所作的任何贡献都是**非常感谢**的。


1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



## 版本控制

该项目使用Git进行版本管理。您可以在repository参看当前可用版本。

## 作者

博客

 *您也可以在贡献者名单中参看所有参与该项目的开发者。*

## 版权说明

该项目签署了MIT 授权许可，详情请参阅 [LICENSE.txt](https://github.com/RobKing9/RobKing_Goroutine_ChatRoom/blob/master/LICENSE.txt)

## 鸣谢


- [GitHub Emoji Cheat Sheet](https://www.webpagefx.com/tools/emoji-cheat-sheet)
- [Img Shields](https://shields.io)
- [Choose an Open Source License](https://choosealicense.com)
- [GitHub Pages](https://pages.github.com)
- [Animate.css](https://daneden.github.io/animate.css)

<!-- links -->

[your-project-path]:RobKing9/RobKing_Goroutine_ChatRoom
[contributors-shield]: https://img.shields.io/github/contributors/RobKing9/RobKing_Goroutine_ChatRoom.svg?style=flat-square
[contributors-url]: https://github.com/RobKing9/RobKing_Goroutine_ChatRoom/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/RobKing9/RobKing_Goroutine_ChatRoom.svg?style=flat-square
[forks-url]: https://github.com/RobKing9/RobKing_Goroutine_ChatRoom/network/members
[stars-shield]: https://img.shields.io/github/stars/RobKing9/RobKing_Goroutine_ChatRoom.svg?style=flat-square
[stars-url]: https://github.com/RobKing9/RobKing_Goroutine_ChatRoom/stargazers
[issues-shield]: https://img.shields.io/github/issues/RobKing9/RobKing_Goroutine_ChatRoom.svg?style=flat-square
[issues-url]: https://img.shields.io/github/issues/RobKing9/RobKing_Goroutine_ChatRoom.svg
[license-shield]: https://img.shields.io/github/license/RobKing9/RobKing_Goroutine_ChatRoom.svg?style=flat-square
[license-url]: https://github.com/RobKing9/RobKing_Goroutine_ChatRoom/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/shaojintian
