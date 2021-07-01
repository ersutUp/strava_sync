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
	
	var server_host = "http://127.0.0.1:800";
		
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
				url: server_host+'/v1/upload',
				data: fd,
				synchronous: true,
				onload: function(res){
					if(res.status === 200){
						console.log("上传响应",res);
					}else{
						console.log('失败')
						console.log(res)
					}
				},
				onerror : function(err){
					console.log('error')
					console.log(err)
				}
			  });
		  },
		});
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
				//循环转数组字符串
				val.forEach(function(d,i){
					data.push(JSON.stringify(d))
				})
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
					}else{
						console.log('失败')
						console.log(res)
					}
				},
				onerror : function(err){
					console.log('error')
					console.log(err)
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
			
			//todo 从cookie中获取最后一次同步日期
			console.info(handle_data());

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
					console.log('获取用户信息异常',res.status)
					console.log(res)
				}
			},
			onerror : function(err){
				console.log('获取用户信息出错')
				console.log(err)
			}
		});
	}
	
	
	getUserInfo();
})();