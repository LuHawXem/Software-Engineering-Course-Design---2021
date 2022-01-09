<template>
  <div class="home">
    <div>
      <h1 style="text-align: center; padding-top: 2rem; padding-bottom: 1.25rem">图书管理系统</h1>
      <v-container>
        <v-row>
          <v-col>
            <v-text-field
                label="书号"
                counter="4"
                maxlength="4"
                clearable
                v-model="form.book_num"
            ></v-text-field>
          </v-col>
          <v-col>
            <v-text-field
                label="书名"
                counter="10"
                maxlength="10"
                clearable
                v-model="form.book_name"
            ></v-text-field>
          </v-col>
          <v-col>
            <v-text-field
                label="作者"
                counter="4"
                maxlength="4"
                clearable
                v-model="form.author"
            ></v-text-field>
          </v-col>
          <v-col>
            <v-text-field
                label="出版社"
                counter="10"
                maxlength="10"
                clearable
                v-model="form.publisher"
            ></v-text-field>
          </v-col>
          <div style="display: flex;align-items: center;">
            <v-btn
                color="info"
                class="ma-2 white--text"
                @click="getData"
            >
              查询
              <v-icon
                  right
                  dark
              >
                mdi-magnify
              </v-icon>
            </v-btn>
          </div>
        </v-row>
      </v-container>
      <v-data-table
          :headers="headers"
          :items="desserts"
          :items-per-page="5"
          class="elevation-5"
          calculate-widths
      >
        <template v-slot:item.actions="{ item }">
          <v-icon
              small
              class="mr-2"
              color="success"
              v-if="!item.borrow"
              @click="borrow(item, 1)"
          >
            mdi-tray-arrow-down
          </v-icon>
          <v-icon
              small
              class="mr-2"
              color="error"
              v-if="item.borrow"
              @click="borrow(item, -1)"
          >
            mdi-tray-arrow-up
          </v-icon>
          <v-icon
              small
              class="mr-2"
              color="error"
              v-if="$store.state.role_id==1 && item.lend == 0"
              @click="remove(item)"
          >
            mdi-trash-can-outline
          </v-icon>
        </template>
      </v-data-table>
    </div>
    <v-btn
        v-if="$store.state.role_id == 1"
        color="info"
        class="ma-2 white--text side"
        fab
        to="/manage"
    >
      <v-icon dark>
        mdi-account-cog
      </v-icon>
    </v-btn>
  </div>
</template>

<script>
export default {
  name: 'Home',
  data: () => ({
    headers: [
      {
        text: 'Id',
        align: 'start',
        value: 'id',
        width: '120',
      },
      {
        text: '书号',
        value: 'book_num',
        width: '120',
      },
      {
        text: '书名',
        value: 'book_name',
        sortable: false,
        width: '150',
      },
      {
        text: '作者',
        value: 'author',
        sortable: false,
        width: '150',
      },
      {
        text: '出版社',
        value: 'publisher',
        sortable: false,
        width: '150',
      },
      {
        text: '分类',
        value: 'category',
        sortable: false,
        width: '120',
      },
      {
        text: '藏书量',
        value: 'collection',
        width: '120',
      },
      {
        text: '借出数',
        value: 'lend',
        width: '120',
      },
      {
        value: 'actions',
        width: '120',
        sortable: false,
      }
    ],
    desserts: [],

    form: {
      book_num: null,
      book_name: null,
      author: null,
      publisher: null,
    }
  }),
  methods: {
    getData: function () {
      this.axios({
        url: '/booklist',
        method: 'post',
        data: {
          username: this.$store.state.username,
          password: this.$store.state.password,
          book_num: this.form.book_num,
          book_name: this.form.book_name,
          author: this.form.author,
          publisher: this.form.publisher,
        }
      }).then(res => {
        for(let i = 0; i < res.data.borrow.length; i++) {
          for(let j = 0; j < res.data.booklist.length; j++) {
            if(res.data.borrow[i].book_id == res.data.booklist[j].id) {
              res.data.booklist[j].borrow = true;
              break;
            }
          }
        }
        this.desserts = res.data.booklist;
      }).catch(err => {
        console.log(err);
      })
    },
    borrow: function (item, count) {
      this.axios({
        url: '/borrow',
        method: 'post',
        data: {
          username: this.$store.state.username,
          password: this.$store.state.password,
          book_id: item.id,
          borrow: count,
        }
      }).then(() => {
        this.getData()
      }).catch(err => {
        console.log(err);
      })
    },
    remove: function (item) {
      this.axios({
        url: '/manage',
        method: 'post',
        data: {
          username: this.$store.state.username,
          password: this.$store.state.password,
          book_id: item.id
        }
      }).then(() => {
        this.getData()
      }).catch(err => {
        console.log(err);
      })
    }
  },
  beforeMount() {
    this.getData();
  }
}
</script>

<style scoped>
.home {
  height: 100%;
  width: 100%;
  display: flex;
  justify-content: center;
}
.side {
  position: absolute;
  right: 1rem;
  bottom: 1rem;
}
</style>
