# Backend de una app clon de X(Twitter)

### 1. Registro de Usuario

**Endpoint:** `/registration`  
**Método:** `POST`  
**Descripción:** Crea una nueva cuenta de usuario.

---

#### **Formato de solicitud**
```json
{
  "email": "ejemplo@gmail.com",
  "password": "12345678"
}
```

---

#### **Validaciones**
1. **Email**:
   - Es un campo obligatorio.
   - Debe ser un correo electrónico válido.
   - No debe estar registrado previamente.

2. **Contraseña**:
   - Es un campo obligatorio.
   - Debe tener al menos 6 caracteres.

---

#### **Respuestas**
| Código | Descripción                                                                 |
|--------|-----------------------------------------------------------------------------|
| 201    | Usuario creado exitosamente.                                               |
| 400    | Error en los datos enviados (correo inválido, contraseña no válida, o correo ya registrado). |
| 500    | Error interno al crear el usuario.                                          |

---

### 2. Inicio de Sesión

**Endpoint:** `/login`  
**Método:** `POST`  
**Descripción:** Permite a un usuario autenticarse en el sistema.

---

#### **Formato de solicitud**
```json
{
  "email": "ejemplo@gmail.com",
  "password": "12345678"
}
```

---

#### **Validaciones**
1. **Email**:
   - Es un campo obligatorio.
   - Debe ser un correo electrónico válido registrado en el sistema.

2. **Contraseña**:
   - Es un campo obligatorio.
   - Debe coincidir con la contraseña registrada para el correo proporcionado.

---

#### **Respuestas posibles**

| Código | Descripción                                                                 |
|--------|-----------------------------------------------------------------------------|
| 201    | Inicio de sesión exitoso. Incluye un token JWT en la respuesta.            |
| 400    | Error en los datos enviados (email faltante, credenciales incorrectas, o formato inválido). |
| 500    | Error interno al generar el token JWT.                                      |

---

#### **Ejemplo de respuesta exitosa**
```json
{
  "user_id": "67832c3c0e086b9a18bb6adb",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```


### 2. Ver Perfil de Usuario

**Endpoint:** `/viewprofile`  
**Método:** `GET`  
**Descripción:** Obtiene los datos de un usuario en base al `id` proporcionado como parámetro en la query string.

---


### **Parámetros de la solicitud**

| Parámetro | Tipo   | Requerido | Descripción                     |
|-----------|--------|-----------|---------------------------------|
| `id`      | String | Sí        | Identificador único del usuario |

---

### **Respuestas posibles**

| Código | Descripción                                                                                      |
|--------|--------------------------------------------------------------------------------------------------|
| 201    | Perfil del usuario encontrado exitosamente. Devuelve el JSON con los datos del usuario.          |
| 400    | No se proporcionó el parámetro `id` o el formato del mismo es inválido.                          |
| 400    | Ocurrió un error al buscar el perfil del usuario (por ejemplo, el usuario no existe en la base). |

---

### **Ejemplo de solicitud válida**

#### Solicitud
```http
GET /viewprofile?id=64b7b2f4c3e88b0f5a1d9e7a HTTP/1.1
Host: localhost:8080
```

#### Respuesta
```http
HTTP/1.1 201 Created
Content-Type: application/json

{
  "id": "64b7b2f4c3e88b0f5a1d9e7a",
  "name": "Juan",
  "lastName": "Pérez",
  "dateBirth": "1990-05-15T00:00:00Z",
  "email": "juan.perez@example.com",
  "avatar": "https://example.com/avatars/juan.png",
  "banner": "https://example.com/banners/juan.png",
  "bibliography": "Desarrollador de software con 10 años de experiencia.",
  "ubication": "Buenos Aires, Argentina",
  "webSite": "https://juanperez.dev"
}
```



### 3. Modificar Perfil de Usuario

**Endpoint:** `/modifyprofile`  
**Método:** `PUT`  
**Descripción:** Permite modificar los datos del perfil de un usuario autenticado.

---

### **Formato del cuerpo de la solicitud**

El cliente debe enviar un JSON con los datos que desea modificar. Solo los campos enviados serán actualizados.

```json
{
  "name": "string",
  "lastName": "string",
  "dateBirth": "string (formato ISO 8601)",
  "email": "string",
  "avatar": "string (URL de la imagen)",
  "banner": "string (URL de la imagen)",
  "bibliography": "string",
  "ubication": "string",
  "webSite": "string (URL del sitio web)"
}
```

---

