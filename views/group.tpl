<!DOCTYPE html>
<html lang="en">
  <head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>光合樂早餐Menu</title>
        <script type="text/javascript" src="/static/js/jquery3.3.1.min.js"></script>

		<script>
		var index=0;
		var lastDate=new Date('{{$.firstDay}}');
		function addOneDate(){
			$('#addPlus').remove();
			$('#addMinus').remove();
			var thisDate=addOneDay(lastDate);
			$('#myTable tr:last').after('<tr><td><input type=hidden name=dates value='+thisDate+'>'+thisDate+'</td><td><img id=addPlus src=/static/img/plus.png width=32 onclick=addOneDate()><img id=addMinus src=/static/img/minus.png width=32 onclick=removeOneDate()></td></tr>');
			index++;
			
			$('#dateTitle').attr("rowspan", $('#dateTitle').attr("rowspan")+1);
		}
		function addOneDay(thisDate){
			Date.prototype.yyyymmdd = function() {
				// getMonth() is zero-based
				var mm = this.getMonth() + 1;
  				var dd = this.getDate();
  				return [this.getFullYear(),'-',	(mm>9 ? '' : '0') + mm,'-',	(dd>9 ? '' : '0') + dd].join('');
			};
			var today = new Date(thisDate);
			var tomorrow = new Date(today);
			tomorrow.setDate(today.getDate()+1);
			lastDate = tomorrow.toLocaleDateString();
			//return tomorrow.toLocaleDateString().replace('/','-').replace('/','-');
			return tomorrow.yyyymmdd();
		}
		function removeOneDate(){
			if(index>0){
				$('#myTable tr:last').remove();
				$('#myTable tr:last').find('td:last-child').remove()
				if(index>1)
					$('#myTable tr:last td:last').after('<td><img id=addPlus src=/static/img/plus.png width=32 onclick=addOneDate()><img id=addMinus src=/static/img/minus.png width=32 onclick=removeOneDate()></td></tr>');
				else
					$('#myTable tr:last td:last').after('<td><img id=addPlus src=/static/img/plus.png width=32 onclick=addOneDate()></td></tr>');
				index--;
				var yesterday = new Date(today);
				yesterday.setDate(today.getDate()-1);
				lastDate = yesterday.toLocaleDateString();
			}else{				
			}
		}
		</script>
	</head>
	<body>
    <center><a href='/summary'><img src="/static/img/home.png" width=40></a>　　光合樂早餐</center><br>
    <center><h2>團體編輯</h2></center><br>

    <center>
    <form action="" method=post>{{ .xsrfdata }}
	<table border=1 id="myTable">
	<tr>
	<td>名稱</td>
	<td colspan=2><input type=text name=gName value='' size=24></td>
	</tr>
	<tr>
	<td id=dateTitle rowspan=1 class="rowspan">日期</td>
	<td valign=middle><input type=hidden name=dates value='{{$.firstDay}}'>{{$.firstDay}}</td><td><img id=addPlus src=/static/img/plus.png width=32 onclick=addOneDate()></td>
	</tr>
	</table><br><input type=submit value=確認 name=sub>
	</form>
    </center><br>
    {{ block "content" .}}{{end}}
  </body>
</html>

{{ define "content" }}
  {{if $.groupsBool}}
  <center>
  <h2>已登記</h1>

<table>
    {{range $index, $group := .groups}}
      <tr align=left><td><li>{{$group.Name}}</li></td><td width=40></td><td>{{$group.Date}}</td><td width=40></td><td><a href=/api/group/remove/{{$group.Id}}><img src=/static/img/minus.png width=32></a></td></tr>
    {{end}}
  </table>
  </center>
  {{end}}
{{ end }}
