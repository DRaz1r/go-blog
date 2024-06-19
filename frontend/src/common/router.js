import { createRouter, createWebHashHistory } from "vue-router";

// useRouter 提供的是路由器实例，用于导航和全局路由操作。
// useRoute 提供的是当前的路由信息，用于访问当前路由的各种细节。

let routes = [
    { path: "/login", component: () => import("../views/Login.vue") },
    { path: "/register", component: () => import("../views/Register.vue") },
    { path: "/", component: () => import("../views/MainFrame.vue") },
    { path: "/publish", component: () => import("../views/Publish.vue") },
    { path:"/myself", component: () => import("../views/Myself.vue") },
    { path:"/others", component: () => import("../views/Others.vue") },
    { path:"/detail", component: () => import("../views/Detail.vue") },
    { path:"/update", component: () => import("../views/Update.vue") },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export { router, routes }

