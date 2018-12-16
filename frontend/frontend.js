function requestApi() {
    console.log("hogehoge");
    let request = new XMLHttpRequest();
    request.open('get', '/v1/lottery?num=5&type1=0.5&type2=0.5');
    request.onload = function(event) {
        let result = document.getElementById('result');
        result.innerHTML = "piyopiyo";

        if (request.readyState == 4) {
            if (request.status == 200) {
                console.log("SUCCESS");
                result.innerHTML = request.responseText;
            } else {
                console.log("ERROR");
            }
        }
    };

    request.send(null);
}