### **Requisitos de la solicitud**

- El cuerpo de la solicitud debe estar en formato JSON.
- El usuario debe estar autenticado (requiere un token JWT válido).
- Los datos enviados deben cumplir con los formatos esperados.

---

### **Respuestas posibles**

| Código | Descripción                                                                                      |
|--------|--------------------------------------------------------------------------------------------------|
| 201    | Perfil modificado exitosamente.                                                                  |
| 400    | El cuerpo de la solicitud contiene datos inválidos o faltantes.                                  |
| 400    | Ocurrió un error al intentar modificar el registro.                                              |
| 400    | No se logró modificar el registro (por ejemplo, el usuario no existe o no se aplicaron cambios). |

---

### **Ejemplo de solicitud válida**

#### Solicitud
```http
PUT /modifyprofile HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Authorization: Bearer <token-jwt>

{
  "name": "Juan",
  "lastName": "Pérez",
  "email": "juan.perez@example.com",
  "ubication": "Buenos Aires, Argentina"
}
```

#### Respuesta
```http
HTTP/1.1 201 Created
```

### 4. Registrar un Tweet

**Endpoint:** `/rectweet`  
**Método:** `POST`  
**Descripción:** Permite registrar un nuevo tweet para un usuario autenticado.

---

### **Formato del cuerpo de la solicitud**

El cliente debe enviar un JSON con el contenido del tweet.

```json
{
  "message": "string"
}
```

---

### **Requisitos de la solicitud**

- El cuerpo de la solicitud debe estar en formato JSON.
- El campo `message`:
  - Es obligatorio.
  - Debe contener el texto del tweet.
- El usuario debe estar autenticado (requiere un token JWT válido).

---

### **Respuestas posibles**

| Código | Descripción                                                                                      |
|--------|--------------------------------------------------------------------------------------------------|
| 201    | Tweet registrado exitosamente.                                                                   |
| 400    | El cuerpo de la solicitud es inválido o contiene datos faltantes.                                |
| 400    | Ocurrió un error al intentar insertar el registro.                                               |
| 400    | No se logró insertar el registro (por ejemplo, error en la base de datos o problemas internos).  |

---

### **Ejemplo de solicitud válida**

#### Solicitud
```http
POST /rectweet HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Authorization: Bearer <token-jwt>

{
  "message": "Este es mi primer tweet!"
}
```

#### Respuesta
```http
HTTP/1.1 201 Created
```

### 5. Leer Tweets

**Endpoint:** `/readtweets`  
**Método:** `GET`  
**Descripción:** Permite obtener una lista de tweets publicados por un usuario específico, paginados.

---

### **Parámetros de la solicitud**

- **Query Parameters:**
  - `id` (obligatorio): Identificador único del usuario cuyos tweets se desean consultar.
  - `page` (obligatorio): Número de página para la paginación (debe ser un número mayor a 0).

---

### **Formato de respuesta**

La respuesta será un arreglo de objetos JSON con el siguiente formato:

```json
[
  {
    "_id": "string",
    "userId": "string",
    "message": "string",
    "date": "string (ISO 8601)"
  }
]
```

**Descripción de los campos:**

- `_id`: Identificador único del tweet.
- `userId`: Identificador único del usuario que publicó el tweet.
- `message`: Contenido del tweet.
- `date`: Fecha y hora en la que se publicó el tweet (en formato ISO 8601).

---

### **Respuestas posibles**

| Código | Descripción                                                                                     |
|--------|-------------------------------------------------------------------------------------------------|
| 201    | Tweets obtenidos exitosamente.                                                                  |
| 400    | Falta el parámetro `id` en la URL.                                                              |
| 400    | Falta el parámetro `page` en la URL o no es un número mayor a 0.                                |
| 400    | Error al leer los tweets (por ejemplo, problemas con la base de datos o el usuario no existe).  |

---

### **Ejemplo de solicitud válida**

#### Solicitud
```http
GET /readtweets?id=64abcf12345d67890ef12345&page=1 HTTP/1.1
Host: localhost:8080
Authorization: Bearer <token-jwt>
```

#### Respuesta
```http
HTTP/1.1 201 Created
Content-Type: application/json

[
  {
    "_id": "64abcf12345d67890ef12345",
    "userId": "64abcdef6789012345678901",
    "message": "Este es mi primer tweet.",
    "date": "2025-01-11T14:23:45Z"
  },
  {
    "_id": "64abcf678901234567890123",
    "userId": "64abcdef6789012345678901",
    "message": "Este es otro tweet.",
    "date": "2025-01-10T10:15:30Z"
  }
]
```




