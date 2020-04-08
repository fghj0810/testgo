<template>
    <el-dialog title="" width="30%"
               :visible.sync="dVisible"
               :close-on-click-modal=false
               :close-on-press-escape=false
               :show-close=false
               center>
        <span style="white-space: pre-wrap;font-family: 'Consolas'">{{bodyMessage}}</span>
        <span slot="footer" class="dialog-footer">
             <el-button type="primary" @click="handleOKClicked" :disabled="!canClose">确 定</el-button>
        </span>
    </el-dialog>
</template>

<script>
    export default {
        name: "WaitingForMe",
        props: {
            dialogVisible: Boolean,
            bodyMessage: String,
            canClose: Boolean
        },
        data() {
            return {
                dVisible: this.dialogVisible
            }
        },
        methods: {
            handleOKClicked() {
                this.dVisible = false
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