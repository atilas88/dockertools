import {readonly, ref} from 'vue'

export interface ConfirmOptions {
  title: string
  text: string
  confirmLabel?: string
}

// Module-level state: a single dialog instance lives in the layout.
const options = ref<ConfirmOptions | null>(null)
let settle: ((ok: boolean) => void) | undefined

export function useConfirm() {
  // Resolves to true when the user confirms, false on cancel/close.
  function confirm(opts: ConfirmOptions): Promise<boolean> {
    settle?.(false)
    options.value = opts
    return new Promise(resolve => (settle = resolve))
  }

  function answer(ok: boolean) {
    settle?.(ok)
    settle = undefined
    options.value = null
  }

  return {options: readonly(options), confirm, answer}
}
