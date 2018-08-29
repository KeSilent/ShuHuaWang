<template>
  <div class="art-nav">
    <el-row :gutter="40">
      <el-col :span="6" :xs="24" class="art-logo">
        <a href="#"><img src="http://www.sctx.com/shop/templates/default/images/art/artlogo1.png"/></a>
      </el-col>
      <el-col :span="12" :xs="24" hidden-xs-only class="search">
        <el-input placeholder="请输入内容">
          <el-button slot="append" icon="el-icon-search"></el-button>
        </el-input>
        <ul>
          <li><a href="#">国画</a></li>
          <li><a href="#">国画</a></li>
          <li><a href="#">国画</a></li>
        </ul>
      </el-col>
      <el-col :span="6" class="hidden-xs-only" >
        <a href="#">
          <img src="http://www.sctx.com/shop/templates/default/images/art/topadd.jpg"/>
        </a>
      </el-col>
    </el-row>
  </div>
</template>
<script>
export default {
  data () {
    return {
      formInline: {
        user: '',
        region: ''
      }
    }
  },
  methods: {
    onSubmit () {
      console.log('submit!')
    }
  }
}
</script>
<style>
.art-nav {
    padding: 20px 0;
}
.art-logo{
  padding-top: 9px;
}
.search ul li{
  float: left;
  padding: 6px 0 0 14px;
}
</style>
