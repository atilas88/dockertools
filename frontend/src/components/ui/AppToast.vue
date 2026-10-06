<script lang="ts" setup>
import {useToast} from '../../composables/useToast'

const {message, visible} = useToast()
</script>

<template>
  <div class="toast" :class="{show: visible}" role="status" aria-live="polite">{{ message }}</div>
</template>

<style scoped>
.toast {
  position: fixed; right: 24px; bottom: 24px;
  background: var(--bg-100); border: 1px solid var(--line); border-radius: var(--radius-md);
  box-shadow: var(--shadow-popover); padding: var(--space-3) var(--space-4);
  font-size: 14px; opacity: 0; transform: translateY(8px); transition: .2s; pointer-events: none;
}
.toast.show { opacity: 1; transform: none; }
</style>
