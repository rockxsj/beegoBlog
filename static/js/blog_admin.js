var BLOG_ADMIN = {}

BLOG_ADMIN.delLink = function(t, id) {
	$.post("/admin/DelLink",{"id" : id}, function (data) {
		if (data.success == "1") {
			t.parent().parent().remove()
		} else {
			alert("删除失败！")
		}
	}, "json")
}

BLOG_ADMIN.delCate = function (t, id) {
	$.post("/admin/DelCate",{"id" : id}, function (data) {
		if (data.success == "1") {
			t.parent().parent().remove()
		} else {
			alert("删除失败！")
		}
	}, "json")
}

BLOG_ADMIN.renameCate = function (t, id) {
	var newName = t.val()
	$.post("/admin/RenameCate", {"id" : id, "NewName" : newName}, function (data) {
		if (data.success == "1") {
			alert("更新成功！")
		} else {
			alert("更新失败！")
		}
	}, "json")
}

BLOG_ADMIN.changePassword = function() {
	var oldPassword = $("#oldPassword").val()
	var newPassword = $("#newPassword").val()
	var reNewPassword = $("#reNewPassword").val()
	if (newPassword != reNewPassword) {
		alert("两次新密码输入不同，请确认！")
		return false
	}
	$.post("/admin/ChangePassword", {"oldPassword" : oldPassword, "newPassword": newPassword}, function(data) {
		if (data.success == "1") {
			alert("更新成功！")
			location.reload()
		} else {
			alert("更新失败！")
		}
	}, "json")
}