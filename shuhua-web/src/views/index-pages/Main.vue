<template>
    <div>
        <el-row>
            <el-col :span="24">
                <carousel></carousel>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="24">
                <fine></fine>
            </el-col>
        </el-row>
        <el-row>
            <fineguohua></fineguohua>
        </el-row>
        <el-row>
            <fineshufa></fineshufa>
        </el-row>
    </div>
</template>
<script>
import bus from '@/config/event-bus.js'
import CarouselVue from './Carousel'
import FineVue from './Fine'
import FineShuFaVue from './FineShuFa'
import FineGuoHuaVue from './FineGuoHua.vue'
export default {
  data () {
    return {
    }
  },
  components: {
    'carousel': CarouselVue,
    'fine': FineVue,
    'fineshufa': FineShuFaVue,
    'fineguohua': FineGuoHuaVue
  },
  created: function () {
    bus.$emit('menuIsShowEvent', true)
  }
}
</script>
