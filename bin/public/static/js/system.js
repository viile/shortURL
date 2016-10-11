var apiURL = "http://pturl.cn/";
var sys = {
	generate: function(){
		var httpURL =  $('#httpURL').val();
		var customize =  $('#customize').val();
		if(!httpURL.length)
		{
			swal("wrong!", "网址不能为空!", "error");
			return false;
		}
		var url = apiURL + "generate?url=" + encodeURIComponent(httpURL);
		if(customize.length > 1)
		{
			url += "&customize=" + customize;
		}
		$.get(url,  function (result) {
			if (result.error_code == 0) {
				//console.log(result);
				$('#gtext').html("http://pturl.cn/" + result.short_url);
			} else {
				swal("wrong!", "network error", "error");
			}
		}, 'json');
	},
	convert: function(){
		var shortURL =  $('#shortURL').val();
		if(!shortURL.length)
		{
			swal("wrong!", "网址不能为空!", "error");
			return false;
		}
		var url = apiURL + "convert?url=" + encodeURIComponent(shortURL);
		$.get(url, function (result) {
			if (result.error_code == 0) {
				$('#ctext').html(result.url);
			} else {
				swal("wrong!", "network error", "error");
			}
		}, 'json');
	},
}
