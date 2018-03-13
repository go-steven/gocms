// 初始化左边菜单tree
$(document).ready(function(){
    $.post('/loadMenu', {}, function(result){
        // zNodes = result
        $.fn.zTree.init($('#left_menu_tree'), {
            data: {
                simpleData: {
                    enable: true
                }
            }
            // ,
            // callback: {
            // 	onClick: zTreeOnClick
            // }
        }, result);
    });
});
