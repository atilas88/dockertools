<script lang="ts" setup>
import {computed} from 'vue'
import {containerLogs} from '../../mocks/containers'
import type {Container} from '../../types/docker'
import Badge from '../ui/Badge.vue'
import BaseButton from '../ui/BaseButton.vue'

const props = defineProps<{ container: Container }>()
defineEmits<{ download: [] }>()

// TODO etapa 5: subscribe to a log stream (useLogs) instead of mock lines.
const lines = computed(() => containerLogs(props.container))
</script>

<template>
  <section class="dt-log" :aria-label="`Logs de ${container.name}`">
    <div class="dt-log__bar">
      <span class="bar-title">Logs · {{ container.name }}</span>
      <Badge v-if="container.state === 'running'" variant="running">siguiendo</Badge>
      <Badge v-else variant="exited">detenido</Badge>
      <BaseButton icon="down" @click="$emit('download')">Descargar</BaseButton>
    </div>
    <pre class="dt-log__body"><template v-for="(l, i) in lines" :key="i"><span class="dt-log__ts">{{ l.ts }}</span><span :class="{'dt-log__err': l.stream === 'stderr'}">{{ l.text }}</span>{{ '\n' }}</template></pre>
  </section>
</template>

<style scoped>
.dt-log { margin-top: var(--space-4); }
.dt-log__body { height: 180px; }
.bar-title { margin-right: auto; }
</style>
