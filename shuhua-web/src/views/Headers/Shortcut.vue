<template>
  <div class="wrapper">
    <ul class="fl welcome-txt">
      <li>您好，欢迎来到</li>
      <li><a href="#">书画326</a></li>
      <li><a class="red" href="#">注册</a></li>
      <li><a href="#">登陆</a></li>
    </ul>
    <ul class="fr hidden-xs-only">
      <li>
          <a>收藏</a>
      </li>
      <li class="spacer"></li>
      <li>
        <a>
          关注微信
        </a>
      </li>
    </ul>
  </div>
</template>
<style>
.wrapper .fl {
  float: left;
  margin-left: 10px;
}
.wrapper .fr{
  float:right;
  margin-right: 10px;
}

.wrapper ul li {
  line-height: 30px;
}
.wrapper li {
  margin: 0 4px;
}
.wrapper {
  position: relative;
  clear: both;
  margin: auto;
  color: #999; /* 设置文字颜色 */
  width: 100%;
}

.wrapper .red {
  color: #e33333;
  text-decoration: none;
}

.wrapper .topbar .spacer {
  overflow: hidden;
  margin: 11px 5px 0;
  width: 1px;
  height: 10px;
  background-color: #ccc;
}
</style>
