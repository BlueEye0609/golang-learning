作业
---
我们在数据库操作的时候，比如 DAO 层中遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个error，抛给上层。为什么，应该怎么做请写出代码。

答
---
当 DAO 层中遇到 sql.ErrNoRows 的错误时，我认为不应该将这个错误 wrap 抛给上层。

ErrNoRows 发生原因是因为 https://pkg.go.dev/database/sql#pkg-variables
> ErrNoRows is returned by Scan when QueryRow doesn't return a row. In such a case, QueryRow returns a placeholder *Row value that defers this error until a Scan.

QueryRow是获取至多单行的结果。Go 中定义了一个名为 sql.ErrNowRows 的特殊错误常量。当结果集为空时，QueryRow 就会返回它。
对于业务层来讲，它关心的是我 query 到的值。一个空的结果集往往不应该被认为是应用的错误。

### 数据库准备
```
# docker run --name testdb-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=hyj123 -d mysql:latest
# docker exec -it testdb-mysql bash

# mysql -h localhost -u root -phyj123 

mysql> CREATE DATABASE gotestdb;
Query OK, 1 row affected (0.03 sec)

mysql> USE gotestdb;
Database changed

mysql> CREATE TABLE books(
         ID INT AUTO_INCREMENT NOT NULL,
         title  VARCHAR(128) NOT NULL,
         Author VARCHAR(255) NOT NULL,
         price  DECIMAL(5,2) NOT NULL,
         PRIMARY KEY (`ID`)
     );
Query OK, 0 rows affected (0.07 sec)

mysql> INSERT INTO books (title, Author,price) VALUES
        ('Go Programming', 'Author X', 78.00),
        ('AIOps', 'auth ai', 128.00),
        ('Deep learning', 'robert',58.50);
```

### 模块
main.go -> models -> db

error 处理: 
