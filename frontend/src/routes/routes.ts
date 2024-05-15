import { RouteRecordRaw } from "vue-router";

export default <RouteRecordRaw[]> [
    {
        path: "/",
        name: "Home",
        component: () => import("../views/Home.vue"),
        meta: {
            keepAlive: true, // 需要被缓存
        }
    },{
        path: "/login",
        name: "Login",
        component: () => import("../views/Login.vue"),
    },{
        path:"/setting",
        name:"Setting",
        component:()=>import("../views/Setting.vue")
    }

]