{{if .NoPre}}
  <li><a class="btn btn-default btn-lg disabled" role="button" href="javascript:void(0)">没有了</a></li>
{{else}}
  <li><a class="btn btn-default btn-lg" role="button" href="{{.BaseUrl}}/{{.PrePage}}">上一页</a></li>
{{end}}
{{if .NoNext}}
  <li><a class="btn btn-default btn-lg disabled" role="button" href="javascript:void(0)">没有了</a></li>
{{else}}
  <li><a class="btn btn-default btn-lg" role="button" href="{{.BaseUrl}}/{{.NextPage}}">下一页</a></li>
{{end}}