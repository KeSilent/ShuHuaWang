<template>
  <div class="menu">
    <el-menu :default-active="activeIndex2"
    class="el-menu-demo"
    mode="horizontal"
    @select="handleSelect"
    background-color="#4b382f"
    text-color="#fff"
    active-text-color="#ffd04b">
        <el-menu-item index="0" id="navinside" disabled>
          全部商品分类
        </el-menu-item>
        <el-menu-item  :xs="24" index="1">首页</el-menu-item>
        <el-menu-item  :xs="24" index="2">选画中心</el-menu-item>
        <el-menu-item index="3">艺术家</el-menu-item>
    </el-menu>
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
.el-menu-demo {
  height: 48px;
  width: 93%;
  display: table;
  margin-left: auto;
  margin-right: auto;
}
.el-menu-demo .el-menu-item {
  height: 48px;
  line-height: 50px;
  padding: 0px 40px;
  font-size: 16px;
}
#navinside {
  background-color: #cd2e10 !important;
  padding-left: 55px;
  padding-right: 55px;
}
#navinside a {
  background-color: #cd2e10;
  color: #fff !important;
}
.el-menu-item.is-disabled {
  opacity:1;
  cursor: pointer;
  background: #cd2e10 !important;
}
</style>
