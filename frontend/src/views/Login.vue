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

import { EventsOnce} from '../../wailsjs/runtime'
import { LoginBilibili} from '../../wailsjs/go/app/App'
import { EventsEmit } from '../../wailsjs/runtime';

let text = ref("")
let logined = ref(false)
const $router = useRouter()
const $message = useMessage()
onMounted(async() => {
  EventsOnce("auth", auth)
  text.value = await LoginBilibili()
})

const start = async () => {
  EventsEmit ("start")
  $router.push({name: 'Dashboard', query: {from: 'login'}})
}


const auth = (cookie: string,token:string) => {
  localStorage.setItem("cookie", cookie)
  localStorage.setItem("token", token)
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