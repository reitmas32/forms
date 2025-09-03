# Forms API — README

API para **crear formularios dinámicos** y **registrar/rescatar respuestas**. Ideal para flujos tipo onboarding, encuestas, consentimientos y perfiles de usuario.

## ✨ Características

* Definición de formularios con **secciones**, **preguntas** y **tipos de campo**.
* Soporte para opciones (select, radio, checkbox).
* Registro de respuestas por `form_id` y `user_id`.
* Consulta de formularios y de respuestas por formulario o por id de respuesta.
* Endpoints REST simples bajo `/v1`.

## Typos de questions

* text-long
* text-short
* text-email
* radio
* file
* boolean
* select
* checkbox
* dropdown
* date


---

## 🚀 Endpoints

Base path sugerido: `https://<host>/v1`

### 1) Forms

| Método | Ruta                 | Descripción                        |
| -----: | -------------------- | ---------------------------------- |
|   POST | `/forms`             | Crear un nuevo formulario          |
|    GET | `/forms`             | Listar formularios                 |
|    GET | `/forms/:id`         | Obtener un formulario por `id`     |
|    GET | `/forms/:id/answers` | Listar respuestas de un formulario |

### 2) Answers

| Método | Ruta           | Descripción                             |
| -----: | -------------- | --------------------------------------- |
|   POST | `/answers`     | Enviar (crear) respuestas de un usuario |
|    GET | `/answers/:id` | Recuperar una respuesta por `id`        |

> Rutas según el enrutado provisto:
>
> ```go
> formsGroup := router.Group("/v1/forms")
> formsGroup.POST("", formsController.Create)
> formsGroup.GET("", formsController.List)
> formsGroup.GET("/:id", formsController.Retrieve)
> formsGroup.GET("/:id/answers", formsController.Answers)
>
> answers := r.Group("/v1/answers")
> answers.POST("", controller.Create)
> answers.GET("/:id", controller.Retrieve)
> ```

---

## 🧱 Modelos de datos

### Form

```json
{
  "id": "68b79f5505894042cd8fff59",
  "title": "Datos Personales",
  "description": "Formulario para la obtención de datos personales del usuario",
  "questions": [
    {
      "id": "88754e58-f567-4f09-bbfa-603741b58687",
      "title": "Nombre",
      "description": "Especifica tu nombre",
      "type": "text-short",
      "required": true,
      "section": "Datos Básicos",
      "metadata": null
    }
    // ...
  ],
  "created_at": "2025-09-01T12:00:00Z",
  "updated_at": "2025-09-01T12:00:00Z"
}
```

### Question types soportados

* `text-short`, `text-long`, `text-email`
* `date`
* `radio`, `select`, `checkbox` *(usar `metadata.options` como arreglo de strings)*
* `boolean`
* `file` *(la respuesta suele ser una URL segura o path al recurso)*

**metadata** (opcional):

```json
{ "options": ["Opción A", "Opción B", "..."] }
```

### Answer payload

```json
{
  "id": "a1b2c3...",
  "form_id": "68b79f5505894042cd8fff59",
  "user_id": "e746ee25-ec41-4159-b4d1-169720d5ef15",
  "responses": [
    { "question_id": "uuid-pregunta", "answer": "valor" }
  ],
  "created_at": "2025-09-01T12:00:00Z"
}
```

> **Notas**
>
> * `answer`:
>
>   * `checkbox`: puede ser string separado por comas **o** arreglo de strings (definir en contrato).
>   * `boolean`: `"true"/"false"` o boolean real (recomendado).
>   * `file`: URL del archivo.
> * Validar que cada `question_id` pertenezca al `form_id`.

---

## 📦 Ejemplos de uso

### Crear Formulario

**Request**

```http
POST /v1/forms
Content-Type: application/json
```