### 7. Eliminar Tweet

**Endpoint:** ` /deleteTweet`  
**Método:** `DELETE`  
**Descripción:** Elimina un tweet específico de la base de datos, asegurándose de que pertenezca al usuario que realiza la solicitud.

---

### **Parámetros de entrada**

| Parámetro | Tipo   | Ubicación | Descripción                                                         |
|-----------|--------|-----------|---------------------------------------------------------------------|
| `id`      | string | Query     | Identificador único del tweet a eliminar.                          |
| `userId`  | string | Query     | Identificador único del usuario que solicita la eliminación.       |

---

### **Requisitos**

1. **`id`:** Debe ser un identificador único válido en formato hexadecimal de 24 caracteres.
2. **`userId`:** Debe coincidir con el creador del tweet que se desea eliminar.

---

### **Formato de solicitud**

#### **Ejemplo de solicitud**
```http
DELETE /deleteTweet?id=64abcf12345d67890ef12345&userId=64abcdef6789012345678901
```

---

### **Respuestas posibles**

| Código | Descripción                                                                                      |
|--------|--------------------------------------------------------------------------------------------------|
| 200    | El tweet fue eliminado exitosamente.                                                            |
| 400    | Error en los parámetros enviados o el tweet no existe.                                          |
| 500    | Error interno del servidor al intentar realizar la operación de eliminación.                    |

---



### 8. Subir Avatar

**Endpoint:** `/uploadAvatar`  
**Método:** `POST`

**Descripción:** Permite a un usuario subir una imagen de perfil (avatar) al servidor. El archivo se almacena en el sistema de archivos y la referencia se guarda en la base de datos.

---

### **Parámetros de entrada**

| Parámetro  | Tipo   | Ubicación | Descripción                                 |
|------------|--------|-----------|---------------------------------------------|
| `avatar`   | file   | FormData  | Archivo de imagen que será subido.          |
| `IDUser`   | string | Contexto  | Identificador único del usuario autenticado.|

---

### **Requisitos**

1. **`avatar`:** 
   - Debe ser un archivo válido de imagen con una extensión reconocida (e.g., `.jpg`, `.png`).
2. **`IDUser`:** 
   - Proporcionado a través del contexto de la solicitud, representa el usuario autenticado.

---

### **Formato de solicitud**

#### **Ejemplo de solicitud**
```http
POST /uploadAvatar
Content-Type: multipart/form-data

FormData:
  avatar: [archivo de imagen]
```

---

### **Respuestas posibles**

| Código | Descripción                                                                                      |
|--------|--------------------------------------------------------------------------------------------------|
| 201    | El avatar fue subido y registrado exitosamente.                                                  |
| 400    | Error al subir, copiar o registrar el avatar.                                                    |





### **Ejemplo de código**

#### **Solicitud exitosa**
```http
POST /uploadAvatar
Content-Type: multipart/form-data

FormData:
  avatar: avatar.png
```

**Respuesta:**
```json
{
  "message": "Avatar subido exitosamente."
}
```

### 9. Subir Banner

**Endpoint:** `/uploadBanner`  
**Método:** `POST`
**Descripción:** Permite a un usuario subir una imagen de banner al servidor. El archivo se almacena en el sistema de archivos y la referencia se guarda en la base de datos.

---

### **Parámetros de entrada**

| Parámetro  | Tipo   | Ubicación | Descripción                                 |
|------------|--------|-----------|---------------------------------------------|
| `banner`   | file   | FormData  | Archivo de imagen que será subido.          |
| `IDUser`   | string | Contexto  | Identificador único del usuario autenticado.|

---

### **Requisitos**

1. **`banner`:** 
   - Debe ser un archivo válido de imagen con una extensión reconocida (e.g., `.jpg`, `.png`).
2. **`IDUser`:** 
   - Proporcionado a través del contexto de la solicitud, representa el usuario autenticado.

---

### **Formato de solicitud**

#### **Ejemplo de solicitud**
```http
POST /uploadBanner
Content-Type: multipart/form-data

FormData:
  banner: [archivo de imagen]
```

---

### **Respuestas posibles**

| Código | Descripción                                                                                      |
|--------|--------------------------------------------------------------------------------------------------|
| 201    | El banner fue subido y registrado exitosamente.                                                  |
| 400    | Error al subir, copiar o registrar el banner.                                                    |

