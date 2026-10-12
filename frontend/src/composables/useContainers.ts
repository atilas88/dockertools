import {computed} from 'vue'
import {mockResponse} from '../mocks/delay'
import {ListContainers} from '../../wailsjs/go/main/App'
import type {Container, ContainerAction, ContainerState} from '../types/docker'
import {createListStore} from './createListStore'

const useStore = createListStore<Container>(ListContainers)

// State each action leaves the container in (mock only).
const targetState: Record<Exclude<ContainerAction, 'remove'>, ContainerState> = {
  start: 'running',
  stop: 'exited',
  restart: 'running',
  pause: 'paused',
  unpause: 'running',
}

export function useContainers() {
  const store = useStore()
  const {items} = store

  const runningCount = computed(() => items.value.filter(c => c.state === 'running').length)

  // TODO etapa 3: call the Go binding for each action and then refresh().
  async function runAction(id: string, action: ContainerAction) {
    const c = items.value.find(c => c.id === id)
    if (!c) return

    if (action === 'remove') {
      await mockResponse(null)
      items.value = items.value.filter(x => x.id !== id)
      return
    }
    if (action === 'pause' || action === 'unpause') {
      c.state = targetState[action]
      return
    }
    c.state = 'restarting'
    await mockResponse(null, 900)
    c.state = targetState[action]
    c.exitCode = c.state === 'exited' ? 0 : undefined
  }

  return {...store, runningCount, runAction}
}