```json
{
  "title": "Datos Personales",
  "description": "Formulario para la obtención de datos personales del usuario",
  "questions": [
    { "title": "Nombre", "description": "Especifica tu nombre", "type": "text-short", "required": true, "section": "Datos Básicos" },
    { "title": "Apellido", "description": "Especifica tu apellido", "type": "text-short", "required": true, "section": "Datos Básicos" },
    { "title": "Fecha de nacimiento", "description": "Selecciona tu fecha de nacimiento", "type": "date", "required": true, "section": "Datos Básicos" },
    { "title": "Género", "description": "Selecciona tu género", "type": "radio", "metadata": { "options": ["Masculino", "Femenino", "Otro", "Prefiero no decirlo"] }, "required": true, "section": "Datos Básicos" },
    { "title": "Correo electrónico", "description": "Introduce tu email de contacto", "type": "text-email", "required": true, "section": "Contacto" },
    { "title": "Teléfono", "description": "Introduce tu número de teléfono", "type": "text-short", "required": false, "section": "Contacto" },
    { "title": "País de residencia", "description": "Selecciona tu país de residencia", "type": "select", "metadata": { "options": ["México","Estados Unidos","España","Otro"] }, "required": true, "section": "Contacto" },
    { "title": "Hobbies", "description": "Marca las actividades que te interesan", "type": "checkbox", "metadata": { "options": ["Deporte","Lectura","Viajar","Música","Cine"] }, "required": false, "section": "Información Adicional" },
    { "title": "Biografía", "description": "Cuéntanos algo sobre ti", "type": "text-long", "required": false, "section": "Información Adicional" },
    { "title": "Foto de perfil", "description": "Sube una imagen para tu perfil", "type": "file", "required": false, "section": "Información Adicional" },
    { "title": "Acepto los términos y condiciones", "description": "Debes aceptar para continuar", "type": "boolean", "required": true, "section": "Consentimiento" }
  ]
}
```

**Response (sugerido)**

```json
{
  "id": "68b79f5505894042cd8fff59",
  "title": "Datos Personales",
  "description": "Formulario para la obtención de datos personales del usuario",
  "questions": [
    { "id": "88754e58-f567-4f09-bbfa-603741b58687", "title": "Nombre", "type": "text-short", "required": true, "section": "Datos Básicos" }
    // ...
  ],
  "created_at": "2025-09-01T12:00:00Z"
}
```

**cURL**

```bash
curl -X POST https://<host>/v1/forms \
  -H "Content-Type: application/json" \
  -d @form.json
```

---

### Listar Formularios

```http
GET /v1/forms
```

**Response (ejemplo)**

```json
{
  "items": [
    { "id": "68b79f5505894042cd8fff59", "title": "Datos Personales", "description": "Formulario para la obtención...", "questions_count": 11, "created_at": "2025-09-01T12:00:00Z" }
  ],
  "total": 1
}
```

**cURL**

```bash
curl https://<host>/v1/forms
```

---

### Obtener un Formulario

```http
GET /v1/forms/:id
```

**cURL**

```bash
curl https://<host>/v1/forms/68b79f5505894042cd8fff59
```

---

### Listar Respuestas de un Formulario

```http
GET /v1/forms/:id/answers
```

**Response (ejemplo)**

```json
{
  "items": [
    {
      "id": "a1b2c3",
      "form_id": "68b79f5505894042cd8fff59",
      "user_id": "e746ee25-ec41-4159-b4d1-169720d5ef15",
      "responses": [
        { "question_id": "88754e58-f567-4f09-bbfa-603741b58687", "answer": "Rafa" }
      ],
      "created_at": "2025-09-01T12:34:56Z"
    }
  ],
  "total": 1
}
```

**cURL**

```bash
curl https://<host>/v1/forms/68b79f5505894042cd8fff59/answers
```

---

### Enviar Respuestas

```http
POST /v1/answers
Content-Type: application/json
```

**Request**

```json
{
  "form_id": "68b79f5505894042cd8fff59",
  "user_id": "e746ee25-ec41-4159-b4d1-169720d5ef15",
  "responses": [
    { "question_id": "88754e58-f567-4f09-bbfa-603741b58687", "answer": "Rafa" },
    { "question_id": "e1fd8d2a-aa28-48cc-b164-6b705a44a6b0", "answer": "Zamora" },
    { "question_id": "3271228b-e1da-4ddb-98c8-0d4a467e790f", "answer": "2000-01-11" },
    { "question_id": "08c22683-4762-4539-90fc-78fc5f1a2", "answer": "Masculino" },
    { "question_id": "80a3ab0d-4142-46a5-b56d-7987622776d9", "answer": "rafa.zamora@example.com" },
    { "question_id": "11898615-6416-4455-9bcc-dfda1c5257b8", "answer": "+52 1 234 567 8901" },
    { "question_id": "17369abc-9801-4016-8c90-f3ab48cd7350", "answer": "México" },
    { "question_id": "006beeb0-7320-4fb5-8009-6dcbf471d769", "answer": "Deporte, Música" },
    { "question_id": "3571eb9e-c4f7-4b35-b8c5-0d7adb8a137c", "answer": "Desarrollador entusiasta de Go y Python; me gusta aprender y construir productos útiles." },
    { "question_id": "52edffe9-39f2-41a7-add6-19daa1fca94b", "answer": "https://s3-minio-dev.konectus.tech/assets/public/rafa.png" },
    { "question_id": "5feaa1f2-1697-4c1e-b5fd-f1bfc08e42ca", "answer": "true" }
  ]
}
```

