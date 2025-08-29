import { type RouteRecordRaw } from "vue-router";

import HomeView from "../views/home.vue";
import AboutView from "../views/about.vue";

export const routes: RouteRecordRaw[] = [
    { path: "/", component: HomeView },
    { path: "/about", component: AboutView }
]