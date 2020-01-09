
# Vk Parser Service.

 
## Description.

  
UNDER CONSTRUCTION.

  

## List of endpoints.


[1. **POST**  /getProfiles](#get-profiles)
  

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
		"token": ""
	},
	"server": {
		"host": "",
		"port": ""
	}
}
```

#### Default:

* vk.version --- Version of [vk api](https://vk.com/dev/versions). Type: *String.*  **Default:** "5.103".
* vk.token --- Your private api [token](https://vk.com/dev/permissions?f=1.%20%D0%9F%D1%80%D0%B0%D0%B2%D0%B0%20%D0%B4%D0%BE%D1%81%D1%82%D1%83%D0%BF%D0%B0%20%D0%B4%D0%BB%D1%8F%20%D1%82%D0%BE%D0%BA%D0%B5%D0%BD%D0%B0%20%D0%BF%D0%BE%D0%BB%D1%8C%D0%B7%D0%BE%D0%B2%D0%B0%D1%82%D0%B5%D0%BB%D1%8F). Type: *String.* 
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
	"id": string // User id
	"intersect_number": int //current minimum number(N) of occurrences, N > 1
	"sex": int // 1 - woman, 2 - man
	"Message": bool // can write private message
}
```  

#### Responce .

`1. StatusCode = 200:`

*Body:*
```
{
	"text": "",
	"responce": 
}
```
#### Cases.
1. - "text": **"The list is empty"**. Means that intersection is empty with current params from request.
	- "responce": ""
2. - "text": **"We found N people"**. Means that intersection is not empty and contains N people with current params from request.
	- "responce": [id_1, id_2, ..., id_N]. Contains people's id from intersection. *Type* **[]int64**.

#### Examples.
 Returned group members with ids: **[1, 2, 4, 6, 1, 3, 1, 2, 3]**.
 - IF *intersect_number* = **2** in request ---> intersection is **[1, 2, 3]**.
- ELSE IF *intersect_number* = **3** in request ---> intersection is **[1]**.
- ELSE IF *intersect_number* = **4** in request ---> intersection is **[ ]**.

**Other parameters from the request impose additional constraints.**

## TODO
-  [x] Parsing members.
-  [x] Intersection groups.
-  [ ] **Tests** for intersection methods.
-  [ ] Parse posts.
-  [ ] **Tests** for posts methods.
-  [ ] Parse comments.
-  [ ] **Tests** for comments methods.
-  [ ] Parse what user likes.
-  [ ] **Tests** for like methods.
