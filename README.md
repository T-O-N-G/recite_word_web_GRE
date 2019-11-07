# 背单词 网页版
想要弄个能自定义词表的

> Do not gentle into that good night.


> 抱着置之死地和后生的心态，相信重复的力量，单词其实那么简单！



## 体验 Demo
http://blog.tongxx.top:4000
![demo](https://raw.githubusercontent.com/IvyB/recite_word_web/master/demo.JPG)


## 功能
- 上传了31个list
- 可以随机抽取20个词
- 每道题做完显示正确答案选项 (绿色 加粗)
- 背错的单词加入到背词队列的最后
- 当所有词都答对了，显示答错过的单词和释义

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

`0.0.0.0:4000/words` 
随机获取20个单词

`0.0.0.0:4000/list/28` 
获取list28的单词

`0.0.0.0:4000/means_r/100` 
获取指定数量(100)的随机单词意思，用于生成选项


