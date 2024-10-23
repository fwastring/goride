import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import PrimeVue from 'primevue/config';
import Lara from '@primevue/themes/lara';
import Button from 'primevue/button';
import InputText from 'primevue/inputtext'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(PrimeVue, {
    theme: {
        preset: Lara
    }
});
app.component('Button', Button)
app.component('InputText', InputText)

app.mount('#app')
