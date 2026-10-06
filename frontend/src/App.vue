<script lang="ts" setup>
import AppToast from './components/ui/AppToast.vue'
import ConfirmDialog from './components/ui/ConfirmDialog.vue'
</script>

<template>
  <header class="titlebar">
    <span class="brand"><i></i>dockertools</span>
  </header>
  <div class="shell">
    <main>
      <RouterView/>
    </main>
  </div>
  <ConfirmDialog/>
  <AppToast/>
</template>
