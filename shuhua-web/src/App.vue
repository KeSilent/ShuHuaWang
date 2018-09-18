<template>
  <div id="app" style="100%">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
html,body{
    height: 100%;
    margin: 0;
    padding: 0;
}
#app {
  background: #FFF none repeat scroll 0 0;
  font: 12px/20px "Hiragino Sans GB","Microsoft Yahei",arial,"Helvetica Neue",Helvetica,STHeiTi,sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  height: 100%;
}

ul,
ol {
  list-style: none outside none;
  padding: 0;
  margin: 0;
}
ul li a {
  color: #999; /* 设置文字颜色 */
  text-decoration: none; /* 去掉下划线 */
}

ul li {
  float: left; /* 往左浮动 */
}
</style>
