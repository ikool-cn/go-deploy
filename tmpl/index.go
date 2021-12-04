package tmpl

func GetIndexTpl() string {
	return `<!doctype html>
<html xmlns="http://www.w3.org/1999/html">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width,initial-scale=1, minimum-scale=1.0, maximum-scale=1, user-scalable=no"/>
    <meta content="telephone=no" name="format-detection"/>
    <title></title>
    <script>
        (function (_D) {
            var _self = {};
            _self.resizeEvt = 'orientationchange' in window ? 'orientationchange' : 'resize';
            _self.Html = _D.getElementsByTagName("html")[0];
            _self.widthProportion = function () {
                var p = Number(_self.Html.offsetWidth / 750).toFixed(3);
                //return p;
                return p > 1.024 ? 1.024 : p < 0.427 ? 0.427 : p;
            };
            _self.changePage = function () {
                _self.Html.setAttribute("style", "font-size:" + _self.widthProportion() * 100 + "px");
            };
            _self.changePage();
            if (!document.addEventListener) return;
            window.addEventListener(_self.resizeEvt, _self.changePage, false);
            document.addEventListener('DOMContentLoaded', _self.changePage, false);
        })(document);
    </script>
</head>
<style>
    *{font-size:.24rem}
    *{margin:0;padding:0}
    :root {--font-family:"Helvetica Neue",Helvetica,Arial,"PingFang SC","Hiragino Sans GB","Heiti SC","Microsoft YaHei","WenQuanYi Micro Hei",sans-serif;}
    blockquote,body,button,dd,dl,dt,fieldset,form,h1,h2,h3,h4,h5,h6,hr,input,legend,li,ol,p,pre,td,textarea,th,ul{margin:0;padding:0;box-sizing:border-box}
    button,img,input,select,textarea{vertical-align:middle;outline:0;border:0}
    body{font-family:var(--font-family);overflow-x:hidden;background:#edeef2}
    body,html{min-height:100%}
    a{text-decoration:none;color:#333;text-align:center}
    em,i{font-style:normal;font-weight:500}
    li,ol,ul{list-style:none}
    img{font-size:0;line-height:0;border:0;vertical-align:middle}
    *{outline:0;-webkit-tap-highlight-color:transparent;-webkit-tap-highlight-color:transparent}
    input[type=button]{-webkit-appearance:none}
    .main{margin:0 auto;max-width:7.5rem;min-width: 7.5rem;width:100%;margin-bottom:.3rem}
    .title{margin-top:.2rem;text-align:center}
    .title a{font-size:.3rem;font-weight:700;color:#333;line-height:.6rem;}
    .status_box{margin:auto .2rem;display:flex;flex-wrap:wrap;justify-content:space-between;align-items:center}
    .online li{display:flex;align-items:center;float:left;margin-left:.2rem;font-size:.26rem;font-family:Helvetica;font-weight:700;color:green;line-height:.6rem}
    .sign{margin:auto .1rem;width:.08rem;height:.08rem;background:green;border-radius:50%;display:inline-block}
    .off{background:red}
    .offline{color:red!important}
    .seize_seat{width:1.2rem;height:.4rem}
    .show_logs{background:#333;border-radius:0.06rem;font-size:0.26rem;font-weight:400;color:#fff;line-height:.4rem;text-align:center}
    .item_box{margin:0.2rem 0.2rem 0 0.2rem;padding:0.2rem;background:#fff;box-shadow:0 .06rem .1rem 0 #dcdde0;border-radius:.16rem}
    .item_info{display:flex;align-items:center}
    .item_name{width:1.7rem;font-size:.26rem;font-weight:700;color:#333;line-height:.6rem;align-items:center}
    .item_type img{margin-left:.2rem;width:.3rem;height:.3rem}
    .item_node{flex-grow:1;max-width:5rem;width: 5rem;overflow: hidden;}
    .item_node ul{display:flex;flex-direction:row;flex-wrap:wrap;align-items:center}
    .item_node li{display:flex;align-items:center;flex:0 0 33.3333%}
    .item_node li span{width:1.2rem;font-size:.26rem;font-weight:400;color:green;line-height:.6rem;overflow:hidden;white-space:nowrap;display:block}
    .item_node li img{width:.3rem;height:.3rem}
    .item_node select{flex-grow:1;background: transparent;}
    .item_node option{font-size:.26rem;font-weight:400;color:#333;}
    .item_button{margin:auto 1.6rem;display:flex;justify-content:space-evenly;flex-wrap:wrap;align-items:center;padding-top:.2rem;padding-bottom:.2rem}
    .button{width:1.4rem;height:.4rem;padding:0.06rem 0; background:#333;border-radius:0.06rem;line-height:.4rem;text-align:center;font-size:.26rem;font-weight:400;color:#fff}
    .button,.show_logs{cursor: pointer;}
    .svn, .git{width: 0.5rem;height: 0.5rem;background-repeat: no-repeat;background-size: contain; }
    .svn{background-image: url("data:image/png;base64,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") }
    .git{background-image: url("data:image/png;base64,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")}
    .server{width: 0.3rem;height: 0.3rem;background-repeat: no-repeat;background-size: contain; background-image: url("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAB4AAAAeCAYAAAA7MK6iAAAAAXNSR0IArs4c6QAAAYVJREFUSEvtl81Kw0AQx/9DclWE9BVEbw00B22vSnvWV+hDiIUcAhUfonkEPafotdXDLuSo+AalAb+OgZEJraAHu1tDaqVzysdMfjuzs+E/BAC+7++6rnvJzEcAtuWZgb0S0W2e52dpmj4Z+H9xoSAI9pj5HsCObfDM/5mIDpRSjzbxAr5i5pNms4kwDFGr1Yzip9MpoijCeDwGEV0rpU6NAmdO1Gg0XqS8SZIYQ+cAgXc6Hbl911pv2YJZApRSNnGfvkEQFNdaa7L5gGT8N8Ddbhdpmv64eN/3MRgMCp/SMjYB1+t1xHFcLthmn0rNeGVgk1Jv9ni9z/HKmmsDXlSB0n6Zi0Df3/8vsKgLURliok5EpVSSsagLURliIolEpVQO9jwPw+GwGvBoNCpK7TgOer0eWq1WNWCTDl/Lrl5a3mZZhna7vZy8nQt62T85MtJEJjaZTNDv9yF9sKyg3wdw95sRBsCh1vrBZMFzn0KEy/wE4IKZjwGYTgRvRHQD4Nx2bhLmBwCtKC6u6Y69AAAAAElFTkSuQmCC");}
    .dialog{position:fixed;left:0;top:0;z-index:10001;width:100%;height:100%;display: none;}
    .dialog-overlay{position:absolute;top:0;left:0;z-index:10002;width:100%;height:100%;background-color:rgba(0,0,0,0.5);display: flex;justify-content: center;align-items: center;}
    .dialog-content{width:2rem;height:2rem;padding:0.2rem;border-radius:0.1rem;background-color:rgba(0,0,0,0.8);display:flex;flex-direction: column;flex-wrap: wrap;justify-content: center;align-items: center;}
    .info-icon{width: 1.28rem;height: 1.28rem;background-repeat: no-repeat;background-size: contain;}
    .loading{background-image: url("data:image/gif;base64,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")}
    .loading-success{background-image: url("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGAAAABgBAMAAAAQtmoLAAAAJ1BMVEVMaXH////////////////////////////////////////////////c+C/6AAAADHRSTlMAn8/vMBBgr3BQQN8gQ8DgAAABT0lEQVRYw+3WOwoCMRAG4ATUWnsbCztLG8FC8AJiY+sBPIR38SQqxMI5lPvwkUweMz+IjTv1Dh9ZJn/GmK66+mn1J2DDloYYMKMzCBBEVABBRA0gRAMgRAvoiSdAdNtgANEUA4guGwzQER6gI3xAQwSAhggBmejNWMNFaFiz76X56C15w/C/gREI2JgrA+PoQBIQjYkE8EmUAUbIQEhogIDQAD6hAzxCB3yIGDgmgTcRAS4NvIgY2GeAJ6EHWgIBGgIBaqIMGJ5UND0VgSgL6U5FgKdtohy7WTupYc/v7hgD4h8rAQLhEoFlQaBIuGQmWhAoEC4TuxYEsoTLJrsFgQzhCo+HBYEk4YrvkwWBBLEQtgROXFfS4sKIubgaWRBghAyEhAYICA3gEzrAI3TAh9ACb0ILGDMAAWMOGNASCNAQCFATGFARGFARINDVN+oB9LwfKgs3/BgAAAAASUVORK5CYII=");}
    .loading-faild{background-image: url("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGAAAABgBAMAAAAQtmoLAAAAJFBMVEVMaXH///////////////////////////////////////////9tKdXLAAAAC3RSTlMA7zCfzxBgj3DfgCUCmPUAAAG/SURBVFjD7dcxTsQwEAVQyxGBErmgSbO5QXo3NK5oKHKApU7DpTgABK2E/uUodjdrj//YuKCLSyujp83a8yfG7Gtf/7B6x3btk1qwfB7I7oij8vw9QAg74KQUeIAQI6AQHUAIOwAK4QFCjIBCdAAh7AAohD8XCOIMMOICCOICMMJfCxLiCuTEBiTEBuSEvxVExA2QRARERARIwscFGxEDgngBIRIAeI4LZhAiBb5f44I7QgggpG+JECWAEWWAEGUgJ2pARrxVgIxADZBEFSgSgXeNuQ0oEEHrfHMboBJB78ZzG6AQodTw5zaAEqGcKXMbQIhQiy1BrNWce0gLvqoFC0B7lBqk8ke7NqBK9Pkf59qACtGzw+fagCLR8wvk2oAC0WtNwLUBKtHrjcy1AcDnsQ6kzftUBVZxL441YBJX71QDRNbnhARk2GfEkt/kIkGAMrGwVlEiBgJI4tHQySvuRZ1+oJJUnuhU5IxRiJXOXfLERsREBy9njEKsdLTLr8RGTHS2c8YoxErH058D/U7IgStBU2KkedDpMWQHApwJJYZGGjidnnN2IIAxXs+5kSZapwepfafbH/sX7r7+sH4BzbolrX0lSA8AAAAASUVORK5CYII=");}
    .info-text{display:block;margin:0.12rem auto 0;font-size:0.26rem;color:#ffffff}
    .logwin{position:fixed;left:0;top:0;z-index:10001;width:100%;height:100%;display: none;}
    .logwin-overlay{position:absolute;top:0;left:0;z-index:10002;width:100%;height:100%;background-color:rgba(0,0,0,0.5);display: flex;justify-content: center;align-items: center;}
    .logwin-content{width:6rem;height:80%;padding:0.2rem;border-radius:0.1rem;background-color:#ffffff;display:flex;flex-direction: column;flex-wrap: wrap;align-items: center;}
    .log-title, .log-btn{display: flex;width: 100%;height: 0.6rem;justify-content: center;align-items: center}
    .log-title-left{width: 0.35rem;height: 0.35rem;}
    .log-title-center{font-size: 0.26rem;font-weight:700;flex-grow: 1;text-align: center;}
    .log-title-close{width: 0.35rem;height: 0.35rem;cursor:pointer;background-repeat: no-repeat;background-size: contain;background-image: url("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAE4AAABOCAYAAACOqiAdAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyFpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTQyIDc5LjE2MDkyNCwgMjAxNy8wNy8xMy0wMTowNjozOSAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENDIChXaW5kb3dzKSIgeG1wTU06SW5zdGFuY2VJRD0ieG1wLmlpZDoxMkJENkM2MkUwRDMxMUU4OUQzOTlDMTU4RDhFNDEwQSIgeG1wTU06RG9jdW1lbnRJRD0ieG1wLmRpZDoxMkJENkM2M0UwRDMxMUU4OUQzOTlDMTU4RDhFNDEwQSI+IDx4bXBNTTpEZXJpdmVkRnJvbSBzdFJlZjppbnN0YW5jZUlEPSJ4bXAuaWlkOjEyQkQ2QzYwRTBEMzExRTg5RDM5OUMxNThEOEU0MTBBIiBzdFJlZjpkb2N1bWVudElEPSJ4bXAuZGlkOjEyQkQ2QzYxRTBEMzExRTg5RDM5OUMxNThEOEU0MTBBIi8+IDwvcmRmOkRlc2NyaXB0aW9uPiA8L3JkZjpSREY+IDwveDp4bXBtZXRhPiA8P3hwYWNrZXQgZW5kPSJyIj8+iD/bFQAACItJREFUeNrsXGtsFFUU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")}
    .log-main{width: 100%;height: calc(100% - 1.3rem);}
    .log-main textarea{font-size: 0.2rem;font-family: var(--font-family);width: 100%;height: 100%;border: none;resize: none;}
</style>
<body>
<div class="main">
    <div class="title"><a target="_blank" href="https://github.com/ikool-cn/go-deploy">Go-deploy</a></div>
    <div class="status_box">
        <div class="seize_seat"></div>
        <div class="online">
            <ul>
                <li><span class="sign"></span>online</li>
                <li class="offline"><span class="sign off"></span>offline</li>
            </ul>
        </div>
        <div class="seize_seat show_logs">Logs</div>
    </div>

    <div class="item-list"></div>
</div>

<!--loading-->
<div class="dialog">
    <div class="dialog-overlay">
        <div class="dialog-content">
            <div class="info-icon loading"></div>
            <div class="info-text">Loading</div>
        </div>
    </div>
</div>

<!--log window-->
<div class="logwin">
    <div class="logwin-overlay">
        <div class="logwin-content">
            <div class="log-title">
                <div class="log-title-left"></div>
                <div class="log-title-center">Log List</div>
                <div class="log-title-close"></div>
            </div>
            <div class="log-main">
                <textarea spellcheck="false" id="log"></textarea>
            </div>
            <div class="log-btn">
                <div class="button clear">Clear</div>
            </div>
        </div>
    </div>
</div>

<script src="https://cdn.bootcss.com/zepto/1.2.0/zepto.min.js"></script>
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
                if (json != null) {
                    var html = '';
                    $.each(json, function (index, item) {
                        html += '<div class="item_box group-'+item.groupid+'" data-id="' + item.groupid + '">' + "\n"
                        html += '<div class="item_info">\n' +
                            '                <div class="item_name">Type</div>\n' +
                            '                <div class="item_node '+(item.type =="git" ? 'git' : 'svn')+'"></div>\n' +
                            '            </div>\n'

                        html += '<div class="item_info">\n' +
                            '                <div class="item_name">Name</div>\n' +
                            '                <div class="item_node">'+item.name+'</div>\n' +
                            '</div>\n';

                        html += '<div class="item_info">\n' +
                            '                <div class="item_name">Node</div>\n' +
                            '                <div class="item_node">\n' +
                            '                    <ul>\n' + listNode(item.node) +
                            '                    </ul>\n' +
                            '                </div>\n' +
                            '            </div>';
                        html += '<div class="item_info">\n' +
                            '                <div class="item_name">Commit Log</div>\n' +
                            '                <div class="item_node">\n' +
                            '                    <select class="reversion">\n' +
                            '                    </select>\n' +
                            '                </div>\n' +
                            '            </div>';
                        html += '            <div class="item_button">\n' +
                            '                <div class="button deploy">Update</div>\n' +
                            '                <div class="button rollback">Rollback</div>\n' +
                            '            </div>\n';
                        html += '</div>\n';
                    });
                    $('.item-list').html(html);

                    $.each(json, function (index, item) {
                        showCommitLog(item.groupid)
                    });
                } else {
                    console.log(json);
                    loadingFaild('Resp err', true);
                }
            }
        })
    }

    function listNode(nodeList) {
        var html = '';
        $.each(nodeList, function (index, item) {
            if (typeof item.online != 'undefined' && item.online) {
                html += '<li title="' + item.addr + '"><em class="server"></em><span>'+item.alias+'</span></li>\n';
            } else {
                html += '<li title="' + item.addr + '"><em class="server"></em><span class="offline">'+item.alias+'</span></li>\n';
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
                if (json != null && typeof json.Status != 'undefined') {
                    if (json.Status == true) {
                        var option = '';
                        $.each(json.Data, function (index, item) {
                            option += '<option value="' + item.Reversion + '">' + item.Reversion + ' | ' + item.Author + ' | ' + item.Time + ' | ' + item.Content + '</option>' + "\n";
                        });
                        $('.group-' + groupid).find('.reversion').html(option)
                    } else {
                        stdlog(groupid, json.Msg, "")
                    }
                } else {
                    console.log(json);
                    loadingFaild('Resp err', true);
                }
            }
        });
    }

    function stdlog(groupid, msg, elapsed) {
        msg = msg || "Empty response from server\n";
        var elapsed_time = (elapsed =='' ? '' : '===================\nElapsed Time: ' + elapsed + 's\n===================\n');
        $('#log').prepend(elapsed_time + msg);
    }

    //clear logs
    $('.clear').on('click', function () {
        $('#log').empty();
    });

    //show logs
    $('.show_logs').on('click', function (){
        $('.logwin').show();
    });

    //close logs
    $('.log-title-close').on('click', function (){
        $('.logwin').hide();
    });

    $(document).on('click', '.deploy', function () {
        var groupid = $(this).closest('.item_box').data('id');
        showLoading();
        $.ajax({
            url: "/deply",
            type: "POST",
            dataType: "json",
            data: {"groupid": groupid},
            success: function (json) {
                if (json != null && typeof json.Status != 'undefined') {
                    loadingSuccess('Success', true)
                    if (json.Status == true) {
                        stdlog(groupid, json.Data, json.Elapsed)
                    } else {
                        stdlog(groupid, json.Msg, "")
                    }
                } else {
                    console.log(json);
                    loadingFaild('Resp err', true);
                }
            },
            complete: function () {
                window.setTimeout(function () {
                    closeLoading();
                }, 1500);
            },
            error: function (xhr, textStatus, errorThrown) {
                console.log(xhr)
                try {
                    var jsonobj = JSON.parse(xhr.responseText);
                    loadingFaild(jsonobj.error, true);
                } catch (e) {
                    console.log(xhr)
                }
            }
        })
    });

    $(document).on('click', '.rollback', function () {
        obj = $(this).closest('.item_box');
        var groupid = obj.data('id');
        var reversion = $.trim(obj.find('.reversion').val());
        if ( reversion != "" && confirm('confirm to rollback?')) {
            showLoading();
            $.ajax({
                url: "/rollback",
                type: "POST",
                dataType: "json",
                data: {"groupid": groupid, "reversion": reversion},
                success: function (json) {
                    if (json != null && typeof json.Status != 'undefined') {
                        loadingSuccess('Success', true)
                        if (json.Status == true) {
                            stdlog(groupid, json.Data, json.Elapsed);
                        } else {
                            stdlog(groupid, json.Msg, "");
                        }
                    } else {
                        console.log(json);
                        loadingFaild('Resp err', true);
                    }
                },
                complete: function () {
                    window.setTimeout(function () {
                        closeLoading();
                    }, 1500);
                },
                error: function (xhr, textStatus, errorThrown) {
                    console.log(xhr)
                    try {
                        var jsonobj = JSON.parse(xhr.responseText);
                        loadingFaild(jsonobj.error, true);
                    } catch (e) {
                        console.log(xhr)
                    }
                }
            });
        }
    });

    //onload
    loadData();

    function showLoading(msg, autoClose) {
        msg = msg || 'Loading';
        autoClose = autoClose || false;
        $('.info-text').html(msg);
        $('.info-icon').removeClass('loading-success loading-faild').addClass('loading');
        $('.dialog').show();
        if (autoClose == true) {
            window.setTimeout(function () {
                closeLoading();
            }, 1500);
        }
    }

    function loadingSuccess(msg, autoClose) {
        $('.info-text').html(msg);
        autoClose = autoClose || false;
        $('.info-icon').removeClass('loading loading-faild').addClass('loading-success');
        $('.dialog').show();
        if (autoClose == true) {
            window.setTimeout(function () {
                closeLoading();
            }, 1500);
        }
    }

    function loadingFaild(msg, autoClose) {
        $('.info-text').html(msg);
        autoClose = autoClose || false;
        $('.info-icon').removeClass('loading loading-success ').addClass('loading-faild');
        $('.dialog').show();
        if (autoClose == true) {
            window.setTimeout(function () {
                closeLoading();
            }, 1500);
        }
    }

    function closeLoading() {
        $('.dialog').hide();
    }
</script>
</body>
</html>`
}
