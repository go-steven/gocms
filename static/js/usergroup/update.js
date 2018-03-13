
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

//初始化tree
$(document).ready(loadTree());
function loadTree() {
    var group_user_id = $("input[name='group_user_id']").val();
    var url = "/usergroup/loadtreechecked";
    var data = {
        group_user_id: group_user_id
    };
    $.post(url, data, function (result) {
        // zNodes = result
        $.fn.zTree.init($("#modify_group_role_tree"), user_group_setting, result);
    });
}

/**
 * 修改管理员组
 */
function submitModifyUserGroupForm() {
    var zTree = $.fn.zTree.getZTreeObj("modify_group_role_tree");
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

    var url = "/usergroup/modifyusergroup";

    var data = {
        ids: ids,
        id: $("input[name='group_user_id']").val(),
        group_name: $("input[name='ag_m_name']").val(),
        remarks: $("input[name='ag_m_remarks']").val()
    };

    $.post(url, data, function (result) {
        if (result === "success") {
            $('#modify_user_group').window("close");
            $.messager.alert('操作提示', "修改成功", 'info');
            loadUserGroupDatagrid();
        } else {
            $.messager.alert('操作提示', result, 'info');
        }
    });
}

function clearModifyUserGroupForm() {
    $('#modify_user_group').form('clear');
}


