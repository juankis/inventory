$(document).ready(function(){
    $("#submit_product").on('click', function(){
        console.log(getFormData($("#product")))
        $.ajax({
            url: 'http://localhost:8081/product',
            contentType: "application/json",
            type : "POST",
            dataType : 'json',
            data : getFormData($("#product")),
            success : function(result) {
                console.log(result)
                window.location.href = 'http://localhost:8081/products'
            },
            error: function(xhr, resp, text, data) {
                $("#messages").text("usuario no valido")
                console.log(xhr, resp, text);
            }
        })
    });
});
