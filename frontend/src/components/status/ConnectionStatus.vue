<script lang="ts" setup>
import {computed, onMounted} from 'vue'
import {useDockerStatus} from '../../composables/useDockerStatus'
import Badge from '../ui/Badge.vue'

const {state, info, error, host, refresh} = useDockerStatus()

onMounted(refresh)

const badge = computed(() => {
  switch (state.value) {
    case 'connected': return {variant: 'running', text: 'conectado'} as const
    case 'disconnected': return {variant: 'error', text: 'desconectado'} as const
    default: return {variant: 'restarting', text: 'conectando…'} as const
  }
})

const tooltip = computed(() =>
  state.value === 'disconnected' ? `${error.value}\nClic para reintentar` : 'Clic para comprobar de nuevo')
</script>

<template>
  <button class="status" type="button" :title="tooltip" :disabled="state === 'checking'" @click="refresh">
    <Badge :variant="badge.variant">{{ badge.text }}</Badge>
    <span>{{ host }}</span>
    <span v-if="info">· Engine {{ info.version }} · API {{ info.apiVersion }}</span>
  </button>
</template>

<style scoped>
.status {
  display: flex; align-items: center; gap: var(--space-2);
  padding: 0; border: 0; background: none; color: inherit; font: inherit; cursor: pointer;
}
.status:disabled { cursor: progress; }
.status:focus-visible { outline: 2px solid var(--focus); outline-offset: 2px; border-radius: var(--radius-sm); }
</style>
