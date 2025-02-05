import './assets/main.css'
import 'vue3-toastify/dist/index.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import Vue3Toastify from 'vue3-toastify'
import { createPinia } from 'pinia'

const app = createApp(App)

app.use(router)
app.use(Vue3Toastify, {autoClose: 3000})

const pinia = createPinia()
app.use(pinia)

app.mount('#app')
