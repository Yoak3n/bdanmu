import { createApp } from 'vue'
import './style.less'
import './assets/styles/variables.less'
import App from './App.vue'

const app = createApp(App)
import router from  './routes'
app.use(router)
app.mount('#app')
