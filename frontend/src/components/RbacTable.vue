<template>
  <b-container fluid v-bind:style="tablePadding">
    <h2> {{ title }} </h2>
    <table class="rbactable table table-striped table-bordered">
      <thead>
        <tr>
          <th v-for="field in fields">
            <span v-if="field.key == 'name'">
              {{ field.label }}
            </span>
            <span v-else>
              <div>{{ field.label }}</div>
            </span>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(item, index) in items">
          <td v-for="field in fields">
            <div v-if="field.key == 'name'">
              <div>{{ item.name }}
              </div>
              <div v-if="item.namespace">Namespace: {{item.namespace}}</div>
              <b-btn v-b-modal="modalId(index)">Subjects</b-btn>
              <b-modal :id="'modal' + title + index" :title="'Subjects for ' + item.name" size="lg" ok-only >
                <div v-for="subject in item.subjects">
                  {{ subject.kind }} - {{ subject.name }}
                </div>
              </b-modal>

            </div>
            <div v-else>
              <div v-for="action in item.objects[field.key]">
                <actions :action="action"></actions>
              </div>
            </div>
          </td>
        </tr>
      </tbody>
    </table>
  </b-container>
</template>

<script>
import Actions from './Actions.vue'

export default {
  name: 'RbacTable',
  props: ['rbactable', 'title'],
  data () {
    return {
      items: [],
      fields: [],
      longestField: 0,
      tablePadding: {},
    }
  },
  components: {
    Actions
  },
  methods: {
    setTableHeaders: function (headers) {
      var newheaders = []
      var vm = this;
      newheaders = headers.map(function(val) {
        if (val.length > vm.longestField) {
          vm.longestField = val.length
        }
        return {key: val, label:val}
      });
      newheaders.sort(function(a, b){
        if(a.key < b.key) return -1;
        if(a.key > b.key) return 1;
        return 0;
      });
      var padding = Math.sqrt(Math.pow(this.longestField, 2)/2)
      this.tablePadding = {
        paddingTop: padding + "ex",
      }
      return newheaders
    },
    modalId(i) {
      console.log(this.title)
      return 'modal' + this.title +  i;
    }
  },
  computed: {
  },
  watch: {
    rbactablePadding: function () {
      return {
        'padding-top': Math.sqrt(Math.pow(this.longestField, 2)/2)
      }
    },
    rbactable: function (val) {
      this.fields = this.setTableHeaders(val.objects)
      this.fields.unshift({key: 'name', label:'RoleName'})
      this.items = val.roles
    }
  }
}
</script>

<style>
.rbactable {
  border-collapse: collapse;
}

.rbactable td {
  width: 0px;
}

.table-bordered.rbactable {
  border: none;
}

.table-bordered.rbactable th {
  border: none;
}

.rbactable td {
  text-align: center;
}

.rbactable th {
  white-space: nowrap;
}
.rbactable th > span > div {
  -webkit-transform: translate(10px, -5px) rotate(315deg);
  -ms-transform: translate(10px, -5px) rotate(315deg);
  transform: translate(10px, -5px) rotate(315deg);
  width: 30px;
}

.rbactable th > div > span {
  padding: 5px 10px;
}

.rbactable th {
  padding: 0 10px;
}
.rbactable > tbody > tr > td > div {
  text-align: left;
}

</style>
