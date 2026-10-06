import {createApp} from 'vue'
import App from './App.vue'
import {router} from './router'
import {initTheme} from './composables/useTheme'

import '@fontsource/ibm-plex-sans/400.css'
import '@fontsource/ibm-plex-sans/500.css'
import '@fontsource/ibm-plex-sans/600.css'
import '@fontsource/jetbrains-mono/400.css'
import '@fontsource/jetbrains-mono/500.css'
import './styles/tokens.css'
import './styles/dockertools.css'
import './style.css'

initTheme()
createApp(App).use(router).mount('#app')
