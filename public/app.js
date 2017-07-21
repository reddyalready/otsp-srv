var request = new XMLHttpRequest();
request.open('GET', '/projects', true);

request.onload = function() {
    if(request.status >= 200 && request.status < 400) {
        var data = JSON.parse(request.responseText)
        let html = ""
        for(let row of data) {
            html += row.name+"<br/>"
        }
        document.getElementById("app").innerHTML = html
    }
}

request.onerror = function(err) {
    console.log(err)
}

function main() {
    request.send()
}