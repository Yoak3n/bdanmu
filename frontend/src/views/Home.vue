<template>
  <div class="home-wrapper" >
    <div v-show = "danmus.length > 0" class="super-chat" >

    </div>
    <n-infinite-scroll  class="danmu-box">
      <div v-for="(danmu,index) in danmus"  :id="index ==  danmus.length-1  ?'bottom':'' " :key="danmu.message_id">
        <Transition name="fade" mode="out-in">
          <Danmubox :danmu="danmu" />
        </Transition>
      </div>

    </n-infinite-scroll>
  </div>
</template>
<script setup lang="ts">
import { ref,onMounted ,nextTick} from 'vue';
import {NInfiniteScroll} from 'naive-ui'
import {useRouter} from  'vue-router'

import Danmubox  from '../components/Danmu/index.vue'
import type {Danmu} from '../components/Danmu/danmu'
import type {User} from '../components/types'
import { EventsOn} from '../../wailsjs/runtime'
import {NeedLogin} from '../../wailsjs/go/app/App'
import { EventsEmit } from '../../wailsjs/runtime';

const $router = useRouter()
let danmus =  ref<Array<Danmu>>([])
onMounted(async() => {
  const cookie = localStorage.getItem("cookie")
  if (!cookie || cookie == "") {
    await $router.push("/login")
  }else{
    const need = await NeedLogin(cookie)
    if (need) {
      localStorage.removeItem("cookie")
      localStorage.removeItem("token")
      $router.push("/login")
    }else{
      EventsOn("danmu", pushDanmu)
      EventsOn("user", updateDanmu)
      EventsEmit("start")
    }
  }
})

const pushDanmu = (danmu:Danmu) => {
    if (danmus.value.length > 200) {
        danmus.value.shift()
    }
    danmus.value.push(danmu)
    nextTick(() => {
      const bottom = document.getElementById("bottom")
      bottom?.scrollIntoView({ behavior: "smooth", block: "center", inline: "end" });
    })
    
}

const updateDanmu = (user:User) => {
  danmus.value.forEach(danmu => {
      if (danmu.user.uid == user.uid) {
          danmu.user.avatar = user.avatar
          danmu.user.fans_count = user.fans_count
          danmu.user.sex = user.sex
      }
  })
    
}

</script>

<style scoped lang="less">

.home-wrapper {
  height: 100%;
  width: 100%;
  margin: 0 2rem;
  padding: 0 2rem;
  .danmu-box{
    height: 100%;
    overflow-y:scroll;
  }
}
</style>../components/types