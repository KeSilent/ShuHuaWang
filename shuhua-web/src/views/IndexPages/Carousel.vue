<template>
  <el-row>
    <el-col :span="19">
      <el-carousel :interval="5000" arrow="always">
        <el-carousel-item v-for="item in 4" :key="item">
          <h3>{{ item.ntitle }}</h3>
        </el-carousel-item>
      </el-carousel>
    </el-col>
    <el-col :span="5" class="consultation-panl">
        <h2 class="block">书画快讯
                 <el-button type="text">更多&gt;</el-button>
            </h2>
            <ul>
                <li><a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchshuhua&amp;keyword=毕政" target="_blank">【热点】书协主席-毕政 火热推荐</a></li>
                <li><a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchshuhua&amp;keyword=毕政" target="_blank">【热点】书协主席-毕政 火热推荐</a></li>
                <li><a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchshuhua&amp;keyword=毕政" target="_blank">【热点】书协主席-毕政 火热推荐</a></li>
            </ul>
         <h2 class="block">最新活动
                 <el-button type="text">更多&gt;</el-button>
            </h2>
            <ul>
                <li><a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchshuhua&amp;keyword=毕政" target="_blank">【热点】书协主席-毕政 火热推荐</a></li>
                <li><a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchshuhua&amp;keyword=毕政" target="_blank">【热点】书协主席-毕政 火热推荐</a></li>
                <li><a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchshuhua&amp;keyword=毕政" target="_blank">【热点】书协主席-毕政 火热推荐</a></li>
            </ul>
    </el-col>
  </el-row>
</template>
<script>
export default {
  data() {
    return {
      newlist: null
    };
  },
  mounted() {
    // this.$http.get(this.GLOBAL.serverSrc + this.GLOBAL.news + "?number=2").then(
    //   response => {
    //     console.log(response.data)
    //     this.newlist = response.data
    //   })
  }
};
</script>

<style>
.el-carousel__container {
  height: 470px;
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

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}

.consultation-panl h2 {
  height: 36px;
  border-bottom: 1px dotted #dcdcdc;
  color: #555;
  text-align: left;
  text-indent: 15px;
  font-size: 14px;
  line-height: 36px;
}
.consultation-panl .el-button {
  float: right;
  margin-right: 15px;
  color: #c9ad88;
  font-weight: 400;
  font-size: 9pt;
  font-family: "SimSun";
}

.consultation-panl ul {
  display: inline-block;
}
</style>
