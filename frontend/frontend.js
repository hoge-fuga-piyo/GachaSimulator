function requestApi() {
    console.log("hogehoge");
    let request = new XMLHttpRequest();
    request.open('get', '/v1/lottery?num=5&type1=0.5&type=0.5');
    request.onload = function(event) {
        let result = document.getElementById('result');
        result.innerHTML = "piyopiyo";

        if (req.readyState == 4) {
            if (req.status == 200) {
                console.log("SUCCESS");
                result.innerHTML = req.responseText;
            } else {
                console.log("ERROR");
            }
        }
    };

    request.send(null);
}