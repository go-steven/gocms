function submitAddRoleForm() {
    var url = "/role/addrole";
    var pid = $("input[name='search_role_pid']").val();
    var data = {
        pid: pid,
        name: $("input[name='role_name']").val(),
        role_url: $("input[name='role_url']").val(),
        module: $("input[name='module']").val(),
        action: $("input[name='action']").val(),
        is_menu: $("#role_is_menu").combobox('getValue'),
        remarks: $("input[name='role_remarks']").val()
    };

    $.post(url, data, function (result) {
        if (result === "success") {
            clearAddRoleForm();
            loadTree(pid);
            loadDataGrid(pid);
            $.messager.alert('操作提示', "添加成功", 'info');
        }
    });
}

function clearAddRoleForm() {
    $('#add_role').form('clear');
    $("#role_is_menu").combobox({value:"1"})
}