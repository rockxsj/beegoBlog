var BLOG_FRONT = {}

BLOG_FRONT.cateActive = function(cid) {
	$(".blog-nav-item.cate").each(function(){
		if (cid == $(this).attr("cid")) {
			$(this).addClass("active")
			return true
		}
	})
}

$(document).ready(function() {
	BLOG_FRONT.cateActive($("#blog-top").attr("cid"))	//设置当前标签为选中状态
})