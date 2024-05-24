import { createApp } from 'vue'
import './style.less'
import App from './App.vue'

const app = createApp(App)
import router from  './routes'
app.use(router)
import { createPinia } from 'pinia'
const pinia = createPinia()
app.use(pinia)
app.mount('#app')
