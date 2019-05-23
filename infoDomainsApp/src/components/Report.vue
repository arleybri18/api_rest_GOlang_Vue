<template>
  <div class="report">
    <div class="container">
      <h1 class="text-light">Report of Queries</h1>
      <div v-if="items.length != 0">
        <!-- first information -->
        <div class="card bg-light m-3 col-7 mx-auto">
          <div class="card-header">
            <div class="table-responsive text-nowrap border-secondary">
              <div class="table table-striped">
                <thead class="bg-secondary text-white">
                  <tr>
                    <th>Date</th>
                    <th>Item</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="item in items">
                    <td>{{item.CreatedAt}}</td>
                    <td>{{item.Item}}</td>
                  </tr>
                </tbody>
              </div>
            </div>
          </div>
        </div>
        <!--   end information -->
      </div>
    </div>
  </div>


</template>

<script>
  export default {
    name: 'App',
    data() {
        return {

            // object items
            items: [],
        }
    },

    created(){
        // call endpoint
        var url = 'http://localhost:5000/report/';
        this.$http.get(url)
        .then(res=> this.items = res.data);

        console.log(url);

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
