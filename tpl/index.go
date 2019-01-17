package tpl

func GetIndexTpl() string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>go-deploy</title>
    <style>
        table {border-collapse: collapse;}
        table, th, td {border: 1px solid gray; padding: 3px;}
        .reversion{width: 120px;}
        .title {width: 100%;margin: 0 auto;text-align: center; margin-bottom: 0;}
    </style>
    <script src="https://zeptojs.com/zepto.min.js"></script>
</head>
<body>
<div class="title"><h2>go-deploy</h2></div>
<table style="margin: 0 auto">
    <thead>
        <th>Id</th>
        <th>Name</th>
        <th>Node</th>
        <th>Update</th>
        <th>Rollback</th>
    </thead>
    <tbody class="list">
    </tbody>
    <tfoot>
        <tr><td colspan="5" style="text-align: center;"><span style="color: green">● online</span> <span style="color: red">● offline</span></td></tr>
        <tr><td colspan="5"><textarea id="log" style="width: 99%; height: 200px;"></textarea></td></tr>
    </tfoot>
</table>
<script>
    function loadData() {
        $.ajax({
            url: "/list",
            dataType: "json",
            success: function (json) {
                var html = '';
                $.each(json, function (index, item) {
                    html += '<tr class="gid' + item.groupid + '" data-id="' + item.groupid + '">';
                    html += '<td>' + item.groupid + '</td>';
                    html += '<td>' + item.name + '</td>';
                    html += '<td>' + listNode(item.node) + '</td>';
                    html += '</td>';
                    html += '<td><button class="deploy">update</button></td>';
                    html += '<td><select class="reversion"></select><button class="rollback">rollback</button></td>' + "\n";
                });
                $('.list').html(html);

                $.each(json, function (index, item) {
                    showSvnLog(item.groupid)
                });
            }
        })
    }

    function listNode(json) {
        var html = '';
        $.each(json, function (index, item) {
            if (typeof item.online != 'undefined' && item.online) {
                html += '<span style="color: green">' + item.ip + '●</span> ';
            } else {
                html += '<span style="color: red">' + item.ip + '●</span> ';
            }
        });
        return html;
    }

    function showSvnLog(groupid) {
        $.ajax({
            url: "/showlog",
            type: "POST",
            dataType: "json",
            data: {"groupid": groupid},
            success: function (json) {
                if (json.Status == true) {
                    var option = '';
                    $.each(json.Data, function (index, item) {
                        option += '<option value="' + item.Reversion + '">r' + item.Reversion + ' | ' + item.Author + ' | ' + item.Time + ' | ' + item.Content + '</option>' + "\n";
                    });
                    $('.gid' + groupid).find('.reversion').html(option)
                } else {
                    log(groupid, json.Msg)
                }
            }
        })
    }

    function log(groupid, msg) {
        var tr = $('.gid' + groupid);
        console.log(tr)
        var name = tr.find('td').eq(1).text();
        $('#log').prepend('[' + name + '] ' + msg);
    }

    $(document).on('click', '.deploy', function () {
        var id = $(this).closest('tr').data('id');
        $.ajax({
            url: "/deply",
            type: "POST",
            dataType: "json",
            data: {"groupid": id},
            success: function (json) {

            }
        })
    });

    $(document).on('click', '.rollback', function () {
        var id = $(this).closest('tr').data('id');
        var reversion = $(this).siblings('.reversion').val();
        if (reversion > 0 && confirm('确定要执行回滚吗?')) {
            $.ajax({
                url: "/rollback",
                type: "POST",
                dataType: "json",
                data: {"groupid": id, "reversion": reversion},
                success: function (json) {

                }
            })
        }
    });

    loadData();
</script>
</body>
</html>`
}
