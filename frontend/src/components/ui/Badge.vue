<script lang="ts" setup>
import type {BadgeVariant} from './types'

defineProps<{ variant: BadgeVariant }>()
</script>

<template>
  <span class="dt-badge" :class="`dt-badge--${variant}`"><slot/></span>
</template>
