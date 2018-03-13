//此方法 用于测试 不用每次进来都要点一次左侧菜单
$(document).ready(function(){
    //addTab('系统用户管理auto','user/listview')
    //addTab('权限管理auto','role/listview')
    //addTab('系统用户组管理auto','usergroup/listview')
    //addTab('添加系统用户组管理auto','usergroup/addview')
    //addTab('添加系统用户auto','/user/addview')
});
//添加标签页的方法
function addTab(title, url) {
    if ($('#main_tab').tabs('exists', title)) {
        $('#main_tab').tabs('select', title);
    } else {
        //var content = '<iframe scrolling="auto" frameborder="0"  src="'+url+'" style="width:100%;height:100%;"></iframe>';
        $('#main_tab').tabs('add', {
            title: title,
            //content:content,
            href: url,
            closable: true
        });
    }
}

/**
 * 格式化时间
 */
function dataformatter(value, row, index) {
    return value ? phpjs.date('Y-m-d H:i:s', phpjs.strtotime(value)) : value;
}