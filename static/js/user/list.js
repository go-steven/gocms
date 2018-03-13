$(function () {
    searchUserObj = {
        search: function () {
            $('#user_list').datagrid('load', {
                name: $('input[name="search_name"]').val(),
                realname: $('input[name="search_realname"]').val(),
                mail: $('input[name="search_email"]').val(),
                phone: $('input[name="search_phone"]').val(),
                id: $('input[name="search_user_id"]').val()
            });
        }
    };
    //datagrid初始化
    $('#user_list').datagrid({
        url: 'user/gridlist',
        iconCls: 'icon-edit',//图标
        width: 700,
        height: 'auto',
        nowrap: false,
        striped: true,
        border: true,
        collapsible: false,//是否可折叠的
        fit: true,//自动大小
        //sortName: 'code',
        //sortOrder: 'desc',
        remoteSort: false,
        idField: 'id',
        singleSelect: false,//是否单选
        pagination: true,//分页控件
        rownumbers: true,//行号
        fitColumns: true,//列宽自适应（列设置width=100）
        frozenColumns: [[
            {field: 'ck', checkbox: true}
        ]],//设置表单复选框
        toolbar: user_toolbar
    });
});

//添加修改按钮
function userOpt(val, row, index) {
    return '<a href="#" onclick="openModifyUserWin(' + row.id + ')">修改</a>';
}

//打开添加管理员组窗口
function openAddUserWin() {
    $('#add_user').window({
        width: 800,
        height: 600,
        modal: true,
        // maximizable: false,
        minimizable: false,
        collapsible: false,//是否可折叠的
        href: '/user/addview'
    });
}

//打开修改管理员组窗口
function openModifyUserWin(userId) {
    $('#modify_user').window({
        width: 800,
        height: 600,
        modal: true,
        // maximizable: false,
        minimizable: false,
        collapsible: false,//是否可折叠的
        href: '/user/updateview?user_id=' + userId
    });
}

//删除方法
function deleteUser() {
    var selections = $('#user_list').datagrid('getSelections');
    if (selections.length === 0) {
        alert('请先选择要删除的记录');
        return false;
    }

    if (!confirm('确定要删除选中的数据吗？')) {
        return false;
    }
    var idArray = new Array(selections.length);
    for (var i = 0; i < selections.length; i++) {
        idArray[i] = selections[i].id;
    }

    $.post('/user/delete', {user_ids: idArray.join(',')}, function(result){
        if (result === 'success') {
            loadModifyUserGrid();
            $.messager.alert('操作提示', '删除成功', 'info');
        } else {
            $.messager.alert('操作提示', result, 'warning');

        }
        selections = 0;
    });
}

function loadModifyUserGrid() {
    $('#user_list').datagrid('load', {});
}