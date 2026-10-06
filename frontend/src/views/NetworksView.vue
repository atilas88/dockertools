<script lang="ts" setup>
import {computed} from 'vue'
import NetworkList from '../components/networks/NetworkList.vue'
import BaseButton from '../components/ui/BaseButton.vue'
import PageHeader from '../components/ui/PageHeader.vue'
import {useFilter} from '../composables/useFilter'
import {useNetworks} from '../composables/useNetworks'
import {useToast} from '../composables/useToast'

const {items, loading, error} = useNetworks()
const {query, filtered} = useFilter(items)
const toast = useToast()

const emptyText = computed(() =>
  error.value ?? (query.value ? `Ninguna red coincide con «${query.value}».` : 'No hay redes.'))
</script>

<template>
  <PageHeader v-model:filter="query" title="Redes" :meta="`${items.length} redes`">
    <template #actions>
      <BaseButton variant="primary" icon="plus" @click="toast.show('Demo: esta acción abriría un diálogo')">Crear red</BaseButton>
    </template>
  </PageHeader>

  <NetworkList :networks="filtered" :loading="loading" :empty-text="emptyText"/>
</template>
