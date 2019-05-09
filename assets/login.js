

$(document).ready(function(){
    // click on button submit
    $("#submit").on('click', function(){
        // send ajax
        $.ajax({
            url: 'loginn', // url where to submit the request
            contentType: "application/json",
            type : "POST", // type of action POST || GET
            dataType : 'json', // data type
            data : getFormData($("#login")), // post data || get data
            success : function(result) {
                // you can see the result from the console
                // tab of the developer tools
                console.log(result);
                window.location.href = 'transactions'
            },
            error: function(xhr, resp, text, data) {
                $("#messages").text("usuario no valido")
                console.log(xhr, resp, text);
            }
        })
    });
});

function getFormData($form){
    var unindexed_array = $form.serializeArray();
    var indexed_array = {};

    $.map(unindexed_array, function(n, i){
        indexed_array[n["name"]] = n['value'];
    });

    return JSON.stringify(indexed_array);
}