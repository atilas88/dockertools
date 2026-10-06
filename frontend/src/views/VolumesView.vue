<script lang="ts" setup>
import {computed} from 'vue'
import VolumeList from '../components/volumes/VolumeList.vue'
import BaseButton from '../components/ui/BaseButton.vue'
import PageHeader from '../components/ui/PageHeader.vue'
import {useFilter} from '../composables/useFilter'
import {useToast} from '../composables/useToast'
import {useVolumes} from '../composables/useVolumes'

const {items, loading, error} = useVolumes()
const {query, filtered} = useFilter(items)
const toast = useToast()

const emptyText = computed(() =>
  error.value ?? (query.value ? `Ningún volumen coincide con «${query.value}».` : 'No hay volúmenes.'))
</script>

<template>
  <PageHeader v-model:filter="query" title="Volúmenes" :meta="`${items.length} volúmenes`">
    <template #actions>
      <BaseButton variant="primary" icon="plus" @click="toast.show('Demo: esta acción abriría un diálogo')">Crear volumen</BaseButton>
    </template>
  </PageHeader>

  <VolumeList :volumes="filtered" :loading="loading" :empty-text="emptyText"/>
</template>
