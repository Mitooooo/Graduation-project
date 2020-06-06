<template>
  <div style="width: 70%" class="t1">
    <div style="margin-top: 15px;">
      <el-input placeholder="请输入检测网址" v-model="input" class="input-with-select">
        <el-select  v-model="select"  slot="prepend" placeholder="请选择检测插件" style="width: 150px">
          <el-option
            v-for="item in PocNameAll"
            :label="item.pocname"
            :value="{key:item.pocname,value:item.pocfilePath}"
          ></el-option>
        </el-select>
        <el-button slot="append" icon="el-icon-search" @click="check()">检测</el-button>

      </el-input>
    </div>
    <div class="t2">
      <el-button type="text" @click="dialogTableVisible = true;viewResults()">查看执行结果</el-button>

<el-dialog title="执行结果" :visible.sync="dialogTableVisible" width="70%">
  <el-table :data="SingletonCheckResult">
    <el-table-column property="pluginname" label="插件名称" width="160"></el-table-column>
    <el-table-column property="pluginpath" label="插件路径" width="290"></el-table-column>
    <el-table-column property="returninfo" label="检测结果" width="460"></el-table-column>
    <el-table-column property="urladdress" label="检测链接" width="100"></el-table-column>
  </el-table>
</el-dialog>
    </div>
    <div class="t3">
        <el-button type="text" @click="table = true;viewAllResults()">查看历史执行结果</el-button>
        <el-drawer
          title="检测结果!"
          :visible.sync="table"
          direction="rtl"
          size="70%">
           <el-table :data="SelectSingkeAllInfo">
              <el-table-column property="pluginname" label="插件名称" width="160"></el-table-column>
              <el-table-column property="pluginpath" label="插件路径" width="290"></el-table-column>
             <el-table-column property="returninfo" label="检测结果" width="490"></el-table-column>
             <el-table-column property="urladdress" label="检测链接" width="100"></el-table-column>

           </el-table>
        </el-drawer>
    </div>
  </div>
</template>

<script>
export default {
  name: 'singletonDetection',
  data() {
    return {
      table: false,
      dialog: false,
      loading: false,
      dialogTableVisible: false,
      input: '',
      select: "",
      PocNameAll:{},
      result:'',
      resulthistory:'',
      SelectSingkeAllInfo:'',
      SingletonCheckResult:'',
    }
  },
  mounted: function () {
    this.GetPocNameAll()
  },
  methods: {
    // handleClose(done) {
    //   if (this.loading) {
    //     return;
    //   }
    //   this.$confirm('确定要提交表单吗？')
    //     .then(_ => {
    //       this.loading = true;
    //       this.timer = setTimeout(() => {
    //         done();
    //         // 动画关闭需要一定的时间
    //         setTimeout(() => {
    //           this.loading = false;
    //         }, 400);
    //       }, 2000);
    //     })
    //     .catch(_ => {});
    // },
    // cancelForm() {
    //   this.loading = false;
    //   this.dialog = false;
    //   clearTimeout(this.timer);
    // },
    viewResults(){
      let data =  new FormData()
      data.append("pluginname",this.select['key'])
      data.append("urladdress",this.input)
      data.append("pluginpath",this.select['value'])
      this.$axios.post(`api/api/v1/SingletonCheckResult`,data).then((data)=>{
        this.SingletonCheckResult = data.data.data
      })
    },
    viewAllResults(){
      this.$axios.post(`api/api/v1/SelectSingkeAllInfo`).then((data)=>{
        this.SelectSingkeAllInfo = data.data.data
        // this.SingletonCheckResult = ""
      })
    },
    check(){
      let data =  new FormData()
      data.append("pluginname",this.select['key'])
      data.append("urladdress",this.input)
      data.append("pluginpath",this.select['value'])
      this.$axios.post(`api/api/v1/SingletonCheck`,data).then((data)=>{
        console.log(2)
      })
    },
    changeLocationValue(val){
      let obj = {};
      obj = this.PocNameAll.find((item)=>{
        return item.id === val;
      });
      console.log(obj)
    },
    GetPocNameAll () {
      this.$axios.post(`api/api/v1/GetPocNameAll`).then((data) => {
        this.PocNameAll = data.data.data
      })
      console.log(this.PocNameAll)
    },
    checkscript() {
      alert(this.pocname)
      if (this.pocfilepath == "" || this.input == "") {
        this.$message({
          message: "插件名称或输入链接为空",
          type: 'warning',
          duration: 800
        });
      } else {
        let single = new FormData()
        single.append('pluginname', this.pocname)
        single.append('pluginpath', this.pocfilepath)
        single.append('urladdress', this.input)
        //SingletonCheck
        this.$axios.post(`api/api/v1/SingletonCheck`, single).then((data) => {

        })
      }
    }
    }
}
</script>

<style scoped>
  .el-select .el-input {
    width: 130px;
  }
  .input-with-select .el-input-group__prepend {
    background-color: #fff;
  }
  .t1{
    margin-top: 10%;
  }
  .t2{
    margin-top: 20%;
  }
  .t3{
    margin-top: 3%;
  }
</style>
