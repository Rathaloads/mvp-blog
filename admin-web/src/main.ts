import { createApp } from 'vue';
import ElementPlus from 'element-plus'
import "./styles/theme.scss";
import "./styles/index.scss";
import { router } from "./routes/router";
import { createPinia } from "pinia";
import App from './App.vue';

function bootstart() {
    const pinia = createPinia();
    const app = createApp(App);
    
    app.use(router);
    app.use(pinia);
    app.use(ElementPlus);
    app.mount('#app');
}

bootstart();