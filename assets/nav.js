function changeDivContent() { 
    document.getElementById("nav").innerHTML = "<nav class='nav nav-pills nav-justified'>"+
    "<a class='nav-item nav-link' href='http://localhost:8080/stock'><ion-icon size='large' name='podium'></ion-icon></a>"+
    "<a class='nav-item nav-link' href='http://localhost:8080/transactions'><ion-icon size='large' name='paper-plane'></ion-icon></a>"+
    "<a class='nav-item nav-link' href='http://localhost:8080/products'><ion-icon size='large' name='logo-buffer'></ion-icon></a>"+
    "<a class='nav-item nav-link' href='http://localhost:8080/users'><ion-icon size='large' name='contacts'></ion-icon></a>"+
    "<a class='nav-item nav-link' onclick='logout()' ><ion-icon size='large' name='log-out'></ion-icon></a>"+
"</nav>";
};

changeDivContent()