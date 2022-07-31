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

## getUser

<a id="opIdgetUser"></a>

`GET /{username}`

*Get user*

<h3 id="getuser-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|username|path|string|true|none|

> Example responses

> 200 Response

```json
{
  "username": "string",
  "fname": "string",
  "lname": "string",
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
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[User](#schemauser)|
|default|Default|Unexpected Error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## createUser

<a id="opIdcreateUser"></a>

`POST /{username}`

*Create user*

> Body parameter

```json
{
  "username": "string",
  "fname": "string",
  "lname": "string",
  "displayIDs": [
    0
  ],
  "photoURL": "string",
  "socialLinks": [
    "string"
  ]
}
```

<h3 id="createuser-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|username|path|string|true|none|
|body|body|[User](#schemauser)|true|none|

> Example responses

> 200 Response

```json
{
  "username": "string",
  "fname": "string",
  "lname": "string",
  "displayIDs": [
    0
  ],
  "photoURL": "string",
  "socialLinks": [
    "string"
  ]
}
```

<h3 id="createuser-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[User](#schemauser)|
|default|Default|Unexpected Error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## updateUser

<a id="opIdupdateUser"></a>

`PATCH /{username}`

*Update user*

> Body parameter

```json
{
  "username": "string",
  "fname": "string",
  "lname": "string",
  "displayIDs": [
    0
  ],
  "photoURL": "string",
  "socialLinks": [
    "string"
  ]
}
```

<h3 id="updateuser-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|username|path|string|true|none|
|body|body|[User](#schemauser)|true|none|

> Example responses

> 200 Response

```json
{
  "username": "string",
  "fname": "string",
  "lname": "string",
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
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[User](#schemauser)|
|default|Default|Unexpected Error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## createDisplay

<a id="opIdcreateDisplay"></a>

`POST /d`

*Create display*

<h3 id="createdisplay-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|display|query|[Display](#schemadisplay)|false|none|

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

<h3 id="createdisplay-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[Display](#schemadisplay)|
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
|displayID|path|integer(uint64)|true|none|

> Example responses

> 200 Response

```json
[
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
]
```

<h3 id="getdisplay-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|Inline|
|default|Default|Unexpected Error|[Error](#schemaerror)|

<h3 id="getdisplay-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[[Display](#schemadisplay)]|false|none|none|
|» id|integer(uint64)|false|none|none|
|» username|string|false|none|none|
|» title|string|false|none|none|
|» description|string|false|none|none|
|» photoURL|string|false|none|none|
|» itemIDs|[integer]|false|none|none|

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

<h3 id="updatedisplay-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|displayID|path|integer(uint64)|true|none|
|body|body|[Display](#schemadisplay)|true|none|

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
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[Display](#schemadisplay)|
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
|displayID|path|integer(uint64)|true|none|

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
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|None|
|default|Default|Unexpected Error|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## createItem

<a id="opIdcreateItem"></a>

`POST /i`

*Create Item*

<h3 id="createitem-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|item|query|[Item](#schemaitem)|false|none|

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
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[Item](#schemaitem)|
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
|itemID|path|integer(uint64)|true|none|

> Example responses

> 200 Response

```json
[
  {
    "id": 0,
    "externalLink": "string",
    "socialPostLink": "string",
    "photoURL": "string",
    "username": "string",
    "displayID": 0
  }
]
```

<h3 id="getitem-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|Inline|
|default|Default|Unexpected Error|[Error](#schemaerror)|

<h3 id="getitem-responseschema">Response Schema</h3>

Status Code **200**

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|*anonymous*|[[Item](#schemaitem)]|false|none|none|
|» id|integer(uint64)|false|none|none|
|» externalLink|string|false|none|none|
|» socialPostLink|string|false|none|none|
|» photoURL|string|false|none|none|
|» username|string|false|none|none|
|» displayID|integer(uint64)|false|none|none|

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
  "id": 0,
  "externalLink": "string",
  "socialPostLink": "string",
  "photoURL": "string",
  "username": "string",
  "displayID": 0
}
```

<h3 id="updateitem-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|itemID|path|integer(uint64)|true|none|
|body|body|[Item](#schemaitem)|true|none|

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
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[Item](#schemaitem)|
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
|itemID|path|integer(uint64)|true|none|

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

<h2 id="tocS_User">User</h2>
<!-- backwards compatibility -->
<a id="schemauser"></a>
<a id="schema_User"></a>
<a id="tocSuser"></a>
<a id="tocsuser"></a>

```json
{
  "username": "string",
  "fname": "string",
  "lname": "string",
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
|username|string|false|none|none|
|fname|string|false|none|none|
|lname|string|false|none|none|
|displayIDs|[integer]|false|none|none|
|photoURL|string|false|none|none|
|socialLinks|[string]|false|none|none|

<h2 id="tocS_Display">Display</h2>
<!-- backwards compatibility -->
<a id="schemadisplay"></a>
<a id="schema_Display"></a>
<a id="tocSdisplay"></a>
<a id="tocsdisplay"></a>

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
|id|integer(uint64)|false|none|none|
|username|string|false|none|none|
|title|string|false|none|none|
|description|string|false|none|none|
|photoURL|string|false|none|none|
|itemIDs|[integer]|false|none|none|

<h2 id="tocS_Item">Item</h2>
<!-- backwards compatibility -->
<a id="schemaitem"></a>
<a id="schema_Item"></a>
<a id="tocSitem"></a>
<a id="tocsitem"></a>

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
|id|integer(uint64)|false|none|none|
|externalLink|string|false|none|none|
|socialPostLink|string|false|none|none|
|photoURL|string|false|none|none|
|username|string|false|none|none|
|displayID|integer(uint64)|false|none|none|

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

