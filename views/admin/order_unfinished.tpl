{{template "admin/base.tpl" .}}
{{define "meta"}}
<title>未完成订单-商城管理系统</title>
{{end}}

{{define "css"}}
<link href="/web-static/admin/css/font-awesome.css?v=4.4.0" rel="stylesheet">

<!-- Data Tables -->
<link href="/web-static/admin/css/plugins/dataTables/dataTables.bootstrap.css" rel="stylesheet">

<link href="/web-static/admin/css/animate.css" rel="stylesheet">
<link href="/web-static/admin/css/style.css?v=4.1.0" rel="stylesheet">
<style>
</style>
{{end}}

{{define "content"}}
<div class="row">
    <div class="col-sm-12">
        <div class="ibox float-e-margins">
            <div class="ibox-title">
                <h5>待处理订单 <small>分类，查找</small></h5>
                <div class="ibox-tools">
                    <a class="collapse-link">
                        <i class="fa fa-chevron-up"></i>
                    </a>
                    <a class="dropdown-toggle" data-toggle="dropdown" href="table_data_tables.html#">
                        <i class="fa fa-wrench"></i>
                    </a>
                    <ul class="dropdown-menu dropdown-user">
                    </ul>
                    <a class="close-link">
                        <i class="fa fa-times"></i>
                    </a>
                </div>
            </div>
            <div class="ibox-content">
                <table class="table table-striped table-bordered table-hover dataTables-example">
                    <thead>
                    <tr>
                        <th>序号</th>
                        <th>购买者</th>
                        <th>总金额</th>
                        <th>订单信息</th>
                        <th>下单时间</th>
                        <th>收货地址</th>
                        <th>查询物流</th>
                        <th>删除</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range .order_unfinished}}
                    <tr class="gradeX">
                        <th>{{.Id}}</th>
                        <th>{{.User.Name}}</th>
                        <th>￥{{.Price}}</th>
                        <th>
                            <p>商品代码</p>
                            <!--<p>商品名称 ×1</p>-->
                            <!--<p>商品名称 ×3</p>-->
                            <!--<p>商品名称 ×1</p>-->
                        </th>
                        <th>{{date .CreateTime "Y-m-d h:i:s"}}</th>
                        <th>{{.Address}}</th>
                        <!--发货输入物流单号-->
                        <td class="center" style="text-align: center"><button class="btn btn-warning btn-sm" onclick="openDriverModal({{.Id}});"><i class="fa fa-truck"></i> 发货</button></td>
                        <td class="center" style="text-align: center"><button class="btn btn-danger btn-sm" onclick="deleteOrder({{.Id}})"><i class="fa fa-trash"></i> 删除</button></td>
                    </tr>
                    {{end}}
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>

{{end}}



<div class="modal fade" id="DriverModal" tabindex="-1" role="dialog"  aria-hidden="true">
    <div class="modal-dialog modal-lg">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">修改分类信息</h4>
            </div>
            <div class="modal-body">
                <form role="form">
                    <div class="form-group">
                        <label>输入已发货物流单号</label>
                        <input type="text" placeholder="请输入物流单号" class="form-control" id="oddNumber" value="">
                    </div>
                    <input type="text" class="hidden" id="hiddenOrderId" value="">
                </form>

            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-primary" onclick="submitDriver();">提交</button>
            </div>
        </div>
    </div>
</div>

{{define "js"}}

<!-- 全局js -->
<script src="/web-static/admin/js/jquery.min.js?v=2.1.4"></script>
<script src="/web-static/admin/js/bootstrap.min.js?v=3.3.6"></script>



<script src="/web-static/admin/js/plugins/jeditable/jquery.jeditable.js"></script>

<!-- Data Tables -->
<script src="/web-static/admin/js/plugins/dataTables/jquery.dataTables.js"></script>
<script src="/web-static/admin/js/plugins/dataTables/dataTables.bootstrap.js"></script>

<!-- 自定义js -->
<script src="/web-static/admin/js/content.js?v=1.0.0"></script>
<script>
    function deleteOrder(id) {

        $.ajax({
            url: '/order/deleteOrder',
            type: 'post',
            dataType: 'json',
            data:{
                'oId':id
            },
            contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
            cache: false,
            success: function(data) {
                if(data.status==10000){
                    window.location.href="/admin/order_unfinished";
                }else if(data.status==10001){
                    alert(data.message);
                }else{
                    alert('您的网络异常!!！');
                }

            },
            error : function() {
                alert('服务器失去响应，请检查您的网络！');


            }
        })
    }
    function openDriverModal(id){
            $('#hiddenOrderId').val(id);
            $('#DriverModal').modal('show');

    }
    function submitDriver() {
        var oId = $('#hiddenOrderId').val();
        var oddNumber = $('#oddNumber').val();
        $.ajax({
            url: '/order/finishingOrder',
            type: 'post',
            dataType: 'json',
            data:{
                'oId':oId,
                'oddNumber':oddNumber
            },
            contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
            cache: false,
            success: function(data) {
                if(data.status==10000){
                    window.location.href="/admin/order_unfinished";
                }else if(data.status==10001){
                    alert(data.message);
                }else if(data.status==10002){
                    alert(data.message);
                }else{
                    alert('您的网络异常!!！');
                }

            },
            error : function() {
                alert('服务器失去响应，请检查您的网络！');


            }
        })
    }
</script>

<!-- Page-Level Scripts -->
<script>
    $(document).ready(function () {
        $('.dataTables-example').dataTable();

        /* Init DataTables */
        var oTable = $('#editable').dataTable();

        /* Apply the jEditable handlers to the table */
        oTable.$('td').editable('../example_ajax.php', {
            "callback": function (sValue, y) {
                var aPos = oTable.fnGetPosition(this);
                oTable.fnUpdate(sValue, aPos[0], aPos[1]);
            },
            "submitdata": function (value, settings) {
                return {
                    "row_id": this.parentNode.getAttribute('id'),
                    "column": oTable.fnGetPosition(this)[2]
                };
            },

            "width": "90%",
            "height": "100%"
        });


    });

    function fnClickAddRow() {
        $('#editable').dataTable().fnAddData([
            "Custom row",
            "New row",
            "New row",
            "New row",
            "New row"]);

    }
</script>


{{end}}