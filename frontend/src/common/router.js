import { createRouter, createWebHashHistory } from "vue-router";

// useRouter 提供的是路由器实例，用于导航和全局路由操作。
// useRoute 提供的是当前的路由信息，用于访问当前路由的各种细节。

let routes = [
    { path: "/login", component: () => import("../views/Login.vue") },
    { path: "/register", component: () => import("../views/Register.vue") },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export { router, routes }

