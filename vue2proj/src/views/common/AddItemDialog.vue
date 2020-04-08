<template>
    <el-dialog title="添加道具" :visible.sync="dVisible" width="30%">
        <div><span>道具ID：</span>
            <el-input v-model="itemId" placeholder="请输入道具ID"></el-input>
        </div>
        <div><span>数量：</span>
            <el-input v-model="itemNum" placeholder="请输入道具数量"></el-input>
        </div>
        <span slot="footer" class="dialog-footer">
             <el-button @click="handleCancelClicked">取 消</el-button>
             <el-button type="primary" @click="handleOKClicked" :disabled="dataInvalid">确 定</el-button>
        </span>
    </el-dialog>
</template>

<script>
    export default {
        name: "AddItemDialog",
        props: {
            dialogVisible: Boolean
        },
        data() {
            return {
                dVisible: this.dialogVisible,
                itemId: '',
                itemNum: '',
                // dataInvalid: false
            }
        },
        methods: {
            handleCancelClicked() {
                this.$emit('update:dialogVisible', false)
            },
            handleOKClicked() {
                this.$emit('update:dialogVisible', false)
                let arg = {itemId: this.itemId, itemNum: this.itemNum}
                this.$emit('onAddItem', arg)
            }
        },
        computed: {
            dataInvalid: function () {
                return !(parseInt(this.itemId) > 0 && parseInt(this.itemNum) > 0);
            }
        },
        watch: {
            dialogVisible(val) {
                if (val != this.dVisible) {
                    this.dVisible = val
                }
            },
            dVisible(val) {
                this.$emit('update:dialogVisible', val)
            }
        }
    }
</script>

<style scoped>

</style>