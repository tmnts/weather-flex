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

---
*Made with ❤️ for the 2026 Tech Scene.*# sv

Everything you need to build a Svelte project, powered by [`sv`](https://github.com/sveltejs/cli).

## Creating a project

If you're seeing this, you've probably already done this step. Congrats!

```sh
# create a new project
npx sv create my-app
```

To recreate this project with the same configuration:

```sh
# recreate this project
pnpm dlx sv create --template minimal --types ts --add tailwindcss="plugins:typography,forms" prettier --install pnpm ./
```

## Developing

Once you've created a project and installed dependencies with `npm install` (or `pnpm install` or `yarn`), start a development server:

```sh
npm run dev

# or start the server and open the app in a new browser tab
npm run dev -- --open
```

## Building

To create a production version of your app:

```sh
npm run build
```

You can preview the production build with `npm run preview`.

> To deploy your app, you may need to install an [adapter](https://svelte.dev/docs/kit/adapters) for your target environment.
