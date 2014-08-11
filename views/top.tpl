<div class="blog-masthead">
  <div class="container">
    <nav class="blog-nav" id="blog-top" cid="{{.Cid}}">
      <a class="blog-nav-item{{if .Is_index}} active{{end}}" href="/">首页</a>
      {{range .Cates}}
        <a class="blog-nav-item cate" href="/cate/{{.Id}}/1" cid="{{.Id}}">{{.Name}}</a>
      {{end}}
      <a class="blog-nav-item cate" href="/cate/0/1" cid="0">未分类</a>
      {{if .Is_login}}
      <a class="blog-nav-item fright" href="/login/logout">退出</a>
      <a class="blog-nav-item fright" href="/admin">{{.Username}}</a>
      {{else}}
      <a class="blog-nav-item fright" href="/login">登录</a>
      {{end}}
    </nav>
  </div>
</div>