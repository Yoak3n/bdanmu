import { createRouter, createWebHashHistory } from "vue-router";
import routes from './routes'
const router =  createRouter({ 
    history: createWebHashHistory(),
    routes,
})

router.beforeEach((to, _, next) => {
    if (to.meta.requireAuth) {
        // 判断该路由是否需要登录权限
        const token = localStorage.getItem('token')
        const cookie = localStorage.getItem('cookie')
        if (token == '' || cookie == '') {
            next({
                path: '/login',
                query: {
                    redirect: to.fullPath
                }
            })
        }else{
            next()
        }
    }else{
        next()
    }
})


export default router