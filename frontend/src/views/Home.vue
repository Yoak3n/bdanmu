<template>

</template>


<script setup lang="ts">
import { useRouter } from "vue-router";
import { onMounted } from "vue";
import { NeedLogin } from '../../wailsjs/go/app/App'

const $router = useRouter()
onMounted(async () => {
    const cookie = localStorage.getItem("cookie")
    if (!cookie || cookie == "") {
        await $router.push("/login")
    } else {
        const need = await NeedLogin(cookie)
        if (need) {
            localStorage.removeItem("cookie")
            localStorage.removeItem("token")
            await $router.push("/login")
        }
        await $router.push({name:"Dashboard",query:{from:"home"}})
    }
}
)
</script>