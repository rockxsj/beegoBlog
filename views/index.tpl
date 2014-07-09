{{range .Posts}}
	<div class="jumbotron">
	  <h1>{{.Title}}</h1>
	  <p>{{.Content}}</p>
	  <p>
	    <a class="btn btn-lg btn-primary" href="/post/{{.Id}}" role="button">阅读全文 &raquo;</a>
	  </p>
	</div>
{{end}}
<ul class="pager">
  <li class="previous{{if .NoPre}} disabled{{end}}">
  	<a href={{if .NoPre}}"javascript:void(0)"{{else}}"/list/{{.PrePage}}"{{end}}>
  		{{if .NoPre}}没有了{{else}}上一页{{end}}
  	</a>
  </li>
  <li class="next{{if .NoNext}} disabled{{end}}">
  	<a href={{if .NoNext}}"javascript:void(0)"{{else}}"/list/{{.NextPage}}"{{end}}>
  		{{if .NoNext}}没有了{{else}}下一页{{end}}
  	</a>
  </li>
</ul>