<template>
  <div>
    <div  width = "50%" class="rside">
<el-form
  ref="form"
  :model="form"
  label-width="80px"
  style="width: 90%"
>
  <el-form-item label="插件名称">
    <el-input v-model="form.name"></el-input>
  </el-form-item>
  <el-form-item label="插件类别" align="left">
    <el-select v-model="form.region" placeholder="请选择插件类别">
      <el-option label="web相关" value="applicationOfWeb"></el-option>
      <el-option label="主机相关" value="host"></el-option>
    </el-select>
  </el-form-item>
  <el-form-item label="插件描述">
    <el-input type="textarea" v-model="form.describe" rows="4"></el-input>
  </el-form-item>
  <el-form-item label="漏洞危害">
    <el-input type="textarea" v-model="form.harm" rows="4"></el-input>
  </el-form-item>
  <el-form-item label="文件上传">
  <el-upload
    align="left"
    class="upload-demo"
    :auto-upload= false
    drag
    action="api/Fileupload"
    :data="form"
    ref="uploadData"
    multiple>
      <i class="el-icon-upload"></i>
    <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
  </el-upload>
  </el-form-item>

  <el-form-item>
    <el-button type="primary" @click="onSubmit">立即创建</el-button>
    <el-button>取消</el-button>
  </el-form-item>
</el-form>
  </div>
  <div width = "50%" class="side">
    <el-table
      :data="plugininfo"
      style="width: 100%">
      <el-table-column label="已存在插件" align="center">
        <el-table-column
          label="插件名称"
          align="center"
          prop="pocname"
          width="130px"
        >
        </el-table-column>
        <el-table-column
          label="插件类别"
          align="center"
          prop="pocclass">
        </el-table-column>
        <el-table-column
          label="下载插件"
          align="center"
        >
          <template slot-scope="scope">
          <el-button type="primary" plain="true" size="small "  @click="downloadScript(scope.row)">下载<i class="el-icon-download
 el-icon--right"></i></el-button>
          </template>
        </el-table-column>
        <el-table-column
          label="移除插件"
          align="center"
        >
          <template slot-scope="scope">
          <el-button type="danger" plain size="small" @click="removePlugininfo(scope.row)">移除<i class="el-icon-delete
 el-icon--right"></i></el-button>
          </template>
        </el-table-column>
      </el-table-column>
    </el-table>
  </div>
  </div>

</template>
<script>
  export default {
    mounted: function () {
      //生命周期自动执行的方法
        this.getpocinfo()
    },
    data() {
      return {
        plugininfo:{},
        form: {
          name: '',
          region: '',
          delivery: false,
          type: [],
          resource: '',
          describe: '',
          harm: '' //危害
        }
      }
    },
    methods: {
      getpocinfo(){
        this.$axios.post(`api/api/v1/GetPocInfo`).then((data) =>{
          this.plugininfo = data.data.data
        })
      },
      onSubmit() {
        console.log('submit!');
        this.$refs.uploadData.submit()
        this.getpocinfo()
      },
      downloadScript(row){
          // alert(row.pocfilePath)
        let data =  new FormData()
        data.append("scriptname",row.pocfilePath)
        const fileName = row.pocfilePath.split("/")[2]
        this.$axios({
          method: 'post',
          url:  `api/api/v1/download`,
          data: data,
          responseType: 'blob'
        }).then(response => {
          this.download(response.data,fileName)
        }).catch((error) => {

        })
      },
      download (data,fileName) {
        if (!data) {
          return
        }
        let url = window.URL.createObjectURL(new Blob([data]))
        let link = document.createElement('a')
        link.style.display = 'none'
        link.href = url

        link.setAttribute('download', fileName)

        document.body.appendChild(link)
        link.click()
      },
      removePlugininfo(row){
        const pluginname = row.pocname
        let data =  new FormData()
        data.append("pocname",pluginname)
        this.$axios.post(`api/api/v1/RemovePlugin`,data).then((res)=>{
          console.log("移除成功")
        })
        this.getpocinfo()
      }
    },

  }
</script>
<style scoped>
  .rside{
    float: left;
    width: 50%;
  }
  .side{
    margin-left: 0%;
    width: 50%;
    float: right;
  }
</style>
