function run() { 
    line =  document.getElementById('inputreq').value.split(' ')
    action = line[0]
    table = line[1]
    values = [{}]
    len = line[2].split(',').length
    array = line[2].split(',')
    for(i = 0; i < len; i++){
        prop = array[i].split(':')[0]
        val = array[i].split(':')[1]
        values[i] = JSON.stringify({ 
            [prop]: val 
        });
    }
    var result = {
            Action: action,
            Table: table,
            Value: values
        }
    var xmlhttp = new XMLHttpRequest();   // new HttpRequest instance 
    var theUrl = "http://localhost:5000/json";
    xmlhttp.open("POST", theUrl);
    xmlhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xmlhttp.send(JSON.stringify(result));
}

// window.onload= function (){
//     editor = CodeMirror.fromTextArea(document.getElementById('editor'), {
//         mode: "python",
//         theme: "darcula"
//     });

// }