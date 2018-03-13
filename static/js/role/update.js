function submitModifyRoleForm() {
    var url = "/role/modify";
    var pid = $("input[name='role_pid']").val();
    var data = {
        pid: pid,
        id: $("input[name='m_role_id']").val(),
        name: $("input[name='m_role_name']").val(),
        role_url: $("input[name='m_role_url']").val(),
        module: $("input[name='m_module']").val(),
        action: $("input[name='m_action']").val(),
        is_menu: $("#m_role_is_menu").combobox('getValue'),
        remarks: $("input[name='m_role_remarks']").val()
    };

    $.post(url, data, function (result) {
        if (result === "success") {
            clearModifyRoleForm();
            loadTree(pid);
            loadDataGrid(pid);
            $('#modify_role').window('close');
            $.messager.alert('操作提示', "修改成功", 'info');
        }
    });
}

function clearModifyRoleForm() {
    $('#modify_role').form('clear');
}