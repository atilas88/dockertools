<script lang="ts" setup>
import {ref, onMounted} from 'vue'
import {GetDockerInfo} from '../../wailsjs/go/main/App'

const docker = ref({ version: '', apiVersion: '' })
const loading = ref<boolean>(true)
const error = ref<unknown | null>(null)

const fetchDockerInfo = async () => {
  loading.value = true
  error.value = null
  try {
    // Llamada directa al método de Go como si fuera una función de JS
    const data = await GetDockerInfo()
    docker.value = data
  } catch (err) {
    // Si Go devuelve un error, se captura aquí en el bloque catch
    error.value = err
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchDockerInfo()
})

</script>

<template>
  <main>
    <h2>Información de Docker</h2>
    
    <div v-if="loading">Cargando datos de Docker...</div>
    
    <div v-else-if="error" class="error-message">
      {{ error }}
    </div>
    
    <div v-else class="version-info">
      <p><strong>Versión de Docker:</strong> {{ docker.version }}</p>
      <p><strong>Versión de la API:</strong> {{ docker.apiVersion }}</p>
    </div>

    <button @click="fetchDockerInfo">Actualizar</button>
  </main>
</template>

<style scoped>
.error-message {
  color: #ff4d4d;
  font-weight: bold;
}
.version-info p {
  font-size: 1.1em;
  margin: 8px 0;
}
</style>
