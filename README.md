
🚀 Project_CRUD_Go-Next  
**Full-Stack CRUD Web App with Go (Gin) & Next.js (TypeScript)**

![GitHub last commit](https://img.shields.io/github/last-commit/ArifRosandika/Project_CRUD_Go-Next?color=blue)  
![GitHub repo size](https://img.shields.io/github/repo-size/ArifRosandika/Project_CRUD_Go-Next)  
![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)  
![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)  
![Next.js](https://img.shields.io/badge/Next.js-black?logo=nextdotjs&logoColor=white)  
![TypeScript](https://img.shields.io/badge/TypeScript-3178C6?logo=typescript&logoColor=white)

---

## 📘 Table of Contents  
- [Overview](#overview)  
- [Features](#features)  
- [Technologies](#technologies)  
- [Getting Started](#getting-started)  
  - [Prerequisites](#prerequisites)  
  - [Installation](#installation)  
  - [Usage](#usage)  
- [Project Structure](#project-structure)  
- [Contributing](#contributing)  
- [License](#license)

---

## 🧭 Overview  
This project demonstrates a full-stack CRUD (Create, Read, Update, Delete) application using the **Go** language with the **Gin** framework on the backend, and **Next.js** with **TypeScript** on the frontend.  
It’s designed to help you (and me) understand how to build a modular and scalable web application with a clean separation between frontend and backend.

---

## ✨ Features  
- ✅ Full CRUD operations via a RESTful API  
- 🗄️ Modular backend architecture (controllers, routes, database layer)  
- 📋 Frontend built with Next.js, TypeScript and clean UI pages  
- 🔒 Input validation and clean routing logic  
- 🔧 Ready to extend — easy to add authentication, roles, etc.

---

## 🧰 Technologies  
| Layer       | Technology             | Description                                  |
|-------------|------------------------|----------------------------------------------|
| Backend     | Go + Gin               | REST API, routing, business logic             |
| Frontend    | Next.js + TypeScript   | SSR/SSG pages, React components               |
| Database    | (configurable)         | Persistent storage for CRUD operations        |

---

## 🚀 Getting Started  

### Prerequisites  
- Go (version 1.x)  
- Node.js + npm or yarn  
- Database (e.g., PostgreSQL / MySQL)  
- Git  

### Installation  
```bash
git clone https://github.com/ArifRosandika/Project_CRUD_Go-Next.git
cd Project_CRUD_Go-Next

Backend setup

cd backend
# create/config .env (DB connection, etc.)
go mod tidy
go run main.go

Frontend setup

cd ../frontend
npm install
npm run dev

Usage

Open your browser at http://localhost:3000 (or your configured port)

Use the frontend UI to create, view, update, and delete records.

The backend exposes REST API endpoints for CRUD operations (check routes folder).

Feel free to extend the project with authentication, pagination, search, etc.



---

🗂️ Project Structure

Project_CRUD_Go-Next/
├── backend/
│   ├── main.go
│   ├── controllers/
│   ├── routes/
│   └── database/
└── frontend/
    ├── pages/
    ├── components/
    ├── services/
    ├── package.json
    └── tsconfig.json


---

🤝 Contributing

Contributions are welcome! To contribute:

1. Fork the repository


2. Create a new feature branch (git checkout -b feature/my-feature)


3. Commit your changes with clear messages


4. Push your branch and open a Pull Request


5. I’ll review and merge if everything looks good



Please ensure your code follows the existing project style and you add any necessary documentation.


---

📄 License

This project is licensed under the MIT License — feel free to use, modify, and distribute. See the LICENSE file for details.

