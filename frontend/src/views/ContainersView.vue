<script lang="ts" setup>
import {computed, ref, watch} from 'vue'
import ContainerList from '../components/containers/ContainerList.vue'
import LogPanel from '../components/containers/LogPanel.vue'
import BaseButton from '../components/ui/BaseButton.vue'
import PageHeader from '../components/ui/PageHeader.vue'
import {useConfirm} from '../composables/useConfirm'
import {useContainers} from '../composables/useContainers'
import {useFilter} from '../composables/useFilter'
import {useToast} from '../composables/useToast'
import type {ContainerAction} from '../types/docker'
import BaseToggle from '../components/ui/BaseToggle.vue'

const {items, loading, error, refresh, runningCount, runAction, onlyRunning} = useContainers()
const {query, filtered} = useFilter(items)
const {confirm} = useConfirm()
const toast = useToast()

const selectedId = ref<string | null>(null)
const selected = computed(() => items.value.find(c => c.id === selectedId.value))

// Select the first container when there are items but none is selected:
// once data arrives, or when the selected one disappears from the list.
watch(() => items.value.length > 0 && !selected.value, needSelection => {
  if (needSelection) selectedId.value = items.value[0].id
}, {immediate: true})

const emptyTextFilter = computed(() => onlyRunning.value ? 'No hay contenedores en ejecución.' : 'No hay contenedores.')

const meta = computed(() => onlyRunning.value
  ? `${items.value.length} en ejecución`
  : `${runningCount.value} en ejecución · ${items.value.length} en total`)

const scope = computed(() => onlyRunning.value ? 'contenedor en ejecución' : 'contenedor')

const emptyText = computed(() =>
  error.value ?? (query.value ? `Ningún ${scope.value} coincide con «${query.value}».` : emptyTextFilter.value))

const doneMessage: Record<ContainerAction, string> = {
  start: 'iniciado', stop: 'detenido', restart: 'reiniciado',
  pause: 'pausado', unpause: 'reanudado', remove: 'eliminado',
}

async function onAction(id: string, action: ContainerAction) {
  const c = items.value.find(c => c.id === id)
  if (!c) return
  const name = c.name
  if (action === 'remove' && !await confirm({
    title: `¿Eliminar ${name}?`,
    text: 'Se perderán los datos que no estén montados en volúmenes.',
  })) return

  try {
    await runAction(id, action)
    toast.show(`${name} ${doneMessage[action]}`)
  } catch (err) {
    toast.show(`Error: ${String(err)}`)
  }
}
</script>

<template>
  <PageHeader v-model:filter="query" title="Contenedores" :meta="meta">
    <template #actions>
      <BaseButton variant="ghost" icon="refresh" icon-only label="Refrescar" @click="refresh()" :disabled="loading"/>
      <BaseToggle v-model="onlyRunning">Solo en ejecución</BaseToggle>
      <BaseButton variant="primary" icon="plus" @click="toast.show('Demo: esta acción abriría un diálogo')">Ejecutar</BaseButton>
    </template>
  </PageHeader>

  <ContainerList
    :containers="filtered" :selected-id="selectedId" :loading="loading" :empty-text="emptyText"
    @select="selectedId = $event" @action="onAction"
  />

  <LogPanel v-if="selected" :container="selected" @download="toast.show('Demo: descargaría los logs')"/>
</template>
