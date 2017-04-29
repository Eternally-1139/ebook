{{template "admin/base.tpl" .}}
{{define "meta"}}
<title>书籍列表-商城管理系统</title>
{{end}}

{{define "css"}}
<link href="/web-static/admin/css/font-awesome.css?v=4.4.0" rel="stylesheet">

<!-- Data Tables -->
<link href="/web-static/admin/css/plugins/dataTables/dataTables.bootstrap.css" rel="stylesheet">

<link href="/web-static/admin/css/animate.css" rel="stylesheet">
<link href="/web-static/admin/css/style.css?v=4.1.0" rel="stylesheet">

<!--picture upload-->
<style>
</style>
{{end}}

{{define "content"}}
<div class="row">
    <div class="col-sm-12">
        <div class="ibox float-e-margins">
            <div class="ibox-title">
                <h5>书籍列表 <small>分类，查找</small></h5>
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
                        <th>封面</th>
                        <th>书籍名称</th>
                        <th>所属分类</th>
                        <th>价格</th>
                        <th>上架时间</th>
                        <th>书籍介绍</th>
                        <th>修改</th>
                        <th>删除</th>
                        <th>上传图片</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{range .product}}
                    <tr class="gradeX">
                        <th>{{.Id}}</th>
                        <th><img class="img-rounded" src="{{.Image}}" width="40" height="40" alt=""></th>
                        <th>{{.Name}}</th>
                        <th>{{if eq .Category.Status 0}}普通分类{{else}}特殊分类{{end}}-{{.Category.Name}}</th>
                        <th>{{.Price}}</th>
                        <th>{{date .CreateTime "Y-m-d h:i:s"}}</th>
                        <th>{{.Content}}</th>
                        <td class="center" style="text-align: center"><button class="btn btn-default btn-sm" onclick="openChangeModal({{.Name}},{{.Id}},{{.Image}},{{.Price}},{{.Content}})"><i class="fa fa-edit"></i> 编辑信息</button></td>
                        {{if eq .IsImage 0}}
                        <td class="center" style="text-align: center"><button class="btn btn-warning btn-sm" onclick="deleteHtml();openUploadModal({{.Id}})"><i class="fa fa-upload"></i> 未上传产品图片</button></td>
                        {{else}}
                        <td class="center" style="text-align: center"><button class="btn btn-success btn-sm" onclick="showImageModal({{.Id}})"><i class="fa fa-eye"></i> 查看产品图片</button></td>
                        {{end}}
                        <td class="center" style="text-align: center"><button class="btn btn-danger btn-sm" onclick="deleteProduct({{.Id}})"><i class="fa fa-trash"></i> 删除</button></td>

                    </tr>
                    {{end}}
                    </tbody>
                </table>
                <div style="text-align: center;">
                    <button class="btn btn-info" data-toggle="modal" data-target="#myModal5">新增书籍 <span class="glyphicon glyphicon-plus"></span></button>
                </div>
            </div>
        </div>
    </div>
</div>

