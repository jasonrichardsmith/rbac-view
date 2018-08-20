<template>
    <div v-if="listactions" class="header-wrap">
        <p v-for="(value, key) in actionTypes" class="badges-wrapper">
            <span class="action-badge" v-bind:style="actionColors(key)">  {{ actionLabels(key) }} </span>  - {{ key }}
        </p>
    </div>
    <div v-else class="table-badges">
        <p class="action-badge" v-bind:style="actionColor" v-b-tooltip
           :title="action"><span>{{actionLabel}}</span>
        </p>
    </div>
</template>

<script>
    export default {
        name: 'Actions',
        props: ['action', 'listactions'],
        data() {
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
                if (typeof this.actionTypes[this.action]) {
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
                if (typeof this.actionTypes[this.action]) {
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

<style lang="scss">

    .header-wrap {
        display: flex;
        .badges-wrapper {
            display: flex;
            margin: 0 0.43rem;
            justify-content: center;
            align-items:center;
        }
    }

    .action-badge {
        font-size: 60%;
        font-weight: 700;
        text-align: center;
        white-space: nowrap;
        vertical-align: baseline;
        border-radius: .50rem;
        min-width: 1.5rem;
        display: flex;
        margin:0  0.2rem 0 auto;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        height: 1rem;
    }
    .table-badges {
        .action-badge {
            margin-top: 0.2rem;
            max-width: 1.2rem;
        }
    }

</style>

