import {computed} from 'vue'
import {mockResponse} from '../mocks/delay'
import {images as mockImages} from '../mocks/resources'
import type {Image} from '../types/docker'
import {createListStore} from './createListStore'

// TODO etapa 8: return ListImages() from wailsjs/go/main/App.
const useStore = createListStore<Image>(() => mockResponse(mockImages))

export function useImages() {
  const store = useStore()
  const {items} = store

  const dangling = computed(() => items.value.filter(i => !i.inUse))
  const danglingSize = computed(() => dangling.value.reduce((sum, i) => sum + i.size, 0))

  // TODO etapa 8: call PruneImages() and refresh(); return reclaimed bytes from Go.
  async function prune(): Promise<number> {
    const reclaimed = danglingSize.value
    await mockResponse(null)
    items.value = items.value.filter(i => i.inUse)
    return reclaimed
  }

  return {...store, dangling, danglingSize, prune}
}
