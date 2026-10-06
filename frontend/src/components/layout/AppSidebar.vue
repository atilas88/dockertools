<script lang="ts" setup>
import {useContainers} from '../../composables/useContainers'
import {useImages} from '../../composables/useImages'
import {useNetworks} from '../../composables/useNetworks'
import {useVolumes} from '../../composables/useVolumes'
import NavItem from './NavItem.vue'

// Stores are shared, so these counts update when a view changes the data.
const {items: containers} = useContainers()
const {items: images} = useImages()
const {items: volumes} = useVolumes()
const {items: networks} = useNetworks()
</script>

<template>
  <nav class="dt-sidebar" aria-label="Recursos">
    <div class="dt-sidebar__section">Recursos</div>
    <NavItem :to="{name: 'containers'}" icon="box" label="Contenedores" :count="containers.length"/>
    <NavItem :to="{name: 'images'}" icon="layers" label="Imágenes" :count="images.length"/>
    <NavItem :to="{name: 'volumes'}" icon="disk" label="Volúmenes" :count="volumes.length"/>
    <NavItem :to="{name: 'networks'}" icon="net" label="Redes" :count="networks.length"/>
  </nav>
</template>
