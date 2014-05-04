<div class="page-header">
	<h1>{{.Post.Title}}</h1>
	<p class="lead">{{str2html .Post.Content}}</p>
</div>
<ul class="pager">
  <li class="previous{{if .NoPre}} disabled{{end}}">
  	<a href={{if .NoPre}}"javascript:void(0)"{{else}}"/post/{{.Pre.Id}}"{{end}}>
  		{{if .NoPre}}没有了{{else}}&larr; {{.Pre.Title}}{{end}}
  	</a>
  </li>
  <li class="next{{if .NoNext}} disabled{{end}}">
  	<a href={{if .NoNext}}"javascript:void(0)"{{else}}"/post/{{.Next.Id}}"{{end}}>
  		{{if .NoNext}}没有了{{else}}&rarr; {{.Next.Title}}{{end}}
  	</a>
  </li>
</ul>