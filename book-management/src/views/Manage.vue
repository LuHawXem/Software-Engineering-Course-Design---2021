<template>
  <div class="users">
    <div>
      <h1 style="text-align: center; padding-top: 2rem; padding-bottom: 1.25rem">借出书籍查看</h1>
      <v-data-table
          :headers="headers"
          :items="desserts"
          :items-per-page="5"
          class="elevation-5"
          calculate-widths
      >
      </v-data-table>
    </div>
    <div>
      <v-dialog
          v-model="show"
          persistent
          max-width="600px"
      >
        <v-card>
          <v-card-title>
            <span class="text-h5">添加书籍</span>
          </v-card-title>
          <v-card-text>
            <v-container>
              <v-row>
                <v-col
                    cols="12"
                    sm="6"
                    md="4"
                >
                  <v-text-field
                      label="书名"
                      counter="12"
                      maxlength="12"
                      clearable
                      v-model="book.book_name"
                  ></v-text-field>
                </v-col>
                <v-col
                    cols="12"
                    sm="6"
                    md="4"
                >
                  <v-text-field
                      label="书号"
                      counter="4"
                      maxlength="4"
                      type="number"
                      clearable
                      v-model="book.book_num"
                  ></v-text-field>
                </v-col>
                <v-col
                    cols="12"
                    sm="6"
                    md="4"
                >
                  <v-text-field
                      label="分类"
                      counter="3"
                      maxlength="3"
                      clearable
                      v-model="book.category"
                  ></v-text-field>
                </v-col>
                <v-col
                    cols="12"
                    sm="6"
                    md="4"
                >
                  <v-text-field
                      label="作者"
                      counter="4"
                      maxlength="4"
                      clearable
                      v-model="book.author"
                  ></v-text-field>
                </v-col>
                <v-col
                    cols="12"
                    sm="6"
                    md="4"
                >
                  <v-text-field
                      label="出版社"
                      counter="4"
                      maxlength="4"
                      clearable
                      v-model="book.publisher"
                  ></v-text-field>
                </v-col>
                <v-col
                    cols="12"
                    sm="6"
                    md="4"
                >
                  <v-text-field
                      label="数量"
                      counter="3"
                      maxlength="3"
                      type="number"
                      clearable
                      v-model="book.collection"
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
                @click="addBook"
            >
              添加
            </v-btn>
            <v-btn
                color="blue darken-1"
                text
                @click="show = false"
            >
              取消
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
    <div class="side opbtn">
      <v-btn
          color="info"
          class="ma-2 white--text"
          fab
          @click="show = true"
      >
        <v-icon dark>
          mdi-plus-circle
        </v-icon>
      </v-btn>
      <v-btn
          color="blue-grey"
          class="ma-2 white--text"
          fab
          to="/home"
      >
        <v-icon dark>
          mdi-arrow-left-circle
        </v-icon>
      </v-btn>
    </div>
    <v-snackbar :color="tipColor" top :value="showTips">{{ tipContent }}</v-snackbar>
  </div>
</template>

<script>
export default {
  name: "Manage",
  data: () => ({
    headers: [
      {
        text: '用户名',
        value: 'name',
        sortable: false,
        width: '150',
      },
      {
        text: '书号',
        value: 'book_num',
        sortable: true,
        width: '150',
      },
      {
        text: '书名',
        value: 'book_name',
        sortable: false,
        width: '150',
      }
    ],
    desserts: [],
    book: {
      book_num: null,
      book_name: null,
      author: null,
      publisher: null,
      category: null,
      collection: null,
    },
    show: false,
    tipColor: "success",
    showTips: false,
    tipContent: null,
  }),
  methods: {
    getData: function () {
      this.axios({
        url: '/manage',
        method: 'post',
        data: {
          username: this.$store.state.username,
          password: this.$store.state.password,
        }
      }).then(res => {
        this.desserts = res.data
      }).catch(err => {
        console.log(err);
      })
    },
    addBook: function () {
      if(!this.book.book_num || !this.book.book_name || !this.book.author || !this.book.publisher || !this.book.category || !this.book.collection) {
        this.tipContent = "表单不完整";
        this.tipColor = "error";
        this.show = false;
        this.showTips = true;
        setTimeout(()=>{
          this.showTips = false;
        }, 1500)
        return;
      }
      this.axios({
        url: '/manage',
        method: 'post',
        data: {
          username: this.$store.state.username,
          password: this.$store.state.password,
          book_num: this.book.book_num,
          book_name: this.book.book_name,
          author: this.book.author,
          publisher: this.book.publisher,
          category: this.book.category,
          collection: this.book.collection,
        }
      }).then(() => {
        this.tipContent = "添加成功";
        this.tipColor = "success";
        this.show = false;
        this.showTips = true;
        setTimeout(()=>{
          this.showTips = false;
        }, 1500);
        return;
      }).catch(err => {
        console.log(err);
        this.tipContent = "出错了";
        this.tipColor = "error";
        this.show = false;
        this.showTips = true;
        setTimeout(()=>{
          this.showTips = false;
        }, 1500);
        return;
      })
    }
  },
  beforeMount() {
    this.getData();
  }
}
</script>

<style scoped>
.users {
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
}
.opbtn {
  display: flex;
  flex-direction: column;
}
</style>
