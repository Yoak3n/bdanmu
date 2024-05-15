<template>
  <div class="setting-wrapper">
    <n-card 
    title="直播间设置" 
    style="width: 512px;color: white"
    :bordered="false" 
    size="huge" 
    role="dialog" 
    embedded 
    :segmented="{ content: true, footer: 'soft' }"
    >

      <n-form v-model="model" style="color: white;">
        <n-form-item label="直播间ID" path="username">
          <n-input placeholder="请输入用户名" />
        </n-form-item>
        <n-form-item label="密码" path="password">
          <n-input type="password" placeholder="请输入密码" />
        </n-form-item>
      </n-form>
    </n-card>
  </div>


</template>

<script setup lang="ts">
import { ref } from 'vue';
import { NCard, NForm, NFormItem, NInput } from 'naive-ui';

let model = ref({})


</script>

<style scoped lang="less"></style>