<template>
  <div id="app">
    <b-alert :show="showLoading" variant="primary"><strong>Roles are being retrieved from Kubernetes</strong></b-alert>
    <b-alert :show="loadingError" variant="danger"><strong>There was an error retrieving roles from Kubernetes</strong></b-alert>
    <nav class="navbar navbar-default">
      <div class="container-fluid">
        <div class="navbar-header">
          <span class="navbar-brand">
              <page-header></page-header>
          </span>
          <actions listactions="true"></actions>
        </div>
      </div>
    </nav>
    <b-tabs>
      <b-tab title="Cluster Roles">
          <rbac-table :rbactable="ClusterRoles" title="ClusterRoles"></rbac-table>
      </b-tab>
      <b-tab title="Roles">
          <rbac-table :rbactable="Roles" title="Roles"></rbac-table>
      </b-tab>
    </b-tabs>
  </div>
</template>

<script>

import RbacTable from './components/RbacTable.vue'
import PageHeader from './components/PageHeader.vue'
import Actions from './components/Actions.vue'
import axios from 'axios';
export default {
  name: 'app',
  data: function () {
    return {
      ClusterRoles: [] ,
      Rolels: [],
      showLoading: false,
      loadingError: false
    };

  },
  components: {
    RbacTable,
    PageHeader,
    Actions
  },
  methods: {
    getTableItems() {
      this.showLoading=true;
      var vm = this;
      axios({ method: "GET", "url": "/allroles.json",
        headers: {
          'Content-Type': 'application/json',
          'Cache-Control' : 'no-cache'}
      }).then(result => {
          var items
          items = result.data
          this.ClusterRoles = items.clusterRoles
          this.Roles = items.roles
          this.showLoading=false;
      }).catch (error => {
          this.showLoading=false;
          this.loadingError=true;
          console.log(error);
      });
    }
  },
  mounted: function() {
    this.getTableItems();
  }
}
</script>
<style lang="scss">
  .navbar-header {
    display: flex;
    justify-content: center;
    align-items: center;
  }

  </style>
