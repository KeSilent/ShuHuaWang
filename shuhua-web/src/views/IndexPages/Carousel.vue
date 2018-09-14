<template>
  <el-row>
    <el-col :span="20">
      <el-carousel :interval="5000" arrow="always">
        <el-carousel-item v-for="item in 4" :key="item">
          <h3>{{ item.ntitle }}</h3>
        </el-carousel-item>
      </el-carousel>
    </el-col>
    <el-col :span="4">
      <el-card class="box-card">
        <div>
          {{newlist}}
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>
<script>
export default {
  data () {
    return {
      newlist: null
    }
  },
  mounted () {
    this.$http.get(this.GLOBAL.serverSrc + this.GLOBAL.news + "?number=2").then(
      response => {
        console.log(response.data)
        this.newlist = response.data
      })
  }
}
</script>

<style>
  .box-card{
    height: 300px;
  }
  .el-carousel__item h3 {
    color: #475669;
    font-size: 18px;
    opacity: 0.75;
    line-height: 300px;
    margin: 0;
  }

  .el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
  }

  .el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
  }
</style>
