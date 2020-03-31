package tmpl

func GetIndexTpl() string {
	return `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>go-deploy</title>
    <style>
        body{background-color:#fff;color:#24292e;font-family:-apple-system,BlinkMacSystemFont,Segoe UI,Helvetica,Arial,sans-serif,Apple Color Emoji,Segoe UI Emoji,Segoe UI Symbol;font-size:16px;line-height:1.5}
        button,input,select,textarea{font-family:inherit;font-size:inherit;line-height:inherit}
        table{border-collapse:collapse;border-spacing:0;display:table}
        tr{display:table-row;vertical-align:inherit;border-color:inherit}
        td,th{vertical-align:top;padding:5px;border:1px solid #aaa}
        textarea{border:1px solid #ccc;border-radius:4px;transition:border-color ease-in-out .15s,box-shadow ease-in-out .15s}
        .reversion{width:200px}
        .title{width:100%;margin:0 auto;text-align:center}
        .title a{color:#000;text-decoration:none}
        .clearlog{float:right;font-size:10px}
        .svn{width:24px; height:24px; margin: 0 20px; display: inline-block; background-image: url("data:image/png;base64,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") }
        .git{width:24px; height:24px; margin: 0 20px;  display: inline-block;  background-image: url("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAYAAADgdz34AAAACXBIWXMAAAsTAAALEwEAmpwYAAAD2ElEQVRIiaWVW2hcVRiFv31m4mQSejqpvSG16kMssb6KLQojanOpCqWiRQQt1ZKmtNSohfpW8KmoFI2xFAtWQTEUCiY2EfFBqWi8VRSL9VUFFTqZyUxm5pw5e/+/DzOZmZjJBTyw2S//XmuvtfbiGFb5FY4eTIu4M0ZE2+LJQ8mRM5dXc86sFlyjYFwrFR9RjBfL09n5sD9ybkUSb6WBuZMneqvgkY8CKGojXwv5j/OHD9z7vwjmTp7oTQ4dm4rd1u2DgjYtG/lamL20EsmSBIXhobT+89eF4N23veShY8Rv76FKQmO3kc9sdiJ/6Ol7lsJpmUFheCitQXkCG61REeLd20gOPYcGASQSlM+O4H66AiqoCCYWz5lU10P+W+98taKCwvBQWsNgHButmbfD/vYrGgR4GzfhrU1xwzNDBOLqdmlUSWnm+qX84FM7lyUovHjkPq2E4zjrA6gq895rVGnIDgKK1hGKNEhsJaWZzGT+wBM7WxIUjh9Na1D+CGv9BWHWPA/fO4dkruNmc5jOTpLd2yhZh22aU1tJSS63gMQ0wEvjRFEVXKreogKi9T2x5zFKYUh47Sobhk+QOfUy4bVf6PAMcaU+Z7xYllTX7rXnx6ZN4fjRNEEwrrZaIrQqW928xw2SxJ7HKYUh+YtjJHu2s/75l5BSEdqThG++hr3yLYiiKnixtizr/AFPg/KoutrNm996NYSFC8UY8FSpXP0ZKZeI37SF+LobSRx5ASeK1p6wRGGXzGRHPbGR1r2uAdfD/e+7VzAKBsVgMO3JRvDGYFUQaVxUwki8ucgedKK5RU3VVgqqqRkFo0ru9Vdw2RlcdobC6VN4gKiiqjghU7B20AD8vX/fjkRb21QMTS0OtymDvfsII8vchQ/wVDA1YqOCp4oRwSg4kUzJ2l23fvb1jx7A5vNj02EUDVhxs81Ps1lFrGc7bekH6bi/l8Qdd2LmBani1XajBlWtgy/owebzY9OloNIvIrlWPWjfP4i3fgPxjZvwnz2MMVoFr1sGajRTdNI3D76oyVvevzhdjOyAqOQaKmo2rPEb7fTXglK/uVedzBQj13fLp1/+0LLJzSRzYdTvtKakRlIePd0I9I1Xq+AonhpEdKYY2UXg9Sa3+v588pEdHepNxVRTWmu3E8E5hwGMVIN14rKlit3VCrylgoaSiemiRP1WtJ5JDIgbU7fFqWTLQu9S4MsSANz84eQ3ZRgQ0RwANeAY4ESyZaRv6+Tn3y+Hsaqf/h+P7r67A/dJDFKI4pxkixr1bZ344ruVzq6KAOD3vQ/c1S7eWVQIxA6uBhzgX8aAx+FcfeXWAAAAAElFTkSuQmCC")}
    </style>
    <script src="https://cdn.bootcss.com/zepto/1.2.0/zepto.min.js"></script>
</head>
<body>
<div class="title"><h2><a href="https://github.com/ikool-cn/go-deploy" target="_blank">Go-deploy</a></h2></div>
<table style="margin: 0 auto">
    <thead>
        <th>Type</th>
        <th>Name</th>
        <th>Node</th>
        <th>Update</th>
        <th>Commit Log</th>
    </thead>
    <tbody class="list">
    </tbody>
    <tfoot>
        <tr><td colspan="5" style="text-align: center;"><span style="color: green">● online</span> <span style="color: red">● offline</span><span class="clearlog"><a href="javascript:void(0);">clear</a></span></td></tr>
        <tr><td colspan="5"><textarea spellcheck="false" id="log" style="width: 99%; height: 300px;"></textarea></td></tr>
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
                    html += '<td>' + (item.type =="git" ? '<span class="git"></span>' : '<span class="svn"></span>') + '</td>';
                    html += '<td>' + item.name + '</td>';
                    html += '<td>' + listNode(item.node) + '</td>';
                    html += '</td>';
                    html += '<td><button class="deploy">update</button></td>';
                    html += '<td><select class="reversion"></select><button class="rollback">rollback</button></td>' + "\n";
                });
                $('.list').html(html);

                $.each(json, function (index, item) {
                    showCommitLog(item.groupid)
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

    function showCommitLog(groupid) {
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
                            option += '<option value="' + item.Reversion + '">' + item.Reversion + ' | ' + item.Author + ' | ' + item.Time + ' | ' + item.Content + '</option>' + "\n";
                        });
                        $('.gid' + groupid).find('.reversion').html(option)
                    } else {
                        log(groupid, json.Msg, "")
                    }
                } else {
                    console.log(json);
                    alert('network err');
                }
            }
        })
    }

    function log(groupid, msg, elapsed) {
        var tr = $('.gid' + groupid);
        var name = tr.find('td').eq(1).text();
        msg = msg || "Empty response from server\n";
        var elapsed_time = (elapsed =='' ? '' : '------------------elapsed time: ' + elapsed + 's ------------------' + "\n");
        $('#log').prepend(elapsed_time + msg);
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
                        log(groupid, json.Data, json.Elapsed)
                    } else {
                        log(groupid, json.Msg, "")
                    }
                } else {
                    console.log(json);
                    alert('network err');
                }
            }
        })
    });

    $(document).on('click', '.rollback', function () {
        var groupid = $(this).closest('tr').data('id');
        var reversion = $.trim($(this).siblings('.reversion').val());
        if ( reversion != "" && confirm('confirm to rollback?')) {
            $.ajax({
                url: "/rollback",
                type: "POST",
                dataType: "json",
                data: {"groupid": groupid, "reversion": reversion},
                success: function (json) {
                    if (typeof json.Status != 'undefined') {
                        if (json.Status == true) {
                            log(groupid, json.Data, json.Elapsed);
                        } else {
                            log(groupid, json.Msg, "");
                        }
                    } else {
                        console.log(json);
                        alert('network err');
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
