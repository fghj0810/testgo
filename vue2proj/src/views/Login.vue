<template>
    <div id="bg" class="bg">
        <div class="login" @keyup.13="doLogin">
            <div class="form-horizontal login">
                <div class="logo">客服工具</div>
                <div class="form-group input-group input-group-lg ">
                    <el-input v-model="userInfo.username" placeholder="请输入用户名"></el-input>
                </div>
                <div class="form-group input-group input-group-lg">
                    <el-input placeholder="请输入密码" v-model="userInfo.password" show-password></el-input>
                </div>
                <div class="form-group">
                    <el-button type="primary" round @click="doLogin">登 录</el-button>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import axios from 'axios';

    export default {
        name: 'login',
        data() {
            return {
                userInfo: {
                    username: '',
                    password: '',
                },
                show: false,
            }
        },
        methods: {
            doLogin() {
                if (this.userInfo.username == '') {
                    alert('用户名不能为空');
                    return false
                }
                if (this.userInfo.password == '') {
                    alert('密码名不能为空');
                    return false
                }
                axios.post('/api/auth/do_auth', JSON.stringify(this.userInfo))
                    .then(res => {
                        console.log(res)
                        if (res.status == 200) {
                            if (res.data.code == 0) {
                                this.$store.commit('setUserInfo', {username: res.data.username, token: res.data.token})
                                this.$router.push({path: '/'})
                            } else if (res.data.code == 10005) {
                                alert("用户名或密码不正确")
                            } else {
                                alert("登录失败")
                            }
                        } else {
                            alert("登录失败")
                        }
                    })
                    .catch(err => {
                        alert("登录异常")
                        console.log(err)
                    })
            }
        },
        mounted() {
            var wi = window.screen.width;
            var hi = window.screen.height;
            document.getElementById("bg").style.width = wi + "px";
            document.getElementById("bg").style.height = hi + "px";
        },
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    /*.bg {*/
    /*!*background-color: aqua;*!*/
    /*background-image: url("../assets/bj.jpg");*/
    /*background-size:100% 100%*/
    /*}*/
    .login {
        position: absolute;
        top: 50%;
        left: 50%;
        -webkit-transform: translate(-50%, -50%);
        -moz-transform: translate(-50%, -50%);
        -ms-transform: translate(-50%, -50%);
        -o-transform: translate(-50%, -50%);
        transform: translate(-50%, -50%);
        width: 400px;

    }

    .login-btn {
        background-color: whitesmoke;
    }

    .logo {
        font-family: "DejaVu Sans Mono";
        color: lightblue;
        font-size: 50px;
    }
</style>