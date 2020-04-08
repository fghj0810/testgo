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
        </div>
        <add-item-dialog :dialog-visible.sync="addItemDialogVisible" v-on:onAddItem="handleAddItem"/>
        <waiting-for-me :dialog-visible.sync="waitingDialog.visible"
                        :body-message="waitingDialog.message"
                        :can-close="waitingDialog.canClose"/>
    </div>
</template>

<script>
    import AddItemDialog from "../common/AddItemDialog";
    import WaitingForMe from "../common/WaitingForMe";
    import axios from "axios";

    export default {
        name: "SendMailToOneUser",
        components: {AddItemDialog, WaitingForMe},
        data() {
            return {
                userId: '',
                title: '',
                context: '',
                hasItems: false,
                items: [],
                addItemDialogVisible: false,
                waitingDialog: {
                    visible: false,
                    message: '',
                    canClose: false
                }
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
                this.waitingDialog.canClose = false
                this.waitingDialog.visible = true
                this.waitingDialog.message = "正在执行操作……"
                axios.post('/api/auth/do_auth', "")
                    .then(res => {
                        console.log(res)
                        if (res.status == 200) {
                            this.waitingDialog.message = JSON.stringify(res, null, 4)
                        } else {
                            this.waitingDialog.message = "登录失败"
                        }
                        this.waitingDialog.canClose = true
                    })
                    .catch(err => {
                        this.waitingDialog.message = JSON.stringify(err, null, 4)
                        this.waitingDialog.canClose = true
                    })
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