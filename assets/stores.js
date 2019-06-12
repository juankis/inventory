function getStores(){
    $.ajax({
        url: 'http://3.215.116.162:8081/store',
        contentType: "application/json",
        type : "GET",
        dataType : 'json',
        success : function(result) {
            console.log(result)
            loadStores(result)
        },
        error: function(xhr, resp, text, data) {
            console.log(xhr, resp, text);
        }
    })
}

function loadStores(stores){
    var datatable = $( '#stores' ).DataTable();
    var lista  = new Array()
    datatable.clear()
    jQuery.each(stores, function(i, val) {
        l = new Array()
        l.push(val.id)
        l.push(val.name)
        //l.push("<button type='submit' class='btn btn-info btn-sm btn-block'>edit</button>")
        //l.push("<button type='submit' class='btn btn-danger btn-sm btn-block'>delete</button>")
        lista.push(l)
    });
    datatable.rows.add(lista)
    datatable.draw();
}
