<script lang="ts" setup>
import type {RouteLocationRaw} from 'vue-router'
import BaseIcon from '../ui/BaseIcon.vue'
import type {IconName} from '../ui/icons'

// RouterLink sets aria-current="page" on the active link, which
// dockertools.css uses to highlight it.
defineProps<{ to: RouteLocationRaw; icon: IconName; label: string; count?: number }>()
</script>

<template>
  <RouterLink :to="to" class="dt-nav-item">
    <BaseIcon :name="icon"/>
    {{ label }}
    <span v-if="count !== undefined" class="dt-nav-item__count">{{ count }}</span>
  </RouterLink>
</template>
