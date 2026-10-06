<script lang="ts" setup>
import AppSidebar from '../components/layout/AppSidebar.vue'
import TitleBar from '../components/layout/TitleBar.vue'
import ConnectionStatus from '../components/status/ConnectionStatus.vue'
import AppToast from '../components/ui/AppToast.vue'
import ConfirmDialog from '../components/ui/ConfirmDialog.vue'
</script>

<template>
  <TitleBar>
    <ConnectionStatus/>
  </TitleBar>
  <div class="shell">
    <AppSidebar/>
    <main>
      <RouterView/>
    </main>
  </div>
  <ConfirmDialog/>
  <AppToast/>
</template>
