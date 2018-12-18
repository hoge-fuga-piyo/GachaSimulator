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
	let write_data = '';
	for(let i=0; i < result.length; i++) {
		if (result[i] == 0) {
			write_data += "N<br>"
		} else if (result[i] == 1) {
			write_data += "R<br>"
		} else {
			write_data += "SR<br>"
		}
	}
	$('#result').html(write_data);
}
