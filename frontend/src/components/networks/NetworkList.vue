<script lang="ts" setup>
import type {Network} from '../../types/docker'
import DataPanel from '../ui/DataPanel.vue'

defineProps<{ networks: Network[]; loading?: boolean; emptyText?: string }>()
</script>

<template>
  <DataPanel :columns="4" :loading="loading && !networks.length" :empty="!networks.length" :empty-text="emptyText">
    <template #head>
      <th>Nombre</th><th>Driver</th><th>Scope</th><th>Subred</th>
    </template>
    <tr v-for="n in networks" :key="n.id">
      <td>{{ n.name }}</td>
      <td class="dt-mono">{{ n.driver }}</td>
      <td class="dt-mono dt-muted">{{ n.scope }}</td>
      <td class="dt-mono" :class="{'dt-muted': !n.subnet}">{{ n.subnet || '—' }}</td>
    </tr>
  </DataPanel>
</template>
