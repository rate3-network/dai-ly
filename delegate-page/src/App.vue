<template>
  <div id="app">
    <Delegate msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
import Delegate from './components/Delegate.vue'

export default {
  name: 'app',
  components: {
    Delegate
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
