<template>
  <span v-if="listactions">
    <span v-for="(value, key) in actionTypes">
      <span class="badge action-badge navbar-text" v-bind:style="actionColors(key)">
        {{ actionLabels(key) }}
      </span>
      -
      <span>
        {{ key }}
      </span>
    </span>
  </span>
  <div v-else class="badge action-badge badge-pill" v-bind:style="actionColor" v-b-tooltip
             :title="action">{{actionLabel}}</div>
</template>

<script>
export default {
  name: 'Actions',
  props: ['action', 'listactions'],
  data () {
    return {
      actionTypes: {
        create: {color: "green", textcolor: "white", label: "C"},
        delete: {color: "red", textcolor: "white", label: "D"},
        get: {color: "yellow", textcolor: "black", label: "G"},
        list: {color: "blue", textcolor: "white", label: "L"},
        watch: {color: "brown", textcolor: "white", label: "W"},
        patch: {color: "gray", textcolor: "white", label: "P"},
        update: {color: "pink", textcolor: "black", label: "U"},
        deletecollection: {color: "black", textcolor: "white", label: "DC"},
        '*': {color: "orange", textcolor: "white", label: "*"}
      }
    }
  },
  computed: {
    actionColor: function () {
      if (typeof this.actionTypes[this.action] != "undefined") {
        return {
          color: this.actionTypes[this.action].textcolor,
          'background-color': this.actionTypes[this.action].color
        }
      } else {
        console.log('no color set for ' + this.action)
        return {
          color: "white",
          'background-color': "black"
        }
      }
    },
    actionLabel: function () {
      if (typeof this.actionTypes[this.action] != "undefined") {
        return this.actionTypes[this.action].label
      } else {
        console.log('no label set for ' + this.action)
        return this.action
      }
    }
  },
  methods: {
    actionColors: function (action) {
      return {
        color: this.actionTypes[action].textcolor,
        'background-color': this.actionTypes[action].color
      }
    },
    actionLabels: function (action) {
      return this.actionTypes[action].label
    }
  }
}
</script>

