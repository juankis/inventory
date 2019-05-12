$(document).ready(function(){
    $('#users').DataTable({
        "searching": false,
        "bLengthChange": false
    });
   getUsers()
});

function getUsers(){
    $.ajax({
        url: 'http://localhost:8080/user',
        contentType: "application/json",
        type : "GET",
        dataType : 'json',
        success : function(result) {
            console.log(result)
            loadUsers(result)
        },
        error: function(xhr, resp, text, data) {
            console.log(xhr, resp, text);
        }
    })
}

function loadUsers(users){
    var datatable = $( '#users' ).DataTable();
    var lista  = new Array()
    datatable.clear()
    jQuery.each(users, function(i, val) {
        l = Object.values(val)
        l.push("<button type='submit' class='btn btn-info btn-sm btn-block'>edit</button>")
        l.push("<button type='submit' class='btn btn-danger btn-sm btn-block'>delete</button>")
        lista.push(l)
        console.log(lista)
    });
    datatable.rows.add(lista)
    datatable.draw();
    console.log(lista)  
}
