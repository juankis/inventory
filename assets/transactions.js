$(document).ready(function(){
    $("#submit_transaction").on('click', function(){
        createTransaction()
    }); 

    $('#transactions').DataTable({
        "searching": false,
        "bLengthChange": false
        //"paging": false,
        //"ordering": false
    });
   getTransactions()
});

function getTransactions(){
    $.ajax({
        url: 'http://localhost:8081/transaction',
        contentType: "application/json",
        type : "GET",
        dataType : 'json',
        success : function(result) {
            console.log(result)
            loadTransactions(result)
        },
        error: function(xhr, resp, text, data) {
            console.log(xhr, resp, text);
        }
    })
}

function loadTransactions(transactions){
    var datatable = $( '#transactions' ).DataTable();
    var lista  = new Array()
    datatable.clear()
    jQuery.each(transactions, function(i, val) {
        l = Object.values(val)
        l.push("<button type='submit' class='btn btn-primary btn-sm btn-block'>Ok</button>")
        lista.push(l)
        console.log(lista)
    });
    datatable.rows.add(lista)
    datatable.draw();
    console.log(lista)  
}

function createTransaction(){
    console.log(getFormData($("#transaction")))
        $.ajax({
            url: 'http://localhost:8081/transaction',
            contentType: "application/json",
            type : "POST",
            dataType : 'json',
            data : getFormData($("#transaction")),
            success : function(result) {
                console.log(result)
                window.location.href = 'http://localhost:8081/transactions'
            },
            error: function(xhr, resp, text, data) {
                $("#messages").text("usuario no valido")
                console.log(xhr, resp, text);
            }
        })
}