<form class="form-inline" role="form" method="post" action="/admin/ChangePassword">
  <div class="row">
    <div class="col-xs-3">
      <input type="password" id="oldPassword" class="form-control" placeholder="老密码">
    </div>
    <div class="col-xs-3">
      <input type="password" id="newPassword" class="form-control" placeholder="新密码">
    </div>
    <div class="col-xs-3">
      <input type="password" id="reNewPassword" class="form-control" placeholder="重复新密码">
    </div>
    <div class="col-xs-2">
      <button type="button" class="btn btn-default" onclick="BLOG_ADMIN.changePassword();">修改</button>
    </div>
  </div>
</form>