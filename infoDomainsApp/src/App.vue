<template>
  <div id="app">
<nav class="navbar navbar-expand-lg navbar-dark bg-dark">
  <router-link to="/" class="navbar-brand"> Home <span class="sr-only">(current)</span></router-link>
  <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNavAltMarkup" aria-controls="navbarNavAltMarkup" aria-expanded="false" aria-label="Toggle navigation">
    <span class="navbar-toggler-icon"></span>
  </button>
  <div class="collapse navbar-collapse" id="navbarNavAltMarkup">
    <div class="navbar-nav">
      <router-link to="/" class="nav-item nav-link active"> Info - Domain <span class="sr-only">(current)</span></router-link>
      <router-link to="/report" class="nav-item nav-link active"> Report <span class="sr-only">(current)</span></router-link>
    </div>
  </div>
</nav>
    <router-view></router-view>

  </div>
</template>

<script>
  import DomainInfo from './components/DomainInfo';
  import Report from './components/Report';
  export default {
      name: 'App',
      components: {
      DomainInfo,
      Report

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
  }
</style>
