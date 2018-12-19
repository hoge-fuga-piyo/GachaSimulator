function runLottery() {
	let num = $('#num').val();
	let api_result = requestApi(num).done(function(result) {
		console.log(JSON.stringify(result));
		writeResult(JSON.stringify(result));
	}).fail(function(result) {
		console.log("failed call api");
	});
}

function requestApi(num) {
	return $.ajax({
		type: 'GET',
		url: '/v1/lottery?type1=0.7&type2=0.2&type3=0.1&num=' + num,
		dataType: 'json',
	});
}

function writeResult(response) {
	let result = JSON.parse(response).result;
	let detail_data = '';
	let num1 = 0;
	let num2 = 0;
	let num3 = 0;
	for(let i=0; i < result.length; i++) {
		if (result[i] == 0) {
 			num1++;
			detail_data += '<p>N</p>';
		} else if (result[i] == 1) {
			num2++;
			detail_data += '<p>R</p>';
		} else {
			num3++;
			detail_data += '<p>SR</p>';
		}
	}
	let overview_data = '<h4>結果</h4>SR ' + num3 + '<br>R ' + num2 + '<br>N ' + num1;
	$('#resultDetail').show();
	$('#resultOverview').html(overview_data);
	$('#resultContent').html(detail_data);
}
