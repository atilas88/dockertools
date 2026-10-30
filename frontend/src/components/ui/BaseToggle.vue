<script lang="ts" setup>
const model = defineModel<boolean>({default: false})

defineProps<{ disabled?: boolean }>()
</script>

<template>
  <label class="dt-toggle" :class="{'dt-toggle--disabled': disabled}">
    <input v-model="model" class="dt-toggle__input" type="checkbox" role="switch" :disabled="disabled">
    <slot/>
  </label>
</template>
