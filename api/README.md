<!-- Generator: Widdershins v4.0.1 -->

<h1 id="facade">Facade v0.0.1</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

This is the public API for Facade.

Base URLs:

* <a href="https://api.facadeapp.dev/v0">https://api.facadeapp.dev/v0</a>

<h1 id="facade-default">Default</h1>

## me

<a id="opIdme"></a>

`GET /@me`

*Get Session*

> Example responses

> 200 Response

```json
{
  "id": "string",
  "active": true,
  "expires_at": "string",
  "authenticated_at": "string",
  "authenticator_assurance_level": "string",
  "authentication_methods": [
    {
      "method": "string",
      "completed_at": "string"
    }
  ],
  "issued_at": "string",
  "identity": {
    "id": "string",
    "schema_id": "string",
    "schema_url": "string",
    "state": "string",
    "state_changed_at": "string",
    "traits": {
      "website": "string",
      "email": "string"
    },
    "verifiable_addresses": [
      {
        "id": "string",
        "value": "string",
        "verified": true,
        "via": "string",
        "status": "string",
        "created_at": "string",
        "updated_at": "string"
      }
    ],
    "recovery_addresses": [
      {
        "id": "string",
        "value": "string",
        "via": "string",
        "created_at": "string",
        "updated_at": "string"
      }
    ],
    "created_at": "string",
    "updated_at": "string"
  }
}
```

<h3 id="me-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[Session](#schemasession)|

<aside class="success">
This operation does not require authentication
</aside>

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
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[Error](#schemaerror)|

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

> 201 Response

```json
{
  "id": 0,
  "userID": "string",
  "title": "string",
  "description": "string",
  "photoURL": "string",
  "items": [
    {
      "id": 0,
      "externalLink": "string",
      "socialPostLink": "string",
      "photoURL": "string",
      "userID": "string",
      "displayID": 0
    }
  ]
}
```

<h3 id="createdisplay-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Created|[GetDisplay](#schemagetdisplay)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[Error](#schemaerror)|

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
  "userID": "string",
  "title": "string",
  "description": "string",
  "photoURL": "string",
  "items": [
    {
      "id": 0,
      "externalLink": "string",
      "socialPostLink": "string",
      "photoURL": "string",
      "userID": "string",
      "displayID": 0
    }
  ]
}
```

<h3 id="getdisplay-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[GetDisplay](#schemagetdisplay)|

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
  "userID": "string",
  "title": "string",
  "description": "string",
  "photoURL": "string",
  "items": [
    {
      "id": 0,
      "externalLink": "string",
      "socialPostLink": "string",
      "photoURL": "string",
      "userID": "string",
      "displayID": 0
    }
  ]
}
```

<h3 id="updatedisplay-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[GetDisplay](#schemagetdisplay)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[Error](#schemaerror)|

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

<h3 id="deletedisplay-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|204|[No Content](https://tools.ietf.org/html/rfc7231#section-6.3.5)|No Content|None|

<aside class="success">
This operation does not require authentication
</aside>

## createItem

<a id="opIdcreateItem"></a>

`POST /d/{displayID}/i`

*Create Item*

> Body parameter

```json
{
  "externalLink": "string"
}
```

