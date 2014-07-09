
<!DOCTYPE html>
<html lang="zh-cn">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="">
    <meta name="author" content="">

    <title>{{.Page_title}}</title>
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/blog.css" rel="stylesheet">
    <style>
      body {
        min-height: 800px;
        padding-top: 70px;
      }
    </style>
  </head>

  <body>

    <!-- Fixed navbar -->
    <div class="navbar navbar-default navbar-fixed-top" role="navigation">
      {{template "top.tpl" .}}
    </div>

    <div class="container">

      {{.LayoutContent}}

    </div> <!-- /container -->
    <div class="blog-footer">
      <p>基于beego开发的博客程序 作者@<a href="http://blog.whoknow.me" target="_blank">北城</a></p>
      <p>
        <a href="#">Back to top</a>
      </p>
    </div>
    <script src="/static/js/jquery.min.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>
  </body>
</html>
