<script lang="ts" setup>
import type {Image} from '../../types/docker'
import {formatBytes, timeAgo} from '../../utils/format'
import Badge from '../ui/Badge.vue'
import DataPanel from '../ui/DataPanel.vue'

defineProps<{ images: Image[]; loading?: boolean; emptyText?: string }>()
</script>

<template>
  <DataPanel :columns="6" :loading="loading && !images.length" :empty="!images.length" :empty-text="emptyText">
    <template #head>
      <th>Repositorio</th><th>Tag</th><th>ID</th><th>Tamaño</th><th>Creada</th><th>Uso</th>
    </template>
    <tr v-for="i in images" :key="i.id">
      <td class="dt-mono">{{ i.repo }}</td>
      <td class="dt-mono">{{ i.tag }}</td>
      <td class="dt-mono dt-muted">{{ i.id }}</td>
      <td class="dt-mono">{{ formatBytes(i.size) }}</td>
      <td class="caption dt-muted">{{ timeAgo(i.created) }}</td>
      <td>
        <Badge v-if="i.inUse" variant="running">en uso</Badge>
        <Badge v-else variant="exited">dangling</Badge>
      </td>
    </tr>
  </DataPanel>
</template>
