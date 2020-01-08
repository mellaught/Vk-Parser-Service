
# Vk Parser Service.

 
## Description.

  
UNDER CONSTRUCTION.

  

## List of endpoints.


[1. **POST:**  /getProfiles](#get-profiles)
  

## TORUN.


-  `go get github.com/mrKitikat/Vk-Parser-Service`.

- cd $GOPATH/src/github.com/mrkitikat/vk-parser-service/src

-  `dep ensure`.
- `go run main.go`

  

### Config file.

Config in `config.json`.

 
```
{
	"name": "Vk Parser Service",
	"vk": {
		"url": "https://api.vk.com/method/",
		"version": "",
		"token": "PUT YOUR TOKEN HERE"
	},
	"server": {
		"host": "",
		"port": ""
	}
}
```

#### Default:

* vk.version --- Version of [vk api](https://vk.com/dev/versions). Type: *String.*  **Default:** "5.103".
* VK.token --- Your private api [token](https://vk.com/dev/permissions?f=1.%20%D0%9F%D1%80%D0%B0%D0%B2%D0%B0%20%D0%B4%D0%BE%D1%81%D1%82%D1%83%D0%BF%D0%B0%20%D0%B4%D0%BB%D1%8F%20%D1%82%D0%BE%D0%BA%D0%B5%D0%BD%D0%B0%20%D0%BF%D0%BE%D0%BB%D1%8C%D0%B7%D0%BE%D0%B2%D0%B0%D1%82%D0%B5%D0%BB%D1%8F). Type: *String.* 
* server.host --- Host of **Vk Parser Service**. Type: *String.*  **Default:** "localhost".
* server.port --- Port of **Vk Parser Service**. Type *String.*  **Default:** "8080".
 
  
<div  id='get-profiles'/>

  

### 1. *POST* /getProfiles

  

#### Description.

UNDER CONSTRUCTION.

#### Request.

- "Content-Type", "application/json"

- **POST** `http://service.host:service.port/getProfiles`

*Body:*
```
{

	"id": string, // User id
	"intersect_number": int, //current minimum number(N) of occurrences
	"sex": int, // 1 - woman, 2 - man
	"Message": bool // can write private message
}
```  

#### Responce .

`1. StatusCode = 200:`

*Body:*
```
{
	"text": "Started to follow tx",
	"responce:"
}
```

## TODO
-  [x] Parsing members.
-  [x] Intersectaion groups.
-  [ ] Parse posts.
-  [ ] Parse comments.
-  [ ] Check likes.
