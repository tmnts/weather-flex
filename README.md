# 🌦️ Weather-Flex: Full-Stack Monorepo MVP

A high-end, real-time weather station built with a **Go** microservice and a **Svelte 5** reactive frontend. Designed for high performance, modularity, and aesthetic minimalism.

## 🏗️ Architecture Overview

- **Monorepo Management:** Powered by **Turborepo** & **pnpm Workspaces** for seamless local development and shared configuration.
- **Backend (Go 1.22):** A lightweight serverless-ready API that proxies requests to **OpenWeatherMap**, handling authentication via secure environment variables.
- **Frontend (Svelte 5 + Tailwind v4):** A cutting-edge reactive UI utilizing Svelte **Runes** (`$state`, `$effect`) and a custom "Glassmorphism" design system.
- **Cloud Infrastructure:** CI/CD via **Vercel** with automatic deployment of Go Serverless Functions.

## 🛠️ The Tech Stack


| Layer | Technology | Why? |
| :--- | :--- | :--- |
| **Frontend** | SvelteKit 5 | Next-gen reactivity and superior DX. |
| **Backend** | Go (Golang) | High-concurrency, type-safety, and tiny binary footprint. |
| **Styling** | Tailwind CSS v4 | Rapid UI development with utility-first CSS. |
| **Deployment** | Vercel | Global Edge Network and native Go runtime support. |
| **Environment** | Turborepo | Optimized build pipelines and task execution. |

## 🚀 Key Technical Challenges Overcome

1.  **CORS & Networking:** Implemented a **Vite Proxy** bridge during development to resolve cross-origin issues between the Svelte dev server (:5173) and the Go API (:3000).
2.  **Infrastructure as Code:** Configured `vercel.json` rewrites to map serverless Go functions to a unified `/api` endpoint, ensuring a clean URL structure.
3.  **Resilience:** Handled regional API limitations and networking "Hells" (IPv6 vs IPv4) by standardizing local connections to `IPv1.0.0.1`.

## 📖 Local Development

1.  Clone the repository.
2.  Add `OPENWEATHER_API_KEY` to your environment.
3.  Run `pnpm install`.
4.  Execute `pnpm dev` to start the Turborepo pipeline.

## 📖 API Documentation (Swagger/OpenAPI)

The backend is self-documenting via **Swagger UI**, providing an interactive sandbox for frontend consumers.

- **Interactive UI:** [https://weather-flex.vercel.app](https://weather-flex.vercel.app)
- **OpenAPI Specification:** Generated from Go source comments using `swag init`.

### 🔄 Data Flow (Sequence)
1. **Client** (Svelte 5) requests `/api/weather`.
2. **Vercel** routes the request to the **Go Serverless Function**.
3. **Go Middleware** checks for the `swagger` path; if absent, it invokes the Weather Handler.
4. **Backend** fetches authenticated data from **OpenWeatherMap** (Server-to-Server).
5. **JSON Response** is proxied back to the client, bypassing regional CORS restrictions.

---
*Made with ❤️ for the 2026 Tech Scene.*# 

## Architecture Scheme (Mermaid)
<pre> ```mermaid
flowchart LR
    subgraph Client
        A[Svelte 5 UI]
    end

    subgraph Backend
        B[Go Microservice]
        C[Domain Layer]
        D[OpenWeather Adapter]
    end

    A -->|REST| B
    B --> C
    C --> D
    D -->|External API| E[(OpenWeather)] 
</pre>``` 
