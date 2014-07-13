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
				  	<input type="text" class="form-control input-sm" value="{{.Name}}" onchange="BLOG_ADMIN.renameCate($(this), {{.Id}})">
				  </div>
	          </td>
	          <td>
	          	<button type="button" onclick="BLOG_ADMIN.delCate($(this), {{.Id}});" class="btn btn-danger btn-xs">删除</button>
	          </td>
	        </tr>
         {{end}}
      </tbody>
</table>
<form class="form-inline" role="form" method="post" action="/admin/AddCate">
  <div class="form-group">
    <label class="sr-only" for="addCate">添加分类</label>
    <input type="text" name="cate" class="form-control" id="addCate" placeholder="分类名">
  </div>
  <button type="submit" class="btn btn-default">添加</button>
</form>
