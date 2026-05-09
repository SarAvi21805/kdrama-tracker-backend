# 🌸 K-Drama & Anime Tracker - API REST

Este es el backend del proyecto Full Stack, encargado de gestionar la lógica de datos, la persistencia en base de datos y la exposición de una API REST para el seguimiento de series.

## 🚀 Tecnologías usadas
- **Lenguaje:** Go (Golang)
- **Framework:** Chi Router (Ligero y rápido)
- **Base de Datos:** SQLite (Persistencia real de datos)
- **Despliegue:** Render

## 📡 Endpoints de la API
- `GET /series`: Listar todas las series guardadas.
- `GET /series/{id}`: Obtener detalles de una serie específica.
- `POST /series`: Crear una nueva serie.
- `PUT /series/{id}`: Editar una serie existente.
- `DELETE /series/{id}`: Eliminar una serie.

## 🔐 Configuración de CORS
**¿Qué es CORS?** Cross-Origin Resource Sharing es una política de seguridad del navegador que bloquea peticiones entre distintos dominios. 
**Configuración:** En este proyecto, configuramos los headers de respuesta para permitir peticiones desde cualquier origen (`*`) y habilitar los métodos GET, POST, PUT y DELETE, permitiendo así que el frontend pueda consumir los datos sin bloqueos.

## 💭 Reflexión sobre la tecnología
Elegir **Go** para este proyecto fue un reto gratificante. Aunque al principio la estructura de tipos y el manejo de CGO con SQLite en Windows fue complejo, la velocidad de ejecución y la claridad del código compensan el esfuerzo. Definitivamente volvería a usar Go para microservicios por su eficiencia, aunque para proyectos que requieren iteraciones visuales extremadamente rápidas, Node.js podría ser más flexible.

## 🛠️ Cómo correr localmente
1. Clonar el repositorio.
2. Ejecutar `go mod tidy` para instalar dependencias.
3. Correr con `go run main.go`.

**Link al Frontend:** [Repositorio Frontend](https://github.com/SarAvi21805/kdrama-tracker-frontend)