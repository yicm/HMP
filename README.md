# 1 HMP

`HMP`(Hexo & mini program)，打通Hexo和小程序的博客框架，附带了博文编译器`HBC`，以及目前我自己写的几个插件（生态还需要感兴趣的同学一起来创建），具体列表如下：

- NewWxComment: 一个打通小程序和PC端评论（Valine）的组件，已开源，https://github.com/yicm/NewWxComment
- WxPoster: 一个小程序端海报制作和分享组件，已开源，https://github.com/yicm/WxPoster
- HBC-CLI：Hexo 博文编译器
- 小白AI.易名：一个打通Hexo静态博客框架和小程序端的经典主题（微信搜索`小白AI` ，PC端访问 https://xiaobaiai.net）

# 2 HBC，一个博文编译器

`HBC`(Hexo build compiler for mini program)是我用golang写的`博文编译器`，可以将当前主流的静态博客框架`Markdown`博文重新编译，最终输出包括：

- 博文按时间排序、置顶并分页输出
- 博文按时间排序、种类分类并分页输出
- 博文按时间排序、标签分类并分页输出
- 博文按时间排序、按年份分类并分页输出
- 博文搜索内容输出

该`博文编译器`支持自定义博文配置参数，包括：

- 博文背景图片
- 描述
- 是否加密
- 是否置顶
- 是否可以评论
- 多标签、多分类支持
- 首部第一张图片配置
- 分享海报背景图配置
- 文章属性，是否原创、翻译等
- ......

**HBC的软件特点：**

- Golang编写, 跨平台（bin目录中仅提供了windows平台可执行文件）
- 协程处理，处理几百篇博文，一秒钟完成
- CLI操作方式(HBC-CLI)，交互友好，容易上手

![image](https://gitee.com/yicm/Images/raw/master/xiaobaiai/blog/1.png)


# 3 小白AI.易名 主题

`小白AI.易名` 是我打通Hexo静态博客和小程序端制作的一个经典主题，UI良心设计，细节方面更是呕心沥血。具体效果可以用微信扫一扫小程序码查看：

![](https://gitee.com/yicm/Images/raw/master/xiaobaiai/news/10004_2.jpg)

可以看几张截图：

![image](https://gitee.com/yicm/Images/raw/master/xiaobaiai/index.png)

![image](https://gitee.com/yicm/Images/raw/master/xiaobaiai/home.png)

# 4 想使用小白AI经典主题？

- 可以微信直接扫码“小白AI”小程序，查看“我的”->“关于”
- 可以到[https://xiaobaiai.net](https://xiaobaiai.net) 直接留言给我
- 可以直接加微信[XEthanm](XEthanm)咨询主题相关问题

# 5 关于开源

- 目前不会开源`小白AI.易名`主题
- 会开源HBC-CLI博文编译器工具，期待更多感兴趣同学开发新的主题
