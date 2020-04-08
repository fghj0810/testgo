<template>
    <div align="left">
        <div>
            <el-input v-model="userId" placeholder="用户ID" style="width: 150px" type="text"></el-input>
            <span v-show="checkUserId">用户ID不合法</span>
        </div>
        <div>
            <el-input v-model="title" placeholder="邮件标题" type="text"></el-input>
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
                    <el-button @click.native.prevent="handleDelete(scope.$index)" type="danger">
                        移除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
        <div style="margin-top: 50px">
            <el-button @click="handleSendMsgClicked" type="primary">发送</el-button>
            <el-popover
                    placement="top"
                    width="160"
                    v-model="visible">
                <p>确定发送吗？</p>
                <div style="text-align: right; margin: 0">
                    <el-button size="mini" type="text" @click="visible = false">取消</el-button>
                    <el-button type="primary" size="mini" @click="visible = false">确定</el-button>
                </div>
                <el-button slot="reference">删除</el-button>
            </el-popover>
        </div>
        <add-item-dialog :dialog-visible.sync="addItemDialogVisible" v-on:onAddItem="handleAddItem"/>
    </div>
</template>

<script>
    import AddItemDialog from "../common/AddItemDialog";

    export default {
        name: "SendMailToOneUser",
        components: {AddItemDialog},
        comments: {AddItemDialog},
        data() {
            return {
                userId: '',
                title: '',
                context: '',
                hasItems: false,
                items: [],
                addItemDialogVisible: false
            }
        },
        methods: {
            handleDelete(index) {
                this.items.splice(index, 1);
            },
            handleAddItem(event) {
                this.items.push({itemId: event.itemId, itemNum: event.itemNum})
            },
            handleSendMsgClicked(event) {
                console.log(event);
            }
        },
        computed: {
            checkUserId: function () {
                let reg = new RegExp("^[0-9]*$")
                return !reg.test(this.userId)
            }
        }
    }
</script>

<style scoped>
</style>