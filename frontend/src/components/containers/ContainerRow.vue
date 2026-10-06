<script lang="ts" setup>
import type {Container, ContainerAction} from '../../types/docker'
import {timeAgo} from '../../utils/format'
import StatusBadge from '../ui/StatusBadge.vue'
import ContainerActions from './ContainerActions.vue'

defineProps<{ container: Container; selected?: boolean }>()
const emit = defineEmits<{ select: []; action: [action: ContainerAction] }>()
</script>

<template>
  <tr :aria-selected="selected" @click="emit('select')">
    <td>{{ container.name }}</td>
    <td><StatusBadge :state="container.state" :exit-code="container.exitCode"/></td>
    <td class="dt-mono">{{ container.image }}</td>
    <td class="dt-mono dt-muted">{{ container.id }}</td>
    <td class="dt-mono" :class="{'dt-muted': !container.ports.length}">{{ container.ports.join(', ') || '—' }}</td>
    <td class="caption dt-muted">{{ timeAgo(container.created) }}</td>
    <td class="dt-actions" @click.stop>
      <ContainerActions :container="container" @action="emit('action', $event)"/>
    </td>
  </tr>
</template>
