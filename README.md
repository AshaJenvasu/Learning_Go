# 🐹 GoAPI Essential - Learning Workspace

Welcome to my personal learning repository! 🚀 This project is a hands-on workspace where I am following the **"GoAPI Essential"** series by **Mikelopster**. 

---

## 📌 Project Overview

This repository serves as my digital notebook and code playground as I transition from writing basic Go syntax to building production-ready RESTful Web APIs using the **Go (Golang)** language and the **Fiber** framework.

### 🛠️ Tech Stack & Tools
* **Language:** Go (Golang)
* **Framework:** Fiber (Express-like web framework for Go)
* **Live Reload:** Nodemon / Air
* **Course Reference:** [GoAPI Essential](https://www.youtube.com/watch?v=nfSbFk8y4D0) by Mikelopster 💻

---

## 🧠 Key Concepts Covered & Practiced

### 🌐 1. HTTP Server Evolution
* Started with Go's native `net/http` package to understand lower-level servers.
* Migrated to **Fiber** to leverage its fast, Express.js-like routing and performance.

### 🎯 2. In-Memory CRUD Operations (Slice Manipulation)
Before connecting to a persistent database, I practiced data management in RAM using Go slices:
* **Create, Read, Update, Delete** routines for a mock `Book` structure.
* For **Deletion**, used the slice-stitching technique: `append(books[:i], books[i+1:]...)` to remove items by their ID dynamically.

### 🔒 3. Middlewares & Authentication
* Implemented logging middlewares to intercept and log requests.
* Learned how to protect routes using **JWT (JSON Web Tokens)** and role-based access control.

### ⚙️ 4. Environment & Tools
* Managed configurations using `.env` files.
* Explored Form Data handling, file uploads, and view templates.
* Integrated **Swagger UI** for automated interactive API documentation.

---

## 💻 How to Run This Project Locally

1. **Clone the repository:**
   ```bash
   git clone https://github.com
   ```
2. **Navigate into the directory:**
   ```bash
   cd YOUR_REPO_NAME
   ```
3. **Install dependencies:**
   ```bash
   go mod tidy
   ```
4. **Run with live reload (using Nodemon):**
   ```bash
   nodemon --exec go run main.go --signal SIGTERM
   ```

---

🛡️ *This repo is updated continuously as I progress through the GoAPI Essential series. Stay tuned for database integration!*
