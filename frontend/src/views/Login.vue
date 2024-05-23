<template>
  <div class="login-wrapper">
    <n-form>
      <n-form-item label="请使用哔哩哔哩App扫码登录">
        <n-card v-if="text !== ''" class="qrcode">
          <n-qr-code :value="text"  :size="250"/>
        </n-card>
      </n-form-item>

      <n-form-item >
        <n-button :disabled="!logined" class="login-button" type="primary" shape="square" @click="start">
          启动
        </n-button>
      </n-form-item>
    </n-form>

  </div>
</template>
<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { NCard, NButton, NQrCode, NForm, NFormItem,useMessage } from 'naive-ui';
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
  .qrcode {
    width: 63%;
    margin: 0 auto;
  }

  .login-button{
    width: 50%;
    margin: 0 auto;
  }

}
</style>