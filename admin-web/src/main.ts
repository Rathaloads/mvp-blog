import { createApp } from 'vue';
import "./styles/index.scss";
import "./styles/theme.scss";
import ElementPlus from 'element-plus'
import App from './App.vue'

function bootstart() {
    const app = createApp(App);
    app.use(ElementPlus)
    app.mount('#app')
}

bootstart();