<template>
    <div align="left">
        <div>
            <el-input v-model="userId" placeholder="用户ID" style="width: 150px"></el-input>
        </div>
        <div>
            <el-input v-model="title" placeholder="邮件标题"></el-input>
        </div>
        <el-input
                type="textarea"
                :rows="10"
                placeholder="请输入邮件内容"
                v-model="context">
        </el-input>
        <el-switch v-model="hasItems" active-text="有附件" inactive-text="无附件">
        </el-switch>
        <el-table :data="items" style="width: 460px" v-show="hasItems">
            <el-table-column prop="itemId" label="道具ID" width="180">
            </el-table-column>
            <el-table-column prop="itemNum" label="道具数量" width="180">
            </el-table-column>
            <el-table-column label="操作" width="100">
                <template slot="header">
                    <el-button @click.native.prevent="addItemDialogVisible = true" type="primary"
                               icon="el-icon-plus">
                    </el-button>
                </template>
                <template slot-scope="scope">
                    <el-button @click.native.prevent="handleDelete(scope.$index, scope.raw)" type="danger">
                        移除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
        <div style="margin-top: 50px">
            <el-button @click="handleSendMsgClicked">发送</el-button>
        </div>
        <add-item-dialog :dialog-visible.sync="addItemDialogVisible" v-on:onAddItem="handleAddItem" />
    </div>
</template>

<script>
    import AddItemDialog from "../common/AddItemDialog";
    export default {
        name: "SendMailToOneUser",
        components: {AddItemDialog},
        comments : {AddItemDialog},
        data() {
            return {
                userId: '',
                title: '',
                context: '',
                hasItems: false,
                items: [{itemId: 1, itemNum: 10}, {itemId: 2, itemNum: 10}, {itemId: 2, itemNum: 10}],
                addItemDialogVisible : false
            }
        },
        methods: {
            handleDelete(index, row) {
                console.log(index, row);
                this.items.splice(index, 1);
            },
            handleSendMsgClicked(event) {
                console.log(event);
            },
            handleAddItem(event) {
                console.log(event);
                event.itemId = 100
            }
        }
    }
</script>