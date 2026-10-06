<script lang="ts" setup>
import {ref, watch} from 'vue'
import {useConfirm} from '../../composables/useConfirm'
import BaseButton from './BaseButton.vue'

const {options, answer} = useConfirm()
const dialog = ref<HTMLDialogElement>()

// Native <dialog>: showModal() gives focus trap, Esc and backdrop for free.
watch(options, opts => {
  if (opts && !dialog.value?.open) dialog.value?.showModal()
  else if (!opts && dialog.value?.open) dialog.value.close()
})
</script>

<template>
  <dialog ref="dialog" @close="answer(false)">
    <template v-if="options">
      <h2 class="title">{{ options.title }}</h2>
      <p class="body">{{ options.text }}</p>
      <div class="btns">
        <BaseButton variant="ghost" @click="answer(false)">Cancelar</BaseButton>
        <BaseButton variant="danger" @click="answer(true)">{{ options.confirmLabel ?? 'Eliminar' }}</BaseButton>
      </div>
    </template>
  </dialog>
</template>

<style scoped>
dialog {
  background: var(--bg-100); color: var(--ink); border: 1px solid var(--line);
  border-radius: var(--radius-lg); box-shadow: var(--shadow-popover);
  padding: var(--space-6); max-width: 420px;
}
dialog::backdrop { background: rgba(0, 0, 0, .45); }
h2 { margin: 0 0 var(--space-2); }
p { margin: 0 0 var(--space-6); color: var(--ink-muted); }
.btns { display: flex; justify-content: flex-end; gap: var(--space-2); }
</style>
