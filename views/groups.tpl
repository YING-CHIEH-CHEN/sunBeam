<!DOCTYPE html>
<html lang="en">
  <head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>光合樂早餐Menu</title>
	</head>
	<body>
    <center><h2>光合樂&nbsp;{{$.toDay}}&nbsp;團體早餐單</h2></center>
    <center>
    {{if $.recordsGroupBool}}
  <center>
    <table border=0>
        <tr>
            <td>
            <ul style='list-style-type: "\1F44D";'>
            {{range $group, $val := .AllGroups}}
              <il>
              <a href="/group/mem/{{$group}}" style="text-decoration: none;"><font style="font-size:28px">{{$val}}</font></a><br><br>
              </il>
            {{end}}
            </ul>
            </td>
        </tr>
    </table>
  {{end}}
    </center>
  </body>
</html>