---



### **Ejemplo de código**

#### **Solicitud exitosa**
```http
POST /uploadBanner
Content-Type: multipart/form-data

FormData:
  banner: banner.jpg
```

**Respuesta:**
```json
{
  "message": "Banner subido exitosamente."
}
```



### 10. Obtener Avatar

**Endpoint:** `/getAvatar`  
**Método:** `GET`
**Descripción:** Permite obtener el archivo de avatar asociado a un usuario específico.

---

### **Parámetros de entrada**

| Parámetro | Tipo   | Ubicación | Descripción                                  |
|-----------|--------|-----------|----------------------------------------------|
| `id`      | string | Query     | Identificador único del usuario.             |

---

### **Requisitos**

1. **`id`:** 
   - Obligatorio. Identificador del usuario para buscar el avatar correspondiente.
   - Debe ser válido y estar registrado en la base de datos.

---

### **Formato de solicitud**

#### **Ejemplo de solicitud**
```http
GET /getAvatar?id=63bfc58e3e8b2a5d6c8a1234
```

---

### **Respuestas posibles**

| Código | Descripción                                                                                      |
|--------|--------------------------------------------------------------------------------------------------|
| 200    | Devuelve el archivo del avatar solicitado.                                                       |
| 400    | Error en la solicitud (parámetros inválidos, archivo no encontrado, etc.).                       |



#### **Solicitud exitosa**
```http
GET /getAvatar?id=63bfc58e3e8b2a5d6c8a1234
```

**Respuesta:**  
- Devuelve el archivo de imagen del avatar.





### 11. Obtener Banner

**Endpoint:** `/getBanner`  
**Método:** `GET`
**Descripción:** Permite obtener el archivo de banner asociado a un usuario específico.

---

### **Parámetros de entrada**

| Parámetro | Tipo   | Ubicación | Descripción                                  |
|-----------|--------|-----------|----------------------------------------------|
| `id`      | string | Query     | Identificador único del usuario.             |

---

### **Requisitos**

1. **`id`:** 
   - Obligatorio. Identificador del usuario para buscar el banner correspondiente.
   - Debe ser válido y estar registrado en la base de datos.

---

### **Formato de solicitud**

#### **Ejemplo de solicitud**
```http
GET /getBanner?id=63bfc58e3e8b2a5d6c8a1234
```

---

### **Respuestas posibles**

| Código | Descripción                                                                                      |
|--------|--------------------------------------------------------------------------------------------------|
| 200    | Devuelve el archivo del banner solicitado.                                                       |
| 400    | Error en la solicitud (parámetros inválidos, archivo no encontrado, etc.).                       |


#### **Solicitud exitosa**
```http
GET /getBanner?id=63bfc58e3e8b2a5d6c8a1234
```

**Respuesta:**  
- Devuelve el archivo de imagen del banner.



### 12. Crear Relación de Usuario

**Endpoint:** `/highRelation`  
**Método:** `POST`
**Descripción:** Permite crear una relación entre el usuario autenticado y otro usuario especificado por su ID.

---

### **Parámetros de entrada**

| Parámetro | Tipo   | Ubicación | Descripción                                  |
|-----------|--------|-----------|----------------------------------------------|
| `id`      | string | Query     | Identificador único del usuario con el que se desea crear la relación. |

---

### **Requisitos**

1. **`id`:** 
   - Obligatorio. Representa el identificador del usuario con el que se establecerá la relación.
   - Debe ser un identificador válido registrado en la base de datos.

---

### **Formato de solicitud**

#### **Ejemplo de solicitud**
```http
POST /highRelation?id=63bfc58e3e8b2a5d6c8a1234
```

---

### **Respuestas posibles**

| Código | Descripción                                                                                      |
|--------|--------------------------------------------------------------------------------------------------|
| 201    | Relación creada exitosamente.                                                                    |
| 400    | Error en la solicitud (parámetros inválidos, error en la base de datos, etc.).                   |


#### **Solicitud exitosa**
```http
POST /highRelation?id=63bfc58e3e8b2a5d6c8a1234
```

**Respuesta:**  
```http
HTTP/1.1 201 Created
```



### 13. Eliminar Relación de Usuario

**Endpoint:** `/downRelation`  
**Método:** `DELETE`
**Descripción:** Permite eliminar una relación existente entre el usuario autenticado y otro usuario especificado por su ID.

---

### **Parámetros de entrada**

