$(document).ready(function(){
    $("#submit_transaction").on('click', function(){
        console.log(getFormData($("#transaction")))
        $.ajax({
            url: 'http://localhost:8080/transaction',
            contentType: "application/json",
            type : "POST",
            dataType : 'json',
            data : getFormData($("#transaction")),
            success : function(result) {
                console.log(result)
                window.location.href = 'http://localhost:8080/transactions'
            },
            error: function(xhr, resp, text, data) {
                $("#messages").text("usuario no valido")
                console.log(xhr, resp, text);
            }
        })
    });
});



function createTransaction(){
    
}