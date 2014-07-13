
<!DOCTYPE html>
<html lang="zh-cn">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="">
    <meta name="author" content="">

    <title>{{.Options.title}}|{{.Options.secondtitle}}</title>
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/blog.css" rel="stylesheet">
  </head>

  <body>

    {{template "top.tpl" .}}
    <div class="container">

      <div class="blog-header">
        <h1 class="blog-title">{{.Options.title}}</h1>
        <p class="lead blog-description">{{.Options.secondtitle}}</p>
      </div>

      <div class="row">

        <div class="col-sm-8 blog-main">
        {{range .Posts}}
          <div class="blog-post">
            <h2 class="blog-post-title">{{.Title}}</h2>
            <p class="blog-post-meta">{{date .Add_time "Y-m-d H:i:s"}}</p>
            {{str2html .Content}}
          </div><!-- /.blog-post -->
        {{end}}

          <ul class="pager">
            {{template "pager.tpl" .}}
          </ul>

        </div><!-- /.blog-main -->

        <div class="col-sm-3 col-sm-offset-1 blog-sidebar">
          {{template "right.tpl" .}}
        </div><!-- /.blog-sidebar -->

      </div><!-- /.row -->

    </div>
    <div class="blog-footer">
      {{template "footer.tpl" .}}
    </div>
    <script src="/static/js/jquery.min.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>
    <script src="/static/js/docs.min.js"></script>
    <script src="/static/js/blog.js"></script>
  </body>
</html>
