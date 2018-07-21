<template>
  <b-container fluid>
    <!-- User Interface controls -->
    <b-row>
      <b-col md="6" class="my-1">
        <b-form-group horizontal label="Filter" class="mb-0">
          <b-input-group>
            <b-form-input v-model="filter" placeholder="Type to Search" />
            <b-input-group-append>
              <b-btn :disabled="!filter" @click="filter = ''">Clear</b-btn>
            </b-input-group-append>
          </b-input-group>
        </b-form-group>
      </b-col>
    </b-row>
    <!-- Main table element -->
    <b-table show-empty
             stacked="md"
             :items="items"
             :fields="fields"
             :filter="filter"
             @filtered="onFiltered"
    >
      <template slot="rolename" slot-scope="row">{{row.value.first}} {{row.value.last}}</template>
      <template slot="isActive" slot-scope="row">{{row.value?'Yes :)':'No :('}}</template>
      <template slot="actions" slot-scope="row">
        <!-- We use @click.stop here to prevent a 'row-clicked' event from also happening -->
        <b-button size="sm" @click.stop="info(row.item, row.index, $event.target)" class="mr-1">
          Info modal
        </b-button>
        <b-button size="sm" @click.stop="row.toggleDetails">
          {{ row.detailsShowing ? 'Hide' : 'Show' }} Details
        </b-button>
      </template>
      <template slot="row-details" slot-scope="row">
        <b-card>
          <ul>
            <li v-for="(value, key) in row.item" :key="key">{{ key }}: {{ value}}</li>
          </ul>
        </b-card>
      </template>
        <template slot="HEAD_*" slot-scope="data">
        <!-- A custom formatted footer cell  for field 'name' -->
        <div>{{data.label}}</div>
      </template>
    </b-table>

    <b-row>
      <b-col md="6" class="my-1">
        <b-pagination :total-rows="totalRows" :per-page="perPage" v-model="currentPage" class="my-0" />
      </b-col>
    </b-row>

    <!-- Info modal -->
    <b-modal id="modalInfo" @hide="resetModal" :title="modalInfo.title" ok-only>
      <pre>{{ modalInfo.content }}</pre>
    </b-modal>

  </b-container>
</template>

<script>
import axios from 'axios';
export default {
  name: 'RbacTable',
  props: ['tabletype'],
  data () {
    return {
      items: [
  { isActive: true, age: 40, rolename: { first: 'Dickerson', last: 'Macdonald' } },
  { isActive: false, age: 21, rolename: { first: 'Larsen', last: 'Shaw' } }
	],
      fields: [
        { key: 'rolename', label: ' '},
        { key: 'age', label: '<div>Person age</div>'},
        { key: 'isActive', label: '<div>is Active</div>' },
        { key: 'actions', label: '<div>Actions</div>' }
      ],
      totalRows: 2,
      filter: null,
      modalInfo: { title: '', content: '' }
    }
  },
  computed: {
    sortOptions () {
      // Create an options list from our fields
      return this.fields
        .filter(f => f.sortable)
        .map(f => { return { text: f.label, value: f.key } })
    }
  },
  methods: {
    info (item, index, button) {
      this.modalInfo.title = `Row index: ${index}`
      this.modalInfo.content = JSON.stringify(item, null, 2)
      this.$root.$emit('bv::show::modal', 'modalInfo', button)
    },
    resetModal () {
      this.modalInfo.title = ''
      this.modalInfo.content = ''
    },
    onFiltered (filteredItems) {
      // Trigger pagination to update the number of buttons/pages due to filtering
      this.totalRows = filteredItems.length
      this.currentPage = 1
    },
    getTableItems() {
      axios({ method: "GET", "url": "/tables/" + this.tabletype + "/items",
        headers: {
          'Content-Type': 'application/json',
          'Cache-Control' : 'no-cache'}
      }).then(result => {
          this.items = result.data
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

<style>
.b-table {
  border-collapse: collapse;
}

.b-table td {
  width: 0px;
}

.b-table th {
  padding: 5px 10px;
  border-top: none;
}

.b-table td {
  text-align: center;
  padding: 10px 5px;
}

.b-table th {
  height: 140px;
  white-space: nowrap;
}
.b-table th > div {
  -webkit-transform: translate(10px, -5px) rotate(315deg);
  -ms-transform: translate(10px, -5px) rotate(315deg);
  transform: translate(10px, -5px) rotate(315deg);
  width: 30px;
}

.b-table th > div > span {
  padding: 5px 10px;
}

.b-table th {
  padding: 0 10px;
}
.b-table > tbody > tr > td > div {
  text-align: left;
}

</style>
