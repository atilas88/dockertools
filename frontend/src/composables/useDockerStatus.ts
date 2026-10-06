import {ref} from 'vue'
import {GetDockerInfo, GetDockerStatus} from '../../wailsjs/go/main/App'
import {docker} from '../../wailsjs/go/models'

export type ConnectionState = 'checking' | 'connected' | 'disconnected'

// Module-level state: the titlebar and any view share one connection status.
const state = ref<ConnectionState>('checking')
const info = ref<docker.DockerInfo | null>(null)
const error = ref<string | null>(null)

// TODO etapa 1: expose cli.DaemonHost() from Go instead of this mock.
const host = ref('unix:///var/run/docker.sock')

export function useDockerStatus() {
  async function refresh() {
    state.value = 'checking'
    error.value = null
    try {
      // Ping first: it fails fast (3s timeout in Go) when the daemon is down.
      await GetDockerStatus()
      info.value = await GetDockerInfo()
      state.value = 'connected'
    } catch (err) {
      // Wails rejects the promise with the Go error message (string).
      info.value = null
      error.value = String(err)
      state.value = 'disconnected'
    }
  }

  return {state, info, error, host, refresh}
}
