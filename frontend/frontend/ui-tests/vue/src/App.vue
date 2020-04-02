<template>
  <div v-if="isDrizzleInitialized" id="app">
    <div class="logo">
      <img alt="Vue logo" src="./assets/logo.png" />
    </div>

    <div class="left">
      <h2>Insertion</h2>
      <Covid />
    </div>

    <div class="right">
      <h2>Consultation</h2>
    </div>

  </div>

  <div v-else>Loading...</div>
</template>

<script>
import Covid from './Covid'
import { mapGetters } from 'vuex'

export default {
  name: 'app',
  components: {
    Covid
  },

  computed: mapGetters('drizzle', ['isDrizzleInitialized'])
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  width: 100%;
}

img {
  width: 100%;
}

.logo {
  width: 100%;
  margin: auto;
}

.logo img {
  width: 20%;
  margin: auto;
}

.left, .right {
  display: inline-block;
  vertical-align: text-top;
  margin: auto 1%;
  padding: 1%;
  border: 2px solid #e67e22;
  background-color: #34495e;
  color: #e67e22;
}

.left {
  width: 25%;
}

.right {
  width: 65%;
}

input {
  width: 90%;
  padding: 2% 1%;
  margin: 1% auto;
}

button {
  background-color: #e67e22;
  width: 30%;
  padding: 2%;
  border-style: none;
  transition: all .7s ease;
  margin-top: 2%;
  cursor: pointer;
  color: white;
}

button:hover {
  background-color: #965216;
}
</style>
