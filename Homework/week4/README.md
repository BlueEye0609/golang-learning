# 作业
按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

# 构想
实现 storage 的增删改查以及命令行
- cmd

component:
- Storage
  - hostname
  - size

### ent
https://entgo.io/docs/getting-started/
```bash
# go get entgo.io/ent/cmd/ent
# go run entgo.io/ent/cmd/ent init Storage

>> change ent/schema/storage.go

# go generate ./ent
```

### api
api 的requestURL:
```
/<package_name>.<version>.<service_name>/{method} 
```
package:
```
<package_name>.<version>; 
```

定义了 storage.proto
```
 protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative storage.proto
```

## 目录
https://github.com/golang-standards/project-layout/blob/master/README_zh.md
- /api <br>

API 协议定义目录，xxapi.proto protobuf 文件，以及生成的 go 文件。我们通常把 api 文档直接在 proto 文件中描述

OpenAPI/Swagger 规范，JSON 模式文件，协议定义文件。有关示例，请参见 [/api](https://github.com/golang-standards/project-layout/blob/master/api/README.md) 目录。


- /build<br>

打包和持续集成

- /cmd<br>

负责程序的：启动、关闭、配置初始化等。

项目的主干。每个应用程序的目录名应该与你想要的可自行文件的名称相匹配（例如,`/cmd/myapp`)。

不要在这个目录中放置太多代码。如果你认为代码可以导入并在其他项目中使用，那么它应该位于 `/pkg` 目录中。如果代码不是可重用的，或者你不希望其他人重用它，请将该代码放到 `/internal` 目录中。

通常有一个小的 main 函数，从 `/internal` 和 `/pkg` 目录导入和调用代码，除此之外没有别的东西。

有关示例，请参阅 [/cmd](https://github.com/golang-standards/project-layout/blob/master/cmd/README.md) 目录。

- /configs

配置文件模板或默认配置。

将你的 confd 或 consul-template 模板文件放在这里。
- /examples

你的应用程序和/或公共库的示例。

有关示例，请参见 /examples 目录。

- /internal

私有应用程序和库代码。这是你不希望其他人在其应用程序或库中导入代码。请注意，这个布局模式是由 Go 编译器本身执行的。有关更多细节，请参阅Go 1.4 [release notes](https://golang.org/doc/go1.4#internalpackages) 。注意，你并不局限于顶级 internal 目录。在项目树的任何级别上都可以有多个内部目录。

你可以选择向 internal 包中添加一些额外的结构，以分隔共享和非共享的内部代码。这不是必需的(特别是对于较小的项目)，但是最好有有可视化的线索来显示预期的包的用途。你的实际应用程序代码可以放在 /internal/app 目录下(例如 /internal/app/myapp)，这些应用程序共享的代码可以放在 /internal/pkg 目录下(例如 /internal/pkg/myprivlib)

- /pkg

外部应用程序可以使用的库代码(例如 /pkg/mypubliclib)。其他项目会导入这些库，希望它们能正常工作，所以在这里放东西之前要三思。

internal 目录是确保私有包不可导入的更好方法，因为它是由 Go 强制执行的。/pkg 目录仍然是一种很好的方式，可以显式地表示该目录中的代码对于其他人来说是安全使用的好方法。
- /test

额外的外部测试应用程序和测试数据。你可以随时根据需求构造 /test 目录。对于较大的项目，有一个数据子目录是有意义的。例如，你可以使用 /test/data 或 /test/testdata (如果你需要忽略目录中的内容)。请注意，Go 还会忽略以“.”或“_”开头的目录或文件，因此在如何命名测试数据目录方面有更大的灵活性。
- /third_party

外部辅助工具，分叉代码和其他第三方工具(例如 Swagger UI)。


