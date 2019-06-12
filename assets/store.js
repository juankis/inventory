$(document).ready(function(){
    $("#submit_store").on('click', function(){
        console.log(getFormData($("#store")))
        $.ajax({
            url: 'http://localhost:8081/store',
            contentType: "application/json",
            type : "POST",
            dataType : 'json',
            data : getFormData($("#store")),
            success : function(result) {
                console.log(result)
                window.location.href = 'http://localhost:8081/stores'
            },
            error: function(xhr, resp, text, data) {
                $("#messages").text("usuario no valido")
                console.log(xhr, resp, text);
            }
        })
    });
});
