<script lang="ts" setup>
// Table wrapper: pass header cells in #head and rows in the default slot.
// When `empty` is true the `emptyText` row is shown instead of the rows.
defineProps<{ columns: number; empty?: boolean; emptyText?: string; loading?: boolean }>()
</script>

<template>
  <div class="dt-panel">
    <table class="dt-table">
      <thead>
        <tr><slot name="head"/></tr>
      </thead>
      <tbody>
        <tr v-if="loading">
          <td :colspan="columns" class="empty">Cargando…</td>
        </tr>
        <tr v-else-if="empty">
          <td :colspan="columns" class="empty">{{ emptyText ?? 'Sin resultados.' }}</td>
        </tr>
        <slot v-else/>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
/* Narrow windows: scroll the table instead of clipping the action column. */
.dt-panel { overflow-x: auto; }
</style>
