import {createRouter, createWebHashHistory} from 'vue-router'

// Hash history: Wails serves the app from an embedded FS, so there is no
// server to fall back to index.html on deep links.
export const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {path: '/', redirect: '/containers'},
    {path: '/containers', name: 'containers', component: () => import('../views/ContainersView.vue')},
    {path: '/images', name: 'images', component: () => import('../views/ImagesView.vue')},
    {path: '/volumes', name: 'volumes', component: () => import('../views/VolumesView.vue')},
    {path: '/networks', name: 'networks', component: () => import('../views/NetworksView.vue')},
  ],
})
