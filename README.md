# 🚀 Airbnb Manta - Dashboard de Inteligencia de Negocios

Sistema de gestión de reservas desarrollado con **Go (Golang)** y **SQL Server**, enfocado en la arquitectura de microservicios y principios de Ingeniería de Software.

## 🛠️ Tecnologías Utilizadas
* **Backend:** Go 1.25 (REST API nativa)
* **Base de Datos:** Microsoft SQL Server
* **Frontend:** HTML5, Tailwind CSS, JavaScript (Fetch API)
* **Arquitectura:** Cliente-Servidor con serialización JSON

## 🌟 Características Principales
1. **8 Servicios Web (REST):** Implementación de CRUD completo, KPIs financieros y monitoreo de estado.
2. **Programación Orientada a Objetos (POO):** Uso de patrones Factory para la gestión de canales de venta (Airbnb, Booking, Directo).
3. **Concurrencia Avanzada:** Generación de reportes financieros en formato CSV mediante **Goroutines**, permitiendo procesos asíncronos sin bloquear la interfaz de usuario.
4. **Validación de Negocio:** Algoritmo de detección de *Overbooking* (cruces de fechas) directamente en SQL Server.
5. **Serialización JSON:** Comunicación desacoplada entre el cliente y el servidor.

## 🧪 Pruebas de Software
El proyecto incluye:
* **Pruebas Unitarias:** Validación de lógica de comisiones y constructores en `models/`.
* **Pruebas de Integración:** Conexión persistente con SQL Server.

---
*Desarrollado como proyecto práctico de Ingeniería de Software.*