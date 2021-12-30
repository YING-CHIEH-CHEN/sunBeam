<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>光合樂早餐</title>
	</head>
	<body>
<center><a href='/summary'><img src="/static/img/home.png" width=40></a></center>
    {{ block "content" .}}{{end}}
  </body>
</html>

{{ define "content" }}
  <center><h2>{{$.groupName}} {{$.toDay}} 已登記</h1></center>
  {{if $.recordsBool}}
    <center>
      <table border=1>
        <tr bgcolor=#ACD6FF align=center><td>姓名</td><td>餐點</td><td>飲料</td></tr>
        {{range $index, $record := .records}}
        <tr bgcolor=#{{if mod $index 2}}FFE6D9{{else}}EFFFD7{{end}}>
          <td>{{$record.Name}}</td>
          <td>{{$record.Main}}</td>
          <td>{{$record.Drink}}</td>
        </tr>
        {{end}}
      </table>
    </center>
  {{end}}
{{ end }}

{{ define "testBlock" }}
<tr>===={{.}}====</tr>
{{end}}