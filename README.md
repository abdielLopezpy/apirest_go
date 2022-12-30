# apirest_go
"Este es un proyecto de ejemplo que muestra cómo crear una API en Go con funcionalidades CRUD. La API se ha desarrollado siguiendo buenas prácticas y está diseñada para ser fácilmente extensible. Si estás interesado en aprender cómo crear APIs en Go, ¡échale un vistazo a este repositorio!

# Recurso: Ítems

## Operación: Obtener lista de ítems

Método: GET

Ruta: /items

Parámetros: N/A

Respuesta:

- Formato: JSON
- Códigos de estado:
  - 200 OK: Se ha obtenido la lista de ítems correctamente.
  - 401 Unauthorized: Se requiere autenticación para acceder a la API.

## Operación: Crear nuevo ítem

Método: POST

Ruta: /items

Parámetros: En el cuerpo de la solicitud, en formato JSON:

- ID (obligatorio): Entero único que identifica al ítem.
- Name (obligatorio): Nombre del ítem.

Respuesta:

- Formato: JSON
- Códigos de estado:
  - 201 Created: Se ha creado el ítem correctamente.
  - 400 Bad Request: El cuerpo de la solicitud no tiene el formato correcto o faltan parámetros obligatorios.
  - 401 Unauthorized: Se requiere autenticación para acceder a la API.

## Operación: Obtener ítem por ID

Método: GET

Ruta: /items/{id}

Parámetros: En la ruta:

- id: Entero único que identifica al ítem.

Respuesta:

- Formato: JSON
- Códigos de estado:
  - 200 OK: Se ha obtenido el ítem correctamente.
  - 400 Bad Request: El ID especificado no es válido.
  - 401 Unauthorized: Se requiere autenticación para acceder a la API.
  - 404 Not Found: No se ha encontrado el ítem con el ID especificado.

## Operación: Actualizar ítem por ID

Método: PUT

Ruta: /items/{id}

Parámetros:

- En la ruta:
  - id: Entero único que identifica al ítem.
- En el cuerpo de la solicitud, en formato JSON:
  - ID (obligatorio): Entero único que identifica al ítem.
  - Name (obligatorio): Nombre del ítem.

Respuesta:

- Formato: JSON
- Códigos de estado:
  - 200 OK: Se ha actualizado el ítem correctamente.
Copy code
## Operación: Eliminar ítem por ID

Método: DELETE

Ruta: /items/{id}

Parámetros: En la ruta:

- id: Entero único que identifica al ítem.

Respuesta:

- Formato: JSON
- Códigos de estado:
  - 200 OK: Se ha eliminado el ítem correctamente.
  - 400 Bad Request: El ID especificado no es válido.
  - 401 Unauthorized: Se requiere autenticación para acceder a la API.
  - 404 Not Found: No se ha encontrado el ítem con el ID especificado.
