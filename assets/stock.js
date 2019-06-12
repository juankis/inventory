function getStock(storeID){
    console.log("storeID:",storeID)
    $.ajax({
        url: 'http://3.215.116.162:8081/store_stock/'+storeID,
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
    var datatable = $( '#stocks' ).DataTable();
    var lista  = new Array()
    datatable.clear()
    jQuery.each(transactions, function(i, val) {
        l = Object.values(val)
        lista.push(l)
        console.log(lista)
    });
    datatable.rows.add(lista)
    datatable.draw();
    console.log(lista)  
}
