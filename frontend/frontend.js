function runLottery() {
	let api_result = requestApi().done(function(result) {
		console.log(JSON.stringify(result));
		writeResult(JSON.stringify(result));
	}).fail(function(result) {
		console.log("failed call api");
	});
}

function requestApi() {
	return $.ajax({
		type: 'GET',
		url: '/v1/lottery?num=10&type1=0.5&type2=0.3&type3=0.2',
		dataType: 'json',
	});
}

function writeResult(response) {
	let result = JSON.parse(response).result;
	let write_data = '';
	for(let i=0; i < result.length; i++) {
		if (result[i] == 0) {
			write_data += "N<br>"
			console.log("N");
		} else if (result[i] == 1) {
			write_data += "R<br>"
			console.log("R");
		} else {
			write_data += "SR<br>"
			console.log("SR");
		}
	}
	console.log(write_data);
	$('#result').html(write_data);
}
