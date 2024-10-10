<script setup lang="ts">
import { ref, onMounted,nextTick } from 'vue'
import type { Danmu } from '../components/Danmu/danmu'
import DanmuBox from '../components/Danmu/index.vue'
import type{User} from '../components/types'
let $websocket = ref<WebSocket|null>(null)
let danmus = ref<Array<Danmu>>([])
onMounted(() => {
  initWebSocket()
})


const initWebSocket = () => {
  //初始化weosocket
  const wsuri = "ws://localhost:8080/ws"; //ws地址
  $websocket.value = new WebSocket(wsuri);
  $websocket.value.onopen = () => {
    console.log('连接成功')
  };
  $websocket.value.onerror = (err) => {
    console.log('连接出错', err)
  };
  $websocket.value.onmessage = (event) => {
    if (event.data == "success") {return}
    const object = JSON.parse(event.data) 
    switch (object.type) {
      case 2:
        const user:User = object.data
        updateDanmu(user)
        break;
      case 3:
        const danmu:Danmu = object.data
        console.log(danmu)
        pushDanmu(danmu)
        break;
    }
  };
  $websocket.value.onclose = () => {
    console.log('连接关闭')
  };
}
const pushDanmu = (danmu: Danmu) => {
  if (danmus.value.length > 200) {
    danmus.value.shift()
  }
  danmus.value.push(danmu)
  nextTick(() => {
    const bottom = document.getElementById("bottom")
    bottom?.scrollIntoView({ behavior: "smooth", block: "center", inline: "end" });
  })

}
const updateDanmu = (user: User) => {
  danmus.value.forEach(danmu => {
    if (danmu.user.uid == user.uid) {
      danmu.user.avatar = user.avatar
      danmu.user.fans_count = user.fans_count
      danmu.user.sex = user.sex
    }
  })

}
</script>

<template>
  <div  class="infinite-list" style="overflow: auto">
    <transition-group name="fade" tag="div">
      <DanmuBox 
      v-for="(danmu, index) in danmus" 
      :id="index == danmus.length - 1 ? 'bottom' : ''" 
      :key="danmu.message_id" 
      class="danmu-item"
       :danmu="danmu" />
  </transition-group>
  </div>
</template>

<style scoped lang="less">
.fade-move{
  transition: all .5s ease;
  }
  .fade-leave-active,
  .fade-enter-active {
    transition: all .5s ease;
  }
  .fade-leave-to,
  .fade-enter-from{
    opacity: 0;
    transform: translateX(-30px);
  }
  .fade-leave-active{
    position: absolute;
  }
  .infinite-list {
    height: 100%;
    width: 100%;
    overflow: hidden;
    scrollbar-width: none;
  
  }

</style>