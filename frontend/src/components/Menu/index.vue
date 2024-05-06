<template>
  <div class="menu-wrapper">
    <n-menu
      :root-indent="9"
      :options="menuOptions">

    </n-menu>
  </div>

</template>
<script setup lang="ts">
import {NMenu} from "naive-ui";
import {menuOptions} from './index'
</script>