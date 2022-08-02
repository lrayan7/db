var editor

function run() {   
    
    // var result = JSON.stringify( { 'code' : editor.getValue() } );
    
    $.ajax({
        type: "POST", //rest Type
        // dataType: 'json', //mispelled
        url: "http://localhost:8080/json",
        // async: false,
        // contentType: "application/json;", //  charset=utf-8
        // data: result
    });

}

window.onload = function (){
    editor = CodeMirror.fromTextArea(document.getElementById('editor'), {
        mode: "python",
        theme: "darcula-dark"
    });
}