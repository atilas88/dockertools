<script lang="ts" setup>
import {computed} from 'vue'
import ImageList from '../components/images/ImageList.vue'
import BaseButton from '../components/ui/BaseButton.vue'
import PageHeader from '../components/ui/PageHeader.vue'
import {useConfirm} from '../composables/useConfirm'
import {useFilter} from '../composables/useFilter'
import {useImages} from '../composables/useImages'
import {useToast} from '../composables/useToast'
import {formatBytes} from '../utils/format'

const {items, loading, error, dangling, danglingSize, prune} = useImages()
const {query, filtered} = useFilter(items)
const {confirm} = useConfirm()
const toast = useToast()

const meta = computed(() => `${items.value.length} imágenes · ${dangling.value.length} sin usar`)
const emptyText = computed(() =>
  error.value ?? (query.value ? `Ninguna imagen coincide con «${query.value}».` : 'No hay imágenes.'))

async function onPrune() {
  if (!await confirm({
    title: '¿Eliminar imágenes sin usar?',
    text: `Se eliminarán las imágenes dangling (${formatBytes(danglingSize.value)}).`,
  })) return
  try {
    toast.show(`${formatBytes(await prune())} liberados`)
  } catch (err) {
    toast.show(`Error: ${String(err)}`)
  }
}
</script>

<template>
  <PageHeader v-model:filter="query" title="Imágenes" :meta="meta">
    <template #actions>
      <BaseButton variant="danger" icon="trash" :disabled="!dangling.length" @click="onPrune">Limpiar sin usar</BaseButton>
      <BaseButton variant="primary" icon="down" @click="toast.show('Demo: esta acción abriría un diálogo')">Descargar imagen</BaseButton>
    </template>
  </PageHeader>

  <ImageList :images="filtered" :loading="loading" :empty-text="emptyText"/>
</template>
