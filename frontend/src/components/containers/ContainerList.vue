<script lang="ts" setup>
import type {Container, ContainerAction} from '../../types/docker'
import DataPanel from '../ui/DataPanel.vue'
import ContainerRow from './ContainerRow.vue'

defineProps<{ containers: Container[]; selectedId?: string | null; loading?: boolean; emptyText?: string }>()
const emit = defineEmits<{ select: [id: string]; action: [id: string, action: ContainerAction] }>()
</script>

<template>
  <DataPanel :columns="7" :loading="loading && !containers.length" :empty="!containers.length" :empty-text="emptyText">
    <template #head>
      <th>Nombre</th><th>Estado</th><th>Imagen</th><th>ID</th><th>Puertos</th><th>Creado</th><th class="dt-actions"></th>
    </template>
    <ContainerRow
      v-for="c in containers" :key="c.id"
      :container="c" :selected="c.id === selectedId"
      @select="emit('select', c.id)" @action="emit('action', c.id, $event)"
    />
  </DataPanel>
</template>
