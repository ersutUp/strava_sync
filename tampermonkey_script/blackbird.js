// ==UserScript==
// @name         blackbird_upload
// @namespace    http://ersut.top
// @version      0.0.1
// @description  从服务端获取fit文件上传至黑鸟
// @author       ersut
// @include      http://www.blackbirdsport.com/*
// @icon         <$ICON$>
// @grant        GM_xmlhttpRequest
// @require https://code.jquery.com/jquery-2.1.4.min.js
// @connect      *
// ==/UserScript==

(function() {
    'use strict';
	
	var server_host = "http://127.0.0.1:81";
	
	var last_sync_key = "last_sync";
	var err_count = "err_count";
		
	var user_info = {};

    console.info("启动脚本");

    function sleep(ms) {
        return new Promise((resolve) => setTimeout(resolve, ms));
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


    var login = function(){
        $("#username").val(user_info.BlackbirdUsername);
        $("#password").val(user_info.BlackbirdPass);
        $("#remember-me")[0].checked = true;
        $("input[type=submit]").click();
		
		setTimeout(function(){
			location.reload();
		},30*1000)
    }
	
	//通知后台上传黑鸟成功
	function noticeServerUploadOk(id,recordId,resolve, reject){

		GM_xmlhttpRequest({
			method: "put",
			url: server_host+'/v1/training/blackbird',
			data: "{\"ID\":"+id+",\"IsUploadBlackbird\":"+recordId+"}",
			binary:true,
			onload: function(res){
				console.log("通知后台上传黑鸟成功->响应",res);
				if(res.status === 200){
					if(res.response != "ok"){
						var err_msg = "通知后台上传黑鸟成功->出错"
						console.log(err_msg,res.response);
						reject(err_msg);
					} else {
						resolve("ok");
					}
				} else {
					var err_msg = "通知后台上传黑鸟成功->状态码异常"
					handleError(err_msg)
					console.error(err_msg,res)
					reject(err_msg);
				}
			},
			onerror : function(err){
				var err_msg = "通知后台上传黑鸟成功->请求异常"
				handleError(err_msg)
				console.error(err_msg,err)
				reject(err_msg);
			}
		})
	}
	
	//上传给黑鸟
	function uploadFitBlackbird(id,blob,resolve, reject){
		var fd = new FormData();
		fd.append('file', blob);
		$.ajax({
			method: "post",
			url: "/api/records/upload",
			data: fd,
			processData : false,
			contentType : false,
			success: function (res) {
				console.log("上传给黑鸟fit文件->响应",res);
				if (res.status == "ok") {
					//通知后台上传黑鸟成功
					noticeServerUploadOk(id,res.recordId,resolve, reject)
				} else {
					var err_msg = "上传给黑鸟fit文件->响应异常"
					console.error(err_msg,res);
					reject(err_msg);
				}
			},
			error:function(err){
				//由于黑鸟不一定上传成功所以不做handleError的处理
				var err_msg = "上传给黑鸟fit文件->状态码异常:["+err.status+"],"+err.responseText
				console.error(err_msg,err)
				reject(err_msg);
			}
		})
	}
	
	
	//上传黑鸟fit
	function uploadFitHandle(id) {
        return new Promise(function(resolve, reject){
			//获取文件
			GM_xmlhttpRequest({
				method: "get",
				url: server_host+'/v1/file/fit?id='+id,
				responseType:"arraybuffer",
				onload: function(res){
					if(res.status === 200){
						console.log("获取fit文件->响应",res);
						let blob = new Blob([res.response], {
							type: "application/octet-stream",
						});
						//上传给黑鸟
						uploadFitBlackbird(id,blob,resolve,reject)
					}else{
						var err_msg = "获取fit文件->状态码异常:["+res.status+"]"
						handleError(err_msg)
						console.error(err_msg,res)
						reject(err_msg);
					}
				},
				onerror : function(err){
					var err_msg = "获取fit文件->请求异常"
					handleError(err_msg)
					console.error(err_msg,err)
					reject(err_msg);
				}
			});
		});
    }
	
	async function uploadFits(ids){
		for(var i = 0 ; i < ids.length ; i++ ){
			var id = ids[i];
			console.info("id["+id+"]开始处理")
			await uploadFitHandle(id).then(function(val){
				console.info("id["+id+"]处理完毕");
			},function(err){
				console.info("id["+id+"]出错:"+err);
			})
		}
		console.info("fit文件处理完毕")
		next_sync(1)
	}

	function handle_data(){
		//请求后台获取上传数据id
		GM_xmlhttpRequest({
			method: "get",
			url: server_host+'/v1/training/blackbird',
			synchronous: true,
			onload: function(res){
				if(res.status === 200){
					console.log("获取需上传fit文件的id->响应",res);
					//遍历 id 获取文件流 并上传黑鸟
					var ids = JSON.parse(res.response);
					if(ids != []){
						uploadFits(ids)
					} else {
						next_sync(1)
					}
				}else{
					var err_msg = "获取需上传fit文件的id->状态码异常:["+res.status+"]"
					handleError(err_msg)
					console.error(err_msg,res)
				}
			},
			onerror : function(err){
				var err_msg = "获取需上传fit文件的id->请求异常"
				handleError(err_msg)
				console.error(err_msg,err)
			}
		});
	}
	
	function handlePath(){
		//获取当前地址
		var path = location.pathname;
		if(path === "/auth/login"){
			console.info("开始登录");
			login();
		} else if(path === "/user/records"){
			console.info("我的记录");
			
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
			location.href = "/user/records"
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