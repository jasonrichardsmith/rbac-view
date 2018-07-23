<template>
  <div id="app">
    <page-header></page-header>
    <b-tabs>
      <b-tab title="Cluster Roles">
          <rbac-table :rbactable="ClusterRoles"></rbac-table>
      </b-tab>
      <b-tab title="Roles">
          <rbac-table :rbactable="ClusterRoles"></rbac-table>
      </b-tab>
    </b-tabs>
  </div>
</template>

<script>

import RbacTable from './components/RbacTable.vue'
import PageHeader from './components/PageHeader.vue'
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
    PageHeader
  },
  methods: {
    getTableItems() {
      axios({ method: "GET", "url": "/test_data/test.json",
        headers: {
          'Content-Type': 'application/json',
          'Cache-Control' : 'no-cache'}
      }).then(result => {
          var items
          items = result.data
          this.ClusterRoles = items.ClusterRoles
          this.Roles = items.ClusterRoles
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
