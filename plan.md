# 方案

目标:~~基于Sauce for Strava插件~~从strava中获取fit文件 ，最终实现代码自动获取文件上传至黑鸟与行者。

获取fit文件：

1. 基于浏览器插件插入js代码与后台交互
2. ~~插入js代码链接服务的socket~~ ，取消socket方式太消耗资源
3. ~~后台通知js刷新页面（防止掉线问题）~~
4. 获取strava账号、密码以及获取近多少天的数据等基础数据
5. 读取localStorage中的最后一次同步日期
	1. 若当前日期比下一次同步日期早
		1. 恢复定时
	2. 若当前日期比下一次同步日期晚，或者没有同步日期，开始同步
		1. js获取近多少天的数据并发送给后台，经后台与数据库比对（未同步的数据）后，响应给js获取那些记录的fit文件
		2. js下载fit文件并上传给后台
			1. 由js获取原数据（佳明数据，导出原来的活动按钮）链接
			2. 下载由ajax获取流并传给后台，这样知道文件的下载进度
			3. 后台接收数据保存文件并存库，后续上传给黑鸟以及行者
		3. 所有文件上传成功后，在localStorage中存储最后一次同步日期，并开启下次同步定时（后台控制同步的时间）
			1. 定时中执行刷新页面
6. ~~后台通知js获取近几天（具体几天后台可配置项）的数据,并由js获取接口数据提交至后台~~
7. ~~后台对比数据后通知js获取哪几天的数据~~
8. ~~由js篡改Sauce for Strava插件中下载文件方法获取其链接（具体看备注1）~~，由于Sauce for Strava插件的fit文件无法导入黑鸟，所以放弃这个方案
9. 所有js网络报错后刷新网页重新执行流程（只尝试10次，10次未果通知微信server酱）

上传至行者：


上传至黑鸟：

难点：

1. ~~如果运行在centos上需要排查火狐是否有Sauce for Strava插件，谷歌有~~
2. 记得有款谷歌插件支持js插入，找找，再确定下火狐是否有（为了兼容centos）
	1. Tampermonkey插件
	2. 火狐有Tampermonkey插件
3. 跨域问题~~改本机host，创建一个二级域名~~，通过油猴的跨域请求解决
4. Tampermonkey插件的使用

备注

1. 文件名唯一处理

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

2. 由于strava引入了google的js导致整个页面加载慢，解决方案：修改本机host将google设置为127.0.0.1

3. zwift的记录黑鸟不接收，行者可以，需处理（骑行数据中有type字段可区分）

