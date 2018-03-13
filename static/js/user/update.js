$(function () {
  modifyUserObj = {
        search: function () {
            $('#update_user_group').datagrid('load', {
                group_name: $('input[name="modify_user_user_group_name"]').val()
            });
        }
    };
    //datagrid初始化
    $('#update_user_group').datagrid({
        url: 'user/gridgrouplist',
        queryParams: {user_id: $("input[name='user_id']").val()},
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
            { field: 'ck', checkbox: true }
        ]],//设置表单复选框
        toolbar: modify_user_toolbar,
        onLoadSuccess:function(row){//当表格成功加载时执行
            $.each(row.rows, function(idx,val){//遍历JSON
                if(val.check){
                    $('#update_user_group').datagrid('selectRow', idx);//如果数据行为已选中则选中改行
                }
            });
        }
    });
});



function submitModifyUserForm() {
    var groupSelections = $('#update_user_group').datagrid('getSelections');
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
        user_id:$("input[name='user_id']").val(),
        name: $("input[name='modify_name']").val(),
        realname: $("input[name='modify_realname']").val(),
        phone: $("input[name='modify_phone']").val(),
        department: $("input[name='modify_department']").val(),
        passwd: $("input[name='modify_passwd']").val(),
        mail: $("input[name='modify_email']").val()
    };

    if (data.name.length < 1 || data.realname.length < 1
        || data.phone.length < 1 || data.department.length < 1
        || data.mail.length < 1) {
        $.messager.alert('操作提示', '信息填写不完整,请补充后重新提交', 'info');
        return false;
    }

    $.post('/user/modifyuser', data, function(result){
        if (result === 'success') {
            $('#modify_user').window('close');
            $.messager.alert('操作提示', '修改成功', 'info');
            loadModifyUserGrid();
        } else {
            $.messager.alert('操作提示', result, 'info');
        }
    });
}

function clearModifyUserForm() {
    $('#modify_user').form('clear');
}

function loadModifyUserGrid() {
    $('#user_list').datagrid('load', {});
}