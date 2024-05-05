<script setup lang="ts">
import {NFloatButton,NIcon} from 'naive-ui'
import {CloseOutline} from '@vicons/ionicons5'
import NaiveProvider from './components/NaiveProvider/index.vue'
import {Quit} from '../wailsjs/runtime'

const quit = ()=> {
  Quit()
}

</script>

<template>
  <naive-provider>
    <router-view/>
  </naive-provider>
  <div >
    <n-float-button :right="0" :bottom="0" shape="square" @click="quit" class="quit">
      <n-icon >
        <close-outline/>
      </n-icon>
    </n-float-button>
  </div>
</template>

<style scoped lang="less">
.quit{
  visibility: hidden;

}
.quit:hover{
  visibility: visible;
}

</style>