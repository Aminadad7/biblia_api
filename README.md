# 📖 Biblia API – Reina Valera 1960

API REST desarrollada en **Go + Gin** para consultar libros, capítulos y versículos de la **Biblia Reina Valera 1960** a partir de un archivo JSON local.

---

## 🚀 Tecnologías usadas

- Go
- Gin Gonic
- JSON
- Biblia Reina Valera 1960 (Español)

---

## ⚙️ Configuración y ejecución

### Clonar el proyecto
```bash
git clone https://github.com/tuusuario/biblia-api.git
cd biblia-api
```

### Estructura del proyecto
```text
.
├── main.go
└── data/
    └── Biblia_Reina_Valera_1960_Esp.json
```

### Ejecutar la API
```bash
go run main.go
```

La API estará disponible en:
```
http://localhost:8000
```

---

## 🌐 Endpoints

### 1. Estado de la API
**GET** `/`

Respuesta:
```json
{
  "status": "Biblia API activa",
  "verses": 31102
}
```

---

### 2. Obtener libros
**GET** `/books`

Respuesta:
```json
["Génesis", "Éxodo", "Levítico"]
```

---

### 3. Capítulos de un libro
**GET** `/books/:book/chapters`

Ejemplo:
```
/books/Génesis/chapters
```

Respuesta:
```json
[1, 2, 3]
```

---

### 4. Versículos por capítulo
**GET** `/books/:book/chapters/:chapter`

Ejemplo:
```
/books/Génesis/chapters/1
```

Respuesta:
```json
[
  {
    "Book": "Génesis",
    "Chapter": 1,
    "Verse": 1,
    "Text": "En el principio creó Dios los cielos y la tierra."
  }
]
```

---

### 5. Versículo específico
**GET** `/verse`

Query params:
- book
- chapter
- verse

Ejemplo:
```
/verse?book=Juan&chapter=3&verse=16
```

Respuesta:
```json
{
  "Book": "Juan",
  "Chapter": 3,
  "Verse": 16,
  "Text": "Porque de tal manera amó Dios al mundo..."
}
```

---

## 📄 Licencia

Uso libre con fines educativos.
