<table class="table table-hover">
  <thead>
        <tr>
          <th>链接名</th>
          <th>链接地址</th>
          <th>删除</th>
        </tr>
      </thead>
      <tbody>
      	{{range .Links}}
	        <tr>
	          <td>
              {{.Name}}
            </td>
	          <td>
              {{.Link}}
	          </td>
	          <td>
	          	<button type="button" onclick="BLOG_ADMIN.delLink($(this), {{.Id}});" class="btn btn-danger btn-xs">删除</button>
	          </td>
	        </tr>
         {{end}}
      </tbody>
</table>
<form class="form-inline" role="form" method="post" action="/admin/AddLink">
  <div class="row">
    <div class="col-xs-4">
      <input type="text" name="name" class="form-control" placeholder="链接名称">
    </div>
    <div class="col-xs-4">
      <input type="text" name="link" class="form-control" placeholder="链接地址，必须http://或者https://开头">
    </div>
    <div class="col-xs-2">
      <button type="submit" class="btn btn-default">添加</button>
    </div>
  </div>
</form>