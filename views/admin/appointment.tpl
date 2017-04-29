{{template "admin/base.tpl" .}}
{{define "meta"}}
<title>预约管理-商城管理系统</title>
{{end}}

{{define "css"}}
<link href="/web-static/admin/css/font-awesome.css?v=4.4.0" rel="stylesheet">
<!-- Data Tables -->
<link href="/web-static/admin/css/plugins/dataTables/dataTables.bootstrap.css" rel="stylesheet">
<link href="/web-static/admin/css/animate.css" rel="stylesheet">
<link href="/web-static/admin/css/style.css?v=4.1.0" rel="stylesheet">
{{end}}

{{define "content"}}
<div class="row">
    <div class="col-sm-12">
        <div class="ibox float-e-margins">
            <div class="ibox-title">
                <h5>用户在线预约 <small>分类，查找</small></h5>
                <div class="ibox-tools">
                    <a class="collapse-link">
                        <i class="fa fa-chevron-up"></i>
                    </a>
                    <a class="dropdown-toggle" data-toggle="dropdown" href="table_data_tables.html#">
                        <i class="fa fa-wrench"></i>
                    </a>
                    <ul class="dropdown-menu dropdown-user">
                        <li><a href="table_data_tables.html#">选项1</a>
                        </li>
                        <li><a href="table_data_tables.html#">选项2</a>
                        </li>
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
                        <th>用户名</th>
                        <th>真实姓名</th>
                        <th>联系方式</th>
                        <th>预约内容</th>
                        <th>预约时间</th>
                        <th>状态</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range .appointment}}
                    <tr class="gradeX">
                        <td>{{.Id}}</td>
                        <td>{{.User.Name}}</td>
                        <td>{{.Name}}</td>
                        <td>{{.Mobile}}</td>
                        <td>{{.Content}}</td>
                        <td>{{.Time}}</td>
                        {{if eq .IsRead 0}}
                        <td class="center" style="text-align: center"><button class="btn btn-warning btn-sm" onclick="readIt({{.Id}})"><i class="fa fa-book"></i> 标记为已读</button></td>
                        {{else}}
                        <td style="text-align: center;color: #999"><p>√</p></td>
                        {{end}}
                    </tr>
                    {{end}}
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>
{{end}}



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
    function readIt(id) {
        $.ajax({
            url: '/appointment/readIt',
            type: 'post',
            data:{
                'mid':id
            },
            dataType: 'json',
            contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
            cache: false,
            success: function(data) {
                if(data.status==10000){
                    window.location.reload()
                }else{
                    alert("操作失败，请重试");
                }
            },
            error : function() {
                alert("无法连接服务器");
            }
        });
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