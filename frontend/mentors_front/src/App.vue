<template>
  <div id="app">
    <MentorsFront />
  </div>
</template>

<script>
import MentorsFront from './components/MentorsFront.vue'

export default {
  name: 'App',
  components: {
    MentorsFront
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #FFf;
}
body{
  background-color: #27292d;
}
</style>
