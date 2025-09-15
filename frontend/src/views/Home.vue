<template>

</template>


<script setup lang="ts">
import { useRouter } from "vue-router";
import { onMounted } from "vue";
import { CheckAuth } from '../../wailsjs/go/app/App'

const $router = useRouter()
onMounted(async () => {
    const authed = await CheckAuth()
    if (!authed) {
        localStorage.removeItem("cookie")
        localStorage.removeItem("token")
        await $router.push("/login")
    }else{
        await $router.push({name:"Dashboard",query:{from:"home"}})
    }

}
)
</script>