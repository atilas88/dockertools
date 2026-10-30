import {computed, ref, watch} from 'vue'
import {ListContainers, StartContainer, StopContainer,
       PauseContainer, UnpauseContainer, RestartContainer,
       RemoveContainer} from '../../wailsjs/go/main/App'
import type {Container, ContainerAction} from '../types/docker'
import {createListStore} from './createListStore'

const onlyRunning = ref(false)

const useStore = createListStore<Container>(() => ListContainers(onlyRunning.value))

watch(onlyRunning, () => useStore().refresh())

const bindings: Record<ContainerAction, (id: string) => Promise<void>> = {
  start: StartContainer, stop: StopContainer, restart: RestartContainer,
  pause: PauseContainer, unpause: UnpauseContainer, remove: RemoveContainer,
}

export function useContainers() {
  const store = useStore()
  const {items} = store

  const runningCount = computed(() => items.value.filter(c => c.state === 'running').length)

  async function runAction(id: string, action: ContainerAction) {
    const c = items.value.find(c => c.id === id)
    if (c) c.state = action === 'remove' ? 'removing' : 'restarting'
    try {
      await bindings[action](id)
    } finally {
      await store.refresh()
    }
  }

  return {...store, runningCount, runAction, onlyRunning}
}
