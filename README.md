
![Awesome](https://img.shields.io/badge/awesome-yes-8a2be2.svg?style=flat-square) ![Last Commit](https://img.shields.io/badge/last%20commit-january%202026-9acd32.svg?style=flat-square) ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=flat-square&logo=go&logoColor=white) ![TypeScript](https://img.shields.io/badge/typescript-%23007ACC.svg?style=flat-square&logo=typescript&logoColor=white) ![JavaScript](https://img.shields.io/badge/javascript-%23323330.svg?style=flat-square&logo=javascript&logoColor=%23F7DF1E) ![HTML5](https://img.shields.io/badge/html5-%23E34F26.svg?style=flat-square&logo=html5&logoColor=white) ![CSS3](https://img.shields.io/badge/css3-%231572B6.svg?style=flat-square&logo=css3&logoColor=white) ![React](https://img.shields.io/badge/react-%2320232a.svg?style=flat-square&logo=react&logoColor=%2361DAFB) ![TailwindCSS](https://img.shields.io/badge/tailwindcss-%2338B2AC.svg?style=flat-square&logo=tailwind-css&logoColor=white) ![Vite](https://img.shields.io/badge/vite-%23646CFF.svg?style=flat-square&logo=vite&logoColor=white) ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=flat-square&logo=docker&logoColor=white) ![PostCSS](https://img.shields.io/badge/postcss-%23DD3A0A.svg?style=flat-square&logo=postcss&logoColor=white) ![JSON](https://img.shields.io/badge/json-%23000000.svg?style=flat-square&logo=json&logoColor=white) ![YAML](https://img.shields.io/badge/yaml-%23ffffff.svg?style=flat-square&logo=yaml&logoColor=black) ![Markdown](https://img.shields.io/badge/markdown-%23000000.svg?style=flat-square&logo=markdown&logoColor=white) ![Git](https://img.shields.io/badge/git-%23F05033.svg?style=flat-square&logo=git&logoColor=white) ![ESLint](https://img.shields.io/badge/eslint-%234B32C3.svg?style=flat-square&logo=eslint&logoColor=white)

# Polymarket Trader

A comprehensive trading application built with a Go backend and a React/TypeScript frontend.

## Project Structure

This project is organized into a clear separation of concerns between the backend and frontend services.

```text
polymarket-trader
├── .gitignore
├── docker-compose.yml
├── backend
│   ├── .env.example
│   ├── go.mod
│   ├── go.sum
│   ├── server.exe
│   ├── cmd
│   │   └── server
│   │       ├── main.go
│   │       └── main_test.go
│   └── internal
│       ├── adapters
│       │   └── polymarket
│       │       ├── client.go
│       │       └── websocket.go
│       ├── api
│       │   ├── handlers.go
│       │   └── router.go
│       ├── config
│       │   └── config.go
│       ├── core
│       │   ├── analytics.go
│       │   ├── copy_engine.go
│       │   └── trader_discovery.go
│       └── models
│           └── models.go
└── frontend
    ├── .gitignore
    ├── components.json
    ├── eslint.config.js
    ├── index.html
    ├── package-lock.json
    ├── package.json
    ├── postcss.config.js
    ├── tailwind.config.js
    ├── tsconfig.app.json
    ├── tsconfig.json
    ├── tsconfig.node.json
    ├── vite.config.ts
    ├── public
    │   └── vite.svg
    └── src
        ├── App.css
        ├── App.test.tsx
        ├── App.tsx
        ├── index.css
        ├── main.tsx
        ├── assets
        │   └── react.svg
        ├── components
        │   ├── ActivePositions.tsx
        │   ├── BotChat.tsx
        │   ├── CopyConfigDialog.tsx
        │   └── ui
        │       ├── avatar.tsx
        │       ├── badge.tsx
        │       ├── button.tsx
        │       ├── card.tsx
        │       ├── dialog.tsx
        │       ├── input.tsx
        │       ├── label.tsx
        │       ├── scroll-area.tsx
        │       ├── table.tsx
        │       └── tabs.tsx
        ├── lib
        │   └── utils.ts
        └── test
            └── setup.ts
```

## Backend (`/backend`)
The backend is built with **Go** and follows a standard clean architecture layout:
- **`cmd/`**: Contains the main entry points for the application.
- **`internal/`**: Private application code.
- **`adapters/`**: implementations of interfaces for external services (e.g., Polymarket).
- **`api/`**: HTTP handlers and router configuration.
- **`core/`**: Business logic and domain services.

## Frontend (`/frontend`)
The frontend is a **React** application powered by **Vite** and **TypeScript**:
- **`src/`**: Source code for the frontend application.
- **`components/`**: Reusable UI components.
- **`lib/`**: Helper functions and utilities.
