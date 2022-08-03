var editor
function run() { 
    line = editor.getValue().split(' ')

    action = line[0]
    table = line[1]
    values = [{}]
    len = line[2].split(',').length
    array = line[2].split(',')
    for(i = 0; i < len; i++){
        prop = array[i].split(':')[0]
        val = array[i].split(':')[1]
        values[i] = { 
            [prop+" "+String(val)]: "" 
        };
    }
    // values is an array of 
    // [{prop: val},{pro: val}...]
    console.log("first js ", values)
   // Name_is_john_and_Age_is_19
    var result = JSON.stringify( 
        {
            'Action' : action,
            'Table'  : table,
            '-'      : values
            // 'More':  
        }
    );
    console.log("from js ", result);
    $.ajax({
        type: "POST", //rest Type
        dataType: 'json', //mispelled
        url: "http://localhost:5000/json",
        async: false,
        contentType: "application/json;", //  charset=utf-8
        data: result
    });
}

window.onload= function (){
    editor = CodeMirror.fromTextArea(document.getElementById('editor'), {
        mode: "python",
        theme: "darcula"
    });

}