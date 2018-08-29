<template>
  <div id="app">
    <el-container style="height:100%">
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
          <mainpage></mainpage>
        </el-main>
        <el-footer>Footer</el-footer>
    </el-container>
  </div>
</template>
<script>
import HeaderShortcut from '../Headers/Shortcut'
import HeaderLogSelect from '../Headers/LogSelect'
import HeaderMenu from '../Headers/Menu'
import MainPage from '../IndexPages/Main'
export default{
  data () {
    return {
    }
  },
  components: {
    'headershortcut': HeaderShortcut,
    'headerlogselect': HeaderLogSelect,
    'headermenu': HeaderMenu,
    'mainpage': MainPage
  }
}
</script>

<style>
.el-main {
  padding: 10px;
}
.el-header{
  padding: 0;
}
.topbar {
  background: #FAFAFA !important;
  border-bottom: 1px #dcdedd solid;
}
</style>
