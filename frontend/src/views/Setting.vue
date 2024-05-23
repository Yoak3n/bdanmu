<template>
  <div class="setting-wrapper">
    <n-config-provider :theme="darkTheme">
      <n-card title="直播间设置" style="width: 512px;color: white" :bordered="false" size="huge" role="dialog" embedded
        :segmented="{ content: true, footer: 'soft' }">

        <n-form style="color: white;">
          <n-form-item label="直播间ID" path="username">
            <n-input v-model:value="id" placeholder="请输入将连接的直播间ID" />
          </n-form-item>
          <n-form-item>
            <n-button @click="saveSettingAndRestart" type="primary" style="width: 50%;margin: 0 auto;">保存</n-button>
          </n-form-item>
        </n-form>
      </n-card>
    </n-config-provider>

  </div>


</template>

<script setup lang="ts">
import { onMounted,ref } from 'vue';
import { EventsEmit } from '../../wailsjs/runtime';
import { useRouter } from 'vue-router'
import { 
  NConfigProvider,
  NCard, 
  NForm, 
  NFormItem, 
  NInput,
  NButton, 
  darkTheme 
} from 'naive-ui';
const $router = useRouter()
let id = ref('')
onMounted(() => {
  id.value = localStorage.getItem("room_id")!
})
const saveSettingAndRestart = () => {
  EventsEmit("change", Number(id.value))
  $router.push({name: 'Dashboard', query: {from: 'setting'}})
}

</script>

<style scoped lang="less"></style>