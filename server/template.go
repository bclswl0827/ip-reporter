package main

const INDEX_HTML = `
<!DOCTYPE html>
<html lang="zh-CN">
    <head>
        <meta charset="utf-8" />
        <title>设备 IP 记录</title>
    </head>
    <body>
        <h1>设备 IP 记录</h1>
        <table border="1" width="100%">
            <thead>
                <tr>
                    <th>设备名称</th>
                    <th>IP 地址</th>
                    <th>更新时间</th>
                </tr>
            </thead>
            <tbody id="device_list"></tbody>
        </table>
    </body>
    <script>
        const getTime = (ts) => {
            const date = new Date(ts);
            const year = date.getFullYear();
            const month = (date.getMonth() + 1).toString().padStart(2, "0");
            const day = date.getDate().toString().padStart(2, "0");
            const hour = date.getHours().toString().padStart(2, "0");
            const minute = date.getMinutes().toString().padStart(2, "0");
            const second = date.getSeconds().toString().padStart(2, "0");
            return (
                year +
                "-" +
                month +
                "-" +
                day +
                " " +
                hour +
                ":" +
                minute +
                ":" +
                second
            );
        };

        const main = async () => {
            const table_list = document.getElementById("device_list");
            table_list.innerText = "";
            const res = await fetch("/devices");
            const { data } = await res.json();
            if (!!data.length) {
                data.forEach((item) => {
                    const tr = document.createElement("tr");
                    const td_device_tag = document.createElement("td");
                    td_device_tag.innerText = item.device_tag;
                    tr.appendChild(td_device_tag);
                    const td_ip_address = document.createElement("td");
                    td_ip_address.innerText = item.ip_address;
                    tr.appendChild(td_ip_address);
                    const td_updated_at = document.createElement("td");
                    td_updated_at.innerText = getTime(item.updated_at);
                    tr.appendChild(td_updated_at);
                    table_list.appendChild(tr);
                });
            } else {
                const tr = document.createElement("tr");
                const td = document.createElement("td");
                td.innerText = "暂无数据";
                td.colSpan = 3;
                td.align = "center";
                tr.appendChild(td);
                table_list.appendChild(tr);
            }
        };

        main();
        setInterval(main, 3000);
    </script>
</html>
`
