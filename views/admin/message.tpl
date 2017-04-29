{{template "admin/base.tpl" .}}
{{define "meta"}}
<title>留言管理-商城管理系统</title>
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
                <h5>用户留言 <small>分类，查找</small></h5>
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
                        <th>留言内容</th>
                        <th>留言时间</th>
                        <th>状态</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range .message}}
                    <tr class="gradeX">
                        <td>{{.Id}}</td>
                        <td>{{.User.Name}}</td>
                        <td>{{.Content}}</td>
                        <td>{{date .CreateTime "Y-m-d h:i:s"}}</td>
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
<div class="modal fade" id="myModal5" tabindex="-1" role="dialog"  aria-hidden="true">
    <div class="modal-dialog modal-lg">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">新增产品分类</h4>
            </div>
            <div class="modal-body">
                <form role="form">
                    <div class="form-group">
                        <label>分类名称</label>
                        <input type="email" placeholder="请输入分类名称" class="form-control">
                    </div>
                    <div class="form-group">
                        <label>分类类型</label>
                        <select class="form-control" name="" id="">
                            <option>请选择分类类别</option>
                            <option value="0">产品分类</option>
                            <option value="1">特殊分类</option>
                        </select>
                    </div>
                </form>

            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-primary">提交</button>
            </div>
        </div>
    </div>
</div>

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
            url: '/message/readIt',
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