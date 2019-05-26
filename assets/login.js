$(document).ready(function(){
    login()
});

function login(){
    $("#submit").on('click', function(){
        $.ajax({
            url: 'login', // url where to submit the request
            contentType: "application/json",
            type : "POST", // type of action POST || GET
            dataType : 'json', // data type
            data : getFormData($("#login")), // post data || get data
            success : function(result) {
                console.log(result);
                sessionStorage.setItem("user_id", "1");
                sessionStorage.setItem("store_id", "1");
                window.location.href = 'transactions'
            },
            error: function(xhr, resp, text, data) {
                $("#messages").text("usuario no valido")
                console.log(xhr, resp, text);
            }
        })
    });
}