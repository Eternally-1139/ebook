{{template "admin/base.tpl" .}}
{{define "meta"}}
<title>轮播图管理-商城管理系统</title>
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
                <h5>轮播图管理 <small>分类，查找</small></h5>
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
                        <th>内容</th>
                        <th>图片</th>
                        <th>删除</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range .carousel}}
                    <tr class="gradeX">
                        <td>{{.Id}}</td>
                        <td>{{.Title}}</td>
                        <td style="text-align: center"><img src="{{.Img}}" alt="" class="img-rounded" width="200px"></td>
                        <td class="center" style="text-align: center"><button class="btn btn-danger btn-sm" onclick="deleteIt({{.Id}})"><i class="fa fa-book"></i> 删除</button></td>
                    </tr>
                    {{end}}
                    </tbody>
                </table>
                <div style="text-align: center;">
                    <button class="btn btn-info" data-toggle="modal" data-target="#addCarousel">添加轮播图</button>
                </div>
            </div>
        </div>
    </div>
</div>

{{end}}

<div class="modal fade" id="addCarousel" tabindex="-1" role="dialog"  aria-hidden="true">
    <div class="modal-dialog modal-lg">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">添加轮播图</h4>
            </div>
            <div class="modal-body">
                <form role="form">
                    <div class="form-group">
                        <label>轮播图备注</label>
                        <input type="email" placeholder="输入轮播图备注" class="form-control" id="content">
                    </div>
                    <div class="form-group">
                        <label>选择图片</label>
                        <input type="file" class="form-control" id="file" onchange="coverUpload()">
                        <input type="hidden" id="cover">
                        <img src="" id="cover_img" style="max-height: 240px">
                    </div>
                    <div class="form-group">
                        <label>选择链接的产品</label>
                        <select class="form-control" id="href">
                            <option>请选择要链接的产品</option>
                            {{range .products}}
                            <option value="{{.Id}}">{{.Name}}-{{.Category.Name}}</option>
                            {{end}}
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

    function createSubmit() {

        content=$('#content').val();
        cover =$('#cover').val();
        href=$('#href').val();

        if (!content||!cover||!href){
            alert("请填写完整信息！");
            return
        }


        $.ajax({
            url: '/admin/addCarousel',
            type: 'post',
            data:{
                'content':content,
                'image':cover,
                'href':href
            },
            dataType: 'json',
            contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
            cache: false,
            success: function(data) {
                if(data.status==10000){
                    location.href='/admin/carousel';
                }else if(data.status==10002){
                    alert(data.message)
                }
                else{
                    alert("上传失败");
                }
            },
            error : function() {
                alert("无法连接服务器");
            }
        });

    }
</script>

<script>
    function deleteIt(id) {
        $.ajax({
            url: '/admin/deleteCarousel',
            type: 'post',
            data:{
                'id':id
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