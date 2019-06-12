$(document).ready(function(){
    $("#submit_store").on('click', function(){
        console.log(getFormData($("#store")))
        $.ajax({
            url: 'http://3.215.116.162:8081/store',
            contentType: "application/json",
            type : "POST",
            dataType : 'json',
            data : getFormData($("#store")),
            success : function(result) {
                console.log(result)
                window.location.href = 'http://3.215.116.162:8081/stores'
            },
            error: function(xhr, resp, text, data) {
                $("#messages").text("usuario no valido")
                console.log(xhr, resp, text);
            }
        })
    });
});
