import {computed, ref, type Ref} from 'vue'

// Case-insensitive text filter over every field of each item.
export function useFilter<T extends object>(items: Ref<T[]>) {
  const query = ref('')

  const filtered = computed(() => {
    const q = query.value.trim().toLowerCase()
    if (!q) return items.value
    return items.value.filter(item =>
      Object.values(item).flat().join(' ').toLowerCase().includes(q))
  })

  return {query, filtered}
}
