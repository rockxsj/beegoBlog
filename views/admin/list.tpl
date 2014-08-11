<table class="table table-striped">
    <thead>
      <tr>
        <th>标题</th>
        <th>分类</th>
        <th>更新</th>
        <th>删除</th>
      </tr>
    </thead>
    <tbody>
      {{range .Posts}}
        <tr>
          <td>{{.title}}</td>
          <td>{{.cate_name}}</td>
          <td><a class="btn btn-default" role="button" href="/admin/update/{{.id}}">更新 &raquo;</a></td>
          <td><a class="btn btn-default" role="button" href="/admin/del/{{.id}}">删除</a></td>
        </tr>
       {{end}}
    </tbody>
</table>

<ul class="pager">
  <li class="previous{{if .NoPre}} disabled{{end}}">
    <a href={{if .NoPre}}"javascript:void(0)"{{else}}"/admin/list/{{.PrePage}}"{{end}}>
      {{if .NoPre}}没有了{{else}}上一页{{end}}
    </a>
  </li>
  <li class="next{{if .NoNext}} disabled{{end}}">
    <a href={{if .NoNext}}"javascript:void(0)"{{else}}"/admin/list/{{.NextPage}}"{{end}}>
      {{if .NoNext}}没有了{{else}}下一页{{end}}
    </a>
  </li>
</ul>