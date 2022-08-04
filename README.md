# db
Fun project, A database server written in Golang 

Supports the following 
  
  * noSQL requests
  * concurrent request handling
 
 ## query request format
 Supports 5 keywords: <br/>
 INIT, ADD, UPDATE, DELETE, READ
 
 Query format should be like this: <br/>
 ```<>``` indicates to ```<write here your custom name>```
 
  * ```INIT <tablename> table```
  * ```ADD <tablename> <primekey>,<rest of data>```  
  * ```READ <primekey> <tablename>```
    
## How it works 
flask app listens on query input, that sends JSON formatted ```{Action:_,Table:_,Value:_}``` request to DB server <br/>
for example, if client writes ```INIT users table``` - ```{Action:INIT,Table:users,Value:_}``` will be sent to the DB <br/>
upon intercepting the request at DB side, request is parsed and unmarshalled into the following struct: <br/>
```golang
type msg struct{
  Action string `json:"Action"`
  Table string `json:"Table"`
  Value map[string]interface{} `json:"Value"`
}
```

