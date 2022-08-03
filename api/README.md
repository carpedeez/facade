<!-- Generator: Widdershins v4.0.1 -->

<h1 id="facade">Facade v0.0.1</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

This is the public API for Facade.

Base URLs:

* <a href="https://facade.carpedeez.io/v0">https://facade.carpedeez.io/v0</a>

# Authentication

* API Key (api_key)
    - Parameter Name: **api_key**, in: header. 

<h1 id="facade-default">Default</h1>

## uploadFile

<a id="opIduploadFile"></a>

`POST /assets`

*Upload file*

> Body parameter

```yaml
component: string
file: string

```

<h3 id="uploadfile-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Upload](#schemaupload)|true|none|

> Example responses

> 200 Response

```json
"string"
```

<h3 id="uploadfile-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|string|
|default|Default|Unexpected Error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## createUser

<a id="opIdcreateUser"></a>

`POST /u`

*Create user*

> Body parameter

```json
{
  "username": "string",
  "firstName": "string",
  "lastName": "string"
}
```

<h3 id="createuser-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[PostUser](#schemapostuser)|true|none|

> Example responses

> default Response

```json
{
  "code": 0,
  "message": "string"
}
```

<h3 id="createuser-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|204|[No Content](https://tools.ietf.org/html/rfc7231#section-6.3.5)|OK|None|
|default|Default|Unexpected Error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## getUser

<a id="opIdgetUser"></a>

`GET /u/{username}`

*Get user*

<h3 id="getuser-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|username|path|string|true|none|

> Example responses

> 200 Response

```json
{
  "id": 0,
  "username": "string",
  "firstName": "string",
  "lastName": "string",
  "displayIDs": [
    0
  ],
  "photoURL": "string",
  "socialLinks": [
    "string"
  ]
}
```

<h3 id="getuser-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[GetUser](#schemagetuser)|
|default|Default|Unexpected Error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## updateUser

<a id="opIdupdateUser"></a>

`PATCH /u/{username}`

*Update user*

> Body parameter

```json
{
  "username": "string",
  "firstName": "string",
  "lastName": "string",
  "socialLinks": [
    "string"
  ]
}
```