{{end}}
<div class="modal fade" id="myModal5" tabindex="-1" role="dialog"  aria-hidden="true">
    <div class="modal-dialog modal-lg" style="margin-top: 7%!important;">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">新增书籍</h4>
            </div>
            <div class="modal-body">
                <form role="form">
                    <div class="form-group">
                        <label>书籍名称</label>
                        <input type="email" placeholder="请输入书籍名称" class="form-control" id="proName">
                    </div>
                    <div class="form-group">
                        <label>所属分类</label>
                        <select class="form-control" name="" id="proType">
                            <option>请选择分类类别</option>
                            {{range .category}}
                            <option value="{{.Id}}">{{if eq .Status 0}}普通分类{{else}}特殊分类{{end}}-{{.Name}}</option>
                            {{end}}
                        </select>
                    </div>
                    <div class="form-group">
                        <label>书籍封面</label>
                        <input type="file" class="form-control" id="file" onchange="coverUpload()">
                        <input type="hidden" id="cover">
                        <img src="" id="cover_img" style="max-height: 240px">
                    </div>
                    <div class="form-group">
                        <label>价格</label>
                        <input type="email" placeholder="请输入书籍价格" class="form-control" id="proPrice">
                    </div>
                    <div class="form-group">
                        <label>书籍介绍</label>
                        <input type="email" placeholder="请输入书籍介绍" class="form-control" id="proContent">
                    </div>
                </form>

            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-primary" onclick="createSubmitProduct();">提交</button>
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
                        <label>书籍名称</label>
                        <input type="email" placeholder="请输入书籍名称" class="form-control" id="changeName" value="">
                    </div>
                    <div class="form-group">
                        <label>书籍封面</label>
                        <input type="file" class="form-control" id="file2" onchange="coverUpload2()">
                        <input type="hidden" id="cover2">
                        <img src="" id="cover_img2" style="max-height: 240px">
                    </div>
                    <div class="form-group">
                        <label>价格</label>
                        <input type="email" placeholder="请输入书籍价格" class="form-control" id="changePrice" value="">
                    </div>
                    <div class="form-group">
                        <label>书籍介绍</label>
                        <input type="email" placeholder="请输入书籍介绍" class="form-control" id="changeContent" value="">
                    </div>
                    <input type="text" class="hidden" id="changeId" value="">
                </form>

            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-white" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-primary" onclick="changeProduct();">提交</button>
            </div>
        </div>
    </div>
</div>

<div class="modal fade" id="UploadModal" tabindex="-1" role="dialog"  aria-hidden="true">
    <div class="modal-dialog modal-lg" style="width: 80%">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">上传产品详细图片</h4>
            </div>
            <div class="modal-body">

                    <div class="uploadBox">
                        <div class="container">
                        <div class="row">
                            <div class="col-xs-6"><h3>已上传 <span id="numbers">0</span> 张图片</h3></div>
                            <div class="col-xs-6"><button class="btn btn-info" id="uploadButton" onclick="uploadImageButton()"><i class="fa fa-upload"></i> 点击上传</button></div>
                        </div>
                        <hr>
                        <div id="showImage">

                        </div>
                        </div>
                    </div>

                <input type="text" class="hidden" value="" id="imageHiddenId">
                <form>
                    <div id="imageInput">
                    </div>
                </form>
            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-primary" onclick="closeUpload()">完成</button>
            </div>
        </div>
    </div>
</div>


<div class="modal fade" id="showPicModal" tabindex="-1" role="dialog"  aria-hidden="true">
    <div class="modal-dialog modal-lg" style="width: 80%">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                <h4 class="modal-title">上传产品详细图片</h4>
            </div>
            <div class="modal-body">
                <div class="container">
                    <div class="uploadBox">
                        <div id="showImage1">

                        </div>
                    </div>
                </div>
            </div>

            <div class="modal-footer">
                <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
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


<!-- Page-Level Scripts -->


