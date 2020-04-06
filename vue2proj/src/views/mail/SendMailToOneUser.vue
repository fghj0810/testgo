<template>
    <div>
        <el-input v-model="userId" placeholder="用户ID"></el-input>
        <el-input v-model="title" placeholder="邮件标题"></el-input>
        <el-input
                type="textarea"
                :rows="20"
                placeholder="请输入邮件内容"
                v-model="context">
        </el-input>
        <div>
            <span>附件</span>
            <el-table :data="items">
                <el-table-column prop="itemId" label="道具ID" width="180">
                </el-table-column>
                <el-table-column prop="itemNum" label="道具数量" width="180">
                </el-table-column>
                <el-table-column fixed="right" label="操作" width="100">
                    <template slot-scope="scope">
                        <el-button size="mini" type="danger" @click="handleDelete(scope.$index, scope.row)">
                            <span>删除</span>
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>
        <el-button>发送</el-button>
    </div>
</template>

<script>
    export default {
        name: "SendMailToOneUser",
        data() {
            return {
                userId: '',
                title: '',
                context: '',
                items: [{itemId: 1, itemNum: 10}, {itemId: 2, itemNum: 10}]
            }
        },
        methods: {
            handleDelete(index, row) {
                console.log(index, row);
            }
        }
    }
</script>

<style scoped>

</style>