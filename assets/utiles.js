function getFormData($form){
    var unindexed_array = $form.serializeArray();
    var indexed_array = {};

    $.map(unindexed_array, function(n, i){
        indexed_array[n["name"]] = n['value'];
    });
    console.log(indexed_array)
    return JSON.stringify(indexed_array);
}

function logout(){
    sessionStorage.clear();
    window.location.href = 'http://3.215.116.162:8081/login'
}

function checkLogin(){
    console.log()
    if(sessionStorage.getItem("user_id") === null){
        window.location.href = 'http://3.215.116.162:8081/login'
    }
}

function loadInSelect(idSelect, elements){
    $.ajax({
        url: 'http://3.215.116.162:8081/'+elements,
        contentType: "application/json",
        type : "GET",
        dataType : 'json',
        success : function(result) {
            console.log(result)
            jQuery.each(result, function(i, val) {
            $(idSelect).append(new Option(val.name, val.id));
            })
        },
        error: function(xhr, resp, text, data) {
            console.log(xhr, resp, text);
        }
    })
}