| Parámetro | Tipo   | Ubicación | Descripción                                  |
|-----------|--------|-----------|----------------------------------------------|
| `id`      | string | Query     | Identificador único del usuario con el que se desea eliminar la relación. |

---

### **Requisitos**

1. **`id`:** 
   - Obligatorio. Representa el identificador del usuario con el que se eliminará la relación.
   - Debe ser un identificador válido registrado en la base de datos.

---

### **Formato de solicitud**

#### **Ejemplo de solicitud**
```http
DELETE /downRelation?id=63bfc58e3e8b2a5d6c8a1234
```

---

### **Respuestas posibles**

| Código | Descripción                                                                                      |
|--------|--------------------------------------------------------------------------------------------------|
| 201    | Relación eliminada exitosamente.                                                                 |
| 400    | Error en la solicitud (parámetros inválidos, error en la base de datos, etc.).                   |





#### **Solicitud exitosa**
```http
DELETE /downRelation?id=63bfc58e3e8b2a5d6c8a1234
```

**Respuesta:**  
```http
HTTP/1.1 201 Created
```



### 14. Consultar Relación entre Usuarios

**Endpoint:** `/consultRelation`  
**Método:** `GET`
**Descripción:** Permite verificar si existe una relación entre el usuario autenticado y otro usuario especificado por su ID.

---

### **Parámetros de entrada**

| Parámetro | Tipo   | Ubicación | Descripción                                  |
|-----------|--------|-----------|----------------------------------------------|
| `id`      | string | Query     | Identificador único del usuario a consultar. |

---

### **Requisitos**

1. **`id`:** 
   - Obligatorio. Representa el identificador del usuario con el que se desea consultar la relación.
   - Debe ser un identificador válido registrado en la base de datos.

---

### **Formato de solicitud**

#### **Ejemplo de solicitud**
```http
GET /consultRelation?id=63bfc58e3e8b2a5d6c8a1234
```

---

### **Respuestas posibles**

| Código | Descripción                                                                                      |
|--------|--------------------------------------------------------------------------------------------------|
| 201    | Consulta realizada exitosamente.                                                                 |
| 400    | Error en la solicitud (parámetros inválidos, error en la base de datos, etc.).                   |

---

### **Cuerpo de la respuesta**

| Campo    | Tipo    | Descripción                                           |
|----------|---------|-------------------------------------------------------|
| `status` | boolean | Indica si existe (`true`) o no (`false`) la relación. |

#### **Ejemplo de respuesta exitosa**
```json
{
  "status": true
}
```

#### **Ejemplo de respuesta fallida**
```json
{
  "status": false
}
```


#### **Solicitud exitosa**
```http
GET /consultRelation?id=63bfc58e3e8b2a5d6c8a1234
```

**Respuesta:**  
```json
{
  "status": true
}
```

### 15. Listar Usuarios

**Endpoint:** `/listUsers` 
**Método:** `GET` 
**Descripción:** Permite obtener una lista de usuarios registrados en la plataforma con la posibilidad de filtrar por tipo de relación y realizar búsquedas.

---

### **Parámetros de entrada**

| Parámetro | Tipo    | Ubicación | Descripción                                                                 |
|-----------|---------|-----------|-----------------------------------------------------------------------------|
| `type`    | string  | Query     | Tipo de relación a filtrar (`followers`, `following`, etc.). Opcional.      |
| `page`    | integer | Query     | Número de página para paginación. Obligatorio y debe ser mayor a 0.         |
| `search`  | string  | Query     | Término de búsqueda para filtrar usuarios por nombre o descripción. Opcional.|

---

### **Requisitos**

1. **`page`:** 
   - Obligatorio. Representa el número de página para la paginación.
   - Debe ser un entero positivo mayor a 0.
2. **`type`:** 
   - Opcional. Especifica el tipo de relación que se desea filtrar.
   - Valores posibles: `followers`, `following`, o vacío para no aplicar filtro.
3. **`search`:** 
   - Opcional. Permite buscar usuarios por términos específicos.

---

### **Formato de solicitud**

#### **Ejemplo de solicitud**
```http
GET /listUsers?page=1&type=followers&search=john
```

---

### **Respuestas posibles**

| Código | Descripción                                                                                  |
|--------|----------------------------------------------------------------------------------------------|
| 201    | Lista de usuarios obtenida exitosamente.                                                     |
| 400    | Error en la solicitud (parámetros inválidos, error en la base de datos, etc.).               |

---

### **Cuerpo de la respuesta**

La respuesta es una lista de usuarios con el siguiente formato:

