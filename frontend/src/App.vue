<template>
  <div id="app">
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
          <rbac-table :rbactable="ClusterRoles"></rbac-table>
      </b-tab>
      <b-tab title="Roles">
          <rbac-table :rbactable="Roles"></rbac-table>
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
      Rolels: []
    };

  },
  components: {
    RbacTable,
    PageHeader,
    Actions
  },
  methods: {
    getTableItems() {
      axios({ method: "GET", "url": "/allroles.json",
        headers: {
          'Content-Type': 'application/json',
          'Cache-Control' : 'no-cache'}
      }).then(result => {
          var items
          items = result.data
          this.ClusterRoles = items.clusterRoles
          this.Roles = items.roles
          console.log(this.items)
      }).catch (error => {
          console.log(error);
      });
    }
  },
  mounted: function() {
    this.getTableItems();
  }
}
</script>
