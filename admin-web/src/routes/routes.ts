import { type RouteRecordRaw } from "vue-router";

import BaseLayout from "../layout/baseLayout.vue";

import HomeView from "../views/home.vue";
import AboutView from "../views/about.vue";

export const routes: RouteRecordRaw[] = [
  {
    path: "/",
    name: "base",
    component: BaseLayout,
    redirect: "/home",
    children: [
      {
        path: "home",
        name: "home",
        component: HomeView,
      },
      {
        path: "about",
        name: "about",
        component: AboutView,
      },
    ],
  },
];
