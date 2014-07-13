<div class="sidebar-module sidebar-module-inset">
  <h4>关于我</h4>
  {{str2html .Options.intro}}
</div>
<div class="sidebar-module">
  <h4>过去博文</h4>
  <ol class="list-unstyled">
  {{range .MonthList}}
    <li><a href="/month/{{.}}/1">{{.}}</a></li>
  {{end}}
  </ol>
</div>
<div class="sidebar-module">
  <h4>友情链接</h4>
  <ol class="list-unstyled">
    {{range .Links}}
    <li><a href="{{.Link}}" target="_blank">{{.Name}}</a></li>
    {{end}}
  </ol>
</div>