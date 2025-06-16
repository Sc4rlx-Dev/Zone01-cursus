﻿

## 🚀 Getting Started

### Run Locally

To run the application locally using Go:

```bash
go run cmd/main/main.go
````

The app will be available at:

```
http://localhost:8080
```

---

## 🔧 Features

### ✅ Home Page Implementation

* Fetch and display all posts with author details.
* Uses SQL `JOIN` between `posts` and `users`.

### 🔗 SQL Integration

* Efficient data retrieval with a join query.
* Clean and structured output for rendering posts.

### 🔁 Dynamic Post Display

* Iterates over combined query results.
* Displays content dynamically on the homepage.

### 🎨 Simple and Responsive UI

* Minimal HTML & CSS design.
* Lightweight and mobile-friendly layout.

---

## 📄 Description

The Home Page fetches posts from the database and joins them with corresponding user data using SQL. The page is styled with minimal HTML and CSS, keeping the interface clean, responsive, and performant.

---

## 🐳 Docker Support

### Build the Docker Image

```bash
docker build -t home-page-app .
```

### Run the Docker Container

```bash
docker run -p 8080:8080 home-page-app
```

Then access the application at:

```
http://localhost:8080
```

---

## 👥 Team Members

| Name                | GitHub Username | Role          |
| ------------------- | --------------- | ------------- |
| Mohamed Ndoumghar   | `@mndoumghar`   | Team Lead 👑  |
| Anas Ayyada         | `@aayyada`      | Active Member |
| Oussama Erraoui     |  [@oerraoui](https://github.com/Sc4rlx-Dev)    | Explorer      |
| Abderrahim Ouachani | `@abouachani`   | Contributor   |

---
---

```
