<template>
    <b-container fluid  >
        <b-row>
            <b-col sm="3">
                <div class="search-wrapper">
                    <label> <font-awesome-icon icon="search" /></label>
                    <input type="text" v-model="searchRoles" placeholder="Search Roles.."/>
                </div>
            </b-col>
        </b-row>
        <b-row v-bind:style="tablePadding">
        <h2> {{ title }} </h2>
        <table class="rbactable table table-striped table-bordered">
            <thead>
            <tr>
                <th v-for="field in fields">
                    <span v-if="field.key === 'name'">
                      {{ field.label }}
                    </span>
                    <span v-else>
                      <label>{{ field.label }}</label>
                    </span>
                </th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="(item, index) in filteredRoles">
                <td v-for="field in fields">
                    <div v-if="field.key === 'name'">
                        <p>{{ item.name }}</p>
                        <p v-if="item.namespace">Namespace: {{item.namespace}}</p>
                        <b-btn v-b-modal="modalId(index)">Subjects</b-btn>
                        <b-modal :id="'modal' + title + index" :title="'Subjects for ' + item.name" size="lg" ok-only>
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
        </b-row>
    </b-container>
</template>

<script>
    import Actions from './Actions.vue'

    export default {
        name: 'RbacTable',
        props: ['rbactable', 'title'],
        data() {
            return {
                items: [],
                fields: [],
                longestField: 0,
                tablePadding: {},
                searchRoles: '',
            }
        },
        components: {
            Actions
        },
        computed:{
            filteredRoles() {
                return this.items.filter(item => {
                    return item.name.toLowerCase().includes(this.searchRoles.toLowerCase())
                })
            },
        },
        methods: {
            setTableHeaders: function (headers) {
                var newheaders = []
                var vm = this;
                newheaders = headers.map(function (val) {
                    if (val.length > vm.longestField) {
                        vm.longestField = val.length
                    }
                    return {key: val, label: val}
                });
                newheaders.sort(function (a, b) {
                    if (a.key < b.key) return -1;
                    if (a.key > b.key) return 1;
                    return 0;
                });
                var padding = Math.sqrt(Math.pow(this.longestField, 2) / 2)
                this.tablePadding = {
                    paddingTop: padding + "ex",
                }
                return newheaders
            },
            modalId(i) {
                return 'modal' + this.title + i;
            }
        },
        watch: {
            rbactablePadding: function () {
                return {
                    'padding-top': Math.sqrt(Math.pow(this.longestField, 2) / 2)
                }
            },
            rbactable: function (val) {
                this.fields = this.setTableHeaders(val.objects)
                this.fields.unshift({key: 'name', label: 'RoleName'})
                this.items = val.roles
            }
        }
    }
</script>

<style lang='scss'>

    .rbactable {
        border: none;
        border-collapse: collapse;
        td {
            width: 0;
            text-align: center;
            div {
                text-align: left;
            }
        }
        th {
            border: none;
            white-space: nowrap;
            padding: 0 10px;
            label {
                -webkit-transform: translate(10px, -5px) rotate(315deg);
                -ms-transform: translate(10px, -5px) rotate(315deg);
                transform: translate(10px, -5px) rotate(315deg);
                width: 30px;

            }
        }
    }


</style>

<style lang='scss' scoped>
    .row {
        margin: 0;
    }
    input {
        border: none;
        width: 80%;
        padding-left: 2%;
        &:focus {
            outline: none;
        }
    }
    .search-wrapper {
        border-bottom: 1px solid silver;
        padding-left: 1%;
        margin: 5% 0 -20% 0;
    }

</style>
