<script lang="ts" setup>
const model = defineModel<string>({default: ''})

withDefaults(defineProps<{ placeholder?: string }>(), {placeholder: 'Filtrar…'})
</script>

<template>
  <input v-model="model" class="dt-input" type="search" :placeholder="placeholder" :aria-label="placeholder">
</template>
