<form class="form-horizontal" role="form" action="/admin/SetOptions" method="post">
  <div class="form-group">
    <label for="siteTitle" class="col-sm-2 control-label">站点标题</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" name="title" value="{{.Options.title}}" id="siteTitle" placeholder="GoBlog">
    </div>
  </div>
  <div class="form-group">
    <label for="siteSecondTitle" class="col-sm-2 control-label">站点副标题</label>
    <div class="col-sm-10">
      <input type="text" class="form-control" name="secondtitle" value="{{.Options.secondtitle}}" id="siteSecondTitle" placeholder="A Blog Of Go">
    </div>
  </div>
  <div class="form-group">
    <div class="col-sm-offset-2 col-sm-10">
      <button type="submit" class="btn btn-default">提交修改</button>
    </div>
  </div>
</form>
