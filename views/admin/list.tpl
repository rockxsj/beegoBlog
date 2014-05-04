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