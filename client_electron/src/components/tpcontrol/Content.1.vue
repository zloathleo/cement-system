<template>

  <div class="columns">

    <div class="column is-6">

      <div class="card card-bg-transparent" style="box-shadow: none;padding-bottom: 2rem;">
        <div class="card-header color-primary-4">
          探头状态显示
        </div>
        <div ref="contentPane" class="card-content-table">
          <table class="table is-bordered  is-fullwidth card-bg-transparent">
            <tbody>
              <tr>
                <td class="table-td">电源启动</td>
                <td class="table-td">就地</td>
                <td class="table-td">风机启动/停止</td>
                <td class="table-td"></td>
              </tr>
              <tr>
                <td class="table-td">吹扫风低压报警</td>
                <td class="table-td">冷却风低压报警</td>
                <td class="table-td">冷却风高压报警</td>
                <td class="table-td">温度场探头温度高报警</td>
              </tr>
              <tr>
                <td class="table-td">温度场探头A进</td>
                <td class="table-td">温度场探头A退</td>
                <td class="table-td">温度场探头B进</td>
                <td class="table-td">温度场探头B退</td>
              </tr>
              <tr>
                <td class="table-td">温度场探头A正常</td>
                <td class="table-td"> </td>
                <td class="table-td">温度场探头B正常</td>
                <td class="table-td"> </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <div class="card card-bg-transparent" style="box-shadow: none;padding-bottom: 2rem;">
        <div class="card-header color-primary-4">
          探头控制操作
        </div>
        <div class="card-content-table">
          <table class="table is-bordered  is-fullwidth card-bg-transparent">
            <tbody>
              <tr>
                <td class="table-td">探头A允许进</td>
                <td class="table-td">不允许原因</td>
                <td class="table-td">
                  <a class="button color-primary-4 is-fullwidth" @click="clickJingA">进</a>
                </td>
                <td class="table-td">
                  <a class="button color-primary-4 is-fullwidth" @click="clickTuiA">退</a>
                </td>
              </tr>
              <tr>
                <td class="table-td">探头B允许进</td>
                <td class="table-td">不允许原因</td>
                <td class="table-td">进</td>
                <td class="table-td">退</td>
              </tr>
              <tr>
                <td class="table-td">探头允许上电</td>
                <td class="table-td">不允许原因</td>
                <td class="table-td">上电</td>
                <td class="table-td">断电</td>
              </tr>
              <tr>
                <td class="table-td">风机允许开</td>
                <td class="table-td">不允许原因</td>
                <td class="table-td">开</td>
                <td class="table-td">关</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

    </div>

    <div class="column is-6">
      <div class="card card-bg-transparent" style="box-shadow: none;padding-bottom: 2rem;">
        <div class="card-header color-primary-4">
          Device
        </div>
        <div class="card-content-table" style="display: -webkit-box;">
          <div v-bind:style="{ width: leftGap + 'px' ,height:'100px'}" />
          <transition name="tui-A-ing">
            <img v-if="tuiA" src="assets/img/1.png" class="tui-A">
          </transition>

          <transition name="jing-A-ing">
            <img v-if="jingA" src="assets/img/1.png" class="jing-A">
          </transition>
          <!-- <img v-if="jingA" src="assets/img/4.png" class="jing-A"> -->
          <img src="assets/img/3.png" class="lu-A" v-bind:style="{ left: lutangAImageLeft + 'px' }">
        </div>
      </div>
    </div>

  </div>

</template>

<style scoped lang="less">
//anim 镜头A 进
.jing-A {
  position: relative;
  width: 100px;
  top: 80px;
  left: 60px;
}

.jing-A-ing-enter-active {
  transition: all 6s ease;
}
.jing-A-ing-leave-active {
  transition: all 0s;
}
.jing-A-ing-enter {
  position: relative;
  width: 100px;
  top: 80px;
  left: 0px;
}
.jing-A-ing-leave-to {
  transform: translateX(-60px);
}

//anim 镜头A 退 在外
.tui-A {
  position: relative;
  width: 100px;
  top: 80px;
  left: 0px;
}
.tui-A-ing-enter-active {
  transition: all 6s ease;
}
.tui-A-ing-leave-active {
  transition: all 0s;
}
.tui-A-ing-enter {
  position: relative;
  width: 100px;
  top: 80px;
  left: 60px;
}
.tui-A-ing-leave-to {
  transform: translateX(60px);
}
//anim 镜头A 退

.lu-A {
  position: absolute;
  top: 120px;
  width: 300px;
}
</style>

<script> 
export default {
  components: {
  },
  data() {

    return {
      leftGap: 40,
      lutangAImageLeft: 160,//A炉膛left 

      jingA: false,//A探头进到位
      tuiA: true,//A探头退到位
      jingB: false,//B探头进到位
      tuiB: true,//B探头退到位
    }
  },
  mounted() {
    let _this = this;
    let contentPane = this.$refs.contentPane;
    let contentPaneWidth = contentPane.clientWidth;
    if (contentPaneWidth > 550) {
      this.leftGap = 40 + (contentPaneWidth - 550) * 0.5;
      this.lutangAImageLeft = 160 + (contentPaneWidth - 550) * 0.5;
    }
    console.log(contentPane);

  },
  methods: {
    clickJingA() {
      this.jingA = true;
      this.tuiA = false;

      let _this = this;

      let rows = [{
        name: "2_00018",
        value: 1
      }];
      let content = {
        rows: rows,
      };
      let _json = JSON.stringify(content, null, '    ');
      console.log("_json:", _json);
      let params = {
        'content': _json,
      };

      _this.$myaxios.post('/control', params).then(function (response) {
        _this.data = response.data;

        console.log("_json:", _json);
        // resolve(response);
      }).catch(function (err) {
        console.error(err);
        // reject(err);
      });

    },
    clickTuiA() {
      this.jingA = false;
      this.tuiA = true;

      let _this = this;

      let rows = [{
        name: "2_00018",
        value: 0
      }];
      let content = {
        rows: rows,
      };
      let _json = JSON.stringify(content, null, '    ');
      console.log("_json:", _json);
      let params = {
        'content': _json,
      };

      _this.$myaxios.post('/control', params).then(function (response) {
        _this.data = response.data;

        console.log("_json:", _json);
        // resolve(response);
      }).catch(function (err) {
        console.error(err);
        // reject(err);
      });
    }
  }
}
</script>
