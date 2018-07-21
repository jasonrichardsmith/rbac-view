<template>
  <div id="app">
    <page-header></page-header>
    <b-tabs>
      <b-tab title="Cluster Roles">
        <rbac-table tabletype="clusterrole"></rbac-table>
      </b-tab>
      <b-tab title="Roles">
        <rbac-table tabletype="role"></rbac-table>
      </b-tab>
    </b-tabs>
  </div>
</template>

<script>

import RbacTable from './components/RbacTable.vue'
import PageHeader from './components/PageHeader.vue'
export default {
  name: 'app',
  components: {
    RbacTable,
    PageHeader
  }
}
</script>
