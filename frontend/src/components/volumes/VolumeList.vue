<script lang="ts" setup>
import type {Volume} from '../../types/docker'
import {formatBytes} from '../../utils/format'
import DataPanel from '../ui/DataPanel.vue'

defineProps<{ volumes: Volume[]; loading?: boolean; emptyText?: string }>()
</script>

<template>
  <DataPanel :columns="4" :loading="loading && !volumes.length" :empty="!volumes.length" :empty-text="emptyText">
    <template #head>
      <th>Nombre</th><th>Driver</th><th>Tamaño</th><th>Usado por</th>
    </template>
    <tr v-for="v in volumes" :key="v.name">
      <td>{{ v.name }}</td>
      <td class="dt-mono">{{ v.driver }}</td>
      <td class="dt-mono">{{ formatBytes(v.size) }}</td>
      <td v-if="v.usedBy.length" class="dt-mono">{{ v.usedBy.join(', ') }}</td>
      <td v-else class="dt-muted caption">sin uso</td>
    </tr>
  </DataPanel>
</template>