| Campo     | Tipo    | Descripción                            |
|-----------|---------|----------------------------------------|
| `id`      | string  | Identificador único del usuario.       |
| `name`    | string  | Nombre del usuario.                   |
| `email`   | string  | Correo electrónico del usuario.        |
| `avatar`  | string  | URL del avatar del usuario.            |
| `banner`  | string  | URL del banner del usuario.            |
| `bio`     | string  | Biografía del usuario.                 |

#### **Ejemplo de respuesta exitosa**
```json
[
  {
    "id": "63bfc58e3e8b2a5d6c8a1234",
    "name": "John Doe",
    "email": "johndoe@example.com",
    "avatar": "/uploads/avatars/63bfc58e3e8b2a5d6c8a1234.jpg",
    "banner": "/uploads/banners/63bfc58e3e8b2a5d6c8a1234.jpg",
    "bio": "Desarrollador apasionado por la tecnología."
  },
  {
    "id": "63bfc58e3e8b2a5d6c8a5678",
    "name": "Jane Smith",
    "email": "janesmith@example.com",
    "avatar": "/uploads/avatars/63bfc58e3e8b2a5d6c8a5678.jpg",
    "banner": "/uploads/banners/63bfc58e3e8b2a5d6c8a5678.jpg",
    "bio": "Diseñadora gráfica y amante del arte digital."
  }
]
```



#### **Solicitud exitosa**
```http
GET /listUsers?page=1&type=followers&search=john
```

**Respuesta:**  
```json
[
  {
    "id": "63bfc58e3e8b2a5d6c8a1234",
    "name": "John Doe",
    "email": "johndoe@example.com",
    "avatar": "/uploads/avatars/63bfc58e3e8b2a5d6c8a1234.jpg",
    "banner": "/uploads/banners/63bfc58e3e8b2a5d6c8a1234.jpg",
    "bio": "Desarrollador apasionado por la tecnología."
  }
]
```


### 16. Leer Tweets de Seguidores

**Endpoint:** `/readTweetsFollowers`  
**Método:** `GET`
**Descripción:** Obtiene una lista de tweets de los usuarios que el usuario autenticado sigue, con soporte para paginación.

---

### **Parámetros de entrada**

| Parámetro | Tipo    | Ubicación | Descripción                                                          |
|-----------|---------|-----------|----------------------------------------------------------------------|
| `page`    | integer | Query     | Número de página para la paginación. Obligatorio y debe ser mayor a 0.|

---

### **Requisitos**

1. **`page`:** 
   - Obligatorio.
   - Debe ser un entero positivo mayor a 0.
   - Representa el número de página de los tweets solicitados.

---

### **Formato de solicitud**

#### **Ejemplo de solicitud**
```http
GET /readTweetsFollowers?page=1
```

---

### **Respuestas posibles**

| Código | Descripción                                                                                  |
|--------|----------------------------------------------------------------------------------------------|
| 201    | Lista de tweets obtenida exitosamente.                                                       |
| 400    | Error en la solicitud (parámetros inválidos, error en la base de datos, etc.).               |

---

### **Cuerpo de la respuesta**

La respuesta es una lista de tweets con el siguiente formato:

| Campo        | Tipo    | Descripción                            |
|--------------|---------|----------------------------------------|
| `id`         | string  | Identificador único del tweet.         |
| `user_id`    | string  | ID del usuario que publicó el tweet.   |
| `name`       | string  | Nombre del usuario.                   |
| `content`    | string  | Contenido del tweet.                   |
| `timestamp`  | string  | Fecha y hora de publicación.           |

#### **Ejemplo de respuesta exitosa**
```json
[
  {
    "_id": "605c72ef1532075f8838d8d4",
    "userId": "63bfc58e3e8b2a5d6c8a1234",
    "userRelationId": "63bfc58e3e8b2a5d6c8a5678",
    "tweet": {
      "_id": "605c72ef1532075f8838d8d5",
      "message": "Este es un tweet de prueba.",
      "date": "2025-01-23T12:34:56Z"
    }
  },
  {
    "_id": "605c72ef1532075f8838d8d6",
    "userId": "63bfc58e3e8b2a5d6c8a1234",
    "userRelationId": "63bfc58e3e8b2a5d6c8a5678",
    "tweet": {
      "_id": "605c72ef1532075f8838d8d7",
      "message": "Otro tweet de ejemplo.",
      "date": "2025-01-22T10:20:30Z"
    }
  }
]

```
