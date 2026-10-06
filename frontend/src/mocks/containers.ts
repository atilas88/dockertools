// TODO etapa 2: replace with the ListContainers binding.
import type {Container, LogLine} from '../types/docker'
import {ago} from './delay'

const HOUR = 3600
const DAY = 86400

export const containers: Container[] = [
  {id: 'a3f9c2e81b04', name: 'postgres-dev', image: 'postgres:16-alpine', state: 'running', ports: ['5432→5432'], created: ago(3 * HOUR)},
  {id: '7c01de55a9f2', name: 'laravel-app', image: 'php:8.3-fpm', state: 'running', ports: ['9000→9000'], created: ago(DAY)},
  {id: '5d2b90c1e7aa', name: 'nginx-proxy', image: 'nginx:1.27-alpine', state: 'running', ports: ['80→80', '443→443'], created: ago(DAY)},
  {id: 'e4b27f0c3d18', name: 'redis-cache', image: 'redis:7', state: 'paused', ports: ['6379→6379'], created: ago(2 * DAY)},
  {id: '19ad6c7e02bb', name: 'mailpit', image: 'axllent/mailpit', state: 'exited', exitCode: 0, ports: [], created: ago(5 * DAY)},
  {id: 'b8e3f14a6c90', name: 'worker-queue', image: 'php:8.3-cli', state: 'dead', ports: [], created: ago(6 * DAY)},
]

// TODO etapa 5: replace with log streaming (ContainerLogs + EventsEmit).
export function containerLogs(c: Container): LogLine[] {
  const ts = (i: number) => `2026-09-29T14:0${i}:1${i}Z`
  const out = (i: number, text: string): LogLine => ({ts: ts(i), text, stream: 'stdout'})
  const err = (i: number, text: string): LogLine => ({ts: ts(i), text, stream: 'stderr'})

  if (c.state === 'dead') {
    return [err(1, 'PHP Fatal error:  Allowed memory size of 134217728 bytes exhausted'), out(2, 'worker exited with code 255')]
  }
  if (c.state === 'exited') {
    return [out(1, `${c.name} shutting down`), out(2, 'bye')]
  }
  return [
    out(1, `starting ${c.image}`),
    out(2, `listening on ${c.ports.join(', ') || '—'}`),
    out(3, 'ready to accept connections'),
    c.name === 'postgres-dev'
      ? err(5, 'FATAL:  password authentication failed for user "app"')
      : out(4, 'GET /health 200 1.2ms'),
    out(6, 'checkpoint complete'),
  ]
}
