<!DOCTYPE html>
<html>

<head>
    {{template "meta" .}}
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="renderer" content="webkit">
    <meta name="keywords" content="">
    <meta name="description" content="">
    <!--[if lt IE 9]>
    <meta http-equiv="refresh" content="0;ie.html" />
    <![endif]-->
    <link rel="shortcut icon" href="favicon.ico">
    <link href="/web-static/admin/css/bootstrap.min.css" rel="stylesheet">
    <link href="/web-static/admin/css/font-awesome.min.css?v=4.4.0" rel="stylesheet">
    <link href="/web-static/admin/css/animate.css" rel="stylesheet">
    <link href="/web-static/admin/css/style.css?v=4.1.0" rel="stylesheet">
    {{template "css" .}}
    <style>
        .modal-backdrop {
            z-index: 0 !important;
            pointer-events: none;
        }
        .modal-dialog{
            margin-top: 13.5% !important;
        }

    </style>
</head>

<body class="fixed-sidebar full-height-layout gray-bg" style="overflow:hidden">
<div id="wrapper">
    <!--左侧导航开始-->
    <nav class="navbar-default navbar-static-side" role="navigation">
        <div class="nav-close"><i class="fa fa-times-circle"></i>
        </div>
        <div class="sidebar-collapse">
            <ul class="nav" id="side-menu">
                <li class="nav-header">
                    <div class="dropdown profile-element">
                        <a data-toggle="dropdown" class="dropdown-toggle" href="#">
                                <span class="clear">
                                    <span class="block m-t-xs" style="font-size:20px;">
                                        <i class="fa fa-shopping-cart"> </i>
                                        <strong class="font-bold"> Shop-Go</strong>
                                    </span>
                                </span>
                        </a>
                    </div>
                    <div class="logo-element">Shop-Go
                    </div>
                </li>
                <li class="hidden-folded padder m-t m-b-sm text-muted text-xs">
                    <span class="ng-scope">统计</span>
                </li>
                <li>
                    <a class="J_menuItem" href="/admin/home">
                        <i class="fa fa-home"></i>
                        <span class="nav-label">主页</span>
                    </a>
                </li>

                <li class="line dk"></li>
                <li class="hidden-folded padder m-t m-b-sm text-muted text-xs">
                    <span class="ng-scope">商品管理</span>
                </li>
                <li>
                    <a href="/admin/category"><i class="fa fa-envelope"></i> <span class="">书籍分类 </span><span class="label label-info pull-right" id="productType"></span></a>
                </li>
                <li>
                    <a href="/admin/product"><i class="fa fa-envelope"></i> <span class="">书籍列表 </span><span class="label label-info pull-right" id="productList"></span></a>
                </li>
                <li class="line dk"></li>
                <li class="hidden-folded padder m-t m-b-sm text-muted text-xs">
                    <span class="ng-scope">订单管理</span>
                </li>
                <li>
                    <a href="#"><i class="fa fa-flask"></i> <span class="nav-label">订单信息</span><span class="label label-warning pull-right" id="orderInfo"></span></a>
                    <ul class="nav nav-second-level">
                        <li>
                            <a class="J_menuItem" href="/admin/order_finished">已完成订单</a>
                        </li>
                        <li>
                            <a class="J_menuItem" href="/admin/order_unfinished">未完成订单 <span class="label label-warning pull-right" id="unFinishedOrder"></span></a>
                        </li>
                    </ul>
                </li>
                <li class="line dk"></li>
                <li class="hidden-folded padder m-t m-b-sm text-muted text-xs">
                    <span class="ng-scope">用户管理</span>
                </li>
                <li>
                    <a class="J_menuItem" href="/admin/message"><i class="fa fa-picture-o"></i> <span class="nav-label">留言管理</span><span class="label label-info pull-right" id="messageCount"></span></a>
                </li>
                <li>
                    <a class="J_menuItem" href="/admin/comment"><i class="fa fa-magic"></i> <span class="nav-label">评论审核</span></a>
                </li>
                <li>
                    <a href="/admin/appointment"><i class="fa fa-cutlery"></i> <span class="nav-label">预约管理 <span class="label label-info pull-right" id="appointmentCount"></span></span></a>

                </li>
                <li>
                    <a class="J_menuItem" href="/admin/carousel"><i class="fa fa-magic"></i> <span class="nav-label">轮播图管理</span></a>

                </li>


            </ul>
        </div>
    </nav>
    <!--左侧导航结束-->
    <!--右侧部分开始-->
    <div id="page-wrapper" class="gray-bg dashbard-1" >
        <div class="row border-bottom">
            <nav class="navbar navbar-static-top" role="navigation" style="margin-bottom: 0">
                <ul class="nav navbar-top-links navbar-right">
                    <li class="dropdown">
                        <a class="dropdown-toggle count-info" data-toggle="dropdown" href="#">
                            <i class="fa fa-envelope"></i> <span class="label label-warning">16</span>
                        </a>
                        <ul class="dropdown-menu dropdown-messages">
                            <li class="m-t-xs">
                                <div class="dropdown-messages-box">
                                    <a href="profile.html" class="pull-left">
                                        <img alt="image" class="img-circle" src="img/a7.jpg">
                                    </a>
                                    <div class="media-body">
                                        <small class="pull-right">46小时前</small>
                                        <strong>小四</strong> 是不是只有我死了,你们才不骂爵迹
                                        <br>
                                        <small class="text-muted">3天前 2014.11.8</small>
                                    </div>
                                </div>
                            </li>
                            <li class="divider"></li>
                            <li>
                                <div class="dropdown-messages-box">
                                    <a href="profile.html" class="pull-left">
                                        <img alt="image" class="img-circle" src="img/a4.jpg">
                                    </a>
                                    <div class="media-body ">
                                        <small class="pull-right text-navy">25小时前</small>
                                        <strong>二愣子</strong> 呵呵
                                        <br>
                                        <small class="text-muted">昨天</small>
                                    </div>
                                </div>
                            </li>
                            <li class="divider"></li>
                            <li>
                                <div class="text-center link-block">
                                    <a class="J_menuItem" href="mailbox.html">
                                        <i class="fa fa-envelope"></i> <strong> 查看所有消息</strong>
                                    </a>
                                </div>
                            </li>
                        </ul>
                    </li>
                    <li class="dropdown">
                        <a class="dropdown-toggle count-info" data-toggle="dropdown" href="#">
                            <i class="fa fa-bell"></i> <span class="label label-primary">8</span>
                        </a>
                        <ul class="dropdown-menu dropdown-alerts">
                            <li>
                                <a href="mailbox.html">
                                    <div>
                                        <i class="fa fa-envelope fa-fw"></i> 您有16条未读消息
                                        <span class="pull-right text-muted small">4分钟前</span>
                                    </div>
                                </a>
                            </li>
                            <li class="divider"></li>
                            <li>
                                <a href="profile.html">
                                    <div>
                                        <i class="fa fa-qq fa-fw"></i> 3条新回复
                                        <span class="pull-right text-muted small">12分钟钱</span>
                                    </div>
                                </a>
                            </li>
                            <li class="divider"></li>
                            <li>
                                <div class="text-center link-block">
                                    <a class="J_menuItem" href="notifications.html">
                                        <strong>查看所有 </strong>
                                        <i class="fa fa-angle-right"></i>
                                    </a>
                                </div>
                            </li>
                        </ul>
                    </li>
                    <li><a href="/admin/logout"><i class="fa fa-sign-out"></i>注销</a></li>

                </ul>
            </nav>
        </div>
        <div class="row J_mainContent" id="content-main" style="overflow: scroll">
            <!--代码部分-->
            <div class="gray-bg">
                <div class="wrapper wrapper-content">
                    {{template "content" .}}

                </div>
            </div>



        </div>
    </div>
    <!--右侧部分结束-->
