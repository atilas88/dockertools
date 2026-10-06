// Shapes the UI expects from Go. When a DTO is added to internal/docker/dto.go
// and the binding is generated, prefer the class from wailsjs/go/models and
// keep these in sync (or replace them).

export type ContainerState = 'created' | 'running' | 'paused' | 'restarting' | 'removing' | 'exited' | 'dead'

export interface Container {
  id: string        // short ID (12 chars)
  name: string      // without the leading "/"
  image: string
  state: ContainerState
  exitCode?: number
  ports: string[]   // "5432→5432"
  created: number   // unix seconds
}

export type ContainerAction = 'start' | 'stop' | 'restart' | 'pause' | 'unpause' | 'remove'

export interface LogLine {
  ts: string        // RFC3339
  text: string
  stream: 'stdout' | 'stderr'
}

export interface Image {
  id: string        // "sha256:" + 8 chars
  repo: string
  tag: string
  size: number      // bytes
  created: number   // unix seconds
  inUse: boolean
}

export interface Volume {
  name: string
  driver: string
  size: number      // bytes, -1 when unknown
  usedBy: string[]  // container names
}

export interface Network {
  id: string
  name: string
  driver: string
  scope: string
  subnet: string
}
