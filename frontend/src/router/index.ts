import { createMemoryHistory, createRouter } from "vue-router";
import HomeView from "../views/HomeView.vue";

const router = createRouter({
  history: createMemoryHistory(),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/t/:infoHash/files",
      name: "files",
      component: () => import("@/views/FilesView.vue"),
    },
    {
      path: "/t/:infoHash/watch/:fileName",
      name: "watch",
      component: () => import("@/views/PlayerView.vue"),
    },
    {
      path: "/about",
      name: "about",
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/AboutView.vue"),
    },
  ],
});

export default router;
