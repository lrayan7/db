# db - (in progress...)
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
upon intercepting the request at DB side, request is parsed and a worker spawns if needed to handle it <br/>

<br/>

### Data Structure ###

every struct is stored in a map with a pointer<br/>
```
DB:   tablename1: *table{      entry1: *entry{} 
                               entry2: *entry{} 
                              .
                              .               } 
                        
      tablename2: *table{} <br/>
          .
          .
          .
          
```
## overview
![alt text](https://github.com/lrayan7/db/blob/main/imgs/newdiag.png) 


