import { RouteRecordRaw } from "vue-router";

export default <RouteRecordRaw[]> [
    {
        path: "/",
        name: "Home",
        component: () => import("../views/Home.vue"),
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