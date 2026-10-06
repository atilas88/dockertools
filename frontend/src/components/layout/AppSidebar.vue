<script lang="ts" setup>
import NavItem from './NavItem.vue'
</script>

<template>
  <nav class="dt-sidebar" aria-label="Recursos">
    <div class="dt-sidebar__section">Recursos</div>
    <NavItem :to="{name: 'containers'}" icon="box" label="Contenedores"/>
    <NavItem :to="{name: 'images'}" icon="layers" label="Imágenes"/>
    <NavItem :to="{name: 'volumes'}" icon="disk" label="Volúmenes"/>
    <NavItem :to="{name: 'networks'}" icon="net" label="Redes"/>
  </nav>
</template>
