# 背单词 网页版
想要弄个能自定义词表的

> Do not gentle into that good night.


> 抱着置之死地和后生的心态，相信重复的力量，单词其实那么简单！



## 体验 Demo
http://blog.tongxx.top:4000
![demo](https://raw.githubusercontent.com/IvyB/recite_word_web/master/demo.JPG)

![demo](https://raw.githubusercontent.com/IvyB/recite_word_web/master/demo1.JPG)



## 功能

- 上传了再要你命3000和救我狗命800，感谢微臣

- 可以随机抽取20个词

- GRE800包含近义词，很重要

- 每道题做完显示正确答案选项 (绿色 加粗)

- 背错的单词加入到背词队列的最后

- 当所有词都答对了，显示答错过的单词和释义

  

## 表结构

``` sql
CREATE TABLE `WC800` (
  `id` int(4) NOT NULL AUTO_INCREMENT,
  `list` int(4) DEFAULT NULL,
  `word` varchar(255) DEFAULT NULL,
  `mean` varchar(255) DEFAULT NULL,
  `mean_e` varchar(255) DEFAULT NULL,
  `exp` varchar(500) DEFAULT NULL,
  `similar` varchar(255) DEFAULT NULL,
  `opposite` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1628 DEFAULT CHARSET=utf8mb4;
```

其中 list word mean 是必要字段  ~~不然背个锤子~~



## 数据库连接信息

在main.go文件中填写：

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

`0.0.0.0:4000/word/WC800/rand` 
随机获取WC800中的20个单词

`0.0.0.0:4000/word/WC1500/list/28` 
获取WC1500中的list28的单词

`0.0.0.0:4000/word_learn/WC800/list/28` 
获取WC800中的list28的单词和近义词等

`0.0.0.0:4000/means_r/100` 
获取指定数量(100)的随机单词意思，用于生成选项