<script>

    var numbers = 0;

    function showImageModal(id) {
        $("#showImage1").html("");
        $.ajax({
            url: '/admin/product/showProduct',
            type: 'get',
            data:{
                'proId':id
            },
            dataType: 'json',
            contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
            cache: false,
            success: function(data) {
                $('#showPicModal').modal('show');
                if(data.status==10000){
                    for(var i=0;i<data.pic.length;i++){
                        var pro = '<div class="col-md-3" style="margin-bottom: 30px;padding: 30px">'+' <img'+' src=' +data.pic[i].Image+' alt="" width="100%" height="auto" class="img-rounded">'+' </div>';
                        $("#showImage1").append(pro);
                    }

                }else{
                    alert("获取产品图片失败");
                }
            },
            error : function() {
                alert("无法连接服务器");
            }
        });

    }

    function deleteHtml() {
        $("#showImage").html(" ");
    }
    function numberShow() {
        numbers++;
        $("#numbers").html(numbers);
    }
    function uploadImageButton() {
        var fileInput='<input'+' class="hidden"'+' id="clickInput"'+' type="file"'+' onchange="coverUpload3()"'  + '>';
        $("#imageInput").html(fileInput);
        $("#clickInput").click()

    }
    function openUploadModal(id) {
        $("#imageHiddenId").val(id);
        $('#UploadModal').modal('show');

    }
    function openChangeModal(name,id,image,price,content) {
        $("#changeName").val(name);
        $("#changeId").val(id);
        $("#changePrice").val(price);
        $("#changeContent").val(content);
        $("#cover_img2").attr("src",image);

        $('#myModal6').modal('show')
    }
    function changeProduct() {
        var changeName = $("#changeName").val();
        var changePrice = $("#changePrice").val();
        var changeId = $("#changeId").val();
        var changeContent = $("#changeContent").val();
        var cover2 = $("#cover2").val();
        $.ajax({
            url: '/admin/changeProduct',
            type: 'post',
            dataType: 'json',
            data:{
                'changeName':changeName,
                'changePrice':changePrice,
                'changeId':changeId,
                'cover2':cover2,
                'changeContent':changeContent
            },
            contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
            cache: false,
            success: function(data) {
                if(data.status==10000){
                    window.location.href="/admin/product";
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
    function createSubmitProduct() {

        proName=$('#proName').val();
        cover =$('#cover').val();
        proContent=$('#proContent').val();
        proPrice=$('#proPrice').val();
        proType=$('#proType').val();

        if (!proName||!proContent||!cover){
            alert("标题和内容不能为空");
            return
        }


        $.ajax({
            url: '/admin/product/upload',
            type: 'post',
            data:{
                'proName':proName,
                'cover':cover,
                'proPrice':proPrice,
                'proContent':proContent,
                'proType':proType
            },
            dataType: 'json',
            contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
            cache: false,
            success: function(data) {
                if(data.status==10000){
                    location.href='/admin/product';
                }else{
                    alert("保存失败");
                }
            },
            error : function() {
                alert("无法连接服务器");
            }
        });

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

    function coverUpload2() {
        var form_data = new FormData();
        var file = $("#file2").prop("files")[0];
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
                    $("#cover2").val(data.src);
                    $("#cover_img2").attr("src",data.src)
                }else{
                    alert("上传错误");
                }
            },
            error : function() {
                alert("无法连接服务器");
            }
        })

    }

    function coverUpload3() {
        var form_data = new FormData();
        var file = $("#clickInput").prop("files")[0];
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
                    //$("#cover2").val(data.src);
                    var pic = '<div class="col-md-3" style="margin-bottom: 30px;padding: 30px">'+' <img'+' src=' +data.src+' alt="" width="100%" height="auto" class="img-rounded">'+' </div>';
                    $("#showImage").append(pic);
                    picUpload(data.src);
                    numberShow();
                    $("#uploadButton").html("继续上传");
                    $("#uploadButton").removeClass();
                    $("#uploadButton").addClass("btn btn-default");
                }else{
                    alert("上传错误");
                }
            },
            error : function() {
                alert("无法连接服务器");
            }
        })

    }
    function picUpload(imgSrc) {
        var proId = $("#imageHiddenId").val();
        $.ajax({
            url: '/admin/picUpload',
            type: 'post',
            dataType: 'json',
            data:{
                'proId':proId,
                'imgSrc':imgSrc
            },
            contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
            cache: false,
            success: function(data) {
                if(data.status==10000){

                }else if(data.status==10001){
                    alert(data.message);
                    // window.location.reload();
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

    function closeUpload() {
        var proId = $("#imageHiddenId").val();
        var isImage = $("#showImage").html();
        if(numbers==0){
            alert("请至少上传一张产品图片！")
        }else{
            $.ajax({
                url: '/admin/closeUpload',
                type: 'post',
                dataType: 'json',
                data:{
                    'proId':proId
                },
                contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
                cache: false,
                success: function(data) {
                    if(data.status==10000){
                        window.location.reload();
                    }else if(data.status==10001){
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

    }

    function deleteProduct(id) {

        $.ajax({
            url: '/admin/deleteProduct',
            type: 'post',
            dataType: 'json',
            data:{
                'id':id
            },
            contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
            cache: false,
            success: function(data) {
                if(data.status==10000){
                    window.location.href="/admin/product";
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

</script>
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