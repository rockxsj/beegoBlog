<div class="container">
        <div class="navbar-header">
          <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="navbar-brand" href="#">{{.Options.title}}</a>
        </div>
        <div class="navbar-collapse collapse">
          <ul class="nav navbar-nav">
            <li{{if .Is_index}} class="active"{{end}}><a href="/">首页</a></li>
            <li class="dropdown">
              <a href="#" class="dropdown-toggle" data-toggle="dropdown">关于我 <b class="caret"></b></a>
              <ul class="dropdown-menu">
                <li><a href="http://yazi.whoknow.me">北城</a></li>
                <li class="divider"></li>
                <li><a href="https://twitter.com/rockxsj" target="_blank">我的Twitter</a></li>
                <li><a href="http://weibo.com/rockxsj" target="_blank">新浪微博</a></li>
                <li><a hrer="http://1.t.qq.com/rockxsj" target="_blank">腾讯微博</a></li>
              </ul>
            </li>
          </ul>
          <ul class="nav navbar-nav navbar-right">
            {{if .Is_login}}
              <li><a href="/admin">{{.Username}}</a></li>
              <li><a href="/login/logout">退出</a></li>
            {{else}}
              <li><a href="/login">登录</a></li>
            {{end}}
          </ul>
        </div><!--/.nav-collapse -->
      </div>