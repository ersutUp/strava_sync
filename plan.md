# 方案

目标:~~基于Sauce for Strava插件~~从strava中获取fit文件 ，最终实现代码自动获取文件上传至黑鸟与行者。

获取fit文件：

1. 基于浏览器插件插入js代码与后台交互
2. 插入js代码链接服务的socket
3. 后台通知js获取近几天（具体几天后台可配置项）的数据,并由js获取接口数据提交至后台
4. 后台对比数据后通知js获取哪几天的数据
5. ~~由js篡改Sauce for Strava插件中下载文件方法获取其链接（具体看备注1）~~，由于Sauce for Strava插件的fit文件无法导入黑鸟，所以放弃这个方案
6. 由js获取原数据（佳明数据，导出原来的活动按钮）链接
7. 下载由ajax获取流并传给后台，这样知道文件的下载进度
8. 后台接收数据保存文件并存库，后续上传给黑鸟以及行者

上传至行者：


上传至黑鸟：

难点：

1. ~~如果运行在centos上需要排查火狐是否有Sauce for Strava插件，谷歌有~~
2. 记得有款谷歌插件支持js插入，找找，再确定下火狐是否有（为了兼容centos）
	1. Tampermonkey插件
	2. 火狐未确认
3. 跨域问题改本机host，创建一个二级域名

备注
1、文件名唯一处理

```js
	sauce.downloadBlob = function(blob, name) {
        const url = URL.createObjectURL(blob);
        try {
            sauce.downloadURL(url,"hahaha.fit");
        } finally {
            URL.revokeObjectURL(url);
        }
    };
```

