<template>
    <el-tabs tab-position="left">
        <el-tab-pane label="向单个用户发送邮件">
            <send-mail-to-one-user/>
        </el-tab-pane>
        <el-tab-pane label="向全体用户发送邮件">
            <send-mail-to-all/>
        </el-tab-pane>
        <el-tab-pane label="角色管理">查询邮件</el-tab-pane>
    </el-tabs>
    <!--    <el-container>-->
    <!--        <el-aside width="200px" style="background-color: rgb(238, 241, 246)">-->
    <!--            <el-menu align="left">-->
    <!--                <el-menu-item index="1"><i class="el-icon-message"/>向单个用户发送邮件</el-menu-item>-->
    <!--                <el-menu-item index="2"><i class="el-icon-message"/>向全体用户发送邮件</el-menu-item>-->
    <!--                <el-menu-item index="2"><i class="el-icon-message"/>查询邮件</el-menu-item>-->
    <!--            </el-menu>-->
    <!--        </el-aside>-->

    <!--        <el-container>-->
    <!--            <el-main style="background-color: rgb(238, 241, 246)">-->
    <!--                <send-mail-to-one-user/>-->
    <!--                <send-mail-to-all/>-->
    <!--            </el-main>-->
    <!--        </el-container>-->
    <!--    </el-container>-->
</template>

<script>
    import SendMailToOneUser from "./mail/SendMailToOneUser.vue";
    import SendMailToAll from "./mail/SendMailToAll";

    export default {
        name: "Mail",
        components: {SendMailToOneUser, SendMailToAll},
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