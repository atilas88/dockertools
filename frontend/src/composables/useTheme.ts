import {readonly, ref} from 'vue'

export type Theme = 'light' | 'dark'

const STORAGE_KEY = 'dt-theme'
const media = window.matchMedia('(prefers-color-scheme: dark)')

function readStored(): Theme | null {
  try {
    const v = localStorage.getItem(STORAGE_KEY)
    return v === 'light' || v === 'dark' ? v : null
  } catch {
    return null
  }
}

function apply(t: Theme) {
  document.documentElement.dataset.theme = t
}

// Module-level state. `explicit` is true once the user picks a theme;
// until then the app follows the OS preference.
const explicit = ref(readStored() !== null)
const theme = ref<Theme>(readStored() ?? (media.matches ? 'dark' : 'light'))

// Call before mounting so the first paint already uses the right theme.
export function initTheme() {
  apply(theme.value)
  media.addEventListener('change', e => {
    if (explicit.value) return
    theme.value = e.matches ? 'dark' : 'light'
    apply(theme.value)
  })
}

export function useTheme() {
  function setTheme(t: Theme) {
    theme.value = t
    explicit.value = true
    apply(t)
    try {
      localStorage.setItem(STORAGE_KEY, t)
    } catch {
      // Storage unavailable: the choice lasts for this session only.
    }
  }

  function toggle() {
    setTheme(theme.value === 'dark' ? 'light' : 'dark')
  }

  return {theme: readonly(theme), setTheme, toggle}
}
