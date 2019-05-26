
function getFormData($form){
    var unindexed_array = $form.serializeArray();
    var indexed_array = {};

    $.map(unindexed_array, function(n, i){
        indexed_array[n["name"]] = n['value'];
    });
    console.log(indexed_array)
    return JSON.stringify(indexed_array);
}

function logout(){
    sessionStorage.clear();
    window.location.href = 'login'
}