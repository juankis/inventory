# Inventory
# Requirements
- go get github.com/go-sql-driver/mysql
- go get github.com/gin-gonic/gin
- go get github.com/jmoiron/sqlx
# Install
- go get github.com/juankis/inventory
# Run
- In inventory folder run in the console: go run api/src/main/main.go

# Endpoints

- [<code>POST</code> transaction](#post-transaction)
- [<code>GET</code> transaction](#get-transaction)
- [<code>GET</code> all transactions](#get-all-transaction)
- [<code>DELETE</code> transaction](#delete-transaction)
- [<code>PUT</code> transaction](#put-transaction)

## Examples
# Post transaction
- **Request**
``` json
{
	"data":{
		"quantity":	50, 
		"movement": 1, 
		"user_creator":3, 
		"product_id": 5
		}
}
```
- **Response**
``` json
{
    "id": 5,
    "quantity": 50,
    "movement": 1,
    "user_creator": 3,
    "user_confirm": null,
    "product_id": 5
}
``` 
# Get all transaction
- **Request**
- <code>localhost:8080/transaction</code>
- **Response**
``` json
[
    {
        "id": 1,
        "quantity": 40,
        "movement": 1,
        "user_creator": 1,
        "user_confirm": 3,
        "product_id": 5,
        "created_at": "2019-05-05 17:15:08",
        "updated_at": "2019-05-05 17:22:45"
    },
    {
        "id": 3,
        "quantity": 50,
        "movement": 1,
        "user_creator": 3,
        "user_confirm": null,
        "product_id": 5,
        "created_at": "2019-05-05 17:27:48",
        "updated_at": "2019-05-05 17:27:48"
    }
]    
```

# Get transaction
- **Request**
- <code>localhost:8080/transaction/1</code>
- **Response**
``` json
{
        "id": 1,
        "quantity": 40,
        "movement": 1,
        "user_creator": 1,
        "user_confirm": 3,
        "product_id": 5,
        "created_at": "2019-05-05 17:15:08",
        "updated_at": "2019-05-05 17:22:45"
    }
```

# Put transaction
- **Request**
- <code>localhost:8080/transaction/1</code>
``` json
{
	"data":{
		"quantity":	40, 
		"movement": 1, 
		"user_creator":1,
		"user_confirm":3,
		"product_id": 5
		}
}
```
- **Response**
``` json
{
    "id": 1,
    "quantity": 40,
    "movement": 1,
    "user_creator": 1,
    "user_confirm": 3,
    "product_id": 5
}
```

# Delete transaction
- **Request**
- <code>localhost:8080/transaction/1</code>
