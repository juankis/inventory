$(document).ready(function(){
    $("#submit_transaction").on('click', function(){
        console.log(getFormData($("#transaction")))
        $.ajax({
            url: 'http://3.215.116.162:8081/transaction',
            contentType: "application/json",
            type : "POST",
            dataType : 'json',
            data : getFormData($("#transaction")),
            success : function(result) {
                console.log(result)
                window.location.href = 'http://3.215.116.162:8081/transactions'
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