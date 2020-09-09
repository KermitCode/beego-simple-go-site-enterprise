<!DOCTYPE html>

<html>
<head>
  <title>青岛市国企列表-青岛市国企名单</title>
  <meta name="description" content="青岛市国企列表,青岛市国企名单,青岛有哪些比较好的国企,青岛国企推荐">
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link href="https://cdn.staticfile.org/twitter-bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
  <script src="https://cdn.staticfile.org/jquery/2.1.1/jquery.min.js"></script>
  <script src="https://cdn.staticfile.org/twitter-bootstrap/3.3.7/js/bootstrap.min.js"></script>
</head>

<body>
  <div class="container">
	<div class="row">
		<div class="span12">
			<ul class="nav nav-tabs">
				<li class="active">
					<a href="/">青岛市国企列表</a>
				</li>
				<li>
					<a href="http://www.04007.cn">04007.cn主站</a>
				</li>
				<li>
					<a href="http://www.04007.cn/article/211.html" target="_blank">关于本站</a>
				</li>
        <li>
					<a href="#">本页面只为体验Go开发而写</a>
				</li>
        
				<li class="dropdown pull-right">
					 <a href="#" data-toggle="dropdown" class="dropdown-toggle">其它链接<strong class="caret"></strong></a>
					<ul class="dropdown-menu">
						<li>
							<a href="http://www.04007.cn/">查看04007主站</a>
						</li>
						<li class="divider"></li>
						<li>
							<a href="http://history.04007.cn/">历史上的今天</a>
						</li>
					</ul>
				</li>
			</ul>
			<h3>
				青岛市国企列表
			</h3>
			<table class="table table-striped">
				<thead>
					<tr>
						<th>序号</th>
            <th>企业名称</th>
            <th>所在区</th>
            <th>企业CEO</th>
            <th>官网地址</th>
            <!--<th>中国排名</th>
            <th>排名年份</th>-->
            <th>更新时间</th>
            <!--<th>更多</th>-->
					</tr>
				</thead>
				<tbody>
          {{range $index, $row := .enterprise}}
					<tr>
          
						<td>{{$index | index}}</td>
						<td>{{.title}}</td>
            <td>{{.area}}</td>
						<td><a href='{{.leaderbaike}}' target="_blank">{{.leader}}</a></td>
            <td><a href='{{.website}}' target="_blank">{{substr .website 0 30}}</a></td>
            <!--<td>{{.sortnum_china}}</td>
            <td>{{.sortyear}}</td>-->
            <td>{{.updateTime}}</td>
						<!--<td>企业详情</td>-->
            
					</tr>
          {{end}}
				</tbody>
			</table>
      <hr>
			<p class="text-center">
				<h6 class="text-center">本页面只为体验Go语言WEB开发而写. 2020-03-01</h6>
			</p>
		</div>
	</div>
</div>
</body>
</html>
