$(document).ready(function(){
    $("#submit_user").on('click', function(){
        console.log(getFormData($("#user")))
        $.ajax({
            url: 'http://3.215.116.162:8081/user',
            contentType: "application/json",
            type : "POST",
            dataType : 'json',
            data : getFormData($("#user")),
            success : function(result) {
                console.log(result)
                window.location.href = 'http://3.215.116.162:8081/users'
            },
            error: function(xhr, resp, text, data) {
                $("#messages").text("usuario no valido")
                console.log(xhr, resp, text);
            }
        })
    });
});
