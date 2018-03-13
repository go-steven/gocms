$(function () {
    addUserObj = {
        search: function () {
            $('#user_group').datagrid('load', {
                group_name: $('input[name="add_user_user_group_name"]').val()
            });
        }
    };
    //datagrid初始化
    $('#user_group').datagrid({
        url: 'usergroup/gridlist',
        iconCls: 'icon-edit',//图标
        width: 700,
        height: 'auto',
        nowrap: false,
        striped: true,
        border: true,
        collapsible: false,//是否可折叠的
        fit: true,//自动大小
        //sortName: 'id',
        //sortOrder: 'desc',
        remoteSort: false,
        idField: 'id',
        singleSelect: false,//是否单选
        pagination: true,//分页控件
        rownumbers: true,//行号
        fitColumns: true,//列宽自适应（列设置width=100）
        frozenColumns: [[
            { field: 'ck', checkbox: true }
        ]],//设置表单复选框
        toolbar: add_user_toolbar
    });
});



function submitAddUserForm() {
    var groupSelections = $('#user_group').datagrid('getSelections');
    if (groupSelections.length === 0) {
        $.messager.alert('操作提示', '请至少选择一个组', 'info');
        return false;
    }

    var groupIdArray = new Array(groupSelections.length);
    for (var i = 0; i < groupSelections.length; i++) {
        groupIdArray[i] = groupSelections[i].id;
    }

    var data = {
        group_ids: groupIdArray.join(','),
        name: $("input[name='name']").val(),
        realname: $("input[name='realname']").val(),
        phone: $("input[name='phone']").val(),
        department: $("input[name='department']").val(),
        passwd: $("input[name='passwd']").val(),
        mail: $("input[name='email']").val()
    };

    if (data.name.length < 1 || data.realname.length < 1
        || data.phone.length < 1 || data.department.length < 1
        || data.passwd.length < 1 || data.mail.length < 1) {
        $.messager.alert('操作提示', '信息填写不完整,请补充后重新提交', 'info');
        return false;
    }


    $.post('/user/adduser', data, function (result) {
        if (result === 'success') {
            $('#add_user').window('close');
            $.messager.alert('操作提示', "添加成功", 'info');
            loadUserGrid();
        } else {
            $.messager.alert('操作提示', result, 'info');
        }
    });
}


function clearAddUserForm() {
    $('#add_user').form('clear');
}

function loadUserGrid() {
    $('#user_list').datagrid('load', {});
}