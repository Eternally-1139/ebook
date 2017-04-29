{{template "admin/base.tpl" .}}
{{define "meta"}}
<title>书籍分类-商城管理系统</title>
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
            <h5>书籍分类 <small>分类，查找</small></h5>
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
                    <th>分类名称</th>
                    <th>分类图片</th>
                    <th>分类类型</th>
                    <th>修改</th>
                    <th>删除</th>
                </tr>
                </thead>
                <tbody>
                {{range .category}}
                <tr class="gradeX">
                    <td>{{.Id}}</td>
                    <td>{{.Name}}</td>
                    <td><img src="{{.Image}}" alt="" class="img-rounded" width="60px"></td>
                    <td>{{if eq .Status 0}}普通分类{{else}}特殊分类{{end}}</td>
                    <td class="center" style="text-align: center"><button class="btn btn-default btn-sm" onclick="openChangeModal({{.Name}},{{.Id}})"><i class="fa fa-edit"></i> 修改</button></td>
                    <td class="center" style="text-align: center"><button class="btn btn-danger btn-sm" onclick="deleteCategory({{.Id}})"><i class="fa fa-trash"></i> 删除</button></td>
                </tr>
                {{end}}
                </tbody>
            </table>
            <div style="text-align: center;">
                <button class="btn btn-info" data-toggle="modal" data-target="#myModal5">新增分类</button>
            </div>
        </div>
    </div>
</div>
</div>

{{end}}

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
                        <input type="email" placeholder="请输入分类名称" class="form-control" id="cateName">
                    </div>
                    <div class="form-group">
                        <label>选择图片</label>
                        <input type="file" class="form-control" id="file" onchange="coverUpload()">
                        <input type="hidden" id="cover">
                        <img src="" id="cover_img" style="max-height: 240px">
                    </div>
                    <div class="form-group">
                        <label>分类类型</label>
                        <select class="form-control" id="cateType">
                            <option>请选择分类类别</option>
                            <option value="0">普通分类</option>
                            <option value="1">特殊分类</option>
                        </select>
                    </div>
                </form>

            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-primary" onclick="createSubmit();">提交</button>
            </div>
        </div>
    </div>
</div>

<div class="modal fade" id="myModal6" tabindex="-1" role="dialog"  aria-hidden="true">
    <div class="modal-dialog modal-lg">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">修改分类信息</h4>
            </div>
            <div class="modal-body">
                <form role="form">
                    <div class="form-group">
                        <label>分类名称</label>
                        <input type="email" placeholder="请输入分类名称" class="form-control" id="changeName" value="">
                        <input type="text" class="hidden" value="" id="changeCateId">
                    </div>
                </form>

            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-primary" onclick="changeCate();">提交</button>
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
    function openChangeModal(name,id) {
        $("#changeName").val(name);
        $("#changeCateId").val(id);
        $('#myModal6').modal('show')
    }
    function changeCate() {
        var changeName = $("#changeName").val();
        var changeCateId = $("#changeCateId").val();
        $.ajax({
            url: '/admin/changeCategory',
            type: 'post',
            dataType: 'json',
            data:{
                'id':changeCateId,
                'changeName':changeName
            },
            contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
            cache: false,
            success: function(data) {
                if(data.status==10000){
                    window.location.href="/admin/category";
                }else if(data.status==10002){
                    alert(data.message);
                    // window.location.reload();
                }else{
                    alert('您的网络异常!!！');
                }

            },
            error : function() {
                alert('服务器失去响应，请检查您的网络！');


            }
        })
    }

    function deleteCategory(id) {

        $.ajax({
            url: '/admin/deleteCategory',
            type: 'get',
            dataType: 'json',
            data:{
                'id':id
            },
            contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
            cache: false,
            success: function(data) {
                if(data.status==10000){
                    window.location.href="/admin/category";
                }else if(data.status==10002){
                    alert(data.message);
                    // window.location.reload();
                }else{
                    alert('您的网络异常!!！');
                }

            },
            error : function() {
                alert('服务器失去响应，请检查您的网络！');


            }
        })
    }
    function createSubmit() {
            var cateName = $("#cateName").val();
            var cateType = $("#cateType").val();
            var cover =$('#cover').val();

        if(cateName==""||cateName==""||cover==""){
                alert('请填写完整信息！');
            }else{
                $.ajax({
                    url: '/admin/addCategory',
                    type: 'get',
                    dataType: 'json',
                    data:{
                        'cateName':cateName,
                        'image':cover,
                        'cateType':cateType
                    },
                    contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
                    cache: false,
                    success: function(data) {
                        if(data.status==10000){
                            window.location.href="/admin/category";
                        }else if(data.status==10002){
                            alert(data.message);
                            // window.location.reload();
                        }else{
                            alert('您的网络异常!!！');
                        }

                    },
                    error : function() {
                        alert('服务器失去响应，请检查您的网络！');
                    }
                });
            }
    }
    function coverUpload() {
        var form_data = new FormData();
        var file = $("#file").prop("files")[0];
        if (!file){
            return;
        }
        form_data.append("file", file);
        $.ajax({
            type: "POST",
            url: '/file/upload',
            dataType : "json",
            processData:false,
            contentType:false,
            data:form_data,
            success:function(data) {
                if(data.status==10000){
                    $("#cover").val(data.src);
                    $("#cover_img").attr("src",data.src)
                }else{
                    alert("上传错误");
                }
            },
            error : function() {
                alert("无法连接服务器");
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