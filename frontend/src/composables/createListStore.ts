import {ref, type Ref} from 'vue'

export interface ListStore<T> {
  items: Ref<T[]>
  loading: Ref<boolean>
  error: Ref<string | null>
  refresh: () => Promise<void>
}

// Builds a shared (module-level) list store around a loader. The loader is
// the only thing that changes when a mock is replaced by a Go binding.
// The returned composable loads on first use.
export function createListStore<T>(loader: () => Promise<T[]>): () => ListStore<T> {
  const items = ref([]) as Ref<T[]>
  const loading = ref(false)
  const error = ref<string | null>(null)
  let loaded = false

  async function refresh() {
    loading.value = true
    error.value = null
    try {
      items.value = await loader()
      loaded = true
    } catch (err) {
      // Wails rejects with the Go error message.
      error.value = String(err)
    } finally {
      loading.value = false
    }
  }

  return () => {
    if (!loaded && !loading.value) void refresh()
    return {items, loading, error, refresh}
  }
}
