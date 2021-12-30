
<!--
{{ define "contentGroups" }}
  {{if $.recordsGroupBool}}
  <center>
  <table border=1>
  {{range $indexG, $group := .groups}}  
    <tr bgcolor=#FFE6D9>
    <td rowspan=4 width=66 bgcolor=#CCFFFF><a href="/group/mem/{{$indexG}}" style="text-decoration: none;">{{$group}}</a></td>
    {{range $meal := $.meals}}
    {{if gt (len $meal.Name) 0}}<td align=center>{{$meal.Name}}</td>{{end}}
    {{end}}
    </tr>


    <tr bgcolor=#FFE6D9>
      <td>0</td><td>0</td><td>0</td><td>0</td><td>0</td><td>0</td>
    </tr>
    <tr bgcolor=#EFFFD7>
      {{range $drink := $.drinks}}
      {{if gt (len $drink.Name) 0}}<td align=center>{{$drink.Name}}</td>{{end}}
      {{end}}
    </tr>
    <tr align=right bgcolor=#EFFFD7>
      <td>0</td><td>0</td><td>0</td>
    </tr>
  {{end}}
  </table>
-->

<!--
    <table border=1>
      <tr align=center bgcolor=#FFE6D9>
        <td rowspan=4 align=left bgcolor=#CCFFFF width=60>陽光潛水夥伴</td>
        <td>巧克力吐司</td><td>香雞蛋吐司</td><td>蔬菜蛋燒餅</td><td>原味蛋餅</td><td>豬肉漢堡</td><td>素菜包+饅頭</td>
      </tr>
      <tr align=right bgcolor=#FFE6D9>
        <td>0</td><td>1</td><td>0</td><td>3</td><td>0</td><td>0</td>
      </tr>
      <tr align=right bgcolor=#EFFFD7>
        <td>冰紅茶</td><td>冰奶茶</td><td>溫豆漿（有糖）</td>
      </tr>
      <tr align=right bgcolor=#EFFFD7>
        <td>1</td><td>1</td><td>2</td>
      </tr>
      <tr align=center bgcolor=#FFE6D9>
        <td rowspan=4 align=left bgcolor=#CCFFFF width=60>宗諺教練</td>
        <td>巧克力吐司</td><td>香雞蛋吐司</td><td>蔬菜蛋燒餅</td><td>原味蛋餅</td><td>豬肉漢堡</td><td>素菜包+饅頭</td>
      </tr>
      <tr align=right bgcolor=#FFE6D9>
        <td>0</td><td>0</td><td>1</td><td>0</td><td>1</td><td>1</td>
      </tr>
      <tr align=right bgcolor=#EFFFD7>
        <td>冰紅茶</td><td>冰奶茶</td><td>溫豆漿（有糖）</td>
      </tr>
      <tr align=right bgcolor=#EFFFD7>
        <td>1</td><td>0</td><td>0</td>
      </tr>
      <tr align=center bgcolor=#FFE6D9>
        <td rowspan=4 align=left bgcolor=#CCFFFF width=60>孟聖教練</td>
        <td>巧克力吐司</td><td>香雞蛋吐司</td><td>蔬菜蛋燒餅</td><td>原味蛋餅</td><td>豬肉漢堡</td><td>素菜包+饅頭</td>
      </tr>
      <tr align=right bgcolor=#FFE6D9>
        <td>0</td><td>0</td><td>0</td><td>0</td><td>0</td><td>0</td>
      </tr>
      <tr align=right bgcolor=#EFFFD7>
        <td>冰紅茶</td><td>冰奶茶</td><td>溫豆漿（有糖）</td>
      </tr>
      <tr align=right bgcolor=#EFFFD7>
        <td>0</td><td>0</td><td>0</td>
      </tr>
    </table>
    -->
    </center>
  {{end}}
{{ end }}

<br><br>

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

{{ define "testBlock" }}
<tr>{{.}}</tr>
{{end}}