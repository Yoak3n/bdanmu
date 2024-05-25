<template>
  <div class="login-wrapper">
    <n-card v-if="text !== ''" class="qrcode" title="请使用哔哩哔哩App扫码登录">
      <n-qr-code :value="text"  :size="250" color="#f69"/>
    </n-card>
  </div>
</template>
<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { NCard, NQrCode,useMessage } from 'naive-ui';
import {useRouter} from  'vue-router'
import { useRoomStore,useAppStore } from '@/store';
import { EventsOnce} from '../../wailsjs/runtime'
import { LoginBilibili} from '../../wailsjs/go/app/App'
import { EventsEmit } from '../../wailsjs/runtime';

let text = ref("")
let logined = ref(false)
const $router = useRouter()
const $message = useMessage()
const appStore = useAppStore()
const roomStore = useRoomStore()
onMounted(async() => {
  EventsOnce("auth", auth)
  text.value = await LoginBilibili()
})

const start = async () => {
  if (roomStore.room_id !== 0) {
    EventsEmit ("change",Number(roomStore.room_id))
    $router.push({name: 'Dashboard', query: {from: 'login'}})
  }else{
    $router.push({name: 'Setting', query: {from: 'login'}})
    window.$message.error("未找到直播间",{keepAliveOnHover: true})
  }

}


const auth = (cookie: string,token:string) => {
  appStore.setCookie(cookie)
  appStore.setToken(token)
  logined.value = true
  $message.success("登录成功",{keepAliveOnHover: true})
  start()
}

</script>

<style scoped lang="less">
.login-wrapper {
  color: #fff;
  margin-top: 3rem;
  .qrcode {
    width: 63%;
    margin: 0 auto;
  }

}
</style>