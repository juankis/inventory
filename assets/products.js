$(document).ready(function(){
    $('#products').DataTable({
        "searching": false,
        "bLengthChange": false
    });
   getProducts()
});

function getProducts(){
    $.ajax({
        url: 'http://3.215.116.162:8081/product',
        contentType: "application/json",
        type : "GET",
        dataType : 'json',
        success : function(result) {
            console.log(result)
            loadProducts(result)
        },
        error: function(xhr, resp, text, data) {
            console.log(xhr, resp, text);
        }
    })
}

function loadProducts(products){
    var datatable = $( '#products' ).DataTable();
    var lista  = new Array()
    datatable.clear()
    jQuery.each(products, function(i, val) {
        l = Object.values(val)
        //l.push("<button type='submit' class='btn btn-info btn-sm btn-block'>edit</button>")
        //l.push("<button type='submit' class='btn btn-danger btn-sm btn-block'>delete</button>")
        lista.push(l)
        console.log(lista)
    });
    datatable.rows.add(lista)
    datatable.draw();
    console.log(lista)  
}
