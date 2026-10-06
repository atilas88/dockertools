<script lang="ts" setup>
import FilterInput from './FilterInput.vue'

const filter = defineModel<string>('filter', {default: ''})

defineProps<{ title: string; meta?: string }>()
</script>

<template>
  <div class="head">
    <h1 class="display">{{ title }}</h1>
    <FilterInput v-model="filter"/>
    <slot name="actions"/>
  </div>
  <p v-if="meta" class="caption meta">{{ meta }}</p>
</template>
