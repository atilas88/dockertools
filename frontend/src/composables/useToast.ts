import {readonly, ref} from 'vue'

// Module-level state: every caller shares the same toast.
const message = ref('')
const visible = ref(false)
let timer: ReturnType<typeof setTimeout> | undefined

export function useToast() {
  function show(text: string, duration = 2200) {
    message.value = text
    visible.value = true
    clearTimeout(timer)
    timer = setTimeout(() => (visible.value = false), duration)
  }

  return {message: readonly(message), visible: readonly(visible), show}
}
