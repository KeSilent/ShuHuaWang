<template>
  <div id="app" class="layout" style="position:absolute; left:0; top:0; width:100%;height:100%;">
    <el-container>
      <el-header class="topbar" style="height:30px">
        <headershortcut></headershortcut>
      </el-header>
      <el-header style="height:auto">
        <headerlogselect></headerlogselect>
      </el-header>
      <el-header style="height:48px">
        <headermenu></headermenu>
      </el-header>
      <el-main>
        <router-view></router-view>
      </el-main>
      <el-footer>
        <footervue></footervue>
      </el-footer>
    </el-container>
  </div>
</template>
<script>
import HeaderShortcut from '../headers/Shortcut'
import HeaderLogSelect from '../headers/LogSelect'
import HeaderMenu from '../headers/Menu'
import FooterVue from './Footers'
export default {
  data () {
    return {
    }
  },
  components: {
    'headershortcut': HeaderShortcut,
    'headerlogselect': HeaderLogSelect,
    'headermenu': HeaderMenu,
    'footervue': FooterVue
  }
}
</script>

<style>
.layout .el-main {
  padding: 1px 0px 0px 0px;
}
.layout .el-header {
  padding: 0;
}
.layout .topbar {
  background: #fafafa !important;
  border-bottom: 1px #dcdedd solid;
}
</style>
