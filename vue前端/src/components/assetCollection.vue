<template>
  <div>
    <el-table
      :data="tableData"
      style="width: 90%">
      <el-table-column
        label="开始时间"
        prop="startTime"
        width="190">
      </el-table-column>
      <el-table-column
        label="结束时间"
        prop="endTime"
        width="190">

      </el-table-column>
      <el-table-column
        prop="assetsName"
        label="资产名称"
        width="200">
      </el-table-column>
        <el-table-column
          prop="assetsAddress"
          label="资产地址"
          width="300">
        </el-table-column>
      <el-table-column label="操作" width="200">
        <template slot-scope="scope">
          <el-button
            type="text"
            @click="dialogTableVisible = true;getGetBaseInfoAll(scope.row)">查看详情</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-dialog title="资产详情" :visible.sync="dialogTableVisible">
<!--      <div style="width: 100%">-->
        <h1 style="width: 100%" align="left" >资产ip</h1>
        <h1 style="width: 100%" align="left" >{{this.infoData[0].ipdomain}}</h1>
        <h1 style="width: 100%" align="left" >资产端口开放情况</h1>
        <div style="width: 100%" align="left">{{this.infoData[0].port}}</div>
        <h1 style="width: 100%" align="left">资产网页内嵌链接</h1>
        <div style="width: 100%" align="left" v-html="infoData[0].urladdress"></div>
        <h1 style="width: 100%" align="left">资产目录开放情况</h1>
        <div style="width: 100%" align="left" v-html="infoData[0].dir"></div>
<!--      </div>-->
    </el-dialog>
  </div>
</template>
<script>
export default {
  mounted: function () {
    console.log('生命周期自动执行的方法')
    this.getAssets()
  },

  data: function () {
    return {
      bb: '123123',

      ip: '',
      address: '',
      dir: '',
      infoData: [{
        "ipdomain":""
      }],
      dialogTableVisible: false,
      dialogFormVisible: false,
      tableData: {},

      formLabelWidth: '120px'
    }
  },
  methods: {
    getAssets (){
      this.$axios.post(`api/api/v1/getAssetsAddress`).then((data) => {
        this.tableData = data.data.data
      }, error => {
        console.log(error)
      })
    },
    getGetBaseInfoAll (row) {
      let data = new FormData()
      data.append("address", row.assetsAddress)
      this.$axios.post(`api/api/v1/getBaseInfo`, data).then(res => {
        this.infoData = res.data.data
      })
    }
  }
}
</script>
