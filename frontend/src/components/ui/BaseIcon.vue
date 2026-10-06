<script lang="ts" setup>
import {icons, type IconName} from './icons'

defineProps<{ name: IconName }>()
</script>

<template>
  <svg
    viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5"
    stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"
    v-html="icons[name]"
  />
</template>
