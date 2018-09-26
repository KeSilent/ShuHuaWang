<template>
  <div>
    Artist
  </div>
</template>
<script>
import bus from '@/config/event-bus.js'
export default {
  data () {
    return {}
  },
  created: function () {
    bus.$emit('menuIsShowEvent', false)
  }
}
</script>
