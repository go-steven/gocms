
//设置tree的初始化参数
var user_group_setting = {
    check: {
        enable: true
        // chkboxType: { "Y": "", "N": "" }
    },
    data: {
        simpleData: {
            enable: true
        }
    }
};




//初始化左边tree
$(document).ready(loadTree());
function loadTree() {
    var url = "/usergroup/loadtreewithoutroot";
    $.post(url, {}, function (result) {
        // zNodes = result
        $.fn.zTree.init($("#add_group_role_tree"), user_group_setting, result);
    });
}

function submitAddUserGroupForm() {
    var zTree = $.fn.zTree.getZTreeObj("add_group_role_tree");
    var nodes = zTree.getCheckedNodes(true);
    var checkCount = nodes.length;
    //判断选中的节点数，如果没有选中节点则提示操作错误
    if (checkCount === 0) {
        $.messager.alert('操作提示', "请至少选择一个权限", 'info');
        return false;
    }
    //获取所有选中的节点ID
    var idArray = new Array(checkCount);
    for (var i = 0; i < nodes.length; i++) {
        idArray[i] = nodes[i].id
    }
    var ids = idArray.join(",");
    var url = "/usergroup/addusergroup";

    var data = {
        ids: ids,
        group_name: $("input[name='group_user_name']").val(),
        remarks: $("input[name='group_user_remarks']").val()
    };

    $.post(url, data, function (result) {
        if (result === "success") {
            $('#add_user_group').window("close");
            $.messager.alert('操作提示', "添加成功", 'info');
            loadUserGroupDatagrid();
        }else{
             $.messager.alert('操作提示', result, 'info');
        }
    });
}

function clearAddUserGroupForm() {
    $('#add_user_group').form('clear');
}