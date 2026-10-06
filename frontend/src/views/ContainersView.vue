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

const {items, loading, error, runningCount, runAction} = useContainers()
const {query, filtered} = useFilter(items)
const {confirm} = useConfirm()
const toast = useToast()

const selectedId = ref<string | null>(null)
const selected = computed(() => items.value.find(c => c.id === selectedId.value))

// Select the first container once data arrives, like the design demo.
watch(() => items.value.length > 0, hasItems => {
  if (hasItems && selectedId.value === null) selectedId.value = items.value[0].id
}, {immediate: true})

const meta = computed(() => `${runningCount.value} en ejecución · ${items.value.length} en total`)
const emptyText = computed(() =>
  error.value ?? (query.value ? `Ningún contenedor coincide con «${query.value}».` : 'No hay contenedores.'))

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
      <BaseButton variant="primary" icon="plus" @click="toast.show('Demo: esta acción abriría un diálogo')">Ejecutar</BaseButton>
    </template>
  </PageHeader>

  <ContainerList
    :containers="filtered" :selected-id="selectedId" :loading="loading" :empty-text="emptyText"
    @select="selectedId = $event" @action="onAction"
  />

  <LogPanel v-if="selected" :container="selected" @download="toast.show('Demo: descargaría los logs')"/>
</template>
