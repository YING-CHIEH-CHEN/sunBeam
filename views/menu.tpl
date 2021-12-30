<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>光合樂早餐Menu</title>
	</head>
	<body>
	<center><a href='./summary'><img src=static/img/home.png width=40></a>　　光合樂早餐Menu<br><br>

    <!--
        {{range $key, $value := .locs}}                
            {{index $.locs $key}}{{$value}}
        {{end}}

        {{with $x := index .meals 1}}{{$x.Loc}}{{end}}{{(index .meals 1).Loc}}
    -->
    
	<form action="/menu" method="POST">{{ .xsrfdata }}
	<table border=1>
        {{range $index, $meal := .meals}}                
                <tr bgcolor=#CCFFFF>
                {{if eq $index 0}}
                    <td rowspan={{len $.meals}}>餐點</td>
                {{end}}
                    <td>{{$meal.Cnt}}</td>
                    <td align=center><input type=text name={{$meal.Loc}} value='{{$meal.Name}}' size=12></td>
                </tr>
        {{end}}

        {{range $index, $drink := .drinks}}                
                <tr bgcolor=#CCFFFF>
                {{if eq $index 0}}
                    <td rowspan={{len $.drinks}}>飲料</td>
                {{end}}
                    <td>{{$drink.Cnt}}</td>
                    <td align=center><input type=text name={{$drink.Loc}} value='{{$drink.Name}}' size=12></td>
                </tr>
        {{end}}
    </table><br><input type=submit value=確認 name=sub>
	</form>
	</center>	
	</body>
</html>