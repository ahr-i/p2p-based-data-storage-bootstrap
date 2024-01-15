$(function() {
    let node_list = $('#node_list')

    $(document).ready(function() {
        /* 살아있는 Node List 출력 */
        function fetchNodeList() {
            $.get('/log_get_node_list', function(data) {
                node_list.empty();

                let cnt = 0;    // Node 순서

                /* Node 정보 Node_list에 추가 */
                data.forEach(function(node) {
                    node_list.append('<li>' + (++cnt) + ".&emsp;&emsp;" + node.address + "&emsp;&emsp;" + "<a href='http://" + node.address + "/' target='_blank'>Go" + "</a>" + '</li>');
                });
            });
        }
        fetchNodeList();
    
        setInterval(fetchNodeList, 4000);   // 4초마다 Update
    });
});