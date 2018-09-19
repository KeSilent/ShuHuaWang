<template>
  <div class="menu">
    <el-menu :default-active="activeIndex2"
    class="el-menu-demo"
    mode="horizontal"
    @select="handleSelect"
    background-color="#4b382f"
    text-color="#fff"
    active-text-color="#ffd04b">
        <el-menu-item index="0" id="navinside" class="hidden-xs-only" disabled>
          全部商品分类
        </el-menu-item>
        <el-menu-item  :xs="24" index="1">首页</el-menu-item>
        <el-menu-item  :xs="24" index="2">选画中心</el-menu-item>
        <el-menu-item index="3">艺术家</el-menu-item>
    </el-menu>
    <div class="card-opacity hidden-xs-only">
      <el-card class="box-card">
        <ul class="dd-inner">
          <li>
            <h2><a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchShuHua&amp;gc_parent_id=332&amp;cate_id=332" target="_blank">书法</a></h2>
            <div class="hot-item">
              <a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchShuHua&amp;gc_parent_id=332&amp;cate_id=616" target="_blank">草书</a>
              <a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchShuHua&amp;gc_parent_id=332&amp;cate_id=615" target="_blank">行书</a>
              <a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchShuHua&amp;gc_parent_id=332&amp;cate_id=613" target="_blank">隶书</a>
              <a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchShuHua&amp;gc_parent_id=332&amp;cate_id=612" target="_blank">篆书</a>
              <a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchShuHua&amp;gc_parent_id=332&amp;cate_id=614" target="_blank">楷书</a>
            </div>
          </li>
          <li>
            <h2><a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchShuHua&amp;gc_parent_id=182&amp;cate_id=182" target="_blank">国画</a></h2>
            <div class="hot-item">
              <a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchShuHua&amp;gc_parent_id=182&amp;cate_id=619" target="_blank">山水</a>
              <a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchShuHua&amp;gc_parent_id=182&amp;cate_id=617" target="_blank">人物</a>
              <a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchShuHua&amp;gc_parent_id=182&amp;cate_id=618" target="_blank">花鸟</a>
              <a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchShuHua&amp;gc_parent_id=182&amp;cate_id=636" target="_blank">动物</a>
              <a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchShuHua&amp;gc_parent_id=182&amp;cate_id=637" target="_blank">瓜果</a>
            </div>
          </li>
          <li>
            <h2><a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchShuHua&amp;gc_parent_id=183&amp;cate_id=183" target="_blank">油画</a></h2>
            <div class="hot-item">
              <a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchShuHua&amp;gc_parent_id=183&amp;cate_id=639" target="_blank">人物</a>
              <a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchShuHua&amp;gc_parent_id=183&amp;cate_id=640" target="_blank">风景</a>
              <a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchShuHua&amp;gc_parent_id=183&amp;cate_id=641" target="_blank">静物</a>
              <a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchShuHua&amp;gc_parent_id=183&amp;cate_id=642" target="_blank">花鸟</a>
              <a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchShuHua&amp;gc_parent_id=183&amp;cate_id=643" target="_blank">动物</a>
            </div>
          </li>
          <li class="lastu">
            <h2><a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchShuHua&amp;gc_parent_id=330&amp;cate_id=330" target="_blank">版画</a></h2>
            <div class="hot-item">
            </div>
          </li>
          <li class="lastd">
            <h2><a href="http://www.sctx.com/artist/index.php?act=artist_new&amp;op=searchArtist" target="_blank">名家</a></h2>
            <div class="hot-item">
              <a href="http://www.sctx.com/artist/index.php?act=artist_blog&amp;op=index&amp;aid=234" target="_blank">官春英</a>
              <a href="http://www.sctx.com/artist/index.php?act=artist_blog&amp;op=index&amp;aid=128" target="_blank">杨建军</a>
              <a href="http://www.sctx.com/artist/index.php?act=artist_blog&amp;op=index&amp;aid=154" target="_blank">吴冬梅</a>
              <a href="http://www.sctx.com/artist/index.php?act=artist_blog&amp;op=index&amp;aid=122" target="_blank">毕政</a>
              <a href="http://www.sctx.com/artist/index.php?act=artist_blog&amp;op=index&amp;aid=123" target="_blank">何满宗</a>
            </div>
          </li>
        </ul>
      </el-card>
    </div>
  </div>
</template>
<script>
export default {
  data () {
    return {
      activeIndex: '1',
      activeIndex2: '1'
    }
  },
  methods: {
    handleSelect (key, keyPath) {
      console.log(key, keyPath)
    }
  }
}
</script>
<style>
.menu {
  width: 100%;
  height: 48px;
  *text-align: center;
  margin: auto;
  background-color: #4b382f;
}
.menu .el-menu-demo {
  height: 49px;
  width: 93%;
  display: table;
  margin-left: auto;
  margin-right: auto;
}
.menu .el-menu-demo .el-menu-item {
  height: 48px;
  line-height: 50px;
  padding: 0px 40px;
  font-size: 16px;
}
.menu #navinside {
  background-color: #cd2e10 !important;
  padding-left: 55px;
  padding-right: 55px;
}
.menu #navinside a {
  background-color: #cd2e10;
  color: #fff !important;
}
.menu .el-menu-item.is-disabled {
  opacity:1;
  cursor: pointer;
  background: #cd2e10 !important;
}
.menu ul.dd-inner li{
  height: 40px;
  border-top: 1px #f0f0f0 solid;
  padding: 0;
  line-height: 0;
  transition: .3s;
  -moz-transition: .3s;
  -webkit-transition: .3s;
}
.menu ul.dd-inner li h2 {
    font-size: 16px;
    font-style: normal;
    font-weight: normal;
}
.menu ul.dd-inner li hot-item{
  height: 48px;
  overflow: hidden;
}

.menu .box-card {
  width: 205px;
  float: right;
  background: #fff;
}
.menu .card-opacity{
  position: absolute;
  z-index: 9;
  width: 251px;
  margin: 0 auto;
  height: 470px;
  overflow: hidden;
  background-color:rgba(0,0,0,0);
}
</style>
