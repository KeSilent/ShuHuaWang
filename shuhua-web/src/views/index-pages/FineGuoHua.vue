<template>
    <div class="fine-card">
        <el-card shadow="never" class="box-card">
            <div slot="header" class="clearfix">
                <h2 class="block">精品国画
                    <el-button type="text">更多&gt;&gt;</el-button>
                </h2>
            </div>
            <el-row :gutter="20">
                <el-col :xs="24" :sm="12" :md="5" :lg="4" :xl="4" v-for="o in 5" :key="o" class="card-list">
                    <el-card shadow="hover">
                    <img src="../../assets/demo.jpg" class="image">
                    <div>
                        好吃的汉堡
                    </div>
                    <em><i>{{ currentDate }}</i></em>
                    </el-card>
                </el-col>
            </el-row>
        </el-card>
    </div>
</template>

<style>
.fine-card {
  padding-top: 0px;
  padding-left: 5px;
  padding-right: 5px;
}
.fine-card .box-card {
  border: 0px;
  border-bottom: 1px solid #ebeef5;
}
@media only screen and (min-width: 1200px){
  .fine-card .box-card .el-col-xl-4 {
    width: 20%;
  }
}
.fine-card ul {
  display: inline-block;
}
.fine-card .block {
  margin-left: 20px;
}
.fine-card .block .el-button {
  float: right;
  margin-right: 15px;
  color: #c9ad88;
  font-weight: 400;
  font-size: 9pt;
  font-family: "SimSun";
}

.fine-card .card-list {
  margin-bottom: 10px;
}
.fine-card .text {
  font-size: 14px;
  text-align: left;
}

.fine-card .item {
  margin-bottom: 18px;
}
.fine-card .item li {
  padding: 10px;
}

.fine-card em i {
  bottom: 17px;
  color: #cd2e10;
  font-size: 18px;
  font-family: Calibri;
}

.fine-card .clearfix:before,
.fine-card .clearfix:after {
  display: table;
  content: "";
}
.fine-card .clearfix:after {
  clear: both;
}

.fine-card .box-card {
  width: 100%;
}

.fine-card .price {
  font-size: 13px;
  color: #999;
}

.fine-card .image {
  width: 100%;
  display: block;
}

.fine-card .el-card__header {
  padding: 0px;
}
</style>
<script>
export default {
  data () {
    return {
      currentDate: '¥2001'
    }
  }
}
</script>
