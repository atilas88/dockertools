<script lang="ts" setup>
import {computed} from 'vue'
import {useTheme} from '../../composables/useTheme'
import BaseButton from '../ui/BaseButton.vue'

const {theme, toggle} = useTheme()

const label = computed(() => theme.value === 'dark' ? 'Cambiar a tema claro' : 'Cambiar a tema oscuro')
</script>

<template>
  <BaseButton variant="ghost" icon="theme" icon-only :label="label" @click="toggle"/>
</template>
