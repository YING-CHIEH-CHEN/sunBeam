<!DOCTYPE html>
<html lang="en">
  <head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>光合樂早餐Menu</title>
	</head>
	<body>
    <center><h2>{{block "home" .}}{{end}}光合樂{{$.toDay}}早餐單</h2></center>
    <center>
    {{block "groupName" .}}{{end}}
    	<form action="" method="POST">{{ .xsrfdata }}
      <table border=1><input type=text id=toDay name=toDay value='{{.toDay}}' hidden>
        <tr bgcolor=#ACD6FF><td>姓名</td><td><input type=text id=name name=name size=26></td></tr>
        <tr bgcolor=#ACD6FF><td>餐點</td>
          <td>
          <table border=0>
            {{range $index, $meal := .meals}}
              {{if mod $index 2}}<tr>{{end}}
                <td><input type=radio name=main value={{$meal.Cnt}}>{{$meal.Name}}</td>
              {{if not (mod $index 2)}}</tr>{{end}}
            {{end}}
            </table>
          </td>
        </tr>
        <tr bgcolor=#ACD6FF><td>飲料</td>
          <td>
          <table border=0>
            {{range $index, $drink := .drinks}}
              {{if mod $index 2}}<tr>{{end}}
              {{if not (eq $drink.Name "")}}
                <td><input type=radio name=drink value={{$drink.Cnt}}>{{$drink.Name}}</td>
              {{end}}
              {{if not (mod $index 2)}}</tr>{{end}}
            {{end}}
            </table>
          </td>
        </tr>
        <tr bgcolor=#ACD6FF align=center>
          <td colspan=2><input type=submit value="確定"></td>
        </tr>		
      </table>
      </form>
    </center><br>
    {{ block "content" .}}{{end}}
  </body>
</html>