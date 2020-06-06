<template>
    <div>
      <button v-on:click="logon">logon</button>
    </div>
</template>

<script>
export default {
  methods: {
    logon () {
      if (this.$store.state.token) {
        this.$store.commit('del_token')
        alert('注销成功')
      }
    }
  }
}
</script>

<style scoped>

</style>
