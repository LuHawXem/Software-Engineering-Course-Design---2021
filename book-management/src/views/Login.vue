<template>
  <div class="login">
    <v-card>
      <v-card-title>
        <span class="text-h5">用户登陆(注册)</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12">
              <v-text-field
                  label="用户名"
                  counter="10"
                  maxlength="10"
                  v-model="username"
              ></v-text-field>
            </v-col>
            <v-col cols="12" style="padding-bottom: 0">
              <v-text-field
                  label="密码"
                  type="password"
                  counter="16"
                  maxlength="16"
                  v-model="password"
              ></v-text-field>
            </v-col>
          </v-row>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn
            color="blue darken-1"
            text
            x-large
            @click="Login"
        >
          登陆
        </v-btn>
        <v-spacer></v-spacer>
      </v-card-actions>
    </v-card>
    <v-snackbar :color="tipColor" top :value="showTips">{{ tipContent }}</v-snackbar>
  </div>
</template>

<script>
export default {
  name: "Login",
  data: () => ({
    username: null,
    password: null,

    tipColor: "success",
    showTips: false,
    tipContent: null,
  }),
  methods: {
    Login: function () {
      if(!this.username || !this.password) {
        this.tipColor = "error";
        this.tipContent = "用户名或密码未填写"
        this.showTips = true;
        setTimeout(() => {
          this.showTips = false
        }, 1500);
        return
      }
      this.axios({
        url: '/login',
        method: 'post',
        data: {
          username: this.username,
          password: this.password,
        }
      }).then(res => {
        console.log(res)
        this.$store.commit('saveData', {
          username: this.username,
          password: this.password,
          role_id: res.role_id
        })
        this.tipColor = "success"
        this.tipContent = "登录成功"
        this.showTips = true;
        setTimeout(() => {
          this.showTips = false;
          this.$router.push({ path: '/home' });
        }, 1500);
      }).catch(err => {
        this.tipColor = "error";
        this.tipContent = err.data.errmsg
        this.showTips = true;
        setTimeout(() => {
          this.showTips = false
        }, 1500);
      })
    }
  }
}
</script>

<style scoped>
.login {
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: darkgrey;
}
.v-card__text {
  padding-bottom: 0 !important;
}
</style>
