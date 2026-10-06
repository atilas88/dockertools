<script lang="ts" setup>
import {computed} from 'vue'
import Badge from './Badge.vue'
import type {BadgeVariant} from './types'

// state/exitCode/health mirror container.State from the Docker SDK.
const props = defineProps<{ state: string; exitCode?: number; health?: string }>()

const variant = computed<BadgeVariant>(() => {
  if (props.health === 'unhealthy') return 'paused'
  switch (props.state) {
    case 'running': return 'running'
    case 'paused': return 'paused'
    case 'restarting':
    case 'removing': return 'restarting'
    case 'dead': return 'error'
    default: return 'exited' // exited, created
  }
})

const label = computed(() => {
  if (props.health === 'unhealthy') return 'unhealthy'
  if (props.state === 'exited') return `exited (${props.exitCode ?? 0})`
  return props.state
})
</script>

<template>
  <Badge :variant="variant">{{ label }}</Badge>
</template>
