// ==UserScript==
// @name         strava_sync
// @namespace    http://ersut.top
// @version      0.0.1
// @description  从strava获取fit文件上传至服务端
// @author       ersut
// @include        https://www.strava.com/*
// @icon         <$ICON$>
// @grant        none
// ==/UserScript==

(function() {
    'use strict';
    var $ = jQuery;

    console.info("启动脚本");

    function sleep(ms) {
        return new Promise((resolve) => setTimeout(resolve, ms));
    }

    var login = function(){
        $("#email").val("83873@qq.com");
        $("#password").val("");
        $("#remember_me")[0].checked = true;
        $("#login-button").click();
    }

    function handle_data(page,page_data){

        var page_data_all = [];
        
        //获取几天内的数据
        var end_date = new Date().setDate(new Date().getDate()-65);
        console.info("截止日期",new Date(end_date));
		
		return next_page_data(page,page_data_all,page_data,end_date,1);
        
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

    //获取当前地址
    var path = location.pathname;
    if(path === "/login"){
        console.info("开始登录");
        login();
    }else if(path === "/athlete/training"){
        console.info("我的活动");
        //分页信息
        var page = activityCollectionAdapter.attributes;
        console.info("page",page);
        //数据信息
        var page_data = activityCollectionAdapter.activities;
        console.info("page_data",page_data);

        var page_data_all_promise = handle_data(page,page_data)
        console.info("page_data_all",page_data_all_promise);


    }else {
        console.info("地址错误["+path+"]进行重定向");
        location.href = "/athlete/training"
    }
    // Your code here...
})();