// ==UserScript==
// @name         strava_sync
// @namespace    http://ersut.top
// @version      0.0.1
// @description  从strava获取fit文件上传至服务端
// @author       ersut
// @include        https://www.strava.com/*
// @icon         <$ICON$>
// @grant        GM_xmlhttpRequest
// @connect      *
// ==/UserScript==

(function() {
    'use strict';
    var $ = jQuery;
	
	var server_host = "http://127.0.0.1:81";
	
	var last_sync_key = "last_sync";
	var err_count = "err_count";
		
	var user_info = {};

    console.info("启动脚本");

    function sleep(ms) {
        return new Promise((resolve) => setTimeout(resolve, ms));
    }

    var login = function(){
        $("#email").val(user_info.StravaEmail);
        $("#password").val(user_info.StravaPass);
        $("#remember_me")[0].checked = true;
        $("#login-button").click();
		
		setTimeout(function(){
			location.reload();
		},30*1000)
    }
	
	//微信通知
	function notice_wechat(msg){
		
		if( user_info.SendKey == undefined){
			console.info("微信通知未配置")
			return
		}
		GM_xmlhttpRequest({
			method: "get",
			url: 'https://sctapi.ftqq.com/'+user_info.SendKey+'.send?title='+msg,
			onload: function(res){
				console.info("微信通知响应",res);
			},
			onerror : function(err){
				console.error("微信通知异常",err)
			}
		});
	}
	
	function handleError(err_msg){
		var err_num = localStorage.getItem(err_count);
		if (err_num == null) {
			err_num = 0;
		} else {
			//计数
			err_num++
		}
		
		//错误重试10次后无果则进入下一次定时
		if (err_num >= 10){
			localStorage.removeItem(err_count);
			//todo 微信通知
			notice_wechat(err_msg)
			
			next_sync(1);
			return
		} else {
			//重试
			localStorage.setItem(err_count,err_num);
			setTimeout(function(){
				location.reload()
			},6*1000)
		}
	}
	
	/*
		1:设置最后一次同步时间 并 开启定时任务
		2:恢复定时
	*/
	function next_sync(type){
		
		var millisecond = 30*1000;
		if (user_info.StravaSyncSecond != undefined) {
			millisecond = millisecond + (user_info.StravaSyncSecond*1000);
		} else {
			//正常情况下StravaSyncSecond有值，若没值增大定时的默认值
			millisecond = 300*1000;
		}
		var now = new Date();
		if (type == 1) {			
			//设置最后一次同步时间
			localStorage.setItem(last_sync_key,now)
		} else if (type == 2){
			
			//当前时间
			var nowTime = now.getTime();
			//获取最后一次同步日期，
			var last_date_time = localStorage.getItem(last_sync_key);
			if (last_date_time != null ) {
				var last_date_data = new Date(last_date_time).getTime()
				//计算定时经过的时间
				var jet_lag = nowTime - last_date_data
				if(jet_lag > 0 && millisecond > jet_lag){
					millisecond = millisecond - jet_lag
				}
			}
			
		}
		
		console.info("距离下次定时："+millisecond)
		//设置下次同步定时（比配置时间多30秒）
		setTimeout(function(){
			location.reload()
		},millisecond)
	}

    //过滤数据
    var parsing_bike_data = function(page_data_all,page_data,end_date){
        var filter_count = 0;
		page_data.models.forEach(function(d,i){
			var data = d.attributes;
			var start_time = new Date(data.start_time);
			//比较时间
			if(end_date <= start_time.getTime() ){
				filter_count++;
				page_data_all.push(data);
			} else {
				console.info("数据超出日期",data.start_date);
			}
		})
		return filter_count;
    }
	
	//下一页数据
	async function next_page_data(page,page_data_all,page_data,end_date,page_NO){
		
		if(page_NO != 1){
			var next = $(".next_page");
			//有 disabled 样式说明最后一页了
			if(next.hasClass("disabled")){
				return page_data_all;
			} else {
				next.click();
			}
		}
		
        //等待数据加载成功
        for (;1==1;) {
            await sleep(3000);
            console.info("page",page);
            if(!$.isEmptyObject(page)){
				if(page_NO == 1 || page.page == page_NO){
					console.info("page_data",page_data);
					break;
				}
            }
        }
		
		//过滤数据
        var filter_count = parsing_bike_data(page_data_all,page_data,end_date);
		
		//过滤后的数据少说明下一页数据都会被过滤，因为数据是按时间排序的
		if(filter_count < page_data.length){
			return page_data_all;
		} else {
			//获取下一页数据集
			return next_page_data(page,page_data_all,page_data,end_date,++page_NO);
		}
	}

	//上传fit文件
	var uploadFit = function(id){
		$.ajax({
		  method: "get",
		  url: "/activities/"+id+"/export_original",
		  xhrFields: {
			responseType: "arraybuffer",
		  },
		  success: function (res, _, xhr) {
			  let blob = new Blob([res], {
				  type: "application/octet-stream",
			  });
			  var fd = new FormData();
			  fd.append('id', ''+id);
			  fd.append('data', blob);
			  
			  GM_xmlhttpRequest({
				method: "post",
				url: server_host+'/v1/file',
				data: fd,
				synchronous: true,
				onload: function(res){
					if(res.status === 200){
						console.log("上传响应",res);
					}else{
						var err_msg = "上传fit文件状态码异常:["+res.status+"]"
						handleError(err_msg)
						console.error(err_msg,res)
					}
				},
				onerror : function(err){
					var err_msg = "上传fit文件请求异常"
					handleError(err_msg)
					console.error(err_msg,err)
				}
			  });
		  },
		});
	}

	async function uploadFits(ids){
		
		for(var i = 0 ; i < ids.length ; i++ ){
			uploadFit(ids[i]);
			await sleep(3000);
		}
		
		//开启定时任务
		next_sync(1);
	}

    function handle_data(){

        var page_data_all = [];
        
        //分页信息
        var page = activityCollectionAdapter.attributes;
        console.info("page",page);
        //数据信息
        var page_data = activityCollectionAdapter.activities;
        console.info("page_data",page_data);
		
        //获取几天内的数据
        var end_date = new Date().setDate(new Date().getDate()-user_info.BeforeDay);
        console.info("截止日期",new Date(end_date));
		
		var page_data_all_promise = next_page_data(page,page_data_all,page_data,end_date,1);
		
		//promise的回调函数 异步的
		page_data_all_promise.then(function(val){
			console.info("success",val);
			var data = [];
			if(val.length > 0){
				//倒序循环（为了调整入库顺序）转数组字符串
				for (var i = (val.length - 1); i >= 0 ; i-- ){
					data.push(JSON.stringify(val[i]))
				}
			}
			console.info("training data:",data)
			//上传骑行数据给后台同步上传
			GM_xmlhttpRequest({
				method: "post",
				url: server_host+'/v1/training',
				data: JSON.stringify(data),
				onload: function(res){
					if(res.status === 200){
						//todo 接收需要上传fit的id数据
						console.log("上传数据响应",res);
						var fit_ids = JSON.parse(res.response);
						if(fit_ids != []){
							//遍历下载fit
							uploadFits(fit_ids);
						} else {
							console.log("无需上传fit文件");
							
							//开启定时任务
							next_sync(1);
						}
					}else{
						var err_msg = "上传骑行数据状态码异常:["+res.status+"]"
						handleError(err_msg)
						console.error(err_msg,res)
					}
				},
				onerror : function(err){
					var err_msg = "上传骑行数据请求失败"
					handleError(err_msg)
					console.error(err_msg,err)
				}
			});
			//todo 获取fit文件并上传
			//uploadFit(5538001926);
		},function(){
			console.info("error");
		})
		
		return page_data_all;
        
    } 

	function handlePath(){
		//获取当前地址
		var path = location.pathname;
		if(path === "/login"){
			console.info("开始登录");
			login();
		} else if(path === "/athlete/training"){
			console.info("我的活动");
			
			//获取最后一次同步日期，
			var last_date_time = localStorage.getItem(last_sync_key);
			
			if (last_date_time == null ) {
				//没有最后一次同步时间 则 开始处理数据
				handle_data();
			} else {
				//下一次同步的时间
				var next_time = new Date(last_date_time).getTime() + user_info.StravaSyncSecond*1000
				if (new Date().getTime() >= next_time) {
					//已经到达同步时间则 开始处理数据
					handle_data();
				} else {						
					//时间没到重新开启定时
					next_sync(2);
				}
			}

		} else {
			console.info("地址错误["+path+"]进行重定向");
			location.href = "/athlete/training"
		}
	}
	
	//获取用户信息
	function getUserInfo(){
		GM_xmlhttpRequest({
			method: "get",
			url: server_host+'/v1/user',
			synchronous: true,
			onload: function(res){
				if(res.status === 200){
					console.log("获取用户信息响应",res);
					user_info = JSON.parse(res.response);
					console.info("user_info",user_info);
					handlePath();
				}else{
					var err_msg = "获取用户信息状态码异常:["+res.status+"]"
					console.error(err_msg,res)
				}
			},
			onerror : function(err){
				var err_msg = "获取用户信息请求失败"
				console.error(err_msg,err)
				setTimeout(function(){
					location.reload();
				},30*1000)
			}
		});
	}
	
	
	getUserInfo();
})();