<h3 id="updateuser-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|username|path|string|true|none|
|body|body|[PatchUser](#schemapatchuser)|true|none|

> Example responses

> 200 Response

```json
{
  "id": 0,
  "username": "string",
  "firstName": "string",
  "lastName": "string",
  "displayIDs": [
    0
  ],
  "photoURL": "string",
  "socialLinks": [
    "string"
  ]
}
```

<h3 id="updateuser-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[GetUser](#schemagetuser)|
|default|Default|Unexpected Error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## createDisplay

<a id="opIdcreateDisplay"></a>

`POST /d`

*Create display*

> Body parameter

```json
{
  "title": "string",
  "description": "string"
}
```

<h3 id="createdisplay-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[PostDisplay](#schemapostdisplay)|true|none|

> Example responses

> 200 Response

```json
0
```

<h3 id="createdisplay-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|integer|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|None|
|default|Default|Unexpected Error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## getDisplay

<a id="opIdgetDisplay"></a>

`GET /d/{displayID}`

*Get display*

<h3 id="getdisplay-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|displayID|path|integer(int64)|true|none|

> Example responses

> 200 Response

```json
{
  "id": 0,
  "username": "string",
  "title": "string",
  "description": "string",
  "photoURL": "string",
  "itemIDs": [
    0
  ]
}
```

<h3 id="getdisplay-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[GetDisplay](#schemagetdisplay)|
|default|Default|Unexpected Error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## updateDisplay

<a id="opIdupdateDisplay"></a>

`PATCH /d/{displayID}`

*Update display*

> Body parameter

```json
{
  "title": "string",
  "description": "string"
}
```

<h3 id="updatedisplay-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|displayID|path|integer(int64)|true|none|
|body|body|[PatchDisplay](#schemapatchdisplay)|true|none|

> Example responses

> 200 Response

```json
{
  "id": 0,
  "username": "string",
  "title": "string",
  "description": "string",
  "photoURL": "string",
  "itemIDs": [
    0
  ]
}
```

<h3 id="updatedisplay-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[GetDisplay](#schemagetdisplay)|
|default|Default|Unexpected Error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## deleteDisplay

<a id="opIddeleteDisplay"></a>

`DELETE /d/{displayID}`

*Delete display*

<h3 id="deletedisplay-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|displayID|path|integer(int64)|true|none|

> Example responses

> default Response

```json
{
  "code": 0,
  "message": "string"
}
```

<h3 id="deletedisplay-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|204|[No Content](https://tools.ietf.org/html/rfc7231#section-6.3.5)|No Content|None|
|default|Default|Unexpected Error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## createItem

<a id="opIdcreateItem"></a>

`POST /i`

*Create Item*

> Body parameter

```json
{
  "externalLink": "string",
  "socialPostLink": "string"
}
```

<h3 id="createitem-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[PostItem](#schemapostitem)|true|none|

> Example responses

> 200 Response

```json
{
  "id": 0,
  "externalLink": "string",
  "socialPostLink": "string",
  "photoURL": "string",
  "username": "string",
  "displayID": 0
}
```

<h3 id="createitem-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[GetItem](#schemagetitem)|
|default|Default|Unexpected Error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## getItem

<a id="opIdgetItem"></a>

`GET /i/{itemID}`

*Get item*

<h3 id="getitem-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|itemID|path|integer(int64)|true|none|

> Example responses

> 200 Response

```json
{
  "id": 0,
  "externalLink": "string",
  "socialPostLink": "string",
  "photoURL": "string",
  "username": "string",
  "displayID": 0
}
```

<h3 id="getitem-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[GetItem](#schemagetitem)|
|default|Default|Unexpected Error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## updateItem

<a id="opIdupdateItem"></a>

`PATCH /i/{itemID}`

*Update item*

> Body parameter

```json
{
  "externalLink": "string",
  "socialPostLink": "string",
  "photoURL": "string"
}
```

<h3 id="updateitem-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|itemID|path|integer(int64)|true|none|
|body|body|[PatchItem](#schemapatchitem)|true|none|

> Example responses

> 200 Response

```json
{
  "id": 0,
  "externalLink": "string",
  "socialPostLink": "string",
  "photoURL": "string",
  "username": "string",
  "displayID": 0
}
```

<h3 id="updateitem-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[GetItem](#schemagetitem)|
|default|Default|Unexpected Error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## deleteItem

<a id="opIddeleteItem"></a>

`DELETE /i/{itemID}`

*Delete item*

<h3 id="deleteitem-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|itemID|path|integer(int64)|true|none|

> Example responses

> default Response

```json
{
  "code": 0,
  "message": "string"
}
```

<h3 id="deleteitem-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|None|
|default|Default|Unexpected Error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_Upload">Upload</h2>
<!-- backwards compatibility -->
<a id="schemaupload"></a>
<a id="schema_Upload"></a>
<a id="tocSupload"></a>
<a id="tocsupload"></a>

```json
{
  "component": "string",
  "file": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|component|string|true|none|none|
|file|string(binary)|true|none|none|

<h2 id="tocS_GetUser">GetUser</h2>
<!-- backwards compatibility -->
<a id="schemagetuser"></a>
<a id="schema_GetUser"></a>
<a id="tocSgetuser"></a>
<a id="tocsgetuser"></a>

```json
{
  "id": 0,
  "username": "string",
  "firstName": "string",
  "lastName": "string",
  "displayIDs": [
    0
  ],
  "photoURL": "string",
  "socialLinks": [
    "string"
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|integer(int64)|true|none|none|
|username|string|true|none|none|
|firstName|string|true|none|none|
|lastName|string|true|none|none|
|displayIDs|[integer]|true|none|none|
|photoURL|string|true|none|none|
|socialLinks|[string]|true|none|none|

<h2 id="tocS_PostUser">PostUser</h2>
<!-- backwards compatibility -->
<a id="schemapostuser"></a>
<a id="schema_PostUser"></a>
<a id="tocSpostuser"></a>
<a id="tocspostuser"></a>

```json
{
  "username": "string",
  "firstName": "string",
  "lastName": "string"
}

```

registration

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|username|string|true|none|none|
|firstName|string|true|none|none|
|lastName|string|true|none|none|

<h2 id="tocS_PatchUser">PatchUser</h2>
<!-- backwards compatibility -->
<a id="schemapatchuser"></a>
<a id="schema_PatchUser"></a>
<a id="tocSpatchuser"></a>
<a id="tocspatchuser"></a>

```json
{
  "username": "string",
  "firstName": "string",
  "lastName": "string",
  "socialLinks": [
    "string"
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|username|string|false|none|none|
|firstName|string|false|none|none|
|lastName|string|false|none|none|
|socialLinks|[string]|false|none|none|

<h2 id="tocS_GetDisplay">GetDisplay</h2>
<!-- backwards compatibility -->
<a id="schemagetdisplay"></a>
<a id="schema_GetDisplay"></a>
<a id="tocSgetdisplay"></a>
<a id="tocsgetdisplay"></a>

```json
{
  "id": 0,
  "username": "string",
  "title": "string",
  "description": "string",
  "photoURL": "string",
  "itemIDs": [
    0
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|integer(int64)|true|none|none|
|username|string|true|none|none|
|title|string|true|none|none|
|description|string|true|none|none|
|photoURL|string|true|none|none|
|itemIDs|[integer]|true|none|none|

<h2 id="tocS_PostDisplay">PostDisplay</h2>
<!-- backwards compatibility -->
<a id="schemapostdisplay"></a>
<a id="schema_PostDisplay"></a>
<a id="tocSpostdisplay"></a>
<a id="tocspostdisplay"></a>

```json
{
  "title": "string",
  "description": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|title|string|true|none|none|
|description|string|true|none|none|

<h2 id="tocS_PatchDisplay">PatchDisplay</h2>
<!-- backwards compatibility -->
<a id="schemapatchdisplay"></a>
<a id="schema_PatchDisplay"></a>
<a id="tocSpatchdisplay"></a>
<a id="tocspatchdisplay"></a>

```json
{
  "title": "string",
  "description": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|title|string|false|none|none|
|description|string|false|none|none|

<h2 id="tocS_GetItem">GetItem</h2>
<!-- backwards compatibility -->
<a id="schemagetitem"></a>
<a id="schema_GetItem"></a>
<a id="tocSgetitem"></a>
<a id="tocsgetitem"></a>

```json
{
  "id": 0,
  "externalLink": "string",
  "socialPostLink": "string",
  "photoURL": "string",
  "username": "string",
  "displayID": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|integer(int64)|true|none|none|
|externalLink|string|true|none|none|
|socialPostLink|string|true|none|none|
|photoURL|string|true|none|none|
|username|string|true|none|none|
|displayID|integer(int64)|true|none|none|

<h2 id="tocS_PostItem">PostItem</h2>
<!-- backwards compatibility -->
<a id="schemapostitem"></a>
<a id="schema_PostItem"></a>
<a id="tocSpostitem"></a>
<a id="tocspostitem"></a>

```json
{
  "externalLink": "string",
  "socialPostLink": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|externalLink|string|true|none|none|
|socialPostLink|string|true|none|none|

<h2 id="tocS_PatchItem">PatchItem</h2>
<!-- backwards compatibility -->
<a id="schemapatchitem"></a>
<a id="schema_PatchItem"></a>
<a id="tocSpatchitem"></a>
<a id="tocspatchitem"></a>

```json
{
  "externalLink": "string",
  "socialPostLink": "string",
  "photoURL": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|externalLink|string|false|none|none|
|socialPostLink|string|false|none|none|
|photoURL|string|false|none|none|

<h2 id="tocS_Error">Error</h2>
<!-- backwards compatibility -->
<a id="schemaerror"></a>
<a id="schema_Error"></a>
<a id="tocSerror"></a>
<a id="tocserror"></a>

```json
{
  "code": 0,
  "message": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|code|integer(int32)|true|none|none|
|message|string|true|none|none|

