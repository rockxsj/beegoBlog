<div class="row">
  {{range .Posts}}
    <div class="col-6 col-sm-6 col-lg-4">
      <h2>{{.Title}}</h2>
      <p>{{.Content}} </p>
      <div class="btn-group btn-group-sm">
        <a class="btn btn-default" role="button" href="/admin/update/{{.Id}}">更新 &raquo;</a>
        <a class="btn btn-default" role="button" href="/admin/del/{{.Id}}">删除</a>
      </div>
    </div><!--/span-->
  {{end}}
</div><!--/row-->
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