<template>
  <div class="info_domain">
    <div class="container">
      <h1 class="text-light">Get Info of Your Domain</h1>

      <!-- form search domain -->
      <div class="card col-6 mx-auto m-3">
        <div class="card-header">
          <h3> <span class="text-scondary" style="font-weight: 300;color: #333;">Type your domain</span></h3>
        </div>
        <div class="card-body">

          <form v-on:submit.prevent="getData" class="form-inline">
            <div class="form-group">
              <input v-model="domain" type="text" class="form-control">
            </div>
            <div class="form-group m-1">
              <input type="submit" value="search" class="btn btn-default text-white" style="background-color: #f0ad4e">
            </div>
          </form>

        </div>
      </div>
      <!-- form search domain -->

      <div v-if="doms.length != 0">
        <!-- first information -->
        <div class="card  col-10 bg-light m-3 mx-auto">
          <div class="card-header">
            <h3 class="card-title rounded-pill bg-default text-white" style="background-color: #3498db;">Information
              about {{domain}}</h3>
            <div class="table-responsive text-nowrap border-secondary">
              <div class="table table-striped">
                <thead class="bg-secondary text-white">
                  <tr>
                    <th>Domain</th>
                    <th>Change</th>
                    <th>Previous SSL Grade</th>
                    <th>Logo</th>
                    <th>Title</th>
                    <th>Down</th>
                  </tr>
                </thead>
                <tbody>
                  <!-- iterate over domains -->
                  <tr v-for="dom in doms">
                    <td>{{dom.Domain}}</td>
                    <td>{{dom.Servers_changed}}</td>
                    <td>{{dom.Previous_ssl_grade}}</td>
                    <td><img :src="dom.Logo" width="40" height="65"/></td>
                    <td>{{dom.Title}}</td>
                    <td>{{dom.Is_down}}</td>
                  </tr>
                </tbody>
              </div>
            </div>
          </div>
        </div>
        <!--   end first information -->
        <!-- servers -->
        <div class="card col-8 bg-light m-3 mx-auto">
          <div class="card-header">
            <h3 class="card-title rounded-pill bg-default text-white" style="background-color: #3498db;">Servers
              {{domain}}</h3>

            <div class="table table-striped">
              <thead class="bg-secondary text-white">
                <tr>
                  <th>Address</th>
                  <th>SSL Grade</th>
                  <th>Country</th>
                  <th>Owner</th>
                </tr>
              </thead>
              <!-- iterate over info servers -->
              <tbody v-for="dom in doms">
                <tr v-for="ser in dom.Servers">
                  <td>{{ser.serverName}}</td>
                  <td>{{ser.grade}}</td>
                  <td>{{ser.country}}</td>
                  <td>{{ser.owner}}</td>
                </tr>
              </tbody>
            </div>
          </div>
        </div>
        <!-- end servers -->
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'App',
    data() {
        return {

            // object doms
            doms: [],
        }
    },

    methods: {
      getData: function () {
        // call endpoint
        var url = 'http://localhost:5000/domain/'+this.domain;
        this.$http.get(url)
        .then(res=> this.doms = res.body);

      }
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