**Response (ejemplo)**

```json
{
  "id": "a1b2c3",
  "form_id": "68b79f5505894042cd8fff59",
  "user_id": "e746ee25-ec41-4159-b4d1-169720d5ef15",
  "saved": 11,
  "created_at": "2025-09-01T12:34:56Z"
}
```

**cURL**

```bash
curl -X POST https://<host>/v1/answers \
  -H "Content-Type: application/json" \
  -d @answers.json
```

---

### Recuperar una Respuesta por ID

```http
GET /v1/answers/:id
```

**Response (ejemplo)**

```json
{
  "id": "a1b2c3",
  "form_id": "68b79f5505894042cd8fff59",
  "user_id": "e746ee25-ec41-4159-b4d1-169720d5ef15",
  "responses": [
    { "question_id": "88754e58-f567-4f09-bbfa-603741b58687", "answer": "Rafa" }
    // ...
  ],
  "created_at": "2025-09-01T12:34:56Z"
}
```

**cURL**

```bash
curl https://<host>/v1/answers/a1b2c3
```

---

## ✅ Validaciones recomendadas

* **Existencia** de `form_id` y de cada `question_id` dentro del formulario.
* **Tipos**:

  * `text-email` → validar formato email.
  * `date` → ISO `YYYY-MM-DD`.
  * `boolean` → `true/false` (string o boolean); normalizar a boolean.
  * `radio/select` → valor dentro de `metadata.options`.
  * `checkbox` → todos los valores dentro de `metadata.options`.
* **Obligatorias**: si `required=true` la respuesta no puede venir vacía.
* **Archivos**: si `type=file`, la respuesta debe ser una URL válida o ID de media.

---

## 🔐 Autenticación y permisos

* No especificado en el enrutado. Sugerencias:

  * **Forms**: protegido por rol admin/editor.
  * **Answers**: creación por usuarios autenticados; lectura restringida por `organization` o por `user_id`.
  * Trazabilidad: incluir `created_by`, `organization_id` si aplica.

---

## 📑 Paginación y filtros (opcional)

* `GET /v1/forms?limit=20&offset=0&query=datos`
* `GET /v1/forms/:id/answers?limit=20&offset=0&user_id=<uuid>`

**Respuesta paginada sugerida**

```json
{ "items": [...], "total": 123, "limit": 20, "offset": 0 }
```

---

## 🧪 Códigos de estado y errores

* `201 Created` → creación de form o answers.
* `200 OK` → lecturas.
* `400 Bad Request` → validación fallida (tipo/required/opciones).
* `404 Not Found` → form o answer inexistente.
* `409 Conflict` → respuestas duplicadas (si se restringe un envío por usuario).
* `500 Internal Server Error` → error inesperado.

**Formato de error sugerido**

```json
{ "error": "validation_error", "message": "Email inválido", "field": "question_id:80a3ab0d-..." }
```

---

## 🛠️ Notas de implementación (Go)

* Los grupos de rutas ya definidos:

  ```go
  // /v1/forms y /v1/answers
  ```
* Recomendado:

  * Normalizar respuestas (DTOs).
  * Validar tipos por `question.type`.
  * Índices en BD: `form_id`, `user_id`, `(form_id, user_id)`, `question_id`.
  * Transacción al guardar múltiples `responses`.

---

## 📚 Ejemplos rápidos (Insomnia/Postman)

* **Crear Form**: importar el JSON del ejemplo de “Crear Formulario”.
* **Enviar Answers**: importar el JSON de “Enviar Respuestas”.

---

## 🧭 Roadmap sugerido

* Versionado de formularios (publicar/borrador).
* Soporte i18n en `title/description/options`.
* Reglas condicionales (mostrar/ocultar preguntas).
* Exportación CSV/Parquet de respuestas.
* Webhooks/Events al recibir respuestas.
* Rate limiting y auditoría.

---

¿Quieres que lo entregue en **Markdown listo para el repo** con una sección de **OpenAPI/Swagger** y ejemplos adicionales por tipo de pregunta?