<h3 id="createitem-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|displayID|path|integer(int64)|true|none|
|body|body|[PostItem](#schemapostitem)|true|none|

> Example responses

> 201 Response

```json
{
  "id": 0,
  "externalLink": "string",
  "socialPostLink": "string",
  "photoURL": "string",
  "userID": "string",
  "displayID": 0
}
```

<h3 id="createitem-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Created|[GetItem](#schemagetitem)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## getItem

<a id="opIdgetItem"></a>

`GET /d/{displayID}/i/{itemID}`

*Get item*

<h3 id="getitem-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|displayID|path|integer(int64)|true|none|
|itemID|path|integer(int64)|true|none|

> Example responses

> 200 Response

```json
{
  "id": 0,
  "externalLink": "string",
  "socialPostLink": "string",
  "photoURL": "string",
  "userID": "string",
  "displayID": 0
}
```

<h3 id="getitem-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[GetItem](#schemagetitem)|

<aside class="success">
This operation does not require authentication
</aside>

## updateItem

<a id="opIdupdateItem"></a>

`PATCH /d/{displayID}/i/{itemID}`

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
|displayID|path|integer(int64)|true|none|
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
  "userID": "string",
  "displayID": 0
}
```

<h3 id="updateitem-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[GetItem](#schemagetitem)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[Error](#schemaerror)|

<aside class="success">
This operation does not require authentication
</aside>

## deleteItem

<a id="opIddeleteItem"></a>

`DELETE /d/{displayID}/i/{itemID}`

*Delete item*

<h3 id="deleteitem-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|displayID|path|integer(int64)|true|none|
|itemID|path|integer(int64)|true|none|

<h3 id="deleteitem-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|204|[No Content](https://tools.ietf.org/html/rfc7231#section-6.3.5)|No Content|None|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_Session">Session</h2>
<!-- backwards compatibility -->
<a id="schemasession"></a>
<a id="schema_Session"></a>
<a id="tocSsession"></a>
<a id="tocssession"></a>

```json
{
  "id": "string",
  "active": true,
  "expires_at": "string",
  "authenticated_at": "string",
  "authenticator_assurance_level": "string",
  "authentication_methods": [
    {
      "method": "string",
      "completed_at": "string"
    }
  ],
  "issued_at": "string",
  "identity": {
    "id": "string",
    "schema_id": "string",
    "schema_url": "string",
    "state": "string",
    "state_changed_at": "string",
    "traits": {
      "website": "string",
      "email": "string"
    },
    "verifiable_addresses": [
      {
        "id": "string",
        "value": "string",
        "verified": true,
        "via": "string",
        "status": "string",
        "created_at": "string",
        "updated_at": "string"
      }
    ],
    "recovery_addresses": [
      {
        "id": "string",
        "value": "string",
        "via": "string",
        "created_at": "string",
        "updated_at": "string"
      }
    ],
    "created_at": "string",
    "updated_at": "string"
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|string|true|none|none|
|active|boolean|true|none|none|
|expires_at|string|true|none|none|
|authenticated_at|string|true|none|none|
|authenticator_assurance_level|string|true|none|none|
|authentication_methods|[object]|true|none|none|
|» method|string|true|none|none|
|» completed_at|string|true|none|none|
|issued_at|string|true|none|none|
|identity|object|true|none|none|
|» id|string|true|none|none|
|» schema_id|string|true|none|none|
|» schema_url|string|true|none|none|
|» state|string|true|none|none|
|» state_changed_at|string|true|none|none|
|» traits|object|true|none|none|
|»» website|string|true|none|none|
|»» email|string|true|none|none|
|» verifiable_addresses|[object]|true|none|none|
|»» id|string|true|none|none|
|»» value|string|true|none|none|
|»» verified|boolean|true|none|none|
|»» via|string|true|none|none|
|»» status|string|true|none|none|
|»» created_at|string|true|none|none|
|»» updated_at|string|true|none|none|
|» recovery_addresses|[object]|true|none|none|
|»» id|string|true|none|none|
|»» value|string|true|none|none|
|»» via|string|true|none|none|
|»» created_at|string|true|none|none|
|»» updated_at|string|true|none|none|
|» created_at|string|true|none|none|
|» updated_at|string|true|none|none|

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

<h2 id="tocS_GetDisplay">GetDisplay</h2>
<!-- backwards compatibility -->
<a id="schemagetdisplay"></a>
<a id="schema_GetDisplay"></a>
<a id="tocSgetdisplay"></a>
<a id="tocsgetdisplay"></a>

```json
{
  "id": 0,
  "userID": "string",
  "title": "string",
  "description": "string",
  "photoURL": "string",
  "items": [
    {
      "id": 0,
      "externalLink": "string",
      "socialPostLink": "string",
      "photoURL": "string",
      "userID": "string",
      "displayID": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|integer(int64)|true|none|none|
|userID|string|true|none|none|
|title|string|true|none|none|
|description|string|true|none|none|
|photoURL|string|true|none|none|
|items|[[GetItem](#schemagetitem)]|true|none|none|

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
  "userID": "string",
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
|userID|string|true|none|none|
|displayID|integer(int64)|true|none|none|

<h2 id="tocS_PostItem">PostItem</h2>
<!-- backwards compatibility -->
<a id="schemapostitem"></a>
<a id="schema_PostItem"></a>
<a id="tocSpostitem"></a>
<a id="tocspostitem"></a>

```json
{
  "externalLink": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|externalLink|string|true|none|none|

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

