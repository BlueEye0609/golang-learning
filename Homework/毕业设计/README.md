微服务架构：
- 服务发现、Admin、Cronjob 等工作可以通过 K8S 去实现
- BFF 可以一定程度上利用 istio 的 gateway 完成

需要将原先的巨石结构拆分成微服务，其中一个功能做成微服务的是 CMDB 的增删改的历史记录。

通过 kafka 及DB本身插件将binlog push 到 kafka 再到noSQL，编写一个 application 利用 elasticsearch 来 query 记录。

promethues 已有的插件可以监测本身 pod 使用情况，如 cpu、memory 等等。对于本身应用的监控，可以编写 application 的代码，通过注入来监测需要的数据，比如有多少的binlog ，多少请求等等

对于查看历史记录，我们需要 1. UI 2. ORM 3. API 。 API 到 ORM 可以使用 grpc http2 实现连接复用

