<template>
    <div>
        <el-input v-model="input" placeholder="用户ID"></el-input>
        <el-input v-model="input" placeholder="邮件标题"></el-input>
        <el-input
                type="textarea"
                :rows="2"
                placeholder="请输入邮件内容"
                v-model="textarea">
        </el-input>
    </div>
</template>

<script>
    import SendMailToAll from "./mail/SendMailToAll";

    export default {
        name: "Mail",
        components: {SendMailToAll},
        data() {
            const item = {
                date: '2016-05-02',
                name: '王小虎',
                address: '上海市普陀区金沙江路 1518 弄'
            };
            return {
                tableData: Array(20).fill(item)
            }
        }
    }
</script>

<style scoped>

</style>