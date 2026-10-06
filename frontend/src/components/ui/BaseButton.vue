<script lang="ts" setup>
import BaseIcon from './BaseIcon.vue'
import type {IconName} from './icons'

// `label` is required for icon-only buttons: it becomes aria-label and tooltip.
const props = withDefaults(defineProps<{
  variant?: 'primary' | 'secondary' | 'ghost' | 'danger'
  icon?: IconName
  iconOnly?: boolean
  label?: string
  disabled?: boolean
}>(), {variant: 'secondary'})
</script>

<template>
  <button
    type="button"
    class="dt-btn"
    :class="[`dt-btn--${props.variant}`, {'dt-btn--icon': iconOnly}]"
    :disabled="disabled"
    :aria-label="iconOnly ? label : undefined"
    :title="iconOnly ? label : undefined"
  >
    <BaseIcon v-if="icon" :name="icon"/>
    <slot v-if="!iconOnly"/>
  </button>
</template>
