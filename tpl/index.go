package tpl

func GetIndexTpl() string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>deploy</title>
    <style>
        table {border-collapse: collapse;}
        table, th, td {border: 1px solid gray; padding: 3px;}
        .reversion{width: 60px;}
    </style>
    <script src="https://zeptojs.com/zepto.min.js"></script>
</head>
<body>
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
        <td colspan="5" style="text-align: center;"><span style="color: green">● online</span> <span style="color: red">● offline</span></td>
    </tfoot>
</table>
<script>
    function LoadData() {
        $.ajax({
            url: "/list",
            dataType: "json",
            success: function (json) {
                var html = '';
                $.each(json, function (index, item) {
                    html += '<tr data-id="' + item.groupid + '">';
                    html += '<td>' + item.groupid + '</td>';
                    html += '<td>' + item.name + '</td>';
                    html += '<td>' + listNode(item.node) + '</td>';
                    html += '</td>';
                    html += '<td><button class="deploy">update</button></td>';
                    html += '<td><input type="text" class="reversion"><button class="rollback">rollback</button></td>' + "\n";
                });
                $('.list').html(html);
            }
        })
    }

    function listNode(json) {
        var html = '';
        $.each(json, function (index, item) {
            if (typeof item.online != 'undefined' && item.online) {
                html += '<span style="color: green">'+item.ip+'●</span> ';
            } else {
                html += '<span style="color: red">'+item.ip+'●</span> ';
            }
        });
        return html;
    }
    LoadData();

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
                data: {"groupid": id, "reversion" : reversion},
                success: function (json) {

                }
            })
        }
    });
</script>
</body>
</html>`
}
