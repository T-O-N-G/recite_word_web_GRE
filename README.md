# 背单词 网页版
想要弄个能自定义词表的

## 表结构
``` sql
CREATE TABLE `WC1500` (
  `id` int(8) NOT NULL AUTO_INCREMENT,
  `list` int(8) DEFAULT NULL,
  `word` varchar(255) DEFAULT '',
  `mean` varchar(255) DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=272 DEFAULT CHARSET=utf8mb4;
```

## 数据库连接信息
```go
const (
	USERNAME = ""
	PASSWORD = ""
	NETWORK  = "tcp"
	SERVER   = ""
	PORT     = 3306
	DATABASE = ""
)
```

## APIs

`0.0.0.0:4000/` 
随机获取20个单词

`0.0.0.0:4000/list/28` 
获取list28的单词

`0.0.0.0:4000/means_r/100` 
获取指定数量(100)的随机单词意思，用于生成选项

## TO-DO
已完成 api
接下来是实现界面