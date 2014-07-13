<div class="col-xs-6 col-sm-3 sidebar-offcanvas" id="sidebar" role="navigation">
  <div class="list-group">
    <a href="/admin" class="list-group-item{{if .Is_index}} active{{end}}">控制面板</a>
    <a href="/admin/GetCates" class="list-group-item{{if .Is_cate}} active{{end}}">分类管理</a>
    <a href="/admin/list" class="list-group-item{{if .Is_list}} active{{end}}">文章列表</a>
    <a href="/admin/add" class="list-group-item{{if .Is_add}} active{{end}}">增加文章</a>
    <a href="/admin/GetLinks" class="list-group-item{{if .Is_link}} active{{end}}">友情链接</a>
    <a href="/admin/option" class="list-group-item{{if .Is_option}} active{{end}}">博客设置</a>
  </div>
</div><!--/span-->