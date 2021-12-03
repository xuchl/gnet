<?php
 $st = "req:需要发送的数据";
 $length = strlen($st);

//创建tcp套接字
$socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
        
        
//设置阻塞模式
 socket_set_block($socket);

//连接tcp地址+端口
socket_connect($socket, '127.0.0.1', 9000);
        
//向打开的套集字写入数据（发送数据）
$msg = socket_write($socket, $st, $length);

//接收返回的数据
$recv_data = socket_read($socket, 1024);
		
echo $msg;//成功返回字符长度，否则为空字符
echo $recv_data;
//关闭连接
socket_close($socket);
?>
