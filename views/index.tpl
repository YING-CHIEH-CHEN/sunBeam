{{ template "base.tpl" . }}

{{ define "content" }}
  {{if $.recordsBool}}
<center>
  <h2>已登記</h1>
  <table border=1>
    <tr bgcolor=#ACD6FF align=center><td>姓名</td><td>餐點</td><td>飲料</td><td>刪除</td></tr>
    {{range $index, $record := .records}}
    <tr bgcolor=#{{if mod $index 2}}FFE6D9{{else}}EFFFD7{{end}}>
      <td>{{$record.Name}}</td>
      <td>{{$record.Main}}</td>
      <td>{{$record.Drink}}</td>
      <td><a href=/api/record/remove/{{$record.Id}}><img src=/static/img/minus.png width=32></a></td>
    </tr>
    {{end}}
  </table>
  </center>
  {{end}}
{{ end }}