<template>
    <div>
        <el-row>
            <el-col :span="24">
                <carousel></carousel>
            </el-col>
        </el-row>
        <el-row>
            <el-col :span="24"><div class="grid-content bg-purple-dark">sdfsf</div></el-col>
        </el-row>
        <el-row>
            <el-col :span="24"><div class="grid-content bg-purple-dark">sdfsf</div></el-col>
        </el-row>
    </div>
</template>
<script>
import Carousel from '../IndexPages/Carousel'
export default{
  data () {
    return {
    }
  },
  components: {
    'carousel': Carousel
  }
}
</script>