</div>

<!-- 全局js -->
<script src="/web-static/admin/js/jquery.min.js?v=2.1.4"></script>
<script src="/web-static/admin/js/bootstrap.min.js?v=3.3.6"></script>
<script src="/web-static/admin/js/plugins/metisMenu/jquery.metisMenu.js"></script>
<script src="/web-static/admin/js/plugins/slimscroll/jquery.slimscroll.min.js"></script>
<script src="/web-static/admin/js/plugins/layer/layer.min.js"></script>

<!-- 自定义js -->
<script src="/web-static/admin/js/hAdmin.js?v=4.1.0"></script>
<script type="text/javascript" src="/web-static/admin/js/index.js"></script>

<!-- 第三方插件 -->
<script src="/web-static/admin/js/plugins/pace/pace.min.js"></script>
<!--引入yifanToast-->
<link rel="stylesheet" href="/web-static/admin/css/yifanToast.css">
<script src="/web-static/admin/js/yifanToast.js" type="text/javascript"></script>

<script>

    window.onload = function (){
        $.ajax({
            url: '/api/getProductAndCategory',
            type: 'post',
            dataType: 'json',
            data:{
            },
            contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
            cache: false,
            success: function(data) {
                if(data.status==10000){
                    $("#productType").html(data.category.length);
                    $("#productList").html(data.product.length);
                    $("#messageCount").html(data.message.length);
                    $("#appointmentCount").html(data.appointment.length);


                }else{
                    alert('您的网络异常!!！');
                }

            },
            error : function() {
                alert('服务器失去响应，请检查您的网络！');


            }
        });

        $.ajax({
            url: '/order/showUnfinishedOrder',
            type: 'post',
            dataType: 'json',
            data:{
            },
            contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
            cache: false,
            success: function(data) {
                if(data.status==10000){
                    if(data.order.length==0){
                        $("#orderInfo").remove();
                        $("#unFinishedOrder").remove();
                    }else{
                        $("#orderInfo").html(data.order.length);
                        $("#unFinishedOrder").html(data.order.length);
                    }

                }else{
                    alert('您的网络异常!!！');
                }

            },
            error : function() {
                alert('服务器失去响应，请检查您的网络！');


            }
        })
    };
</script>
{{template "js" .}}
</body>

</html>
