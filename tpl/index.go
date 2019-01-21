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
        .title a {color: #000; text-decoration:none;}
        .clearlog{float: right;font-size: 10px;}
    </style>
    <script src="https://zeptojs.com/zepto.min.js"></script>
</head>
<body>
<div class="title"><h2><a href="https://github.com/ikool-cn/go-deploy" target="_blank">go-deploy</a></h2></div>
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
        <tr><td colspan="5" style="text-align: center;"><span style="color: green">● online</span> <span style="color: red">● offline</span><span class="clearlog"><a href="javascript:void(0);">clear</a></span></td></tr>
        <tr><td colspan="5"><textarea spellcheck="false" id="log" style="width: 99%; height: 360px;"></textarea></td></tr>
    </tfoot>
</table>
<script>
    Date.prototype.Format = function (fmt) {
        var o = {
            "M+": this.getMonth() + 1,
            "d+": this.getDate(),
            "h+": this.getHours(),
            "m+": this.getMinutes(),
            "s+": this.getSeconds(),
            "q+": Math.floor((this.getMonth() + 3) / 3),
            "S": this.getMilliseconds()
        };
        if (/(y+)/.test(fmt))
            fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
        for (var k in o)
            if (new RegExp("(" + k + ")").test(fmt))
                fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
        return fmt;
    }

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
                html += '<span title="' + item.addr + '" style="color: green">' + item.alias + '●</span> ';
            } else {
                html += '<span title="' + item.addr + '" style="color: red">' + item.alias + '●</span> ';
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
                if (typeof json.Status != 'undefined') {
                    if (json.Status == true) {
                        var option = '';
                        $.each(json.Data, function (index, item) {
                            option += '<option value="' + item.Reversion + '">r' + item.Reversion + ' | ' + item.Author + ' | ' + item.Time + ' | ' + item.Content + '</option>' + "\n";
                        });
                        $('.gid' + groupid).find('.reversion').html(option)
                    } else {
                        log(groupid, json.Msg)
                    }
                } else {
                    console.log(json);
                    alert('网络错误');
                }
            }
        })
    }

    function log(groupid, msg) {
        var tr = $('.gid' + groupid);
        var name = tr.find('td').eq(1).text();
        $('#log').prepend('['+(new Date()).Format("MM-dd hh:mm:ss")+']' + msg);
    }

    $('.clearlog').on('click', function () {
        $('#log').empty();
    })

    $(document).on('click', '.deploy', function () {
        var groupid = $(this).closest('tr').data('id');
        $.ajax({
            url: "/deply",
            type: "POST",
            dataType: "json",
            data: {"groupid": groupid},
            success: function (json) {
                if (typeof json.Status != 'undefined') {
                    if (json.Status == true) {
                        log(groupid, json.Data)
                    } else {
                        log(groupid, json.Msg)
                    }
                } else {
                    console.log(json);
                    alert('网络错误');
                }
            }
        })
    });

    $(document).on('click', '.rollback', function () {
        var groupid = $(this).closest('tr').data('id');
        var reversion = $(this).siblings('.reversion').val();
        if (reversion > 0 && confirm('确定要执行回滚吗?')) {
            $.ajax({
                url: "/rollback",
                type: "POST",
                dataType: "json",
                data: {"groupid": groupid, "reversion": reversion},
                success: function (json) {
                    if (typeof json.Status != 'undefined') {
                        if (json.Status == true) {
                            log(groupid, json.Data)
                        } else {
                            log(groupid, json.Msg)
                        }
                    } else {
                        console.log(json);
                        alert('网络错误');
                    }
                }
            })
        }
    });

    //onload
    loadData();
</script>
</body>
</html>`
}
