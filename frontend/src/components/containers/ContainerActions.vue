<script lang="ts" setup>
import type {Container, ContainerAction} from '../../types/docker'
import BaseButton from '../ui/BaseButton.vue'

defineProps<{ container: Container }>()
const emit = defineEmits<{ action: [action: ContainerAction] }>()
</script>

<template>
  <template v-if="container.state === 'running'">
    <BaseButton variant="ghost" icon="pause" icon-only label="Pausar" @click="emit('action', 'pause')"/>
    <BaseButton variant="ghost" icon="stop" icon-only label="Detener" @click="emit('action', 'stop')"/>
    <BaseButton variant="ghost" icon="restart" icon-only label="Reiniciar" @click="emit('action', 'restart')"/>
  </template>
  <template v-else-if="container.state === 'paused'">
    <BaseButton variant="ghost" icon="play" icon-only label="Reanudar" @click="emit('action', 'unpause')"/>
    <BaseButton variant="ghost" icon="stop" icon-only label="Detener" @click="emit('action', 'stop')"/>
  </template>
  <template v-else-if="container.state !== 'restarting' && container.state !== 'removing'">
    <BaseButton variant="ghost" icon="play" icon-only label="Iniciar" @click="emit('action', 'start')"/>
    <BaseButton variant="ghost" icon="trash" icon-only label="Eliminar" @click="emit('action', 'remove')"/>
  </template>
</template>
