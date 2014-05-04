{{range .Posts}}
	<div class="jumbotron">
	  <h1>{{.Title}}</h1>
	  <p>{{.Content}}</p>
	  <p>
	    <a class="btn btn-lg btn-primary" href="/post/{{.Id}}" role="button">阅读全文 &raquo;</a>
	  </p>
	</div>
{{end}}