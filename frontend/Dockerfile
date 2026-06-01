FROM node:24-alpine AS deps

WORKDIR /app

RUN corepack enable

COPY package.json pnpm-lock.yaml .npmrc ./
RUN pnpm install --frozen-lockfile

FROM node:24-alpine AS build

WORKDIR /app

RUN corepack enable

COPY --from=deps /app/node_modules ./node_modules
COPY . .
RUN pnpm build

FROM node:24-alpine AS runner

WORKDIR /app

ENV NODE_ENV=production
ENV HOSTNAME=0.0.0.0

RUN addgroup --system --gid 1001 nodejs \
    && adduser --system --uid 1001 nextjs

COPY --from=build /app/public ./public
COPY --from=build --chown=nextjs:nodejs /app/.next/standalone ./
COPY --from=build --chown=nextjs:nodejs /app/.next/static ./.next/static

USER nextjs

EXPOSE 3000

CMD ["node", "server.js"]
