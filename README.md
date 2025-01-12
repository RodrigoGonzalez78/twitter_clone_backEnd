# Backend de una app clon de X(Twitter)


### **Formato del modelo de usuario**
El perfil del usuario devuelto tiene el siguiente formato:

```json
{
  "id": "string",
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

