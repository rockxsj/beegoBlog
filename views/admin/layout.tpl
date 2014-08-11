
<!DOCTYPE html>
<html lang="zh-cn">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="">
    <meta name="author" content="">
    <title>{{.Options.title}}-后台</title>
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/offcanvas.css" rel="stylesheet">
    <script src="/static/js/jquery.min.js"></script>
  </head>

  <body>
    {{template "admin/top.tpl" .}}

    <div class="container">

      <div class="row row-offcanvas row-offcanvas-right">

        {{template "admin/left.tpl" .}}

        <div class="col-xs-12 col-sm-9">
          {{.LayoutContent}}
        </div><!--/span-->

      </div><!--/row-->

      <hr>

      {{template "admin/foot.tpl" .}}

    </div><!--/.container-->
  </body>
</html>
