<template>
  <div>
    <el-table
      :data="tableData.filter(data => !search || data.assetsName.toLowerCase().includes(search.toLowerCase()))"
      style="width: 90%">
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
      <el-table-column label="操作报表" width="200">
        <template slot-scope="scope">
          <el-button
            type="text"
            @click="dialogTableVisible = true;getGetBaseInfoAll(scope.row)">预览报表</el-button>
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
      <el-table
        :data="bugInfoData"
        style="width: 100%">
        <el-table-column
          prop="urladdress"
          label="链接地址"
          width="180">
        </el-table-column>
        <el-table-column
          prop="bugname"
          label="漏洞信息"
          width="180">
        </el-table-column>
        <el-table-column
          prop="bugpoc"
          label="检测poc">
        </el-table-column>
      </el-table>
    </el-dialog>
  </div>
</template>
<script>
  export default {
    mounted: function () {
      this.getAssets()
    },

    data: function () {
      return {
        bb: '123123',

        ip: '',
        address: '',
        dir: '',
        infoData: [{
          "ipdomain": ""
        }],
        dialogTableVisible: false,
        dialogFormVisible: false,
        tableData: {},
        bugInfoData: [],
        formLabelWidth: '120px'
      }
    },
    methods: {
      getAssets () {
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
        this.$axios.post(`api/api/v1/GetBugInfo`, data).then((data) => {
          console.log(data)
          this.bugInfoData = data.data.data
          console.log(this.bugInfoData)
        })
      },
      getBugInfo (address) {
        let data = new FormData()
        data.append("address",address)
        this.$axios.post(`api/api/v1/GetBugInfo`, data).then((data) => {
          console.log(data)
          this.bugInfoData = data.data.data
          console.log(this.bugInfoData)
        })
      }
    }
  }
</script>
