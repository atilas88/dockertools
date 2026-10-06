// TODO etapa 8: replace with ListImages / ListVolumes / ListNetworks bindings.
import type {Image, Network, Volume} from '../types/docker'
import {ago} from './delay'

const DAY = 86400
const MB = 1_000_000

export const images: Image[] = [
  {id: 'sha256:4b1f0a2c', repo: 'postgres', tag: '16-alpine', size: 243 * MB, created: ago(14 * DAY), inUse: true},
  {id: 'sha256:9e02c7d1', repo: 'php', tag: '8.3-fpm', size: 494 * MB, created: ago(21 * DAY), inUse: true},
  {id: 'sha256:c8a1e5f3', repo: 'nginx', tag: '1.27-alpine', size: 47.9 * MB, created: ago(30 * DAY), inUse: true},
  {id: 'sha256:71d0b3aa', repo: 'redis', tag: '7', size: 117 * MB, created: ago(32 * DAY), inUse: true},
  {id: 'sha256:02f9dd6e', repo: '<none>', tag: '<none>', size: 512 * MB, created: ago(60 * DAY), inUse: false},
]

export const volumes: Volume[] = [
  {name: 'pgdata', driver: 'local', size: 1200 * MB, usedBy: ['postgres-dev']},
  {name: 'redisdata', driver: 'local', size: 38 * MB, usedBy: ['redis-cache']},
  {name: 'composer-cache', driver: 'local', size: 310 * MB, usedBy: []},
]

export const networks: Network[] = [
  {id: 'f1c2a3b4d5e6', name: 'bridge', driver: 'bridge', scope: 'local', subnet: '172.17.0.0/16'},
  {id: 'a9b8c7d6e5f4', name: 'app_default', driver: 'bridge', scope: 'local', subnet: '172.20.0.0/16'},
  {id: '0a1b2c3d4e5f', name: 'host', driver: 'host', scope: 'local', subnet: ''},
  {id: '6f5e4d3c2b1a', name: 'none', driver: 'null', scope: 'local', subnet: ''},
]
