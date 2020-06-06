<template>
  <div>
    <div style="width: 400px" class="assert">
    <el-form :label-position="labelPosition" ref="ruleForm" :model="ruleForm">
      <el-form-item label="资产名称" prop="name"
       :rules="[
      { required: true, message: '资产名称不能为空'}
    ]">
        <el-input v-model="ruleForm.name" type="name"></el-input>
      </el-form-item>
      <el-form-item label="资产地址" prop="address" :rules="[
      { required: true, message: '资产地址不能为空'}
    ]">
        <el-input v-model="ruleForm.address" type="address"
      ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="dialogVisible = true ; submitForm('ruleForm')">立即创建</el-button>
        <el-button @click="resetForm('ruleForm')">重置</el-button>
      </el-form-item>
    </el-form>
    </div>
    <div class="ass">
<!--      右侧  资产名称 资产链接 执行状态  移除资产-->
      <el-table
        :data="assetsAllinfo"
        style="width: 100%">
        <el-table-column label="已存在任务" align="center">
        <el-table-column
          label="资产名称"
          align="center"
          prop="assetsName">
        </el-table-column>
        <el-table-column
          label="资产链接"
          align="center"
          prop="assetsAddress">
        </el-table-column>
        <el-table-column
          label="执行状态"
          align="center"
          prop="execStatus"
          >
          <template slot-scope="scope">
            <span v-if="scope.row.execStatus==0">未执行</span>
            <span v-if="scope.row.execStatus==1">执行完成</span>
            <span v-if="scope.row.execStatus==-1">正在执行</span>
          </template>
        </el-table-column>
          <el-table-column
            label="操作资产"
            align="center"
            >
            <template slot-scope="scope">
              <el-button
                size="mini"
                type="primary" plain
                @click="execTask(scope.$index,scope.row)"
              >执行</el-button>
              <el-button
                size="mini"
                type="danger" plain
                @click="handleDelete(scope.$index, scope.row)">删除</el-button>
            </template>
            </el-table-column>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>
<script>
export default {
  data () {
    return {
      insertStatus: '',
      labelPosition: 'right',
      dialogVisible: false,
      ruleForm: {
        name: '',
        address: ''
      },
      assetsAllinfo: {}
    }
  },
  mounted: function () {
    //生命周期自动执行的方法
    this.$axios.post(`api/api/v1/getAllAssetsInfo`).then((data) =>{
      this.assetsAllinfo = data.data.data
    })
  },
  methods: {
    handleDelete(index, row) {
      console.log(row.assetsName)
      let assetsName = new FormData()
      assetsName.append('assetsName',row.assetsName)
      this.$axios.post(`api/api/v1/removeAssets`,assetsName).then((data) => {
        this.$message({
          message: data.data.message,
          type: 'success',
          duration: 500
        })
      })
      this.$axios.post(`api/api/v1/getAllAssetsInfo`).then((data) =>{
        console.log(data)
        this.assetsAllinfo = data.data.data
        console.log(data.data)
      })
    },
    submitForm (formName) {
      // console.log('成功触发提交事件')
      let data = new FormData()

      data.append('assetsName', this.ruleForm.name)
      data.append('assetsAddress', this.ruleForm.address)
      // eslint-disable-next-line handle-callback-err
      if (this.ruleForm.name === "" || this.ruleForm.address === ""){
        this.$message({
          message: "请不要输入空值",
          type: 'warning',
          duration: 800
        })
      }else{
        this.$axios.post(`api/api/v1/enterAssets`, data).then((data) => {
          console.log(data)
          this.insertStatus = data.data.insertStatus
          console.log(this.insertStatus)
          if (this.insertStatus == "success"){
            this.$message({
              message: "创建成功",
              type: 'success',
              duration: 800
            })
            this.$axios.post(`api/api/v1/getAllAssetsInfo`).then((data) =>{
              console.log(data)
              this.assetsAllinfo = data.data.data

              console.log(data.data)
            })
          }
          if (this.insertStatus == "fail"){
            this.$message({
              message: data.data.message,
              type: 'warning',
              duration: 800
            });
          }
        }, error => {
          console.log(error)
        })
      }

    },
    resetForm (formName) {
      console.log('成功触发重置事件')
      this.$refs[formName].resetFields()
    },
    execTask (index,row) {
      // console.log(row.assetsAddress)
        let data = new FormData()
        data.append('assetsAddress', row.assetsAddress)
        this.$axios.post(`api/api/v1/execTask`, data).then((data) => {
          console.log("AssetsAddress: " + row.assetsAddress + " Start Exec.....")
          this.$message({
            message : "任务开始执行",
            type: 'success',
            duration: 500
          })
        })

      // if (this.insertStatus === 'success') {
      //   this.$message({
      //     message: '任务开始执行',
      //     type: 'success',
      //     duration: 500
      //   })
      //   let data = new FormData()
      //   data.append('assetsAddress', this.ruleForm.address)
      //   this.$axios.post(`api/api/v1/execTask`, data).then((data) => {
      //     console.log("AssetsAddress: " + this.ruleForm.address + " Start Exec.....")
      //   })
      //   // this.insertStatus = ''
      // } else {
      //   this.$message.error('执行失败')
      // }

    }
  }
}
</script>
<style scoped>
  .assert {
    margin-top: 10%;
    margin-left: 6%;
    float: left;
  }
  .ass{
    margin-top: 10%;
    margin-left: 8%;
    width: 650px;
    float: left;
  }
</style>
