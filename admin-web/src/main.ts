import { createApp } from 'vue';
import "./styles/index.scss";
import "./styles/theme.scss";
import ElementPlus from 'element-plus'
import { router } from "./routes/router"
import App from './App.vue'

function bootstart() {
    const app = createApp(App);
    app.use(router);
    app.use(ElementPlus);
    app.mount('#app');
}

bootstart();