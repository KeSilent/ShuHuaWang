<template>
    <div class="fine-card">
        <el-card shadow="never" class="box-card">
            <div slot="header" class="clearfix">
                <h2 class="block">精品推荐
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
<script>
export default {
  data () {
    return {
      currentDate: '¥2001'
    }
  }
}
</script>
