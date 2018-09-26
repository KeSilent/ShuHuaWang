<template>
    <div>
        <el-card class="box-card">
            <div slot="header" class="clearfix">
            <span>选画中心</span>
            </div>
            <div v-for="o in 4" :key="o" class="text item">
            {{'列表内容 ' + o }}
            </div>
        </el-card>
    </div>
</template>
<script>
import bus from '@/config/event-bus.js'

export default {
  created: function () {
    bus.$emit('menuIsShowEvent', false)
  }
}
</script>
