<table class="table table-hover">
  <thead>
        <tr>
          <th>Id</th>
          <th>分类名</th>
          <th>删除</th>
        </tr>
      </thead>
      <tbody>
      	{{range .Cates}}
	        <tr>
	          <td>{{.Id}}</td>
	          <td>
	          	  <div class="col-xs-4">
				  	<input type="text" class="form-control input-sm" value="{{.Name}}" onchange="renameCate($(this), {{.Id}})">
				  </div>
	          </td>
	          <td>
	          	<button type="button" onclick="delCate($(this), {{.Id}});" class="btn btn-danger btn-xs">删除</button>
	          </td>
	        </tr>
         {{end}}
      </tbody>
</table>

<script>
var delCate = function (t, id) {
	$.post("/admin/DelCate",{"Id" : id}, function (data) {
		if (data.success == "1") {
			t.parent().parent().remove()
		} else {
			alert("删除失败！")
		}
	}, "json")
}

var renameCate = function (t, id) {
	var newName = t.val()
	$.post("/admin/RenameCate", {"Id" : id, "NewName" : newName}, function (data) {
		if (data.success == "1") {
			alert("更新成功！")
		} else {
			alert("更新失败！")
		}
	}, "json")
}
</script>