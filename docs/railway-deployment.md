# Railway deployment

This repository is an isolated monorepo: the Go API lives in `backend/`, and
the Next.js app lives in `frontend/`. Deploy it to Railway as two application
services plus one PostgreSQL service.

## 1. Create Railway services

1. Create a Railway project from this GitHub repository.
2. Add a PostgreSQL database service.
3. Add a backend application service:
   - Root directory: `/backend`
   - Config file: `/backend/railway.json`
   - Builder: Dockerfile
4. Add a frontend application service:
   - Root directory: `/frontend`
   - Config file: `/frontend/railway.json`

The backend Docker image runs embedded migrations automatically before the API
starts. The backend healthcheck is `/health/ready`, so Railway will only route
traffic after the API can reach PostgreSQL.

## 2. Backend variables

Set these variables on the backend service:

| Variable | Value |
| --- | --- |
| `DATABASE_URL` | Reference the Railway PostgreSQL connection URL. `TRAMPLIN_DATABASE_URL` is also supported if you prefer the app-specific name. |
| `TRAMPLIN_TOKEN_SECRET` | A long random production secret. |
| `TRAMPLIN_CORS_ALLOW_ORIGIN` | The public frontend URL, for example `https://<frontend>.up.railway.app`. |
| `TRAMPLIN_COOKIE_SECURE` | `true` |

Do not set `TRAMPLIN_HTTP_ADDR` unless you need to override the listener. If it
is unset, the backend listens on Railway's injected `PORT` variable and falls
back to `:8080` locally.

Object storage is optional. Without the `TRAMPLIN_OBJECT_STORAGE_*` variables,
the API boots normally, but upload/download URL endpoints return
`503 service_unavailable`.

## 3. Frontend variables

Set these variables on the frontend service:

| Variable | Value |
| --- | --- |
| `NEXT_PUBLIC_TRAMPLIN_API_URL` | Public browser-facing API base URL with `/v1`, for example `https://<backend>.up.railway.app/v1`. |
| `TRAMPLIN_API_URL` | Optional server-side API base URL. Use the backend public URL or a Railway private-network URL with `/v1`. |

The frontend includes a standalone `/health` route for Railway healthchecks, so
healthchecks do not depend on the API being reachable from the home page.

## 4. After the first deploy

1. Generate public domains for both app services.
2. Update backend `TRAMPLIN_CORS_ALLOW_ORIGIN` to the final frontend domain.
3. Update frontend `NEXT_PUBLIC_TRAMPLIN_API_URL` to the final backend domain
   plus `/v1`.
4. Redeploy both services.
5. Optionally seed demo data by running the backend image in `seed` mode or by
   executing `/app/tramplin-seed` in a one-off job after migrations have run.
